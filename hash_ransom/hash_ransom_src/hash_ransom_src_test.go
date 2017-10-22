package hash_ransom

import (
	"reflect"
	"testing"
)

func TestCreateMaps(t *testing.T) {

	some_text := []string{"a", "a", "a", "b", "b", "c"}

	text_map := CreateMap(some_text)

	expected_result := make(map[string]int)

	expected_result["a"] = 3

	expected_result["b"] = 2

	expected_result["c"] = 1

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
