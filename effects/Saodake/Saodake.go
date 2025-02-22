package main

func getTitle() string {
	return Title
}

const Title string = "saodake" // 楽曲名
var Song_BPM float64 = 80.0    // 楽曲のテンポ
var Repetitions int = 3        // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// TITLE:竿竹屋さん
// 物売りの売り声	竿竹屋さん
// たけや～、さおだけっ

// 楽譜データ
var Notes = []Note{
	{E5, L8d}, // mi	た
	{G5, L16}, // so	け
	{G5, L2d}, // so	や
	{G5, L8},  // so	さ
	{A5, L8},  // ra	お
	{A5, L8},  // ra	だ
	{A5, L1},  // ra	け
}
