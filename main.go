package main

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
)

func main() {
	log.Config("./log/mu_helper")

	robotgo.Sleep(2)

	//校准地图点位置
	//getGamePos()
	//自动打怪。
	autoKillMonster()
}

func getGamePos() {
	for {
		x, y := robotgo.GetMousePos()
		log.Info("x = %v, y = %v", x, y)
		robotgo.Sleep(1)
	}
}

func autoKillMonster() {
	i := 0
	for {
		i++

		//if i >= 1 {
		autoRecycle()
		i = 0
		//}

		//bingShuang1Golden()
		//killAnNingGolden3()

		killLost4Map()

	}
}

func killLost4Map() {
	//死神骑士
	killMonster(lost4Pos[0])
	//黄金恶魔王
	killMonster(lost4Pos[1])
}

func killAnNingGolden3() {
	//再杀黄金毒树妖
	killMonster(anNing3[0])
	//先杀树妖
	killMonster(anNing3[1])

}

func bingShuang1Golden() {
	killMonster(bingShuang1[0])
	killMonster(bingShuang1[1])
}
