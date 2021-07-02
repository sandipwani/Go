//Trim, TrimFunc, Trimleft, TrimLeftFunc, TrimPrefix, TrimRight, TrimRightFunc, TrimSpace, TrimSuffix

package main
import (
	"fmt"
	s"strings"
	"unicode"
)

var p = fmt.Println

func main() {
	p(s.Trim("iiiHello, Gophers!!!","!i"))

	p(s.TrimFunc("%%%Hello, Gophers!!!",func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	p(s.TrimLeft("$$$Hello, Gophers!!!","$!"))

	p(s.TrimLeftFunc("$$$Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	var str = "$$$Hello, Gophers!!!"
	str = s.TrimPrefix(str,"$$$Hello, ")
	str = s.TrimPrefix(str,"***Howdy, ")
	str = s.TrimPrefix(str,"Go")
	p(str)

	p(s.TrimRight("$$$Hello, Gophers!!!","$!"))

	p(s.TrimRightFunc("$$$Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	p(s.TrimSpace("\t\n Hello, Gophers \n\t\r\n"))  //It remove all leading and trailing white spaces


	var str1 = "@@@Hello, Gophers!!!"
	str1 = s.TrimSuffix(str1,", Gophers!!!")
	str1 = s.TrimSuffix(str1,", Mashroom!!!")
	p(str1)
}
