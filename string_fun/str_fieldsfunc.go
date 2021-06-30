//FieldsFunc splits the string s at each run of Unicode code points s satisfying f(c) and returns an array of slices of s. If all code points in s satisfy f(c) or the string is empty, an empty slice is returned.
//FieldsFunc makes no furantees about the order in which it calls f(c) and assumes that f always returns the same value for a given c.

package main
import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q \n",strings.FieldsFunc("     Foo1;bar2,baz3...",f))
}
