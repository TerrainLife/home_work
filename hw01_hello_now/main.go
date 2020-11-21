package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	fmt.Println("current time: ", time.Now().Round(time.Second))

	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Can`t get ntp time!")
	}
	fmt.Println("exact time: ", ntpTime.Round(time.Second))

}
