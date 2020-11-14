package src

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
)

const (
	HashSHA256 = "sha256"
)

var h hash.Hash

func Hash(method string, obj interface{}) (string, error) {
	input, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	switch method {
	case HashSHA256:
		h = sha256.New()
	default:
		return "", errors.New(fmt.Sprintf("%s hash method not exists", method))
	}

	defer h.Reset()
	h.Write(input)
	return hex.EncodeToString(h.Sum(nil)), nil
}
