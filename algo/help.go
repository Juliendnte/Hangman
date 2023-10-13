package hangman

import (
	"fmt"
)

func Count(m string) {
	for n := 0; n < len(m); n++ {
		fmt.Printf("_ ")
	}
}

func IsInWord(word, s string) bool { // on regarde si c'est dans le mot ou pas
	for l := range word {
		if string(l) == s {
			return true // si ça y est tu peux te le mettre dans le trou
		}
	}
	return false
}

func AfficherLettre(mot, s string) {

}
