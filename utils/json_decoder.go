package utils

import (
	"encoding/json"
	"os"
)

func DecodeJSON(filePath string, model any) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&model); err != nil {
		return err
	}

	return nil
}
