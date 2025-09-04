package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// looping kata
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

// kalkulator sederhana
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

// todo list sederhana
var todos []string

func showTodos() {
	if len(todos) == 0 {
		fmt.Println("belum ada todo")
		return
	}

	fmt.Println("daftar todo:")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
	}
}

func addTodo(scanner *bufio.Scanner) {
	fmt.Print("masukan todo baru: ")
	scanner.Scan()
	todo := strings.TrimSpace(scanner.Text())
	if todo != "" {
		todos = append(todos, todo)
		fmt.Println("todo berhasil ditambahkan")
	} else {
		fmt.Println("todo tidak boleh kosong")
	}
}

func removeTodo(scanner *bufio.Scanner) {
	showTodos()
	if len(todos) == 0 {
		return
	}

	fmt.Print("masukan nomor todo yang ingin dihapus: ")
	scanner.Scan()
	var index int
	fmt.Sscanf(scanner.Text(), "%d", &index)

	if index > 0 && index <= len(todos) {
		todos = append(todos[:index-1], todos[index:]...)
		fmt.Println("todo berhasil dihapus")
	} else {
		fmt.Println("nomor tidak valid")
	}
}

func todoList() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Todo List ===")
		fmt.Println("1. lihat todo")
		fmt.Println("2. tambah todo")
		fmt.Println("3. hapus todo")
		fmt.Println("4. keluar")
		fmt.Print("pilih menu: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			showTodos()
		case "2":
			addTodo(scanner)
		case "3":
			removeTodo(scanner)
		case "4":
			fmt.Println("keluar...")
			return
		default:
			fmt.Println("pilihan tidak valid")
		}
	}
}

func main() {
	todoList()
	// TampilkanKata("Looping", true, 10)
	// fmt.Println(BagiAngka(1,0))
}