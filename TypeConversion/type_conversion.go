package typeconversion

import (
	"fmt"
	"strconv"
)

func TypeConversion() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32)

	var value16 int16 = int16(value64)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)

	var name = "lemoncode21"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)

	//Exercises
	fmt.Println("---Question 1---")
	intVar1 := 42
	floatVar := float64(intVar1)
	fmt.Printf("intVar: %d, floatVar: %f\n", intVar1, floatVar)

	fmt.Println("---Question 2---")
	strVar := "123"
	intVar2, _ := strconv.Atoi(strVar)
	fmt.Printf("strVar: %s, intVar: %d\n", strVar, intVar2)

	fmt.Println("---Question 3 ---")
	floatVar3 := 5.678
	intVar3 := int(floatVar3)
	newFloatVar := float64(intVar3)
	fmt.Printf("floatVar: %f, intVar: %d, newFloatVar: %f\n", floatVar3, intVar3, newFloatVar)
}
