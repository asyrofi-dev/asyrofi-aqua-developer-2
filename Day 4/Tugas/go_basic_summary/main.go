package main

import "fmt"

func main() {

	//	==========================================================================
	//	                              DATA TYPE
	//	==========================================================================
	//	1. Decimal
	//	--------------------------------------------------------------------------

	var decimalNumber = 2.62312

	fmt.Printf("bilangan desimal : %f\n", decimalNumber)
	fmt.Printf("bilangan desimal : %.3f\n", decimalNumber)

	//	--------------------------------------------------------------------------
	//	2. Boolean
	//	--------------------------------------------------------------------------

	var exist bool = false

	fmt.Printf("exist? %t\n", exist)

	//	--------------------------------------------------------------------------
	//	3. String
	//	--------------------------------------------------------------------------

	var message = `Nama saya "Muhammad Asyrofi".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message)

	//	========================= END OF "DATA TYPE" =============================

	//	##########################################################################

	//	                               VARIABLE
	//	==========================================================================
	//	1. Variable Declaration
	//	--------------------------------------------------------------------------

	var firstName string = "Muhammad"

	var lastName string
	lastName = "Asyrofi"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	//	--------------------------------------------------------------------------
	//	2. Variable Declaration Without Data Type
	//	--------------------------------------------------------------------------

	var namaDepan string = "Muhammad"
	namaBelakang := "Asyrofi"

	fmt.Printf("halo %s %s!\n", namaDepan, namaBelakang)

	//	--------------------------------------------------------------------------
	//	3. Multi Variable Declaration
	//	--------------------------------------------------------------------------

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(first, second, third, seventh, eight, ninth)

	//	--------------------------------------------------------------------------
	//	4. Underscore Variable Declaration
	//	--------------------------------------------------------------------------

	_, title := "Belajar Golang", "eFishery Academy"
	subtitle, _ := "Aqua Developer", "Batch 2"

	fmt.Println(title, subtitle)

	//	--------------------------------------------------------------------------
	//	5. Constant
	//	--------------------------------------------------------------------------

	const bookTitle = "The Talent Code"

	fmt.Println(bookTitle)

	//	========================= END OF "VARIABLE" =============================

	//	##########################################################################

	//	                        CONDITIONAL STATEMENT
	//	==========================================================================

	//	--------------------------------------------------------------------------
	//	1. If - Else If - Else
	//	--------------------------------------------------------------------------

	var point = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point >= 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Printf("Tidak Lulus. Nilai anda %d\n", point)
	}

	//	--------------------------------------------------------------------------
	//	2.Switch - Case
	//	--------------------------------------------------------------------------

	var value = 6

	switch value {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}

	//	===================== END OF "CONDITIONAL STATEMENT" =====================

	//	##########################################################################

	//	                                LOOPING
	//	==========================================================================
	//	1. For - Range
	//	--------------------------------------------------------------------------

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("Elemen %d : %s\n", i, fruit)
	}

	//	--------------------------------------------------------------------------
	//	2.For
	//	--------------------------------------------------------------------------

	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//	--------------------------------------------------------------------------
	//	3.For - Break - Continue
	//	--------------------------------------------------------------------------

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	//	=========================== END OF "LOOPING" =============================

	//	##########################################################################

	//	                             DATA STRUCTURE
	//	==========================================================================
	//	1. Map
	//	--------------------------------------------------------------------------

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("Januari", chicken["januari"], "ayam")
	fmt.Println("Februari", chicken["februari"], "ayam")

	//	--------------------------------------------------------------------------
	//	2. Array
	//	--------------------------------------------------------------------------

	var names [4]string
	names[0] = "alphandi"
	names[1] = "betania"
	names[2] = "charlieno"
	names[3] = "deltanudin"

	fmt.Println(names[0], names[1], names[2], names[3])

	//	--------------------------------------------------------------------------
	//	3. Slice
	//	--------------------------------------------------------------------------

	var fruitList = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruitList[0:3]
	var bFruits = fruitList[1:4]

	var aaFruits = fruitList[1:2]
	var bbFruits = fruitList[0:1]

	fmt.Println(fruitList)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	bbFruits[0] = "pinnaple"

	fmt.Println(fruitList)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	//	======================= END OF "DATA STRUCTURE" ==========================

	//	##########################################################################

	//	                                DEFER
	//	==========================================================================
	//	1. Defer
	//	--------------------------------------------------------------------------

	defer fmt.Println("Halo")
	fmt.Println("Selamat Datang")

	//	=========================== END OF "DEFER" ===============================

	//	--------------------------------------------------------------------------
	//	# Function Call
	//	--------------------------------------------------------------------------
	a, b := swap("Hello", "world")
	fmt.Println(a, b)

	callAdd := add(5, 6)
	callSubtract := subtract(25, 4)

	fmt.Println(callAdd)
	fmt.Println(callSubtract)

	//	--------------------------------------------------------------------------
	//	# Struct Use Case
	//	--------------------------------------------------------------------------

	var s1 student
	s1.name = "Asyrofi"
	s1.grade = 12

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	//	--------------------------------------------------------------------------
	//	# Embedded Struct Use Case
	//	--------------------------------------------------------------------------

	var s2 schoolStudent
	s2.name = "Muhammad Asyrofi"
	s2.age = 22
	s2.grade = 12

	fmt.Println("Name :", s2.name)
	fmt.Println("Age :", s2.age)
	fmt.Println("Age :", s2.person.age)
	fmt.Println("Grade :", s2.grade)

	//	--------------------------------------------------------------------------
	//	# Struct with Slice
	//	--------------------------------------------------------------------------

	var allStudents = []person{
		{name: "Alphani", age: 18},
		{name: "Betaria", age: 20},
		{name: "Charliesi", age: 19},
		{name: "Deltaria", age: 21},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

}

//	##############################################################################

//	                                  FUNCTION
//	==============================================================================
//	1. Function Declaration
//	------------------------------------------------------------------------------

func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

//	------------------------------------------------------------------------------
//	2. Multi Return Value
//	------------------------------------------------------------------------------

func swap(x, y string) (string, string) {
	return y, x
}

//	============================== END OF "FUNCTION" =============================

//	##############################################################################

//	                                   STRUCT
//	==============================================================================
//	1. Struct
//	------------------------------------------------------------------------------

type student struct {
	name  string
	grade int
}

//	------------------------------------------------------------------------------
//	2. Embedded Struct
//	------------------------------------------------------------------------------

type person struct {
	name string
	age  int
}

type schoolStudent struct {
	grade int
	person
}
