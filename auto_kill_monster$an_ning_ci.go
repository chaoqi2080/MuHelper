package main

// 黄金之间移动的时间(秒)
var moveTakeTimePos3 = 16

// =============== 安宁池 3 点钟 ===============
var anNing3 = [][]int{
	{3978, -194, 114}, //坐标 + 耗时
	{3994, -238, 116},
}

func killAnNingGolden3() {
	//再杀黄金毒树妖
	killMonster(anNing3[0], moveTakeTimePos3)
	//先杀树妖
	killMonster(anNing3[1], moveTakeTimePos3)

}
