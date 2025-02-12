package main

func getTitle() string {
	return Title
}

const Title string = "KimiGaYo" // 楽曲名
var Song_BPM float64 = 45.0     // 楽曲のテンポ
var Repetitions int = 1         // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// 楽譜データ
// TITLE:国歌 「君が代」
// 'レドレミソミレミソラソラ<レ>シラソ
// T60@50D4C4D4E4G4E4D2E4G4A4G8A8<D4>B4A4G4
// 'ミソラ<レドレ>ミソラソミソレ
// E4G4A2<D4C4D2>E4G4A4G4E4.G8D2
// 'ラ<ドレドレ>ラソラソミレ
// A4<C4D2C4D4>A4G4A4G8E8D2
// 君が代は	千代に八千代に
// さざれ石の	巌（いわお）となりて
// 苔（こけ）のむすまで
var Notes = []Note{
	{R, L2}, // 何故か、最初にノイズが再生されるので、ここに、ダミーの休符を置く。
	{D4, L4},
	{C4, L4},
	{D4, L4},
	{E4, L4},
	{G4, L4},
	{E4, L4},
	{D4, L2},
	{E4, L4},
	{G4, L4},
	{A4, L4},
	{G4, L8},
	{A4, L8},
	{D5, L4},
	{B4, L4},
	{A4, L4},
	{G4, L4},
	{E4, L4},
	{G4, L4},
	{A4, L2},
	{D5, L4},
	{C5, L4},
	{D5, L2},
	{E4, L4},
	{G4, L4},
	{A4, L4},
	{G4, L4},
	{E4, L4 + L8},
	{G4, L8},
	{D4, L2},
	{A4, L4},
	{C5, L4},
	{D5, L2},
	{C5, L4},
	{D5, L4},
	{A4, L4},
	{G4, L4},
	{A4, L4},
	{G4, L8},
	{E4, L8},
	{D4, L2},
}
