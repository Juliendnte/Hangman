package hangman

import (
	"fmt"
	"math/rand"
)

var guess string

func Count(m string) string {
	for n := 0; n < len(m); n++ {
		guess += "_ "
	}
	return guess
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

func indice(mot, s, guess string, nb int, lst []string) (string, int, []string, bool) {
	lst = append(lst, s)
	for i, t := range mot {
		if string(t) == s {
			slc := TransformString(guess)
			slc[i*2] = s
			guess = TransformSlice(slc)
		}
	}
	if IsUnderscore(guess) {
		return guess, nb, lst, true
	}
	return guess, nb, lst, false
}

func letterAleatory(guess, mot string, lst []string) string {
	var w string
	var ale int
	ale = rand.Intn(len(mot))
	w = string(mot[ale])
	if IsInList(w, lst) {
		w = letterAleatory(guess, mot, lst)
	}
	return w
}

func ToLower(s string) string {
	var listf string
	for _, c := range s {
		if c > 64 && c < 91 {
			listf = listf + string(c+32)
		} else {
			listf = listf + string(c)
		}
	}
	return listf
}

func printlist(s []string) {
	fmt.Println("Vous avez testé ")
	for _, c := range s {
		fmt.Print(string(c), " ")
	}
	fmt.Print("\n\n\n\n")
}

func AfficherLettre(mot, s, guess string, nb int, lst []string, prt string) (string, int, []string, bool) {
	if IsInWord(mot, s) {
		fmt.Print("\033[H\033[2J")
		Equalizeprint(prt)
		Graphisme(nb)
		if IsInWord(guess, s) {
			fmt.Println("Vous avez déjà essayez cette lettre")
		} else {
			lst = append(lst, s)
			for i, t := range mot {
				if string(t) == s {
					slc := TransformString(guess)
					slc[i*2] = s
					guess = TransformSlice(slc)
				}
			}
			if IsUnderscore(guess) {
				return guess, nb, lst, true
			}
		}
	} else {
		fmt.Print("\033[H\033[2J")
		Equalizeprint(prt)
		if IsInList(s, lst) {
			fmt.Println("Vous avez déjà essayez cette lettre")
			return guess, nb, lst, false
		}
		lst = append(lst, s)
		nb++
		fmt.Println("ҳ̸  Mauvaise lettre")
		printlist(lst)
		Graphisme(nb)
	}
	return guess, nb, lst, false
}

func TestWord(guess, mot string, nb int) (bool, int) {
	if mot == guess {
		fmt.Println("Vous avez trouvé le mot")
		win = true
	} else {
		nb += 2
		fmt.Println("Ce n'est pas le bon mot")
		Graphisme(nb)
	}
	return win, nb
}

func Equalizeprint(s string) {
	spaced := ""
	for len(s)/2+len(spaced) < 95 {
		spaced += " "
	}
	spaced += s
	fmt.Println(spaced)
}

func IsInList(s string, lst []string) bool {
	for _, c := range lst {
		if string(c) == s {
			return true
		}
	}
	return false
}

func IsUnderscore(s string) bool {
	for _, c := range s {
		if string(c) == "_" {
			return false
		}
	}
	return true
}
func IsInWord(word, s string) bool { // on regarde si c'est dans le mot ou pas
	for _, l := range word {
		if string(l) == s {
			return true // si ça y est tu peux te le mettre dans le trou
		}
	}
	return false
}
