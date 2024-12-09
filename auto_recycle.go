package main

import (
	"muHelper/util/log"
	"muHelper/util/mouse"
)

// 背包回收原理：
// 1. 点背包
// 2. 点整理
// 3. 点回收
// 4. 点确认回收
// 5. 点关闭按钮
const (
	BeiBao        = 0
	ZhengLi       = 1
	HuiShou       = 2
	HuiShouQueRen = 3
)

// 回收对应的几个坐标点。
var recyclePos = [][]int{
	{1014, 259}, //背包
	{984, 537},  //整理
	{926, 540},  //回收
	{896, 515},  //确认回收
}

func autoRecycle() {
	if len(recyclePos) != 4 {
		log.Error("背包坐标不等于 4 个，有问题。")
		return
	}

	//点背包
	log.Info("1--点背包")
	mouse.Move2Position(recyclePos[BeiBao])
	//robotgo.Sleep(1)
	mouse.SignalLeftClick()

	//点整理
	log.Info("2--点整理")
	mouse.Move2Position(recyclePos[ZhengLi])
	//robotgo.Sleep(1)
	mouse.SignalLeftClick()

	//点回收
	log.Info("3--点回收")
	mouse.Move2Position(recyclePos[HuiShou])
	//robotgo.Sleep(1)
	mouse.SignalLeftClick()

	//点确认回收
	log.Info("4--点确认回收")
	mouse.Move2Position(recyclePos[HuiShouQueRen])
	//robotgo.Sleep(1)
	mouse.SignalLeftClick()

	//点关闭按钮
	log.Info("5--点关闭按钮")
	mouse.Move2Position(globalLeftBlankPos)
	//robotgo.Sleep(1)
	mouse.DoubleLeftClick()
}
