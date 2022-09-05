package infrastructure

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Load() {
	filePath := ".env"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("%s", err)
	}

	defer file.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("%s", err)
	}

	for _, line := range lines {
		pair := strings.Split(line, "=")
		_ = os.Setenv(pair[0], pair[1])
	}
}
