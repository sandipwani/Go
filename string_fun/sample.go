package main
import "fmt"
import "os"

func main() {
	fmt.Fprintf(os.Stderr,"an %s\n","error")
}
