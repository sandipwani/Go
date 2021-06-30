package main
import (
	"fmt"
	st"strings"
)

func main() {
	rot13 := func(r rune) rune {
		switch{
		
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(st.Map(rot13,"ATwas brillig and the slithy gopher..."))
}
