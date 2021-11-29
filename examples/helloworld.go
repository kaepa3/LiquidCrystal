package main

import (
	"github.com/Kaepa3/LiquidCrystal"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/gobot/platforms/raspi"
)

func main() {

	adaptor := raspi.NewFirmataAdaptor("firmata", "/dev/ttyACM0")
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
