package hackerrank_utils

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"log"
)

func Read() (input []string){

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		input = append(input, scanner.Text())
	}

	return input
}

func ReadFromStdinLikeFile(path string) (input []string){

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


func TestEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		fmt.Println("One is nil and one is not!")
		return false
	}

	if len(a) != len(b) {
		fmt.Println("Oh noes, lengths differ!")
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			fmt.Println("Oh noes, %v and %v have different elements!", a, b)
			return false
		}
	}

	return true
}

func ListIntToStr(n []int) (s string){

	var ids []string

	for _, i := range n {

		ids = append(ids, strconv.Itoa(i))
	}

	s = strings.Join(ids, " ")

	return
}
