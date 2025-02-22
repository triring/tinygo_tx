package main

func getTitle() string {
	return Title
}

const Title string = "Do Re Mi Fa So Ra Si Do" // 楽曲名
var Song_BPM float64 = 120.0                   // 楽曲のテンポ
var Repetitions int = 5                        // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// TITLE:ドレミファソラシド
// デモ用の楽譜データ
// 4オクターブのドから、1オクターブ高いドまでの音階を再生する

// 楽譜データ
var Notes = []Note{
	{C4, L4}, // ド,　4分音符
	{D4, L4}, // レ,　4分音符
	{E4, L4}, // ミ,　4分音符
	{F4, L4}, // ファ,4分音符
	{G4, L4}, // ソ,　4分音符
	{A4, L4}, // ラ,　4分音符
	{B4, L4}, // シ,　4分音符
	{C5, L2}, // ド,　2分音符
	{R, L1},  // 全休符
}
