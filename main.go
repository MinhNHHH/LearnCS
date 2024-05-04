package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	var (
		flagString  = flag.String("string", "", "A string flag")
		flagInt     = flag.Int("int", 0, "An integer flag")
		flagBool    = flag.Bool("bool", false, "A boolean flag")
		flagStrings flagSlice
	)

	flag.Var(&flagStrings, "strings", "A list of strings")

	// Parse command-line flags
	flag.Parse()

	// Access flag values
	fmt.Println("String flag:", *flagString)
	fmt.Println("Integer flag:", *flagInt)
	fmt.Println("Boolean flag:", *flagBool)
	fmt.Println("String slice flag:", flagStrings)
}

// Define a custom flag type for string slices
type flagSlice []string

func (f *flagSlice) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *flagSlice) Set(value string) error {
	*f = append(*f, value)
	return nil
}
