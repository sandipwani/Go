package main
import "fmt"

type student struct {
	firstname string
	lastname string
}

func main() {
	Peter := student{firstname:"sandip",lastname:"wani"}
	PeterPointer := &Peter

	PeterPointer.changeFirstName("Bhagyesh")
	Peter.print()
}

func (s student) print(){
	fmt.Println("FN-> ",s.firstname,"LN->",s.lastname)
}

func (sp *student) changeFirstName(f string) {
	(*sp).firstname = f
}

