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
	backpackX := gConfigPos.PosBackPack[0]
	backpackY := gConfigPos.PosBackPack[1]
	//点背包
	mouse.Move2Position(
		[]int{backpackX, backpackY},
		"回收--点背包",
	)
	mouse.SignalLeftClick()

	//点整理
	mouse.Move2Position(
		[]int{backpackX - 32, backpackY + 380},
		"回收--点整理",
	)
	mouse.SignalLeftClick()

	//点回收
	mouse.Move2Position(
		[]int{backpackX - 90, backpackY + 380},
		"回收--点回收",
	)
	mouse.SignalLeftClick()

	//点确认回收
	mouse.Move2Position(
		[]int{backpackX - 120, backpackY + 260},
		"回收--点确认回收",
	)
	mouse.SignalLeftClick()

	//点关闭按钮
	mouse.Move2Position([]int{gConfigPos.PosBackPack[0] - 900, gConfigPos.PosBackPack[1]},
		"回收--点关闭按钮",
	)
	mouse.DoubleLeftClick()
}
