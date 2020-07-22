package main

import (
	"boss2_ledscreen/clienttcpa"
	"boss2_ledscreen/clienttcpb"
	"boss2_ledscreen/config"
	"boss2_ledscreen/web"
)

func main () {
	//start.StartProgram()
	go web.QuerySpaces( config.Syscfg.CtInfoA.CounterName, web.DataInchanA, config.Syscfg.CtInfoA.Timeinterval)

	if config.Syscfg.TurnLEDControler {

		go web.QuerySpaces(config.Syscfg.CtInfoB.CounterName, web.DataInchanB, config.Syscfg.CtInfoB.Timeinterval)
		go clienttcpb.TcpClinet()
	}

	clienttcpa.TcpClinet()

}