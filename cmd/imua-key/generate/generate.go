package generate

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
	blscommon "github.com/prysmaticlabs/prysm/v5/crypto/bls/common"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	utils "github.com/imua-xyz/imua-avs-sdk/crypto/bls"
	"github.com/imua-xyz/imua-avs-sdk/crypto/ecdsa"
	"github.com/urfave/cli/v2"
)

const (
	DefaultKeyFolder = "keys"

	PasswordFile      = "password.txt"
	PrivateKeyHexFile = "private_key_hex.txt"
)

var (
	KeyTypeFlag = &cli.StringFlag{
		Name:     "key-type",
		Usage:    "key type to create (ecdsa or bls)",
		Required: true,
	}
	NumKeysFlag = &cli.IntFlag{
		Name:        "num-keys",
		Usage:       "number of keys to create",
		Required:    true,
		DefaultText: "1",
	}
	OutputDirFlag = &cli.StringFlag{
		Name:     "output-dir",
		Usage:    "folder to store keys",
		Required: false,
	}
)

var Command = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Description: `Generate keys for testing purpose.
This command creates ecdsa or bls key pair for testing purposes. It generates
all the relevant files for reading and keys and decrypted it and also gets 
you the private keys in plaintext.

It creates the following artifacts based on arguments
- passwords.txt - contains all passwords to decrypt keys
- private_key_hex.txt - will create plaintext private keys
- keys/* - create all the encrypted json files in this folder
`,
	Action: generate,
	Flags: []cli.Flag{
		KeyTypeFlag,
		NumKeysFlag,
		OutputDirFlag,
	},
}

// KeyGenerator defines the interface for generating keys
type KeyGenerator interface {
	GenerateKeys(numKeys int, path string, passwordFile, privateKeyFile *os.File) error
}

// ECDSAKeyGenerator implements KeyGenerator for ECDSA keys
type ECDSAKeyGenerator struct{}

func (g ECDSAKeyGenerator) GenerateKeys(numKeys int, path string, passwordFile, privateKeyFile *os.File) error {
	return generateECDSAKeys(numKeys, path, passwordFile, privateKeyFile) // Existing logic from generateECDSAKeys
}

// BLSKeyGenerator implements KeyGenerator for BLS keys
type BLSKeyGenerator struct{}

func (g BLSKeyGenerator) GenerateKeys(numKeys int, path string, passwordFile, privateKeyFile *os.File) error {
	return generateBlsKeys(numKeys, path, passwordFile, privateKeyFile) // Existing logic from generateBlsKeys
}

// NewKeyGenerator is the factory function to create a KeyGenerator based on the key type
func NewKeyGenerator(keyType string) KeyGenerator {
	switch keyType {
	case "ecdsa":
		return ECDSAKeyGenerator{}
	case "bls":
		return BLSKeyGenerator{}
	default:
		return nil
	}
}

// Modified generate function using the factory pattern
func generate(c *cli.Context) error {
	keyType := c.String(KeyTypeFlag.Name)
	numKeys := c.Int(NumKeysFlag.Name)

	generator := NewKeyGenerator(keyType)
	if generator == nil {
		return cli.Exit("Invalid key type", 1)
	}

	folder, err := createDir(c, keyType+"-")
	if err != nil {
		return err
	}

	passwordFile, privateKeyFile, err := createPasswordAndPrivateKeyFiles(folder)
	if err != nil {
		return err
	}

	err = generator.GenerateKeys(numKeys, folder, passwordFile, privateKeyFile)
	if err != nil {
		return err
	}

	return nil
}

func createDir(c *cli.Context, prefix string) (fileName string, err error) {
	folder := c.String(OutputDirFlag.Name)
	if folder == "" {
		id, err := uuid.NewUUID()
		if err != nil {
			return "", err
		}
		folder = prefix + id.String()
	}

	_, err = os.Stat(folder)
	if !os.IsNotExist(err) {
		return "", err
	}

	// Clean the path
	cleanFilePath := filepath.Clean(folder + "/" + DefaultKeyFolder)

	err = os.MkdirAll(cleanFilePath, 0755)
	if err != nil {
		return "", err
	}
	return folder, nil
}

func createPasswordAndPrivateKeyFiles(folder string) (passwordFile, privateKeyFile *os.File, err error) {

	passwordFile, err = os.Create(filepath.Clean(folder + "/" + PasswordFile))
	if err != nil {
		return nil, nil, err
	}
	privateKeyFile, err = os.Create(filepath.Clean(folder + "/" + PrivateKeyHexFile))
	if err != nil {
		return nil, nil, err
	}
	return passwordFile, privateKeyFile, nil
}

func generateBlsKeys(numKeys int, path string, passwordFile, privateKeyFile *os.File) error {
	for i := 0; i < numKeys; i++ {
		//key, err := bls.GenRandomBlsKeys()
		key, err := blst.RandKey()
		if err != nil {
			return err
		}

		password := generateRandomPassword()
		if err != nil {
			return err
		}

		privateKeyHex := hex.EncodeToString(key.Marshal())
		fileName := fmt.Sprintf("%d.bls.key.json", i+1)
		err = SaveToFile(key, filepath.Clean(path+"/"+DefaultKeyFolder+"/"+fileName), password)
		if err != nil {
			return err
		}

		_, err = passwordFile.WriteString(password + "\n")
		if err != nil {
			return err
		}

		_, err = privateKeyFile.WriteString(privateKeyHex + "\n")
		if err != nil {
			return err
		}

		if (i+1)%50 == 0 {
			fmt.Printf("Generated %d keys\n", i+1)
		}
	}
	return nil
}

func generateECDSAKeys(numKeys int, path string, passwordFile, privateKeyFile *os.File) error {
	for i := 0; i < numKeys; i++ {
		key, err := crypto.GenerateKey()
		privateKeyHex := hex.EncodeToString(key.D.Bytes())

		// Check if the length of privateKeyHex is 32 bytes (64 characters)
		lenPrivateKey := len(privateKeyHex)
		if lenPrivateKey != 64 {
			fmt.Printf("Private key Ignore: %s %d\n", privateKeyHex, lenPrivateKey)
			// Reset count
			i--
			continue
		}

		if err != nil {
			return err
		}

		password := generateRandomPassword()
		if err != nil {
			return err
		}

		fileName := fmt.Sprintf("%d.ecdsa.key.json", i+1)
		err = ecdsa.WriteKey(filepath.Clean(path+"/"+DefaultKeyFolder+"/"+fileName), key, password)
		if err != nil {
			return err
		}

		_, err = passwordFile.WriteString(password + "\n")
		if err != nil {
			return err
		}

		_, err = privateKeyFile.WriteString("0x" + privateKeyHex + "\n")
		if err != nil {
			return err
		}

		if (i+1)%50 == 0 {
			fmt.Printf("Generated %d keys\n", i+1)
		}
	}
	return nil
}

func generateRandomPassword() string {
	// Seed the random number generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define character sets for the password
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	//specialSymbols := "!@#$%^&*()-_+=[]{}|;:,.<>?/\\"

	// Combine character sets into one
	//allCharacters := uppercaseLetters + lowercaseLetters + digits + specialSymbols
	allCharacters := uppercaseLetters + lowercaseLetters + digits

	// Length of the password you want
	passwordLength := 20

	// Generate the password
	password := make([]byte, passwordLength)
	for i := range password {
		password[i] = allCharacters[random.Intn(len(allCharacters))]
	}
	return string(password)
}
func SaveToFile(key blscommon.SecretKey, path string, password string) error {
	cryptoStruct, err := keystore.EncryptDataV3(
		key.Marshal(),
		[]byte(password),
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	if err != nil {
		return err
	}

	encryptedBLSStruct := utils.EncryptedBLSKeyJSONV3{
		PubKey: hex.EncodeToString(key.PublicKey().Marshal()),
		Crypto: cryptoStruct,
	}
	data, err := json.Marshal(encryptedBLSStruct)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directories:", err)
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
