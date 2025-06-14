package puppy

import (
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
