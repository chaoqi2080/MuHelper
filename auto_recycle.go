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
	{1018, 245}, //背包
	{987, 527},  //整理
	{926, 530},  //回收
	{900, 508},  //确认回收
}

func autoRecycle() {
	if len(recyclePos) != 4 {
		log.Error("背包坐标不等于 4 个，有问题。")
		return
	}

	//点背包
	mouse.Move2Position(recyclePos[BeiBao], "回收--点背包")
	mouse.SignalLeftClick()

	//点整理
	mouse.Move2Position(recyclePos[ZhengLi], "回收--点整理")
	mouse.SignalLeftClick()

	//点回收
	mouse.Move2Position(recyclePos[HuiShou], "回收--点回收")
	mouse.SignalLeftClick()

	//点确认回收
	mouse.Move2Position(recyclePos[HuiShouQueRen], "回收--点确认回收")
	mouse.SignalLeftClick()

	//点关闭按钮
	mouse.Move2Position(globalLeftBlankPos, "回收--点关闭按钮")
	mouse.DoubleLeftClick()
}
