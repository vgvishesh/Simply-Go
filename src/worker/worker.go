package worker
import (
	"fmt"
	"time"
)

type DataType int
const (
	INT DataType = 1
	BOOL DataType = 2
	STRING DataType = 3
	UNKNOWN	DataType = 4
)

func WhatAmI(i interface{}) DataType{
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
		return BOOL
	case int:
		fmt.Println("I'm an int")
		return INT
	case string:
		fmt.Println(i)
		return STRING
	default:
		fmt.Printf("Don't know type %T\n", t)
		return UNKNOWN
	}
}

func SimpleSwitch(i int) {
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

func CheckWeekend() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}

func TellTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}