package main

func getTitle() string {
	return Title
}

const Title string = "Close Encounters of the Third Kind" // 効果音名
var Song_BPM float64 = 80.0                               // 楽曲のテンポ
var Repetitions int = 12                                  // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// Close Encounters of the Third Kind（「第三種接近遭遇」の意）
// 映画「未知との遭遇」
// 未知の飛行物体との交信音
// |オクターブ|音階|音長  |音名 |
// |:--------:|:--:|:-----|:----|
// |   Oct5   |  2 |==    |レ   |
// |   Oct5   |  3 |==    |ミ   |
// |   Oct5   |  1 |==    |ド   |
// |   Oct4   |  1 |==    |ド   |
// |   Oct4   |  5 |======|ソ   |

var Notes = []Note{
	{D5, L8},
	{E5, L8},
	{C5, L8},
	{C4, L8},
	{G4, L4d},
	{R, L2},
}
