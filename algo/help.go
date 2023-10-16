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

func IsInList(s string, lst []string) bool {
	for _, c := range lst {
		if string(c) == s {
			return true
		}
	}
	return false
}

func AfficherLettre(mot, s, guess string, nb int, lst []string) (string, int, []string) {
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
		if IsInList(s, lst) {
			fmt.Println("Vous avez déjà essayez cette lettre")
			return guess, nb, lst
		}
		lst = append(lst, s)
		if Graphisme(nb) {
			fmt.Println("Le mot était ", mot)
		}
		nb++
	}
	return guess, nb, lst
}

func TestWord(guess, mot string, nb int) (bool, int) {
	if mot == guess {
		fmt.Println("Vous avez trouvé le mot")
		win = true
	} else {
		nb += 2
		fmt.Println("Ce n'est pas le bon mot")
	}
	return win, nb
}
