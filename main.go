package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
)

func main() {
	log.Config("./log/mu_helper")

	robotgo.Sleep(2)

	//校准地图点位置
	//GetFramePos()
	//自动打怪。
	DoAutoKillMonster()
}

func GetFramePos() {
	for {
		x, y := robotgo.GetMousePos()
		log.Info("x = %v, y = %v", x, y)
		robotgo.Sleep(1)
	}
}
