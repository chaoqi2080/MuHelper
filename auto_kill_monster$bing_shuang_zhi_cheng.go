package main

// 黄金之间移动的时间(秒)
var moveTakeTimeBing1 = 16

// =============== 冰霜之城 1 ===============
var bingShuang1 = [][]int{
	{638, 201, 75},
	{685, 268, 75},
}

func bingShuang1Golden() {
	killMonster(bingShuang1[0], moveTakeTimeBing1)
	killMonster(bingShuang1[1], moveTakeTimeBing1)
}
