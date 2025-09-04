package main

import (
	"fmt"
)

func TampilkanKata(kata string, isLooping bool, loopingValue ...int) {
	if loopingValue[0] == 0 {
		loopingValue[0] = 10
	}

	if !isLooping {
		fmt.Println(kata)
		return
	}

	for i := 0; i < loopingValue[0]; i++ {
		fmt.Println(kata, i)
	}
}

func TambahAngka(a int, b int) int {
	return a + b
}

func KurangAngka(a int, b int) int {
	return a - b
}

func KaliAngka(a int, b int) int {
	return a * b
}

func BagiAngka(a int, b int) any {
	if b < 1 {
		return "maaf tidak dapat dibagi 0"
	}

	return a / b
}

func main() {
	TampilkanKata("Looping", true, 10)
	// fmt.Println(BagiAngka(1,0))
}