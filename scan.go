package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Println("Digite o primeiro número: ")
	_, erro1 := fmt.Scan(&n1)
	fmt.Println("Digite o segundo número: ")
	_, erro2 := fmt.Scan(&n2)

	if erro1 != nil || erro2 != nil {
		fmt.Println("ERRO")
		return
	}

	soma := n1 + n2

	fmt.Println( n1, "+ ", n2, "=", soma)

}
