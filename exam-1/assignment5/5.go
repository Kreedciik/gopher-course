package fifthassignment

import "fmt"

func IncDecOperation(operations []string) {
	initialValue := 0

	fmt.Println("Initial value: ", initialValue)
	for _, operation := range operations {
		switch operation {
		case "++X", "X++":
			initialValue++
		case "--X", "X--":
			initialValue--
		}

	}
	fmt.Println("After inc/dec operations: ", initialValue)
}
