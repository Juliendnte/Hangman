package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Graphisme(nb int) {
	if nb == 0 {
		nb = 1
	}
	file, err := os.Open("hangman(2).txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer file.Close()
	var lines []string
	for i := 0; i < 7; i++ {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}
	set := nb * 10
	for i := 10 * (nb - 1); i < set; i++ {
		fmt.Println(lines[i])
	}
}
