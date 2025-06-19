package puppy

import (
	"fmt"

	"github.com/Rahulkprajapati/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Wof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUps(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUps(Barks())
}

func Form1() {
	fmt.Println("This is a form from version v1.1.0")
}
