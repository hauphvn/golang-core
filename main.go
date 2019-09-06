package main

import "fmt"

func meoKeu() (string, string) {
	return "meo meo", "gau gau"
}

func main() {
	a, b := meoKeu()
	fmt.Println(a, b)

}
