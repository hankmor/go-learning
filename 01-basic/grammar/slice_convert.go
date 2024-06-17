package main

import "fmt"

func main() {
	ids := []int{1, 2, 3}
	// f1(ids) // cannot use ids (variable of type []int) as []interface{} value in argument to f1
	// f1([]interface{}(ids)) // ./slice_convert.go:6:19: cannot convert ids (variable of type []int) to type []interface{}
	// so i have to loop the ids to convert it
	var dest []interface{}
	for _, v := range ids {
		dest = append(dest, v)
	}
	f1(dest)
}

func f1(ids []interface{}) {
	fmt.Print(ids...)
}
