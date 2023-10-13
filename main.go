package main

import h "Hangman/algo"

func main() {
	gs := h.Count("chenille")
	h.AfficherLettre("chenille", "h", gs)
}
