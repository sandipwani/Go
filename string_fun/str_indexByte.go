package main
import (
	"fmt"
	s"strings"
)

func main() {
	fmt.Println(s.IndexByte("Golang",'g'))
	fmt.Println(s.IndexByte("Gophers",'h'))
	fmt.Println(s.IndexByte("golang",'x'))
}
