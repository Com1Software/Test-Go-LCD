package main

import (
	"fmt"
	"strings"
	"time"

	lcd "github.com/mskrha/rpi-lcd"
)

func main() {
	l := lcd.New(lcd.LCD{Bus: "/dev/i2c-1", Address: 0x27, Rows: 2, Cols: 16, Backlight: true})
	start := time.Now()
	if err := l.Init(); err != nil {
		panic(err)
	}

	for {
		elapsed := fmt.Sprint(time.Since(start))
		e := strings.Split(elapsed, ".")
		if err := l.Print(1, 1, time.Now().Format("15:04:05 2006")); err != nil {
			fmt.Println(err)
			return
		}
		if err := l.Print(2, 1, "Elapsed "+e[0]+"s"); err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}
