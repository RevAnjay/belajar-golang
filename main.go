package main

import (
	"fmt"
)

func TampilkanKata(kata string, isLooping bool) {
	if !isLooping {
		fmt.Println(kata)
		return
	}

	for i := 0; i < 20; i++ {
		fmt.Println(kata, i)
	}
}

func main() {
	TampilkanKata("IamBestCode", true)
}