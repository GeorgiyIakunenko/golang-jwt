package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

var formatting = []byte{',', '\n'}

func Create(filename string, entity interface{}) error {
	jsonBytes, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	jsonBytes = append(jsonBytes, '\n')

	file, err := os.OpenFile(fmt.Sprintf("internal/db/repos/%s.txt", filename), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil

}
