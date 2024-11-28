package helloworld

import "fmt"

func HelloWorld() {
	fmt.Println("Hello, World!")

	example_students()
	example_multi_case_switch(5)
	example_struct_person()
	example_make_map()
	example_difference_using_map_or_not()
}

func example_students() {
	fmt.Println("\n[Example: Students]")
	var student1 string = "John" // type is string
	var student2 = "Jane"        // type is inferred
	x := 10                      // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}

func example_multi_case_switch(day int) {
	fmt.Println("\n[Example: Multi-case Switch]")

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

type Person struct {
	name   string
	age    int
	job    string
	salary float64
}

func (p Person) print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Job:", p.job)
	fmt.Println("Salary:", p.salary)
}

func example_struct_person() {
	fmt.Println("\n[Example: Struct Person]")

	var pers1 = Person{"John", 25, "Software Engineer", 50000.0}
	var pers2 = Person{"Jane", 30, "Data Scientist", 60000.0}

	pers1.print()
	fmt.Println()
	pers2.print()
}

func example_difference_using_map_or_not() {
	fmt.Println("\n[Example: Difference Using Map or Not]")

	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}

func example_make_map() {
	fmt.Println("\n[Example: Make Map]")

	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	delete(a, "year")
	fmt.Printf("a\t%v\n", a)
}
