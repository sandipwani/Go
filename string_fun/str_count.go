package main
import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println(s.Count("Jinggga","g"))  // Count counts the number of non-overlapping instances of substr in s.
	fmt.Println(s.Count("five",""))      // If substr is an empty string. Count returns 1 + the number of Unicode code points in s.
}
