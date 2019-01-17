package main
import "github.com/vishesh/Simply-Go/src/worker"

func main() {
	worker.SimpleSwitch(2)
	worker.CheckWeekend()
	worker.TellTime()
	worker.WhatAmI("this is it.")
}
