package main
import (
	"fmt"
	s "strings"
)

var p = fmt.Println

//ContainsRune takes two argument and second rune argument checked with first string if it present in string then it returns true otherwise false.

func main() {
	p(s.ContainsRune("aardvark",97))
	p(s.ContainsRune("Impossible",74))
}
