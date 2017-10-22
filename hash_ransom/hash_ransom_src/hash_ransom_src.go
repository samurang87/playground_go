package hash_ransom

func CreateMap(text_list []string) map[string]int {

	text_map := make(map[string]int)

	for _, word := range text_list {

		if _, ok := text_map[word]; ok {

			text_map[word]++

		} else {

			text_map[word] = 1

		}
	}

	return text_map
}

func CompareMaps(magazine map[string]int, note map[string]int) (result bool) {

	for key, _ := range note {

		if val, ok := magazine[key]; ok {

			if val < note[key] {

				return false
			}

		} else {

			return false
		}
	}

	return true

}
