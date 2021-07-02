package main
import "fmt"

type student struct {
	firstname string
	lastname string
}

func main() {
	Peter := student{firstname:"Peter",lastname:"parker"}

	Peter.changeFirstName("Peterson")
	//fmt.Println("I am IN MAIN")
	//Peter.print()
}

func (s student) print() {
	fmt.Println("Firstname->",s.firstname,"lastname->",s.lastname)
}

func (s student) changeFirstName(f string){
	//fmt.Println("I am IN")
	s.firstname = f
	s.print()
}
