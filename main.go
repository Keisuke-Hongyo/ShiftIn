package main

import (
	"ShiftIn/ShiftIn"
	"fmt"
	"machine"
	"time"
)

func main() {
	var data uint8

	/* 構造体の宣言 */
	shft := ShiftIn.New(machine.D7, machine.D8, machine.D9)

	for {
		/* シフトレジよりデータを取得 */
		data = shft.GetData()
		fmt.Println(data)
		time.Sleep(time.Millisecond * 200)
	}
}
