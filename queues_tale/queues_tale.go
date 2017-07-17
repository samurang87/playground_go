package main

import (
	"fmt"
	"log"
)

func Read() ([]string, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	list := make([]string, length)

	for i := 0; i < len(list); i++ {
		_, err := fmt.Scanf("%v", &list[i])
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func main() {

	list, err := Read()

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range list {


		fmt.Println(item)

		}

	}
