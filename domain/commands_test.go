package domain

import "testing"

func TestForwardCommand_Run(t *testing.T) {
	c := ForwardCommand{}
	w := &world{
		width:  5,
		height: 3,
		scents: map[coords]bool{},
	}

	t.Run("x: 1, y: 1, o: 0", func(t *testing.T) {
		r := &robot{
			coords:      &coords{x: 1, y: 1},
			orientation: orientation(0),
		}

		c.Run(r, w)

		if r.coords.x != 2 {
			t.Errorf("ForwardCommand.Run(%s); expected 2; got %d", r.coords, r.coords.x)
		}
	})

	t.Run("x: 1, y: 2, o: 90", func(t *testing.T) {
		r := &robot{
			coords:      &coords{x: 1, y: 1},
			orientation: orientation(90),
		}

		c.Run(r, w)

		if r.coords.y != 2 {
			t.Errorf("ForwardCommand.Run(%s); expected 2; got %d", r.coords, r.coords.y)
		}
	})
}
