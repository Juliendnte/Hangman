package hangman

import "fmt"

func Graphisme(nb int) {
	switch nb {
	case 1:
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("=========")
	case 2:

		fmt.Println("		|  ")
		fmt.Println("		|  ")
		fmt.Println("		|  ")
		fmt.Println("		|  ")
		fmt.Println("		|  ")
		fmt.Println("  =========")
	case 3:
		fmt.Println("		+---+ ")
		fmt.Println("		|  ")
		fmt.Println("		|  ")
		fmt.Println("		|  ")
		fmt.Println("		|  ")
		fmt.Println("		|  ")
		fmt.Println("  =========")
	case 4:
		fmt.Println("		+---+ ")
		fmt.Println("		|   |  ")
		fmt.Println("			|  ")
		fmt.Println("			|  ")
		fmt.Println("			|  ")
		fmt.Println("			|  ")
		fmt.Println("	  =========")
	case 5:
		fmt.Println("		+---+  ")
		fmt.Println("		|   |  ")
		fmt.Println("		O   |  ")
		fmt.Println("			|  ")
		fmt.Println("			|  ")
		fmt.Println("			|  ")
		fmt.Println("	  =========")
	case 6:
		fmt.Println("		+---+  ")
		fmt.Println("		|   |  ")
		fmt.Println("		O   |  ")
		fmt.Println("		|   |  ")
		fmt.Println("			|  ")
		fmt.Println("			|  ")
		fmt.Println("	  =========")
	case 7:
		fmt.Println("		+---+  ")
		fmt.Println("		|   |  ")
		fmt.Println("		O   |  ")
		fmt.Println("	   /|   |  ")
		fmt.Println("			|  ")
		fmt.Println("			|  ")
		fmt.Println("	  =========")
	case 8:
		fmt.Println("		+---+  ")
		fmt.Println("		|   |  ")
		fmt.Println("		O   |  ")
		fmt.Println("	   /|\\  |  ")
		fmt.Println("			|  ")
		fmt.Println("			|  ")
		fmt.Println("	  =========")
	case 9:
		fmt.Println("		+---+  ")
		fmt.Println("		|   |  ")
		fmt.Println("		O   |  ")
		fmt.Println("	   /|\\  |  ")
		fmt.Println("	   /    |  ")
		fmt.Println("			|  ")
		fmt.Println("	  =========")
	case 10:
		fmt.Println("		+---+  ")
		fmt.Println("		|   |  ")
		fmt.Println("		O   |  ")
		fmt.Println("	   /|\\  |  ")
		fmt.Println("	   / \\  |  ")
		fmt.Println("			|  ")
		fmt.Println("	  =========")
	}
}
