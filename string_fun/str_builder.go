//Builder is used to efficiently build a string write methods. It minimizes memory copying. The zero value is ready to use. Do not copy a non-zero builder.

package main
import (
	"fmt"
	s"strings"
)

func main() {
	var b s.Builder
	for i:=3;i>=1;i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}
