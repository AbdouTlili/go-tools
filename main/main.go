package main

import (
	"fmt"

	files "github.com/AbdouTlili/go-tools/files"
)

func main() {
	tables, _ := files.GetFloat("/home/spoutnik53/tests/go/main/data.txt")
	fmt.Println(tables)
}
