package main

func getTitle() string {
	return Title
}

const Title string = "Westminster Chimes" // 効果音名
var Song_BPM float64 = 60.0               // 楽曲のテンポ
var Repetitions int = 3                   // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// ウェストミンスターの鐘（ウェストミンスターのかね、英語：Westminster Quarters）
// ウェストミンスター宮殿の時計塔ビッグ・ベンで使われている時鐘のメロディ
// 日本では学校のチャイムとして使われている

var Notes = []Note{
	{F4, L4}, // 	ファ
	{A4, L4}, // 	ラ
	{G4, L4}, // 	ソ
	{C4, L1}, // 	ド
	{R, L8},  // 	休符

	{F4, L4}, // 	ファ
	{G4, L4}, // 	ソ
	{A4, L4}, // 	ラ
	{F4, L1}, // 	ファ
	{R, L8},  // 	休符

	{A4, L4}, // 	ラ
	{F4, L4}, // 	ファ
	{G4, L4}, // 	ソ
	{C4, L1}, // 	ド
	{R, L8},  // 	休符

	{C4, L4}, // 	ド
	{G4, L4}, // 	ソ
	{A4, L4}, // 	ラ
	{F4, L1}, // 	ファ
	{R, L8},  // 	休符
	{R, L1},  // 	休符
}
