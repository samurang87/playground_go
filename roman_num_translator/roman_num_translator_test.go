package roman_num_translator

import "testing"

func TestRoman_num_translator(t *testing.T) {

	cases := []struct {
		in string
		want int
	}{
		{"I", 1},
		{"II", 2},

	}
	for _, c := range cases {
		got := Translator(c.in)
		if got != c.want {
			t.Error("Translator(%q) == %q, want %q", c.in, got, c.want)}
	}

}
