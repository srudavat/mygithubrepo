package dice

import (
	"math/rand"
)

func Rolldice() int {
	var i int = rand.Intn(6)
	if i == 0 {
		return Rolldice()

	} else {
		return i
	}

}
