package json

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type T struct {
	Name Name    `json:"name"`
	Age  int8    `json:"age"`
	Pi   float64 `json:"pi"`
	Date time.Time
}

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func TestToJsonNoType(t *testing.T) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	_ = json.Unmarshal(b, &f)
	fmt.Printf("%T， %v\n", f, f)
	// f对应的具体类型为
	// f = map[string]interface{}{
	// 	"Name": "Wednesday",
	// 	"Age":  6,
	// 	"Parents": []interface{}{
	// 		"Gomez",
	// 		"Morticia",
	// 	},
	// }

	// 然后可以通过断言来访问底层数据
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	// map[string]interface {}， map[Age:6 Name:Wednesday Parents:[Gomez Morticia]]
}

func TestJson(t *testing.T) {
	type FamilyMember struct {
		Name    string
		Age     int
		Parents *[]string
		// Parents []string
	}

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var m FamilyMember
	_ = json.Unmarshal(b, &m)
	fmt.Printf("%v\n", m)
}

func TestToJSON(t *testing.T) {
	s, _ := json.Marshal(&T{Name: Name{First: "Janet", Last: "Prichard"}, Age: 47, Pi: 3.1415926, Date: time.Now()})
	fmt.Println(string(s))
	// {"name":{"first":"Janet","last":"Prichard"},"age":47,"pi":3.1415926,"Date":"2022-05-25T10:28:47.318021+08:00"}
}

func TestFrom(t *testing.T) {
	var tt T
	s := "{\"name\":{\"first\":\"Janet\",\"last\":\"Prichard\"},\"age\":47,\"pi\":3.1415926,\"Date\":\"2022-05-25T10:28:47.318021+08:00\"}"
	err := json.Unmarshal([]byte(s), &tt)
	if err != nil {
		return
	}
	fmt.Printf("%v", tt)
}
