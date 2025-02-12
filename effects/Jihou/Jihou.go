package main

func getTitle() string {
	return Title
}

const Title string = "Time signal" // 効果音名
var Song_BPM float64 = 120.0       // 楽曲のテンポ
var Repetitions int = 5            // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// TITLE:時報
// NHKラジオの時報の音階は、「プ」が440Hz、「ピーン」が880Hz
// これは、1オクターブ上の音程で、「プ」が440Hz、「ピーン」がその倍の880Hz
// この時報は、正時の3秒前から440Hzの予報音を3回、正時に880Hzの正報音を1回、正時の3秒後に正報終了という構成になっている

// 楽譜データ
var Notes = []Note{
	{A4, L8}, // 440Hz
	{R, L8},
	{R, L4},

	{A4, L8}, // 440Hz
	{R, L8},
	{R, L4},

	{A4, L8}, // 440Hz
	{R, L8},
	{R, L4},

	{A5, L4 * 3}, // 880Hz
	{R, L1},
}
