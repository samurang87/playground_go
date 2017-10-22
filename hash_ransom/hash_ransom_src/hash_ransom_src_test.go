package hash_ransom

import (
	"reflect"
	"testing"
)

func TestGetHash(t *testing.T) {

	result := GetHash("aaa")

	expected_result := 3

	if result != expected_result {

		t.Errorf("%v is not %v", result, expected_result)
	}
}

func TestCreateMaps(t *testing.T) {

	some_text := []string{"aaa", "aaa", "aaa", "bb", "bb", "c"}

	text_map := CreateMap(some_text)

	expected_result := make(map[int]map[string]int)


	expected_result[3] = make(map[string]int)
	expected_result[3]["aaa"] = 3

	expected_result[5] = make(map[string]int)
	expected_result[5]["bb"] = 2

	expected_result[7] = make(map[string]int)
	expected_result[7]["c"] = 1

	eq := reflect.DeepEqual(expected_result, text_map)

	if eq == false {

		t.Errorf("Not equal! Expected %v, got %v", expected_result, text_map)
	}

}

func TestCompareMapsHappy(t *testing.T) {

	var magazine = map[string]int{

		"a": 3,
		"b": 2,
		"c": 1,
	}

	var note = map[string]int{

		"a": 3,
		"b": 2,
	}

	result := CompareMaps(magazine, note)

	if result == false {

		t.Errorf("Not equal! Expected %v, got %v", true, result)
	}

}

func TestCompareMapsKeyMissing(t *testing.T) {

	var magazine = map[string]int{

		"a": 3,
		"b": 2,
		"c": 1,
	}

	var note = map[string]int{

		"a": 3,
		"b": 2,
		"d": 3,
	}

	result := CompareMaps(magazine, note)

	if result == true {

		t.Errorf("Not equal! Expected %v, got %v", false, result)
	}

}

func TestCompareMapsInsufficientNumberOfWords(t *testing.T) {

	var magazine = map[string]int{

		"a": 3,
		"b": 2,
		"c": 1,
	}

	var note = map[string]int{

		"a": 3,
		"b": 2,
		"c": 3,
	}

	result := CompareMaps(magazine, note)

	if result == true {

		t.Errorf("Not equal! Expected %v, got %v", false, result)
	}

}
