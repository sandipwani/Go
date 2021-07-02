package main
import "fmt"

type marks interface{
	printMarks()
}

type grade11Marks struct {
	math int
	physics int
}

type grade12Marks struct {
	math int
	computer int
}

func main() {
	ram := grade11Marks{50,80}
	shyam := grade12Marks{60,70}

	print(ram)
	print(shyam)
	//ram.printMarks()
	//shyam.printMarks()
}

func (m grade11Marks) printMarks() {
	fmt.Println("math-:--",m.math,"  physics-:--",m.physics)
}

func (g grade12Marks) printMarks() {
	fmt.Println("math-> ",g.math,"  CS->",g.computer)
}

func print(b marks) {
	b.printMarks()
}
