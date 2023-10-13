package Hangman

func IsInWord(word, s string) bool { // on regarde si c'est dans l'inventaire ou pas
	for l := range word {
		if string(l) == s {
			return true // si ca y'est tu peux te le mettre dans le trou
		}
	}
	return false
}
