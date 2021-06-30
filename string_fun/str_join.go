package main
import (
	"fmt"
	st "strings"
)

func main() {
	s := []string {"foo","bar","baz"}
	fmt.Println(st.Join(s,", "))
}
