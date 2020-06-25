package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Println("Hello, Homebrew!")
	fmt.Printf("version %s, commit %s, built at %s\n", version, commit, date)
}
