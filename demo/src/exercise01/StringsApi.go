package main

import "fmt"
import s "strings"

/**
strings的常用api的使用
*/
func main() {
	testStringsApis()
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
