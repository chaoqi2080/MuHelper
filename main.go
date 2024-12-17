package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
	"time"
)

func main() {
	log.Config("./log/mu_helper")

	robotgo.Sleep(2)

	//校准地图点位置
	//GetFramePos()
	//自动打怪。
	DoAutoKillMonster()
}

func DoAutoKillMonster() {
	robotgo.Sleep(3)
	killMonsterCount := 0

	killBossTime := time.Now().Unix()
	internalTime := int64(40 * 60) // 40 分钟
	for {
		killMonsterCount++

		if killMonsterCount >= 3 {
			autoRecycle()
			killMonsterCount = 0
		}

		// 40 分钟打一次 boss
		bNeedKillBoss := (time.Now().Unix() - killBossTime) > internalTime
		if bNeedKillBoss {
			//重置一下时间
			killBossTime = time.Now().Unix()
			killAnNingBoss9()
		}
		killAnNingGolden9()
	}
}

func GetFramePos() {
	for {
		x, y := robotgo.Location()
		log.Info("x = %v, y = %v", x, y)
		robotgo.Sleep(1)
	}
}
