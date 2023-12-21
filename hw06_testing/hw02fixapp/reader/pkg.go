package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/VladislavLisovenko/hw-vladl/hw06_testing/hw02fixapp/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []types.Employee

	_ = json.Unmarshal(bytes, &data)

	return data, err
}