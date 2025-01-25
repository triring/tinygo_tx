package main

import "fmt"

func Title() {
	fmt.Printf("Do Re Mi Fa So Ra Si Do")
}

var Song_BPM float64 = 120.0 // 楽曲のテンポ
var Repetitions int = 5      // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// TITLE:ドレミファソラシド
// デモ用の楽譜データ
// 4オクターブのドから、1オクターブ高いドまでの音階を再生する

// 楽譜データ
var Notes = []Note{
	{C4, L4}, // ド
	{D4, L4}, // レ
	{E4, L4}, // ミ
	{F4, L4}, // ファ
	{G4, L4}, // ソ
	{A4, L4}, // ラ
	{B4, L4}, // シ
	{C5, L2}, // ド
	{R, L1},  // 休符
}
