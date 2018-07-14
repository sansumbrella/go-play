// Echo 2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for key, val := range os.Args[1:] {
		fmt.Println(strconv.Itoa(key) + ": " + val)
	}
}
