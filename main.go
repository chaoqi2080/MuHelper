package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
	"time"
)

var gConfigPos = LocalPos{}

func main() {
	log.Config("./log/mu_helper")

	if !ReadConfigFromFile("config.json") {
		log.Error("读取坐标配置文件失败")
		return
	}

	for i := 0; i < len(gConfigPos.Golden); i++ {
		log.Info("i = %v, => %v", i, gConfigPos.Golden[i])

		for j := 0; j < len(gConfigPos.Golden[i]); j++ {
			log.Info("j = %v, => %v", j, gConfigPos.Golden[i][j])
		}
	}

	return

	robotgo.Sleep(2)

	if gConfigPos.IsGetPos {
		//校准地图点位置
		GetFramePos()
		return
	} else {
		//自动打怪。
		DoAutoKillMonster()
	}
}

func DoAutoKillMonster() {
	robotgo.Sleep(3)
	killMonsterCount := 0

	killBossTime := time.Now().Unix()
	internalTime := int64(gConfigPos.BossRefreshTime * 60)
	for {
		killMonsterCount++

		if killMonsterCount >= 2 {
			autoRecycle()
			killMonsterCount = 0
		}

		// 打一次 boss
		bNeedKillBoss := (time.Now().Unix() - killBossTime) > internalTime
		if bNeedKillBoss {
			for i := 0; i < len(gConfigPos.Boss); i++ {
				killMonster(
					[]int{
						gConfigPos.Boss[i][0], gConfigPos.Boss[i][1],
					},
					gConfigPos.Boss[i][2],
					gConfigPos.Boss[i][3],
				)

				if i == 0 {
					//重置一下时间
					killBossTime = time.Now().Unix()
				}
			}

			continue
		}

		//杀黄金怪
		for i := 0; i < len(gConfigPos.Golden); i++ {
			killMonster(
				[]int{
					gConfigPos.Golden[i][0], gConfigPos.Golden[i][1],
				},
				gConfigPos.Golden[i][2],
				gConfigPos.Golden[i][3],
			)
		}
	}
}

func GetFramePos() {
	var tipsArray = []string{
		"<最左边位置>的坐标",
		"<最右上的小地图>坐标",
		"<自动杀怪>坐标",
		"<背包>坐标",
		"<整理>坐标",
		"<回收>坐标",
		"<确认回收>坐标",
	}
	for {
		for i := 0; i < len(tipsArray); i++ {
			log.Info("开始校准 ========= %v =========", tipsArray[i])
			for j := 0; j < 10; j++ {
				x, y := robotgo.Location()
				log.Info("%v => x = %v, y = %v", tipsArray[i], x, y)
				robotgo.MilliSleep(500)
			}
			robotgo.Sleep(2)
		}

	}
}
