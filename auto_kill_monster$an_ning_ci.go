package main

// 黄金之间移动的时间(秒)
var moveTakeTimePos3 = 16

// =============== 安宁池 3 点钟 ===============
var anNing3 = [][]int{
	{753, 408, 137},
	{775, 359, 137},
}

func killAnNingGolden3() {
	//再杀黄金毒树妖
	killMonster(anNing3[0], moveTakeTimePos3)
	//先杀树妖
	killMonster(anNing3[1], moveTakeTimePos3)

}
