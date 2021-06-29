package main
import (
	"fmt"
	s "strings"  //Here s is alice for strings package.
)

var p = fmt.Println        //Here p is alice for fmt.Println

func main() {
	p(s.Contains("Sandip","an"))                         //True
	p(s.Contains("Ramnarayanan","Raghupathi"))           //False
	p(s.Contains("Karthik",""))                          //True
	p(s.Contains("",""))                                 //True
	p(s.Contains(" "," "))                               //True
}
