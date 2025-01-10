package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
	"muHelper/util/mouse"
)

// // 黄金之间移动的时间(秒)
func killMonster(goldenMonsterPos []int, timeLeft4KillMonster int, moveTimeBetweenMonster int) {
	if len(goldenMonsterPos) != 2 {
		log.Error("传入的怪坐标数据有问题")
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
	mouse.Move2Position(
		[]int{goldenMonsterPos[0], goldenMonsterPos[1]},
		fmt.Sprintf("在大地图中点怪坐标, 耗时 = %v", moveTimeBetweenMonster),
		true,
	)
	mouse.DoubleLeftClick()

	closeBigMap()
	//等待玩家移动到对应的位置
	//log.Info(
	//	"玩家需要移动到 = [%v,%v], 耗时 = %v 秒",
	//	goldenMonsterPos[0], goldenMonsterPos[1], moveTimeBetweenMonster,
	//)
	robotgo.Sleep(moveTimeBetweenMonster)
}

func clickAutoButton(timeLeft4KillMonster int) {
	//点手动->自动
	//使用这个的目的是避免鼠标移动太快，导致失效
	mouse.Move2Position(
		[]int{gConfigPos.Auto[0], gConfigPos.Auto[1]},
		"点自动打怪",
		true,
	)
	mouse.SignalLeftClick()

	log.Info(
		"玩家开始自动杀怪, 预计耗时 = %v",
		timeLeft4KillMonster,
	)

	robotgo.Sleep(timeLeft4KillMonster)
}

// 理由是点一次就是打开大地图，再点一次就是关闭，如此循环
func closeBigMap() {
	opMap("点小地图--关闭大地图")
}

func openBigMap() {
	opMap("点小地图--打开大地图")
}

func opMap(str string) {
	mouse.Move2Position(
		[]int{gConfigPos.Recycle[0][0], gConfigPos.Recycle[0][1] - 150},
		str,
		false,
	)
	mouse.SignalLeftClick()
}
