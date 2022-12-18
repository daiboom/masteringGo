package main

import "fmt"

func main() {
	fmt.Printf("please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("your name is", name)
}
