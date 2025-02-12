package main

func getTitle() string {
	return Title
}

const Title string = "Thirori Sound" // 効果音名
var Song_BPM float64 = 125.0         // 楽曲のテンポ
var Repetitions int = 16             // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// 効果音データ
// ティロリ!!ティロリ!!ティロリ!!
// 某大手ハンバーガーチェーンで、ポテトが揚がったときに店内で流すタイマー音
// https://www.youtube.com/watch?v=8NjhfastLts

var Notes = []Note{
	{G6, L8}, // So
	{E6, L8}, // Mi
	{G6, L8}, // So
	{R, L8},  // 休符
}
