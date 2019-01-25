package main

import (
	"flag"
	"fmt"
)

/*
命令行标记参数练习
*/
func main() {
	// Basic flag declarations are available for string,
	// integer, and boolean options. Here we declare a
	// string flag `word` with a default value `"foo"`
	// and a short description. This `flag.String` function
	// returns a string pointer (not a string value);
	// we'll see how to use this pointer below.
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// It's also possible to declare an option that uses an
	// existing var declared elsewhere in the program.
	// Note that we need to pass in a pointer to the flag
	// declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()

	// Here we'll just dump out the parsed options and
	// any trailing positional arguments. Note that we
	// need to dereference the pointers with e.g. `*wordPtr`
	// to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// go run src\exercise01\MainArgFlag.go -word=opt -numb=7 -fork -svar=flag
// go run src\exercise01\MainArgFlag.go -word=opt
// go run src\exercise01\MainArgFlag.go -word=opt a1 a2 a3
// go run src\exercise01\MainArgFlag.go -word=opt a1 a2 a3 -numb=7
// go run src\exercise01\MainArgFlag.go -h
