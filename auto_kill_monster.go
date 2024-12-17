package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
	"muHelper/util/mouse"
)

// 小地图的坐标。
var globalSmallMapPos = []int{931, 119}

// 靠左的空白点对应的坐标。
var globalLeftBlankPos = []int{124, 140}

// 自动按钮位置
var globalAutoButtonPos = []int{926, 273}

// // 黄金之间移动的时间(秒)
func killMonster(goldenMonsterPos []int, timeLeft4KillMonster int, moveTimeBetweenMonster int) {
	if len(goldenMonsterPos) != 2 {
		log.Error("传入的黄金怪坐标数据有问题")
		return
	}

	//打开大地图
	openBigMap()

	//移动到黄金怪位置
	clickGoldenMonsterPosInBigMap(goldenMonsterPos, moveTimeBetweenMonster)

	//开启自动刷怪
	clickAutoButton(timeLeft4KillMonster)
}

func clickGoldenMonsterPosInBigMap(goldenMonsterPos []int, moveTimeBetweenMonster int) {
	mouse.Move2Position([]int{goldenMonsterPos[0], goldenMonsterPos[1]}, "在大地图中点黄金怪坐标")
	mouse.DoubleLeftClick()
	//等待玩家移动到对应的位置
	log.Info(
		"玩家需要移动到 = [%v,%v], 耗时 = %v 秒",
		goldenMonsterPos[0], goldenMonsterPos[1], moveTimeBetweenMonster,
	)
	robotgo.Sleep(moveTimeBetweenMonster)
}

func clickAutoButton(timeLeft4KillMonster int) {
	closeBigMap()

	//点手动->自动
	//使用这个的目的是避免鼠标移动太快，导致失效
	mouse.Move2Position(globalAutoButtonPos, "点自动打怪")
	mouse.SignalLeftClick()

	log.Info(
		"玩家开始自动杀怪, 预计耗时 = %v",
		timeLeft4KillMonster,
	)
	robotgo.Sleep(timeLeft4KillMonster)
}

// 理由是点一次就是打开大地图，再点一次就是关闭，如此循环
func closeBigMap() {
	mouse.Move2Position(globalSmallMapPos, "点小地图--关闭大地图")
	mouse.SignalLeftClick()
	robotgo.Sleep(1) //停止1秒，避免误差点击
}

func openBigMap() {
	mouse.Move2Position(globalSmallMapPos, "点小地图--打开大地图")
	mouse.SignalLeftClick()
}
