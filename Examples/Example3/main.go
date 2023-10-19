package main

import (
	"fmt"
	"time"

	"github.com/mskrha/rpi-lcd"
)

func main() {
	l := lcd.New(lcd.LCD{Bus: "/dev/i2c-1", Address: 0x27, Rows: 4, Cols: 20, Backlight: true})

	if err := l.Init(); err != nil {
		panic(err)
	}

	for {
		if err := l.Print(2, 0, time.Now().Format("02.01.")); err != nil {
			fmt.Println(err)
			return
		}
		if err := l.Print(1, 8, time.Now().Format("15:04:05")); err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}
