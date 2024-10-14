package core

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"
	"math/big"
)

type TaskResponse struct {
	TaskID *big.Int
	Msg    string
}

// AbiEncodeTaskResponse this hardcodes abi.encode() for cstaskmanager.IAvsTaskManagerTaskResponse
// unclear why abigen doesn't provide this out of the box...
func AbiEncodeTaskResponse(h *TaskResponse) ([]byte, error) {

	// The order here has to match the field ordering of cstaskmanager.IAVSTaskManagerTaskResponse
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "numberSquared",
			Type: "uint256",
		},
	})
	if err != nil {
		return nil, err
	}
	arguments := abi.Arguments{
		{
			Type: taskResponseType,
		},
	}

	bytes, err := arguments.Pack(h)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GetTaskResponseDigest returns the hash of the TaskResponse, which is what operators sign over
func GetTaskResponseDigest(h *TaskResponse) ([32]byte, error) {

	encodeTaskResponseByte, err := AbiEncodeTaskResponse(h)
	if err != nil {
		return [32]byte{}, err
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}
