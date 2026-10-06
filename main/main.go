package main

import (
	"fmt"
	"myapp/export"
)

func main() {
	fmt.Println(export.Hello()) // ✅ works
	// fmt.Println(greeter.bye()) // ❌ error: unexported
}
