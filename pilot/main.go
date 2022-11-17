package main

import (
	"fmt"

	student ".."
)

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

func main() {
	var donnie student.Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = student.AIRCRAFT1
	fmt.Println(donnie)
}
