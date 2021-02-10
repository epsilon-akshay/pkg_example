package main

import (
	"fmt"
	"github.com/epsilon-akshay/pkg_example/internal"
)
import "github.com/epsilon-akshay/pkg_example/pkg/string"
func main() {
	fmt.Print("hi there")
	fmt.Println(string.Hash())
	fmt.Println(internal.Random())

}