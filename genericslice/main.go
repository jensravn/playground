package genericslice

import "fmt"

func main() {
	mySlice := []MyStruct{{SomeField: "some value"}}
	genericSliceFunc(mySlice)
}

type MyStruct struct {
	SomeField string
}

func genericSliceFunc[T any](tSlice []T) {
	for _, t := range tSlice {
		t := t
		fmt.Println(t)
	}
}
