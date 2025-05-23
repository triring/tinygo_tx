package main

func getTitle() string {
	return Title
}

const Title string = "Ode to Joy" // 楽曲名
var Song_BPM float64 = 120.0      // 楽曲のテンポ
var Repetitions int = 1           // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// 楽譜データ
// TITLE:Ode to Joy
// 作者 Ludwig van Beethoven
// 歓喜の歌:ベートーベン作曲交響曲第9番のフィナーレを飾る独唱および合唱の部分

var Notes = []Note{
	{E5, L4},
	{E5, L4},
	{F5, L4},
	{G5, L4},
	{G5, L4},
	{F5, L4},
	{E5, L4},
	{D5, L4},
	{C5, L4},
	{C5, L4},
	{D5, L4},
	{E5, L4},
	{E5, L4 + L8},
	{D5, L8},
	{D5, L4 + L8},
	{R, L8},
	{E5, L4},
	{E5, L4},
	{F5, L4},
	{G5, L4},
	{G5, L4},
	{F5, L4},
	{E5, L4},
	{D5, L4},
	{C5, L4},
	{C5, L4},
	{D5, L4},
	{E5, L4},
	{D5, L4 + L8},
	{C5, L8},
	{C5, L4 + L8},
	{R, L8},
	{D5, L2},
	{E5, L4},
	{C5, L4},
	{D5, L4},
	{E5, L8},
	{F5, L8},
	{E5, L4},
	{C5, L4},
	{D5, L4},
	{E5, L8},
	{F5, L8},
	{E5, L4},
	{D5, L4},
	{C5, L4},
	{D5, L4},
	{G4, L8},
	{R, L8},
	{E5, L4},
	{E5, L4},
	{E5, L4},
	{F5, L4},
	{G5, L4},
	{G5, L4},
	{F5, L4},
	{E5, L4},
	{D5, L4},
	{C5, L4},
	{C5, L4},
	{D5, L4},
	{E5, L4},
	{D5, L4 + L8},
	{C5, L8},
	{C5, L8},
	{G5, L8},
	{F5, L8},
	{E5, L8},

	{D5, L2},
	{E5, L4},
	{C5, L4},
	{D5, L4},
	{E5, L8},
	{F5, L8},
	{E5, L4},
	{C5, L4},
	{D5, L4},
	{E5, L8},
	{F5, L8},
	{E5, L4},
	{D5, L4},
	{C5, L4},
	{D5, L4},
	{G4, L8},
	{R, L8},
	{E5, L4},
	{E5, L4},

	{F5, L4},
	{F5, L4},
	{G5, L4},
	{G5, L4},
	{F5, L4},
	{E5, L4},
	{D5, L4},
	{C5, L4},
	{C5, L4},
	{D5, L4},
	{E5, L4},
	{D5, L4 + L8},
	{C5, L8},
	{C5, L2},
}
