package storage

import (
	"encoding/json"
	"os"
)

type Experience struct {
	Hits   [][2]int `json:"hits"`
	Misses [][2]int `json:"misses"`
}

func LoadExperience(filename string) (*Experience, error) {
	file, err := os.Open(filename)
	if err != nil {
		return &Experience{}, nil
	}
	defer file.Close()

	exp := &Experience{}
	err = json.NewDecoder(file).Decode(exp)
	if err != nil {
		return &Experience{}, err
	}
	return exp, nil
}

func SaveExperience(filename string, exp *Experience) error {

	dir := "storage"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(exp)
}
