package main

import (
	"fmt"
)

func part1() {
	var intNum64 int64 = 42
	var intNum32 int32 = 42
	var unknwonInt int32
	fmt.Println(int64(intNum32) == intNum64)
	fmt.Println(unknwonInt)

	var text string = `first
	\nsecond
	\nthird
	`
	var unknownText string
	fmt.Println("unknown text -> ", unknownText)
	fmt.Println(text)

	var runesText []byte = []byte{208, 159, 209, 128, 208, 184, 208,
		178, 208, 181, 209, 130, 32, 208,
		188, 208, 184, 209, 128}
	fmt.Println(string(runesText))

	runesText2 := []byte{208, 159, 209, 128, 208, 184, 208,
		178, 208, 181, 209, 130, 32, 208,
		188, 208, 184, 209, 128}
	fmt.Println(string(runesText2))
}

func part2() {
	var gdRes string = fmt.Sprintf("Distance between %d and %d in meters is %d",
		START_POINT_METERS, END_POINT_METERS, get_distance(START_POINT_METERS, END_POINT_METERS))
	fmt.Println(gdRes)

	var personD Person = Person{"Dmitriy", 20}
	var personI Person = Person{"Irina", 20}

	fmt.Println("person names -> ", personD.name, personI.name)
	fmt.Println("is same person -> ", personD == personI)

	var personsList []Person = []Person{personD, personI}
	for i, p := range personsList {
		fmt.Println(i, p)
	}

}
