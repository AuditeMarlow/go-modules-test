package main

import "fmt"

var version = "1.2.0" // x-release-please-version

func main() {
	fmt.Printf("go-modules-test v%s\n", version)
}
