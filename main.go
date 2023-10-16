package main

import h "Hangman/algo"

func main() {

	gs := h.Count("chenille")
	gs = h.AfficherLettre("chenille", "h", gs, 1)
	h.AfficherLettre("chenille", "e", gs, 1)
}
