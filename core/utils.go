package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type TaskResponse struct {
	TaskID            uint64
	NumberToBeSquared uint64
}

// MarshalTaskResponse GetTaskResponseDigestEncodeByjson returns the hash of the TaskResponse, which is what operators sign over
// MarshalTaskResponse marshals the TaskResponse struct into JSON bytes.
func MarshalTaskResponse(h TaskResponse) ([]byte, error) {
	return json.Marshal(h)
}

// UnmarshalTaskResponse unmarshals the JSON bytes into a TaskResponse struct.
func UnmarshalTaskResponse(jsonData []byte) (TaskResponse, error) {
	var taskResponse TaskResponse
	err := json.Unmarshal(jsonData, &taskResponse)
	return taskResponse, err
}

// GetTaskResponseDigestEncodeByjson returns the hash of the TaskResponse, which is what operators sign over.
func GetTaskResponseDigestEncodeByjson(h TaskResponse) ([32]byte, []byte, error) {
	jsonData, err := MarshalTaskResponse(h)
	if err != nil {
		return [32]byte{}, []byte{}, err
	}
	taskResponseDigest := crypto.Keccak256Hash(jsonData)
	return taskResponseDigest, jsonData, nil
}

func UpdateYAMLWithComments(filePath, key, newValue string) error {
	if newValue == "" {
		panic("param is nil")
	}
	// Read the original YAML file content
	data, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	// Parse YAML using yaml.v3 node parser to preserve comments
	var doc yaml.Node
	err = yaml.Unmarshal(data, &doc)
	if err != nil {
		return err
	}
	// Iterate through YAML content to find and update the specified key
	for i := 0; i < len(doc.Content[0].Content); i += 2 {
		if doc.Content[0].Content[i].Value == key {
			doc.Content[0].Content[i+1].Kind = yaml.ScalarNode
			doc.Content[0].Content[i+1].Value = newValue
			doc.Content[0].Content[i+1].Tag = "tag:yaml.org,2002:str"
			break
		}
	}

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	err = encoder.Encode(&doc)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, buf.Bytes(), 0644)
}
func GetFileInCurrentDirectory(filename string) (string, error) {
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Construct full file path
	fullPath := filepath.Join(currentDir, filename)

	// Check if file exists
	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("file %s not found in current directory", filename)
	}

	return fullPath, nil
}
func ConvertToEthAddresses(strArray []string) []common.Address {
	var ethAddresses []common.Address

	if len(strArray) > 0 {
		for _, str := range strArray {
			address := common.HexToAddress(str)
			ethAddresses = append(ethAddresses, address)
		}
	}

	return ethAddresses
}
