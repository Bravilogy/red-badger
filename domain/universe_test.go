package domain

import "testing"

func TestNewUniverseFromString(t *testing.T) {
	t.Run("world with bad params", func(t *testing.T) {
		input := "hello"
		_, err := NewUniverseFromString(input)
		if err == nil {
			t.Errorf("ParseUniverseFromString(%s); expected err; got nil", input)
		}
	})

	t.Run("world with no robots", func(t *testing.T) {
		input := "5 3"
		_, err := NewUniverseFromString(input)
		if err == nil {
			t.Errorf("ParseUniverseFromString(%s); expected err; got nil", input)
		}
	})

	t.Run("world with bad robot params", func(t *testing.T) {
		input := `4 10
hello`
		_, err := NewUniverseFromString(input)
		if err == nil {
			t.Errorf("ParseUniverseFromString(%s); expected err; got nil", input)
		}
	})

	t.Run("world with bad robot commands", func(t *testing.T) {
		input := `4 10
1 1 W
HELLO`
		_, err := NewUniverseFromString(input)
		if err.Error() != "invalid command specified - H" {
			t.Errorf("ParseUniverseFromString(%s); expected 'invalid command' error; got %s", input, err)
		}
	})
}
