package main

import (
	"fmt"
	"muHelper/util/mouse"
)

// 背包回收原理：
// 1. 点背包
// 2. 点整理
// 3. 点回收
// 4. 点确认回收
// 5. 点关闭按钮

var recycleStrArray = []string{
	"背包", "整理", "回收", "确认回收",
}

func autoRecycle() {
	for i := 0; i < len(recycleStrArray); i++ {
		mouse.Move2Position(
			[]int{
				gConfigPos.Recycle[i][0], gConfigPos.Recycle[i][1],
			},
			fmt.Sprintf("回收--%v", recycleStrArray[i]),
			false,
		)
		mouse.SignalLeftClick()
	}

	//点关闭按钮
	mouse.Move2Position(
		[]int{gConfigPos.Recycle[0][0] - 900, gConfigPos.Recycle[0][1]},
		"回收--点关闭按钮",
		false,
	)
	mouse.DoubleLeftClick()
}
