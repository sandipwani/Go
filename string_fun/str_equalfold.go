package main
import "fmt"
import s "strings"

func main() {
	fmt.Println(s.EqualFold("Go","go")) // the both parameters are equal under unicode case-folding, which is more general form of case-insensitivity.
	fmt.Println(s.EqualFold("Krishnan","Krishna"))
}
