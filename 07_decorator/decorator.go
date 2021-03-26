/*
another good example: https://tutorialedge.net/golang/go-decorator-function-pattern-tutorial/


*What is this pattern about?
The Decorator pattern is used to dynamically add a new feature to an
OBJECT without changing its implementation. It differs from
inheritance because the new feature is added only to that particular
object, not to the entire subclass.

	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
*/

package decorator

// IDraw IDraw
type IDraw interface {
	Draw() string
}

// Square 正方形
type Square struct{}

// Draw Draw
func (s Square) Draw() string {
	return "this is a square"
}

// ColorSquare 有颜色的正方形
type ColorSquare struct {
	square IDraw // add IDraw interface
	color  string
}

// NewColorSquare NewColorSquare
func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{color: color, square: square}
}

// Draw Draw
/*
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
*/
func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}

type SizeSquare struct {
	square IDraw // add IDraw interface
	size   string
}

// NewColorSquare NewColorSquare
func NewSizeSquare(square IDraw, size string) SizeSquare {
	return SizeSquare{size: size, square: square}
}

// Draw Draw
func (c SizeSquare) Draw() string {
	return c.square.Draw() + ", size is " + c.size
}
