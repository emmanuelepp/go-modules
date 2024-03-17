package main

import (
	"fmt"
	"toolset"
)

func main() {

	var tools toolset.Tools

	s := tools.RandomString(10)

	fmt.Println("Ramdon string", s)

}
