package parser

import (
	"bufio"
	"os"
	"strings"
)

type Parser struct{}

func (p Parser) Parse(path string, delimiter string) ([]string, error) {
	var results []string

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, delimiter)
		results = append(results, parts...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
