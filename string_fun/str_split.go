//Split, SplitAfter, SplitAfterN, SplitN

package main
import (
	"fmt"
	s"strings"
)

func main() {
	fmt.Printf("%q\n", s.Split("a,b,c",","))
	fmt.Printf("%q\n", s.Split(" xyz ",""))
	fmt.Printf("%q\n", s.Split("","Bernardo O'Higgins"))

	fmt.Printf("%q\n", s.SplitAfter("a,b,c",","))   

	fmt.Printf("%q\n", s.SplitAfterN("a,b,c,d",",",3))

	fmt.Printf("%q\n", s.SplitN("a,b,c", ",", 2))
	z := s.SplitN("a,b,c",",",0)
	fmt.Printf("%q (nil = %v)\n", z,z == nil)
}
