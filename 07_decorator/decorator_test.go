package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)

	csq2 := NewSizeSquare(sq, "2")
	got = csq2.Draw()
	assert.Equal(t, "this is a square, size is 2", got)
}
