package hash_ransom

import (
	"strings"
)

func GetHash(word string) int {

	split_word := strings.Split(word, "")

	if len(word) < 3 {

		for i := 0; i < (3 - len(word)); i++ {

			split_word = append(split_word, "z")

		}

	}

	result := 0

	for _, letter := range split_word[len(split_word)-3:] {

		letter = strings.ToLower(letter)

		if letter <= "i" {

			result = result + 1

		} else if letter <= "q" {

			result = result + 2

		} else {

			result = result + 3
		}

	}

	return result

}


func CreateMap(text_list []string) map[int]map[string]int {

	text_map := make(map[int]map[string]int)

	for _, word := range text_list {

		hash := GetHash(word)

		if _, hash_ok := text_map[hash]; hash_ok {

			if _, word_ok := text_map[hash][word]; word_ok {

				text_map[hash][word]++

			} else {

				text_map[hash][word] = 1
			}

		} else {

			text_map[hash] = make(map[string]int)

			text_map[hash][word] = 1
		}

	}

	return text_map

}

func SubtractMap(magazine map[int]map[string]int, note_list []string) (result bool) {

	for _, word := range note_list {

		hash := GetHash(word)

		if _, hash_ok := magazine[hash]; hash_ok {

			if _, word_ok := magazine[hash][word]; word_ok {

				magazine[hash][word] -= 1

				if magazine[hash][word] < 0 {

					return false
				}

			} else {

				return false
			}

		} else {

			return false
		}
	}

	return true
}

func CompareMaps(magazine map[int]map[string]int, note map[int]map[string]int) (result bool){

	for hash, _ := range note {

		if map_keys, hash_ok := magazine[hash]; hash_ok {

			for key, _ := range note[hash] {

				if map_keys[key] < note[hash][key] {

					return false
				}

			}

		} else {

			return false
		}

	}

	return true

}