package utils

import "testing"

func TestParseUniverseFromString(t *testing.T) {
	t.Run("world with bad params", func(t *testing.T) {
		input := "hello"
		_, err := ParseUniverseFromString(input)
		if err == nil {
			t.Errorf("ParseUniverseFromString(%s); expected err; got nil", input)
		}
	})

	t.Run("world with no robots", func(t *testing.T) {
		input := "5 3"
		_, err := ParseUniverseFromString(input)
		if err == nil {
			t.Errorf("ParseUniverseFromString(%s); expected err; got nil", input)
		}
	})

	t.Run("world with bad robot params", func(t *testing.T) {
		input := `4 10
hello`
		_, err := ParseUniverseFromString(input)
		if err == nil {
			t.Errorf("ParseUniverseFromString(%s); expected err; got nil", input)
		}
	})
}
