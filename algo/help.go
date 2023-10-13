package hangman

import (
	"fmt"
)

func Count(m string) string {
	var guess string
	for n := 0; n < len(m); n++ {
		guess = "_ "
		fmt.Printf("_ ")
	}
	return guess
}

func IsInWord(word, s string) bool { // on regarde si c'est dans le mot ou pas
	for l := range word {
		if string(l) == s {
			return true // si ça y est tu peux te le mettre dans le trou
		}
	}
	return false
}

func AfficherLettre(mot, s, guess string) {
	if IsInWord(mot, s) {
		if IsInWord(guess, s) {
			fmt.Println("Vous avez déjà essayez cette lettre")
		}
	}
}
