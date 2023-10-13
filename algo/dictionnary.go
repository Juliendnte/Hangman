package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func readLines(path string) ([]string, error) {
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

func HorrorCoding1() {
	f, err := readLines("gutenberg.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	ale := rand.Intn(len(f) - 1)
	fmt.Println(f[ale])
}
