package core

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type TaskResponse struct {
	TaskID    uint64
	NumberSum *big.Int
}

// GetTaskResponseDigestEncodeByjson returns the hash of the TaskResponse, which is what operators sign over
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
