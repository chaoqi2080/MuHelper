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

	killBossTime := time.Now().Unix()
	internalTime := int64(gConfigPos.Boss[4] * 60)
	for {
		//杀黄金怪
		for i := 0; i < len(gConfigPos.Golden); i++ {
			killMonster(
				[]int{
					gConfigPos.Golden[i][0], gConfigPos.Golden[i][1],
				},
				gConfigPos.Golden[i][2],
				gConfigPos.Golden[i][3],
			)
			// 打一次 boss
			if (time.Now().Unix() - killBossTime) > internalTime {
				killMonster(
					[]int{
						gConfigPos.Boss[0], gConfigPos.Boss[1],
					},
					gConfigPos.Boss[2],
					gConfigPos.Boss[3],
				)

				//重置一下时间
				killBossTime = time.Now().Unix()
			}
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
