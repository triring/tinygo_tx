package main

func getTitle() string {
	return Title
}

const Title string = "ramen" // 楽曲名
var Song_BPM float64 = 80.0  // 楽曲のテンポ
var Repetitions int = 3      // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// TITLE:ラーメン屋さん
// 屋台のラーメン屋さんが客寄せにチャルメラで吹いていた曲

// 楽譜データ
var Notes = []Note{
	{C5, L8d},  // do
	{D5, L16},  // re
	{E5, L4},   // mi
	{E5, L16d}, // mi
	{D5, L16},  // re
	{C5, L8},   // do
	{C5, L8},   // do
	{D5, L8},   // re
	{E5, L8},   // mi
	{D5, L8},   // re
	{C5, L8},   // do
	{D5, L1},   // re
}
