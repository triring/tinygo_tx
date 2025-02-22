package main

func getTitle() string {
	return Title
}

const Title string = "Che che koolay" // 楽曲名
var Song_BPM float64 = 144.0          // 楽曲のテンポ
var Repetitions int = 3               // 繰返しの回数,0と定義すると、無限ループになり、永久に演奏を繰り返す。

// 楽譜データ
// TITLE:Che che koolay (チェッチェッコリ)
// アフリカ、ガーナ民謡
// CheChe Koolay　(Singing Game from Ghana)
// (solo) Che che koolay, (echo) Che che koolay (Hands on head)
// (solo) Che che kofee sa, (echo) Che che kofee sa. (Hands on shoulders)
// (solo) Kofee salanga ,(echo) Kofee salanga. (Hands on waist)
// (solo) Kakashee langa, (echo) Kakashee langa. (Hands on knees)
// (solo) Koommadyeday, (echo) Koommadyeday. (Hands touch toes)

var Notes = []Note{
	{A5, L8},
	{R, L8},
	{A5, L8},
	{R, L8},
	{G5, L8},
	{A5, L8},
	{R, L4},

	{A5, L8},
	{R, L8},
	{G5, L8},
	{A5, L8},
	{R, L8},
	{E5, L8},
	{R, L4},

	{E5, L8},
	{G5, L8},
	{R, L8},
	{E5, L8},
	{A5, L4},
	{A5, L4},

	{R, L8},
	{G5, L8},
	{R, L8},
	{E5, L8},
	{A5, L4},
	{A5, L4},

	{A5, L4},
	{A5, L4},
	{G5, L8},
	{A5, L8},
	{R, L1},
}
