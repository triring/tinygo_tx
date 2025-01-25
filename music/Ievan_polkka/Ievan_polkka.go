package main

import "fmt"

func Title() {
	fmt.Printf("Ievan Polkka")
}

var Song_BPM float64 = 180.0 // 楽曲のテンポ
var Repetitions int = 3      // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// 楽譜データ
// TITLE:Ievan Polkka (イエヴァン・ポルッカ)ネット上で高い知名度を持つフィンランド民謡。
// 日本語では，イエヴァのポルカ。ボーカロイド 初音ミクでよく使われる曲
var Notes = []Note{
	{E4, L4},
	{A4, L4},
	{A4, L4},
	{A4, L8},
	{B4, L8},
	{C5, L4},
	{A4, L4},
	{A4, L4},
	{C5, L4},

	{B4, L4},
	{G4, L4},
	{G4, L4},
	{B4, L4},
	{C5, L4},
	{A4, L4},
	{A4, L2},

	{E4, L4},
	{A4, L4},
	{A4, L4},
	{A4, L8},
	{B4, L8},
	{C5, L4},
	{A4, L4},
	{A4, L4},
	{C5, L4},

	{E5, L4},
	{D5, L4},
	{C5, L4},
	{B4, L4},
	{B4, L4},
	{A4, L4},
	{A4, L4},
	{C5, L4},

	{E5, L4},
	{E5, L4},
	{D5, L4},
	{C5, L4},
	{B4, L4},
	{G4, L4},
	{G4, L4},
	{B4, L4},

	{D5, L4},
	{D5, L4},
	{C5, L4},
	{B4, L4},
	{B4, L8},
	{A4, L8},
	{A4, L4},
	{A4, L4},
	{C5, L4},

	{E5, L4},
	{E5, L4},
	{D5, L4},
	{C5, L4},
	{B4, L4},
	{G4, L4},
	{G4, L4},
	{B4, L8},
	{D5, L8},

	{D5, L4},
	{D5, L4},
	{C5, L4},
	{B4, L4},
	{C5, L4},
	{A4, L4},
	{A4, L2},
}
