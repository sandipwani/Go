package main
import "fmt"
import s"strings"

//Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode. IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.

func main() {
	fmt.Println("Fields are: %q",s.Fields("   foo bar  baz    "))
}
