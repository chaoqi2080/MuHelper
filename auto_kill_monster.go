package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
	"muHelper/util/mouse"
)

// 小地图的坐标。
var globalSmallMapPos = []int{964, 138}

// 靠左的空白点对应的坐标。
var globalLeftBlankPos = []int{310, 243}

// 自动按钮位置
var globalAutoButtonPos = []int{1013, 305}

func DoAutoKillMonster() {
	robotgo.Sleep(3)
	i := 0
	for {
		i++

		if i >= 3 {
			autoRecycle()
			i = 0
		}

		//bingShuang1Golden()
		killAnNingGolden3()

		//killLost4Map()

	}
}

func killMonster(goldenMonsterPos []int, moveTimeBetweenMonster int) {
	if len(goldenMonsterPos) != 3 {
		log.Error("传入的黄金怪坐标数据有问题")
		return
	}

	//打开大地图
	openBigMap()

	//移动到黄金怪位置
	clickGoldenMonsterPosInBigMap(goldenMonsterPos, moveTimeBetweenMonster)

	//开启自动刷怪
	clickAutoButton(goldenMonsterPos[2])
}

func clickGoldenMonsterPosInBigMap(goldenMonsterPos []int, moveTimeBetweenMonster int) {
	mouse.Move2Position([]int{goldenMonsterPos[0], goldenMonsterPos[1]})
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
	mouse.Move2Position(globalAutoButtonPos)
	mouse.SignalLeftClick()

	log.Info(
		"玩家开始自动杀怪, 预计耗时 = %v",
		timeLeft4KillMonster,
	)
	robotgo.Sleep(timeLeft4KillMonster)
}

// closeBigMap and openBigMap 使用同样的指令
// 理由是点一次就是打开大地图，再点一次就是关闭，如此循环
func closeBigMap() {
	mouse.Move2Position(globalSmallMapPos)
	mouse.SignalLeftClick()
}

func openBigMap() {
	mouse.Move2Position(globalSmallMapPos)
	mouse.SignalLeftClick()
}
