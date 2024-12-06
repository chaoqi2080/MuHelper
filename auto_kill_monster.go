package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
	"muHelper/util/mouse"
)

// 小地图的坐标。
var smallMapPos = []int{4211, -427}

// 靠左的空白点对应的坐标。
var leftBlankPos = []int{4264, -256}

// 自动按钮位置
var autoPos = []int{3613, -295}

func DoAutoKillMonster() {
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

func killMonster(pos []int, moveTimeBetweenMonster int) {
	if len(pos) != 3 {
		log.Error("传入的黄金怪坐标数据有问题")
		return
	}
	//打开小地图进行移动
	mouse.OpenSmallMap(smallMapPos)
	log.Info("开始移动到指定位置")
	//在小地图点击第一个位置，然后移动
	mouse.Move2Position([]int{pos[0], pos[1]})
	mouse.DoubleLeftClick()
	//等待玩家移动到对应的位置
	waitPlayerMove(moveTimeBetweenMonster)
	log.Info("完成移动目标，点完自动。")
	hitAutoKillMonster(pos[2])
}

func waitPlayerMove(moveTimeBetweenMonster int) {
	log.Warn("玩家<开始 移动>")
	robotgo.Sleep(moveTimeBetweenMonster)
	log.Warn("玩家<完成 移动>，耗时 = %v 秒", moveTimeBetweenMonster)
}

func hitAutoKillMonster(takeTime int) {
	//点小地图，关闭大地图。
	mouse.Move2Position(smallMapPos)
	mouse.SignalLeftClick()

	//点手动->自动
	mouse.Move2Position(autoPos)
	mouse.SignalLeftClick()

	log.Info("<开始自动杀怪>, 预计耗时 = %v", takeTime)
	robotgo.Sleep(takeTime)
	log.Info("<完成自动杀怪>, 耗时 = %v", takeTime)
}
