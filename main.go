package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

// todo list menggunakan database mysql

// structure model databasenya
type Todo struct {
	gorm.Model
	Title string
}

var DB *gorm.DB

// koneksi ke database mysql
func connDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_golang2?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gagal terhubung ke database")
	}

	database.AutoMigrate(&Todo{})

	DB = database

	fmt.Println("berhasil terhubung ke database")
}

func showTodoMysql() {
	// query ke database buat ambil semua todo yang ada
	var allTodos []Todo
	DB.Find(&allTodos)

	// validasi kalo gada todonya maka kembalikan
	if len(allTodos) == 0 {
		fmt.Println("belum ada todo")
		return
	}

	// tampilin semua todonya
	fmt.Println("daftar todo:")
	for _, todo := range allTodos {
		fmt.Printf("%d. %s\n", todo.ID, todo.Title)
	}
}

func addTodoMysql(scanner *bufio.Scanner) {
	// input nama todo baru
	fmt.Print("masukan todo baru: ")
	scanner.Scan()

	todo := Todo{Title: scanner.Text()}
	DB.Create(&todo)
	fmt.Println("berhasil menambahkan todo: ", todo.Title)
}

func removeTodoMysql(scanner *bufio.Scanner) {
	// tampilin smua todo yang ada
	showTodoMysql()

	// input id nya
	fmt.Print("masukan nomor todo yang mau dihapus: ")
	scanner.Scan()

	// var id uint
	// fmt.Sscan(scanner.Text(), "%d", &id)
	// ubah string ke integer
	id, err := strconv.Atoi(scanner.Text())
	if err != nil || id <= 0 {
		fmt.Println("id todo tidak valid")
		return
	}

	// query ke database
	var todo Todo
	result := DB.First(&todo, id)
	if result.Error != nil {
		fmt.Println("todo tidak ditemukan")
		return
	}

	// delete todo nya
	DB.Delete(&todo)
	fmt.Println("berhasil menghapus todo: ", todo.Title)
}

func todoListMysql() {
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
			showTodoMysql()
		case "2":
			addTodoMysql(scanner)
		case "3":
			removeTodoMysql(scanner)
		case "4":
			fmt.Println("keluar...")
			return
		default:
			fmt.Println("pilihan tidak valid")
		}
	}
}

func main() {
	connDatabase()

	todoListMysql()
	// TampilkanKata("Looping", true, 10)
	// fmt.Println(BagiAngka(1,0))
}