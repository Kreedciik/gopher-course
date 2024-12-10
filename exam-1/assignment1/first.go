package firstassignment

import (
	"fmt"
	"reflect"
)

func GetValueType(value any) {
	fmt.Printf("The type of %v is %v\n", value, reflect.TypeOf(value))
}
