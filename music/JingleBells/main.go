// tinygo build -o xxxxx.uf2 -target=pico -size short ./xxxxx
//  tinygo flash -target=pico -size short .
// for RP2040

package main

import (
	"fmt"
	"machine"
	"time"
)

// AM放送の周波数帯(526.5～1606.5KHz)では、9KHz間隔で割り振られている。
// これに適合する誤差の少ない周波数を生成できるPWMの設定は、以下の通り。
// これ以外の周波数を使用する場合は、周波数の誤差は大きくなるが、以下のように計算式で定義する。
// var period uint64 = uint64(1000000000 / 1494000)

const AM0612KHz uint64 = 1634 // 1000000000	/	1634	=	 611995.104
const AM0684KHz uint64 = 1462 // 1000000000	/	1462	=	 683994.528
const AM0693KHz uint64 = 1443 // 1000000000	/	1443	=	 693000.693
const AM0738KHz uint64 = 1355 // 1000000000	/	1355	=	 738007.380
const AM0774KHz uint64 = 1292 // 1000000000	/	1292	=	 773993.808
const AM0819KHz uint64 = 1221 // 1000000000	/	1221	=	 819000.819
const AM0999KHz uint64 = 1001 // 1000000000	/	1001	=	 999000.999
const AM1224KHz uint64 = 817  // 1000000000	/	 817	=	1223990.208
const AM1269KHz uint64 = 788  // 1000000000	/	 788	=	1269035.533
const AM1287KHz uint64 = 777  // 1000000000	/	 777	=	1287001.287
const AM1368KHz uint64 = 731  // 1000000000	/	 731	=	1367989.056
const AM1548KHz uint64 = 646  // 1000000000	/	 646	=	1547987.616

// 日本のFM放送の周波数帯は、76.1IntermittentWave～94.9MHzである。
// 世界のFM放送の周波数帯は、87.5～108.0 MHz)が一般的である。
// これらの周波数帯の波長をPWMで生成するのは難しい。
// そこで、逓倍波で、これらに適合する周波数となるPWMの設定を探した。
// ただ、計算した周波数とかなりズレた周波数で受信できた。

const FM04_00MHz uint64 = 250 // 1000000000	/ 	250	=	 4000000.000
// 76.70 MHz:26 dbUv, 80.60 MHz:38 dbUv,  84.60 MHz:26 dbUv,. 88.70 MHz:34 dbUv,
// 92.80 MHz:26 dbUv, 96.80 MHz:30 dbUv, 100.80 MHz:23 dbUv, 104.80 MHz:25 dbUv)

const FM05_20MHz uint64 = 192 // 1000000000	/ 	192	=	 5208333.333
// (78.20 MHz:26 dbUv, 83.40 MHz:30 dbUv, 88.60 MHz:29 dbUv, 93.80 MHz:31 dbUv,
//  99.00 MHz:30 dbUv, 104.10 MHz:24 dbUv)

const FM10_00MHz uint64 = 100 // 1000000000	/ 	100	=	 10000000.000
// 	(83.40 MHz:33 dbUv, 93.80 MHz:49 dbUv, 104.10 MHz:27 dbUv)

const FM20_00MHz uint64 = 50 // 1000000000	/	50	=	 20000000.000	(83.40 MHz:37 dbUv, 104.10 MHz:33 dbUv)
const FM25_00MHz uint64 = 40 // 1000000000	/	40	=	 25000000.000	(80.10 MHz:20 dbUv, 100.00 MHz:35 dbUv)
const FM31_25MHz uint64 = 32 // 1000000000	/	32	=	 31250000.000	(93.80 MHz:40 dbUv)
const FM41_67MHz uint64 = 24 // 1000000000	/	24	=	 41666666.667	(83.40 MHz:38 dbUv)

// 出力する搬送波の設定
// 上記の搬送波定義の中から、出力する周波数を選んで以下のperiodに代入する。
// AM波の場合は、搬送波と同じ周波数
// var period uint64 = AM0999KHz

// FM波の場合は、設定した周波数の逓倍波を受信するので、定義の後に記載された逓倍波の周波数をみて選択すること。
// FM帯の高い周波数では、PWMの設定値の制約から、断続時間が長くなり、激しいノイズが交じる。
// 10MHz前後の設定を推奨する。

var period uint64 = FM10_00MHz // FM41_67MHz 

// もし、上記の定義以外の周波数に設定する場合は、
// 以下のように、1000000000を設定する周波数(単位はHz)で割った値を定義すること。
// var period uint64 = uint64(1000000000 / 999000)

func main() {
	/*
	   搬送波を出力するRaspiPicoのGPIOを設定する。
	   RP2040では8つのスライス、最大16チャンネルのPWMが扱える。
	   今回は、GPIO15をアンテナ出力に設定した。
	   また、このGPIOに対応するチャンネルとしてPWM7を設定した。
	   この端子に導線を接続し、送信アンテナとする。
	*/
	pin := machine.GPIO15
	pwm := machine.PWM7
	/* 他のGPIOを使用する場合は、以下の表から、使用するGPIOに対応するPWM チャンネルを設定すること。

	   GPIO		0	1	2	3	4	5	6	7	8	9	10	11	12	13	14	15
	   PWM	Ch	0A	0B	1A	1B	2A	2B	3A	3B	4A	4B	5A	5B	6A	6B	7A	7B

	   GPIO		16	17	18	19	20	21	22	23	24	25	26	27	28	29
	   PWM	Ch	0A	0B	1A	1B	2A	2B	3A	3B	4A	4B	5A	5B	6A	6B
	*/
	// Configure the PWM with the given period.
	// pwmの出力をAM or FM搬送波の周波数に設定する。
	pwm.Configure(machine.PWMConfig{Period: period})

	ch, err := pwm.Channel(pin)
	if err != nil {
		println(err.Error())
		return
	}

	var CarrierWave uint32 = uint32(pwm.Top() / 2)       // 搬送波の設定
	var IntermittentWave uint32 = uint32(pwm.Top() / 10) // 断続波の設定

	var i int = 0
	for {
		fmt.Println(getTitle()) // 楽曲のタイトルを表示する。
		// 別ファイルで定義している楽譜データを読み込み、1音づつ読み出し、再生していく。
		for _, n := range Notes {
			fmt.Printf("%14.8f,%12.8f\n", n.Tone, n.Duration)
			// 音の高さと長さを算出する。
			tone := (1.0 / (2.0 * n.Tone)) * 1000000.0
			// 楽譜データで定義されているBPMから、tempoを算出する。
			tempo := ((60 / Song_BPM) * (n.Duration * 1000))
			//	Low  := pwm.Set(ch, pwm.Top()>>4) // 音にノイズが交じる。ゼロに設定した時と同じ
			// 休符の場合は、音の再生は行わず、指定時間スリープする。
			if n.Tone == R {
				pwm.Set(ch, CarrierWave)
				time.Sleep(time.Duration(tempo) * time.Millisecond)
				continue
			}
			DUR := time.Duration(tone) * time.Microsecond
			for i := 0.0; i < tempo*1000; i += tone * 2.0 {
				pwm.Set(ch, CarrierWave)
				time.Sleep(DUR)
				pwm.Set(ch, IntermittentWave)
				time.Sleep(DUR)
			}
		}
		// 楽譜データで停止されているRepetitionsが0の場合は、無限ループ
		if 0 == Repetitions {
			continue
		}
		// 楽譜データで停止されているRepetitionsに設定された回数だけ楽曲を再生する。
		i++
		if i >= Repetitions {
			break
		}
		fmt.Println(getTitle()) // 楽曲のタイトルを表示する。
		// 以下はデバッグ用
		fmt.Println("period=", period)
		fmt.Println("pwm.Top()=", pwm.Top())
		fmt.Println("CarrierWave=", CarrierWave)
		fmt.Println("IntermittentWave=", IntermittentWave)
	}
	//	pwm.Set(ch, 0) 完全に出力を止めると、ラジオから未受信の状態でノイズが発生する。
	pwm.Set(ch, CarrierWave) // これを防止するために、搬送波(変調していない電波)だけを出し続ける。
	for {
		time.Sleep(1 * time.Hour)
	}
}
