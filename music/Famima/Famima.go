package main

func getTitle() string {
	return Title
}

const Title string = "Family Mart entrance sound" // 楽曲名
var Song_BPM float64 = 84.0                       // 楽曲のテンポ
var Repetitions int = 3                           // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// TITLE:ファミマ入店音
// コンビニ入店時チャイム(お客様を歓迎しながら)
//正式な曲名は「メロディーチャイムNO.1 ニ長調 作品17「大盛況」」

// 楽譜データ
var Notes = []Note{
	{B4, L8},      // シ
	{G4, L8},      // ソ
	{D4, L8},      // レ
	{G4, L8},      // ソ
	{A4, L8},      // ラ
	{D5, L2},      // レ
	{R, L8},       // 休
	{D4, L8},      // レ
	{A4, L8},      // ラ
	{B4, L8},      // シ
	{A4, L8},      // ラ
	{D4, L8},      // レ
	{G4, L2 + L4}, // ソ
	{R, L1},       // 休
}

/*
var Notes = []Note{
	{G5, L8},
	{CS5, L8},
	{AS4, L8},
	{DS5, L8},
	{F5, L8},
	{AS5, L2},
	{R, L4},
	{F5, L8},
	{G5, L8},
	{F5, L8},
	{AS4, L8},
	{DS5, L1},
	{R, L1},
}
*/
