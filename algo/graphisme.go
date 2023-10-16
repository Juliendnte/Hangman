package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Graphisme(nb int) {
	if nb > 9 {
		return
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
	fmt.Println("ҳ̸  Mauvaise lettre")
	for i := 10 * (nb - 1); i < set; i++ {
		fmt.Println(lines[i])
	}
}
