package main
import(
	"fmt"
	s "strings"
)

var p = fmt.Println

//ContainsAny reports whether any Unicode code points in chats are within s.


func main() {
	p(s.ContainsAny("team","i"))
	p(s.ContainsAny("Sandip","Sandy"))
	p(s.ContainsAny("Everything","Nothing"))
}
