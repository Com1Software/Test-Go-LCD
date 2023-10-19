package main

import (
	"fmt"
	"time"

	"github.com/mskrha/rpi-lcd"
)

func main() {
	l := lcd.New(lcd.LCD{Bus: "/dev/i2c-1", Address: 0x27, Rows: 4, Cols: 20, Backlight: true})
	start := time.Now()
	if err := l.Init(); err != nil {
		panic(err)
	}

	for {
		elapsed := fmt.Sprint(time.Since(start))		
		if err := l.Print(1, 2, time.Now().Format("Monday Jan 2")); err != nil {
			fmt.Println(err)
			return
		}
		if err := l.Print(2, 2, time.Now().Format("15:04:05 2006")); err != nil {
			fmt.Println(err)
			return
		}

		if err := l.Print(3, 2, "Elapsed Time"); err != nil {
			fmt.Println(err)
			return
		}
		if err := l.Print(4, 2, elapsed); err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}
