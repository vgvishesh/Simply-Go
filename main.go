package main
import "github.com/vgvishesh/Simply-Go/worker"

func main() {
	worker.SimpleSwitch(2)
	worker.CheckWeekend()
	worker.TellTime()
	worker.WhatAmI("this is it.")
}
