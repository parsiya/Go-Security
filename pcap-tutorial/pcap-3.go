package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"golang.org/x/net/icmp"
)

var (
	pcapFile string = "conv.pcap"
	handle   *pcap.Handle
	err      error
)

// ICMPv4 == iana.ProtocolICMP
// ICMPv6 == iana.ProtocolIPv6ICMP
// https://go.googlesource.com/net/+/master/icmp/message_test.go
// Because this iana an internal package, we cannot use it so we copy the const
// and use the ones we want.
// https://github.com/golang/net/blob/master/internal/iana/const.go#L42

const (
	ProtocolICMP     = 1  // Internet Control Message
	ProtocolIPv6ICMP = 58 // ICMP for IPv6
)

func main() {
	// Open file instead of device
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "icmp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Filter set to ICMP.")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Do something with a packet here.
		// fmt.Println(packet)

		// ICMP is IP layer so we do not care about other layers
		// Note the difference between IPv4 and IPv6 here, it might come and
		// bite us in the backside later
		ipLayer := packet.Layer(layers.LayerTypeIPv4)

		// fmt.Printf("%+v\n", ipLayer)
		// fmt.Printf("%T\n", ipLayer)

		// Can't do this
		// payload1 := ipLayer.Payload
		// fmt.Println(payload1)

		if ipLayer != nil {
			// Cast it to IPv4 layer
			ip, _ := ipLayer.(*layers.IPv4)

			// fmt.Printf("%+v\n", ip)
			// fmt.Println(ip.Payload)
			// fmt.Println(len(ip.Payload))
			// fmt.Println(string(ip.Payload))

			// Get an ICMP message from raw bytes in payload
			msg, err := icmp.ParseMessage(ProtocolICMP, ip.Payload)
			if err != nil {
				panic(err)
			}

			// fmt.Printf("%+v\n", msg)
			// fmt.Printf("%T\n", msg)
			// fmt.Println()

			// fmt.Printf("%v\n", msg.Type)
			// fmt.Printf("%v\n", msg.Code)
			// fmt.Printf("%v\n", msg.Checksum)
			// fmt.Printf("%+v\n", msg.Body)

			// To use Body we need to cast it to *icmp.Echo
			// or body := message.Body.(*icmp.Echo)
			if body, err := msg.Body.(*icmp.Echo); err {
				// Now we can access Body.Data
				fmt.Println(string(body.Data))
			}
		}
		// This return will process one packet and is useful for debugging
		// return
	}
}
