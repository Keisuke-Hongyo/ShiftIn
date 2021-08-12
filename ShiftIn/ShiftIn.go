package ShiftIn

import (
	"machine"
)

type ShiftIn struct {
	data machine.Pin // シフトレジスタデータ入力ピン
	clk  machine.Pin // シフトレジスタクロックピン
	sh   machine.Pin // データロード制御
}

/* シフトレジスタ初期化 */
func New(data machine.Pin, clk machine.Pin, sh machine.Pin) ShiftIn {

	/* 構造体の宣言 */
	shft := ShiftIn{}

	/* 構造体に格納 */
	shft.data = data
	shft.clk = clk
	shft.sh = sh

	/* 各ピンの設定 */
	shft.data.Configure(machine.PinConfig{Mode: machine.PinInput})
	shft.clk.Configure(machine.PinConfig{Mode: machine.PinOutput})
	shft.sh.Configure(machine.PinConfig{Mode: machine.PinOutput})

	/* 初期状態の設定 */
	sh.High()
	clk.Low()

	return shft
}

/* データ取得関数 */
func (shft *ShiftIn) GetData() uint8 {

	/* 格納変数の宣言 */
	var inData uint8

	/* 無名関数の宣言 */
	setBit := func(in bool) uint8 {
		if shft.data.Get() {
			return 0x01
		} else {
			return 0x00
		}
	}

	// 格納変数クリア
	inData = 0x00

	// MSBを取得
	shft.sh.Low()
	shft.sh.High()
	inData |= setBit(shft.data.Get())

	/* 残りのビットを取得 */
	for i := 1; i < 8; i++ {
		inData <<= 1
		shft.clk.High()
		inData |= setBit(shft.data.Get())
		shft.clk.Low()
	}

	/* データを返す */
	return inData
}
