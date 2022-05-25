package json

import (
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

const str = `{"name":{"first":"Janet","last":"Prichard"},"age":47,"pi":3.1415926}`

func TestGjsonFrom(t *testing.T) {
	value := gjson.Get(str, "name.last")
	println(value.String())
}

func TestGjsonFrom1(t *testing.T) {
	// var m map[string]any = make(map[string]any)
	result := gjson.Parse(str).Get("name").Get("last")
	println(result.String())

	m, ok := gjson.Parse(str).Value().(map[string]any)
	if !ok {
		// not a map
	} else {
		fmt.Printf("%v", m)
	}
}
