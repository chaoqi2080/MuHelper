package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
	"muHelper/util/mouse"
)

func killMonster(pos []int) {
	if len(pos) != 3 {
		log.Error("传入的黄金怪坐标数据有问题")
		return
	}
	//打开小地图进行移动
	mouse.OpenSmallMap(smallMapPos)
	log.Info("开始移动到指定位置")
	mouse.Move2Position([]int{pos[0], pos[1]})
	mouse.DoubleLeftClick()
	waitPlayerMove()
	log.Info("完成移动目标，点完自动。")
	hitAutoKillMonster(pos[2])
}

func waitPlayerMove() {
	log.Warn("玩家<开始 移动>")
	robotgo.Sleep(moveTakeTime)
	log.Warn("玩家<完成 移动>，耗时 = %v 秒", moveTakeTime)
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
