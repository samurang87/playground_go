package main

import (
	"fmt"
	"github.com/samurang87/playground_go/brackets/balanced_brackets"
)

func main() {

	var a int
	var b string

	fmt.Scanf("%d\n", &a)

	for i := 0; i < a; i++ {

		fmt.Scanf("%s\n", &b)

		response := balanced_brackets.Balance(b)

		if response == true {

			fmt.Println("YES")

		} else {

			fmt.Println("NO")
		}

	}

}
