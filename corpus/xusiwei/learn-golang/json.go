package main

import "os"
import "fmt"
import "encoding/json"

type Response1 struct {
	Page int
	Fruits []string
}
type Response2 struct {
	Page int        `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	jb, _ := json.Marshal(true)
	fmt.Println(string(jb))

	ji, _ := json.Marshal(1)
	fmt.Println(string(ji))

	jf, _ := json.Marshal(2.34)
	fmt.Println(string(jf))

	jstr, _ := json.Marshal("gopher")
	fmt.Println(string(jstr))

	slice, _ := json.Marshal([]string{"apple", "peach", "pear"})
	fmt.Println("slice:", slice)

	dict, _ := json.Marshal(map[string]int{"apple":5, "lettuce":7})
	fmt.Println("dict :", dict)

	r1 := &Response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	jr1, _ := json.Marshal(r1)
	fmt.Println("jr1:", jr1)

	r2 := &Response2{
		Page: 2,
		Fruits: []string{"apple", "peach", "pear"}}
	jr2, _ := json.Marshal(r2)
	fmt.Println("jr2:", jr2)
	fmt.Println()


	bs := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var data map[string] interface{}

	if err := json.Unmarshal(bs, &data); err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal from `%s` => %v\n", string(bs), data)

	num := data["num"].(float64) // type assertion
	fmt.Println("num:", num)

	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("strs[0]:", str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Printf("Unmarshal from `%s` => %v", str, res)
	fmt.Println("res.Fruits:", res.Fruits)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
