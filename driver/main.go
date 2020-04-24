package main

import (
	"fmt"

	"Oreilly_Go/Chapter_4/Creating_packages/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("Como estas", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("Como estas", 15))

}
