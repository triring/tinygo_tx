// tinygo build -o xxxxx.uf2 -target=pico -size short ./xxxxx
// for RP2040

package main

import (
	"fmt"
	"machine"
	"time"
)

// 出力する搬送波の設定
var period uint64 = uint64(1000000000 / 999000)

/*
AM放送の周波数帯では、9kHz間隔で割り振られている。
これに適合する誤差の少ない周波数を生成できるPWMの設定は、以下の4つ
1000000000	/	693000	=	1443.001443
1000000000	/	999000	=	1001.001001
1000000000	/	1287000	=	777.000777
1000000000	/	1548000	=	645.994832
*/

func main() {
	/*
	   AM波を出力するRaspiPicoのGPIOを設定する。
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
	// pwmの出力をAM搬送波の周波数に設定する。
	pwm.Configure(machine.PWMConfig{Period: period})

	ch, err := pwm.Channel(pin)
	if err != nil {
		println(err.Error())
		return
	}

	//	for i := 0; i < Repetitions; i++ {
	var i int = 0
	for {
		Title() // 楽曲のタイトルを表示する。
		// 別ファイルで定義している楽譜データを読み込み、1音づつ読み出し、再生していく。
		for _, n := range Notes {
			fmt.Printf("%14.8f,%12.8f\n", n.Tone, n.Duration)
			// 音の高さと長さを算出する。
			tone := (1.0 / (2.0 * n.Tone)) * 1000000.0
			// 楽譜データで定義されているBPMから、tempoを算出する。
			tempo := ((60 / Song_BPM) * (n.Duration * 1000))

			// 休符の場合は、音の再生は行わず、指定時間スリープする。
			if n.Tone == R {
				pwm.Set(ch, pwm.Top()>>1)
				time.Sleep(time.Duration(tempo) * time.Millisecond)
				continue
			}
			for i := 0.0; i < tempo*1000; i += tone * 2.0 {
				pwm.Set(ch, pwm.Top()>>1)
				time.Sleep(time.Duration(tone) * time.Microsecond)
				pwm.Set(ch, 0)
				time.Sleep(time.Duration(tone) * time.Microsecond)
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
		Title() // 楽曲のタイトルを表示する。
	}
	pwm.Set(ch, 0)
	for {
		time.Sleep(1 * time.Hour)
	}

}
