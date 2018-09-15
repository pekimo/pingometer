package main

import (
	"fmt"
	"log"

	"github.com/pekimo/pingometer/trace"
)

func main() {
	traceTime, err := trace.Trace("https://yandex.ru")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DNS time: ", traceTime.GetDnsTime())
	fmt.Println("TLS time: ", traceTime.GetTLSHandshakeTime())
	fmt.Println("Connection time: ", traceTime.GetConnectionTime())
	fmt.Println("TTFB: ", traceTime.GetTTFBTime())
	fmt.Println("Total time: ", traceTime.GetTotalTime())
}
