package main

import (
	"fmt"
	"main/database/src"
	"main/database/src/utils"
)

func main() {
	fmt.Println("Hello, World!")
	data := fmt.Sprintf("Hello World %d", utils.RandomInt())
	src.SaveData2("./data.tmp", []byte(data))
}
