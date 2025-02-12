package main

func getTitle() string {
	return Title
}

const Title string = "Welcome to Jafari Park" // 楽曲名
var Song_BPM float64 = 170.0                  // 楽曲のテンポ
var Repetitions int = 3                       // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// 楽譜データ
// TITLE:ようこそジャパリパークへ

var Notes = []Note{
	{G4, L8},
	{F4, L8},
	{G4, L8},
	{A4, L4},
	{G4, L8},
	{A4, L8},
	{B4, L4},

	{C5, L8},
	{CS5, L8},
	{D5, L4},
	{B4, L8},
	{A4, L8},
	{G4, L8},

	{G4, L4},
	{E5, L4},
	{D5, L4},
	{G4, L4},

	{E4, L4},
	{C5, L4},
	{B4, L4},
	{A4, L4},

	{G4, L4},
	{R, L4},
	{R, L2},
}
