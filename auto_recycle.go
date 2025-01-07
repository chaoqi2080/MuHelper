package main

import (
	"muHelper/util/mouse"
)

// 背包回收原理：
// 1. 点背包
// 2. 点整理
// 3. 点回收
// 4. 点确认回收
// 5. 点关闭按钮

func autoRecycle() {
	//点背包ΩΩ
	mouse.Move2Position(
		[]int{gConfigPos.PosBackPack[0], gConfigPos.PosBackPack[1]},
		"回收--点背包",
	)
	mouse.SignalLeftClick()

	//点整理
	mouse.Move2Position(
		[]int{gConfigPos.ZhengLi[0], gConfigPos.ZhengLi[1]},
		"回收--点整理",
	)
	mouse.SignalLeftClick()

	//点回收
	mouse.Move2Position(
		[]int{gConfigPos.HuiShou[0], gConfigPos.HuiShou[1]},
		"回收--点回收",
	)
	mouse.SignalLeftClick()

	//点确认回收
	mouse.Move2Position(
		[]int{gConfigPos.HuiShou2[0], gConfigPos.HuiShou2[1]},
		"回收--点确认回收",
	)
	mouse.SignalLeftClick()

	//点关闭按钮
	mouse.Move2Position([]int{gConfigPos.PosBackPack[0] - 900, gConfigPos.PosBackPack[1]},
		"回收--点关闭按钮",
	)
	mouse.DoubleLeftClick()
}
