package mouse

import (
	"github.com/go-vgo/robotgo"
	"muHelper/util/log"
)

// Move2Position 移动鼠标到游戏指定位置
func Move2Position(pos []int) {
	if len(pos) != 2 {
		log.Error("传入的坐标点不合法")
		return
	}

	robotgo.MoveMouseSmooth(pos[0], pos[1], 1.0, 5.0)
}

func SignalLeftClick() {
	robotgo.MouseClick("left", false)
}

func DoubleLeftClick() {
	robotgo.MouseClick("left", false)
	robotgo.MouseClick("left", false)
}
