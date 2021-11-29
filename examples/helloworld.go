package main

import (
	"github.com/kaepa3/LiquidCrystal"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {

	adaptor := raspi.NewAdaptor()
	lcd := LiquidCrystal.NewLiquidCrystalDriver(adaptor,
		"LiquidCrystal",
		0x27,
		16,
		2)

	work := func() {
		lcd.Print("Hello World!")
	}

	robot := gobot.NewRobot("LiquidCrystal",
		[]gobot.Connection{adaptor},
		[]gobot.Device{lcd},
		work,
	)

	robot.Start()
}
