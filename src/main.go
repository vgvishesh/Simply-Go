package main
import "./worker"

func main() {
	worker.SimpleSwitch(2)
	worker.CheckWeekend()
	worker.TellTime()
	worker.WhatAmI("this is it.")
}
