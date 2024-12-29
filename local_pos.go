package main

import (
	"encoding/json"
	"io"
	"muHelper/util/log"
	"os"
)

type LocalPos struct {
	IsGetPos            bool  `json:"is_get_pos"`
	Auto                []int `json:"auto"`
	PosBackPack         []int `json:"back_pack"`
	ZhengLi             []int `json:"zheng_li"`
	HuiShou             []int `json:"hui_shou"`
	HuiShou2            []int `json:"hui_shou_2"`
	Golden1             []int `json:"golden_1"`
	Golden2             []int `json:"golden_2"`
	IsKillBoss          bool  `json:"is_kill_boss"`
	Boss                []int `json:"boss"`
	BossMove2GoldenTime int   `json:"boss_move2_golden_time"`
	BossRefreshTime     int   `json:"boss_refresh_time"`
}

func ReadConfigFromFile(fileName string) bool {
	configFile, err := os.Open(fileName)
	if err != nil {
		log.Error("打开配置文件错误 %v", err.Error())
		return false
	}

	defer configFile.Close()

	dataArray, err := io.ReadAll(configFile)
	if err != nil {
		log.Error("读取配置文件数据失败 %v", err.Error())
		return false
	}

	if err = json.Unmarshal(dataArray, &gConfigPos); err != nil {
		log.Error("解析 json 文件失败 %v", err.Error())
		return false
	}

	log.Info("读取文件成功 %v", gConfigPos)
	return true
}
