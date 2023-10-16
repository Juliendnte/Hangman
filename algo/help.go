package hangman

import (
	"fmt"
)

var guess string

func Count(m string) string {
	for n := 0; n < len(m); n++ {
		guess += "_ "
	}
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

func AfficherLettre(mot, s, guess string, nb int) (string, int) {
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
			fmt.Print("\n")
		}
	} else {
		Graphisme(nb)
		nb++
	}
	return guess, nb
}

func TestWord(guess, mot string) bool {
	if mot == guess {
		fmt.Println("Vous avez trouvé le mot")
		win = true
	} else {
		fmt.Println("Ce n'est pas le bon mot")
	}
	return win
}
