package color

import (
	"fmt"
	"io"
)

type escCode string

const reset escCode = "\x1b[0m"

func (c escCode) Fprint(w io.Writer, a ...any) {
	c.fset(w)
	fmt.Fprint(w, a...)
	reset.fset(w)
}

func (c escCode) Fprintf(w io.Writer, format string, a ...any) {
	c.fset(w)
	fmt.Fprintf(w, format, a...)
	reset.fset(w)
}

func (c escCode) Fprintln(w io.Writer, a ...any) {
	c.fset(w)
	fmt.Fprint(w, a...)
	reset.fset(w)
	fmt.Print("\n")
}

func (c escCode) Print(a ...any) {
	c.set()
	fmt.Print(a...)
	reset.set()
}

func (c escCode) Printf(format string, a ...any) {
	c.set()
	fmt.Printf(format, a...)
	reset.set()
}

func (c escCode) Println(a ...any) {
	c.set()
	fmt.Print(a...)
	reset.set()
	fmt.Print("\n")
}

func (c escCode) Paint(a any) string {
	return fmt.Sprintf("%s%v%s", c, a, reset)
}

func (c escCode) set() {
	fmt.Print(c)
}

func (c escCode) fset(w io.Writer) {
	fmt.Fprint(w, c)
}

func makeEscCode(typeCode typeCode, cocolorCode colorCode) escCode {
	return escCode(fmt.Sprintf("\x1b[%d;5;%d;m", typeCode, cocolorCode))
}
