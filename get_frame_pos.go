package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
)

func GetFramePos() {
	for {
		x, y := robotgo.GetMousePos()
		log.Info("x = %v, y = %v", x, y)
		robotgo.Sleep(1)
	}
}
