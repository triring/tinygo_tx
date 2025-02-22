package main

func getTitle() string {
	return Title
}

const Title string = "kingyo" // 楽曲名
var Song_BPM float64 = 80.0   // 楽曲のテンポ
var Repetitions int = 3       // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// TITLE:金魚屋さん
// 物売りの売り声	金魚屋さん
// きんぎょ～、え～、きんぎょ～

// 楽譜データ
var Notes = []Note{
	{E5, L8d}, // mi	き
	{G5, L16}, // so	ん
	{G5, L2},  // so	ぎょ
	{E5, L4},  // mi	え
	{E5, L8},  // mi	ー
	{G5, L16}, // so	き
	{A5, L16}, // la	ん
	{A5, L1},  // la	ぎょ
}
