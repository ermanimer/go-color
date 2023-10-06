package color

import (
	"fmt"
	"io"
)

type colorCode int

const (
	Black colorCode = 0 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

func New(code byte) colorCode {
	return colorCode(code)
}

func (c colorCode) Fprint(w io.Writer, a ...any) {
	makeEscCode(fg, c).Fprint(w, a...)
}

func (c colorCode) Fprintf(w io.Writer, format string, a ...any) {
	makeEscCode(fg, c).Fprintf(w, format, a...)
}

func (c colorCode) Fprintln(w io.Writer, a ...any) {
	makeEscCode(fg, c).Fprintln(w, a...)
}

func (c colorCode) Print(a ...any) {
	makeEscCode(fg, c).Print(a...)
}

func (c colorCode) Printf(format string, a ...any) {
	makeEscCode(fg, c).Printf(format, a...)
}

func (c colorCode) Println(a ...any) {
	makeEscCode(fg, c).Println(a...)
}

func (c colorCode) Paint(a any) string {
	return makeEscCode(fg, c).Paint(a)
}

func (foreground colorCode) With(background colorCode) escCode {
	return escCode(fmt.Sprintf("%s%s", makeEscCode(fg, foreground), makeEscCode(bg, background)))
}
