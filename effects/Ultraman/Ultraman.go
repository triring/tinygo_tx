package main

func getTitle() string {
	return Title
}

const Title string = "Ultraman color timer" // 効果音名
var Song_BPM float64 = 600.0                // 楽曲のテンポ
var Repetitions int = 20                    // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// ピッポォーン,ピッポォーン,ピッポォーン
// ピッ音	高音シ 5B 988Hz	0.085sec
// ぽー音	低音ミ 5E 659Hz	0.100sec+0.250sec音量が半分程度の残響
// tempo 600
// 4分音符	400ms
// 8分音符	200ms
// 16分音符	100ms
// 32分音符	50ms

var Notes = []Note{
	{B5, L4}, // 75ms:ピッ音	高音シ 5B 988Hz
	{R, L8},
	{E5, L1 + L4}, // 350ms:ぽー音	低音ミ 5E 659Hz
	{R, L8},
}
