package main

import "fmt"

type ListaHeterogenea struct {
	Data []interface{}
}

func main() {
	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "Olá")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)

}
