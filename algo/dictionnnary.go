package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var nb int = 1
var win bool

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

func WriteWord() string {
	f, err := ReadLines("gutenberg.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	ale := rand.Intn(len(f) - 1)
	return f[ale]
}

func Menu() {
	fmt.Println("Bienvenue au jeu du pendue")
	var answer int
	gs := WriteWord()
	guess = Count(gs)
	count := 1
	fmt.Println("Voici le mot que tu dois devinez", guess)
	for !win {
		fmt.Println("1-Voulez-vous devinez le mot")
		fmt.Println("2-Voulez-vous devinez une lettre")
		fmt.Scanln(&answer)
		switch answer {
		case 1:
			fmt.Println("Rentrez le mot que vous voulez testez")
			var test string
			fmt.Scanln(&test)
			TestWord(gs, test)
		case 2:
			var test string
			fmt.Println("Rentrez la lettre que vous voulez testez")
			fmt.Scanln(&test)
			guess, count = AfficherLettre(gs, test, guess, count)
		}
	}
}
