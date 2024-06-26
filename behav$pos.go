package main

// 回收对应的几个坐标点。
var recyclePos = [][]int{
	{926, 256}, //背包
	{899, 533}, //整理
	{841, 532}, //回收
	{813, 510}, //确认回收
}

// 小地图的坐标。
var smallMapPos = []int{891, 141}

// 靠左的空白点对应的坐标。
var leftBlankPos = []int{259, 302}

// 自动按钮位置
var autoPos = []int{930, 303}

// 黄金之间移动的时间(秒)
var moveTakeTime = 16

// 坐标说明 {x, y, 杀死对应黄金耗时}
// =============== 失落之塔 4 ===============
var lost4Pos = [][]int{
	{436, 286, 40},
	{525, 321, 38},
}

// =============== 冰霜之城 1 ===============
var bingShuang1 = [][]int{
	{638, 201, 75},
	{685, 268, 75},
}

// =============== 安宁池 3 点钟 ===============
var anNing3 = [][]int{
	{753, 408, 137},
	{775, 359, 137},
}
