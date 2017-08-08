package main

import "fmt"

func MakeAwesome(name string) string {
	if name == "Nirag" {
		return name + ", eh"
	}
	return name + " is awesome!"
}

func main() {
	fmt.Println(MakeAwesome("Sam"))
	fmt.Println(MakeAwesome("Matt"))
	fmt.Println(MakeAwesome("Darrell"))
	fmt.Println(MakeAwesome("Conlin"))
	fmt.Println(MakeAwesome("Nirag"))
}
