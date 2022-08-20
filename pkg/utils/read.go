package utils

import (
	"bufio"
	"os"
)

func ReadFromFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}
	return result
}
