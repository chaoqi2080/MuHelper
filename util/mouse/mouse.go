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

	robotgo.MoveSmooth(pos[0], pos[1], 1.0, 1.0)
}

func SignalLeftClick() {
	robotgo.Click()
}

func DoubleLeftClick() {
	robotgo.Click()
	robotgo.Click()
}
