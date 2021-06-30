//Here i mentioned string functionalities like Index, LastIndex, LastIndexAny, LastIndexByte, LastIndexFunc

package main
import (
	"fmt"
	s"strings"
	"unicode"
)

func main() {
	fmt.Println(s.Index("go gopher","go"))     //0
	fmt.Println(s.LastIndex("go gopher","go"))      //3
	fmt.Println(s.LastIndex("go gopher","rodent"))   //-1

	fmt.Println(s.LastIndexAny("go gopher","go"))    //4
	fmt.Println(s.LastIndexAny("go gopher","rodent"))  //8
	fmt.Println(s.LastIndexAny("Go gopher","fail"))    //-1

	fmt.Println(s.LastIndexByte("Hello, World",'l'))  //10 
	fmt.Println(s.LastIndexByte("Hello, World",'o'))  //8
	fmt.Println(s.LastIndexByte("Hello, world",'x'))  //-1

	fmt.Println(s.LastIndexFunc("go 123",unicode.IsNumber))
	fmt.Println(s.LastIndexFunc("123 go",unicode.IsNumber))
	fmt.Println(s.LastIndexFunc("go",unicode.IsNumber))
}
