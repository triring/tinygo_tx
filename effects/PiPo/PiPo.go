package main

import "fmt"

func Title() {
	fmt.Printf("PC-9801VM11 PiPo")
}

var Song_BPM float64 = 256.0 // 楽曲のテンポ
var Repetitions int = 12     // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// 往年の名機 PC-9801VM11の起動音を再現
// B6	120ms 6オクターブ	シ
// B5	120ms 5オクターブ	シ

var Notes = []Note{
	{R, L1},
	{B6, L4},
	{B5, L4},
	{R, L2},
}
