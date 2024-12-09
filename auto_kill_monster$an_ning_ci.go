package main

// 黄金之间移动的时间(秒)
var moveTakeTimePos3 = 14

// =============== 安宁池 3 点钟 ===============
var anNing3 = [][]int{
	{745, 323, 108}, //坐标 + 耗时
	{730, 367, 108},
	//{829, 392, 116},
}

func killAnNingGolden3() {
	killMonster(anNing3[0], moveTakeTimePos3)

	killMonster(anNing3[1], moveTakeTimePos3)

}
