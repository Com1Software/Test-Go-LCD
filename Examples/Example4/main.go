package main

import (
	"fmt"
	"log"
	"time"

	device "github.com/Com1Software/go-hd44780"
	i2c "github.com/googolgl/go-i2c"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i2c, err := i2c.New(0x27,  "/dev/i2c-1")
	check(err)
	defer i2c.Close()
	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	check(err)
	lcd.BacklightOn()
	lcd.Clear()
	for {
		lcd.Home()
		t := time.Now()
		lcd.SetPosition(0, 3)
		fmt.Fprint(lcd, t.Format("Monday Jan 2"))
		lcd.SetPosition(1, 3)
		fmt.Fprint(lcd, t.Format("15:04:05 2006"))
		time.Sleep(100 * time.Millisecond)
	}
}
