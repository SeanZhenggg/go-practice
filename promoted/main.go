package main

import (
	"fmt"
)

type Animal struct {
	string
	int
}

// error
// field must be unique
// type Req struct {
// 	http.Request
// 	Request int
// }

type Person struct {
	name string
	age  int32
}

func (p Person) IsAdult() bool {
	return p.age >= 18
}

type Employee struct {
	position string
}

func (e Employee) IsManager() bool {
	return e.position == "manager"
}

type Record struct {
	Person
	Employee
}

func main() {
	var a = &Animal{"123", 2}

	// anonymous or embedded fields
	fmt.Printf("🍎🍎🍎🍎🍎🍎 a : %v\n", a)
	fmt.Printf("🍎🍎🍎🍎🍎🍎 a.string : %v\n", a.string)
	fmt.Printf("🍎🍎🍎🍎🍎🍎 a.string : %v\n", a.int)

	record := Record{}
	record.name = "Michał"
	record.age = 29
	record.position = "software engineer"
	fmt.Println(record)             // {{Michał 29} {software engineer}}
	fmt.Println(record.name)        // Michał
	fmt.Println(record.age)         // 29
	fmt.Println(record.position)    // software engineer
	fmt.Println(record.IsAdult())   // true
	fmt.Println(record.IsManager()) // false

	record2 := Record{Person{}, Employee{}}

	record2.name = "Sean"
	record2.age = 25
	record2.position = "software engineer"
	fmt.Println(record2)
	fmt.Println(record2.name)
	fmt.Println(record2.age)
	fmt.Println(record2.position)
	fmt.Println(record2.IsAdult())
	fmt.Println(record2.IsManager())
}
