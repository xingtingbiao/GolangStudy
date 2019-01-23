package main

import (
	"fmt"
	"os"
)
import s "strings"

/**
strings的常用api的使用
*/
func main() {
	//testStringsApis()
	testStringFormat()
}

var p = fmt.Println

func testStringsApis() {
	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.
	_, _ = p("Contains:  ", s.Contains("test", "es"))
	_, _ = p("Count:     ", s.Count("test", "t"))
	_, _ = p("HasPrefix: ", s.HasPrefix("test", "te"))
	_, _ = p("HasSuffix: ", s.HasSuffix("test", "st"))
	_, _ = p("Index:     ", s.Index("test", "e"))
	_, _ = p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	_, _ = p("Repeat:    ", s.Repeat("a", 5))
	_, _ = p("Replace:   ", s.Replace("foo", "o", "0", -1))
	_, _ = p("Replace:   ", s.Replace("foo", "o", "0", 1))
	_, _ = p("Split:     ", s.Split("a-b-c-d-e", "-"))
	_, _ = p("ToLower:   ", s.ToLower("TEST"))
	_, _ = p("ToUpper:   ", s.ToUpper("test"))
	_, _ = p()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.
	_, _ = p("Len: ", len("hello"))
	_, _ = p("Char:", "hello"[1])
}

type point struct {
	x, y int
}

func testStringFormat() {
	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `point` struct.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.
	fmt.Printf("%+v\n", p)

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.
	fmt.Printf("%#v\n", p)

	// To print the type of a value, use `%T`.
	fmt.Printf("%T\n", p)

	// Formatting booleans is straight-forward.
	fmt.Printf("%t\n", true)

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.
	fmt.Printf("%d\n", 123)

	// This prints a binary representation.
	fmt.Printf("%b\n", 14)

	// This prints the character corresponding to the
	// given integer.
	fmt.Printf("%c\n", 33)

	// `%x` provides hex encoding.
	fmt.Printf("%x\n", 456)

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.
	fmt.Printf("%f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// For basic string printing use `%s`.
	fmt.Printf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use `%q`.
	fmt.Printf("%q\n", "\"string\"")

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.
	fmt.Printf("%x\n", "hex this")

	// To print a representation of a pointer, use `%p`.
	fmt.Printf("%p\n", &p)

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the `-` flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// To left-justify use the `-` flag as with numbers.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.
	fmt.Println(fmt.Sprintf("a %s", "string"))

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.
	_, _ = fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
