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

func WriteWord(path string) string {
	f, err := ReadLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	ale := rand.Intn(len(f))
	return f[ale]
}

func Menu() {
	var val int
	var answer string
	var gs string
	var test string
	var lst []string
	var prt string
	Equalizeprint("Bienvenue au jeu du pendu")
	Equalizeprint("Sur quel mode veux-tu jouer au jeu")
	fmt.Print("\n\n\n\n")
	for i := 3; i < 11; i++ {
		fmt.Println("•/", i-2, "Mode ", i, " lettres")
	}
	fmt.Println("•/ 9 Mode 10 lettre ou plus")
	fmt.Println("•/ 10 Mode anglais")
	fmt.Println("•/ 11 Mode difficile")
	fmt.Println("•/ 12 Mode impossible")
	fmt.Scanln(&val)
	fmt.Print("\033[H\033[2J")
	switch val {
	case 1:
		gs = WriteWord("mot3lettres.txt")
		prt = "Je suis julien et je suis nul car je laisse mon pc déverrouillé"
	case 2:
		gs = WriteWord("mot4lettres.txt")
		prt = "Bienvenue dans le mode 4 lettres"
	case 3:
		gs = WriteWord("mot5lettres.txt")
		prt = "Bienvenue dans le mode 5 lettres"
	case 4:
		gs = WriteWord("mot6lettres.txt")
		prt = "Bienvenue dans le mode 6 lettres"
	case 5:
		gs = WriteWord("mot7lettres.txt")
		prt = "Bienvenue dans le mode 7 lettres"
	case 6:
		gs = WriteWord("mot8lettres.txt")
		prt = "Bienvenue dans le mode 8 lettres"
	case 7:
		gs = WriteWord("mot9lettres.txt")
		prt = "Bienvenue dans le mode 9 lettres"
	case 8:
		gs = WriteWord("mot10lettres.txt")
		prt = "Bienvenue dans le mode 10 lettres"
	case 9:
		gs = WriteWord("mot101112lettres.txt")
		prt = "Bienvenue dans le mode 10 lettres ou plus"
	case 10:
		gs = WriteWord("motpenduanglais.txt")
		prt = "Welcome to the english version"
	case 11:
		gs = WriteWord("multilettres.txt")
		prt = "Bienvenue dans le mode difficile"
	case 12:
		gs = WriteWord("gutenberg.txt")
		prt = "Bienvenue dans le mode impossible"
	}
	gs = ToLower(gs)
	guess := Count(gs)
	count := 0
	Equalizeprint(prt)
	for !win {
		fmt.Print("\n\n\n")
		fmt.Println("Le mot que tu dois deviner : ", guess)
		if count > 9 {
			fmt.Println("Le mot était ", gs)
			return
		}
		fmt.Println("1-Veux-tu deviner le mot")
		fmt.Println("2-Veux-tu deviner une lettre")
		fmt.Scan(&answer)
		switch answer {
		case "1":
			fmt.Print("\033[H\033[2J")
			Equalizeprint(prt)
			fmt.Println("\nRentrez le mot que vous voulez tester")
			fmt.Scan(&test)
			win, count = TestWord(gs, test, count)
		case "2":
			fmt.Print("\033[H\033[2J")
			Equalizeprint(prt)
			fmt.Println("\nRentrez la lettre que vous voulez tester")
			fmt.Scan(&test)
			guess, count, lst, win = AfficherLettre(gs, test, guess, count, lst, prt)
		default:
			fmt.Println("No Comprendo")
			continue
		}

	}
	fmt.Println("Le mot était ", gs)
	fmt.Println("Vous avez réussi")

}
