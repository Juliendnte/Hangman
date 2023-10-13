package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteWord() {
	f, err := ReadLines("gutenberg.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	ale := rand.Intn(len(f) - 1)
	fmt.Println(f[ale])
}
