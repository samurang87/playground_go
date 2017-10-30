package main

import (
	"fmt"
	"github.com/samurang87/playground_go/hash_ransom/hash_ransom_src"
	"bufio"
	"os"
	"io"
)

func main() {


	reader := bufio.NewReader(os.Stdin)

	lines := [][]string{}
	currentLine := []string{}
	currentWord := []byte{}

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				if len(currentWord) > 0 {
					currentLine = append(currentLine, string(currentWord))
				}
				if len(currentLine) > 0 {
					lines = append(lines, currentLine)
				}
				break
			}
			panic(err)
		}

		if b == ' ' || b == '\n' {
			currentLine = append(currentLine, string(currentWord))
			currentWord = []byte{}

			if b == '\n' {
				lines = append(lines, currentLine)
				currentLine = []string{}
			}
		} else {
			currentWord = append(currentWord, b)
		}
	}


	magazine := hash_ransom.CreateMap(lines[1])


	if hash_ransom.SubtractMap(magazine, lines[2]) {

		fmt.Println("Yes")

	} else {

		fmt.Println("No")
	}



	}



