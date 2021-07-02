package main
import "fmt"

type contact struct {
	email string
	phone int
}

type student struct {
	firstname string
	lastname string
	contacts contact
}

func main() {
	peter := student{firstname:"Sandip",lastname:"wani",
	contacts: contact{email:"sandipwani1998@mymail.com",phone:8983495962},
	}
	peter.print()
}

func (s student) print() {
	fmt.Println("Firstname is :",s.firstname, "Lastname is :",s.lastname)

	fmt.Printf("%+v\n",s)
}
