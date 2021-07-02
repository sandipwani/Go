package main
import (
	"fmt"
	"os"
)

func main() {
	panic("a program")

	file,err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("File created succesfully")
}
