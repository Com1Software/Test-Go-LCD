package main

import (
	"log"

	device "github.com/d2r2/go-hd44780"
	"github.com/d2r2/go-i2c"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i2c, err := i2c.NewI2C(0x27, 1)
	checkError(err)
	defer i2c.Close()
	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	checkError(err)
	err = lcd.BacklightOn()
	checkError(err)
	err = lcd.ShowMessage("Test Line 1", device.SHOW_LINE_1)
	checkError(err)
	err = lcd.ShowMessage("Test line 2", device.SHOW_LINE_2)
	checkError(err)

}
