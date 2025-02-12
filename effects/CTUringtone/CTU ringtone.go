package main

func getTitle() string {
	return Title
}

const Title string = "24 TWENTY-FOUR CTU ringtone" // 効果音名
var Song_BPM float64 = 95                          // 楽曲のテンポ

var Repetitions int = 10 // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// TITLE:24 TWENTY-FOUR CTU ringtone
// アメリカ合衆国の連続ドラマシリーズ。『24 -TWENTY FOUR-』
// この中で登場した架空の政府機関「テロ対策ユニット」（CTU）内で使われていた着信音

// 楽譜データ
var Notes = []Note{
	{F6, L64},
	{F6, L64},
	{R, L16},
	{F6, L64},
	{F6, L64},
	{R, L8},
	{R, L64},
	{B6, L64},
	{B6, L64},
	{B6, L64},
	{B6, L64},
	{R, L64},
	{B5, L64},
	{B5, L64},
	{B5, L64},
	{B5, L64},
	{B5, L64},
	{B5, L64},
	{B5, L64},
	{B5, L64},
	{R, L2},
}
