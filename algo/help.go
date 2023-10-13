package hangman

import (
	"fmt"
)

func Count(m string) string {
	var guess string
	for n := 0; n < len(m); n++ {
		guess += "_ "
		fmt.Printf("_ ")
	}
	fmt.Println("")
	return guess
}

func IsInWord(word, s string) bool { // on regarde si c'est dans le mot ou pas
	for _, l := range word {
		if string(l) == s {
			return true // si ça y est tu peux te le mettre dans le trou
		}
	}
	return false
}

func TransformString(s string) []string {
	slice := []string{}
	for _, c := range s {
		slice = append(slice, string(c))
	}
	return slice
}

func TransformSlice(s []string) string {
	var str string
	for _, c := range s {
		str += c
	}
	return str
}

func AfficherLettre(mot, s, guess string) {
	if IsInWord(mot, s) {
		if IsInWord(guess, s) {
			fmt.Println("Vous avez déjà essayez cette lettre")
		} else {
			for i, t := range mot {
				if string(t) == s {
					slc := TransformString(guess)
					slc[i*2] = s
					guess = TransformSlice(slc)
					fmt.Print(s, " ")
				} else {
					fmt.Print(string(guess[i*2]), " ")
				}
			}
		}
	} else {
		Graphisme(1)
	}
}
