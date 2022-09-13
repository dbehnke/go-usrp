package usrp

import (
	"encoding/binary"
	"log"
	"net"
	"time"
)

func Blah() int {
	return 1
}

func logEndTransmission(callsign string, startTime time.Time, loss string) {
	log.Printf("END TX: %s transmitted for %v with a %s BER", callsign, time.Since(startTime), loss)
}

func findCall(b []byte) string {
	if b[0] == 0 {
		return "UNKNOWN"
	}
	findZero := -1
	for i := 1; i < len(b); i++ {
		if b[i] == 0 {
			findZero = i
			break
		}
	}
	if findZero == -1 {
		return "UNKNOWN"
	}
	return string(b[0:findZero])
}

func RxUSRP(config *Config) {
	protocol := "udp"
	addr, err := net.ResolveUDPAddr(protocol, config.RXPort)
	if err != nil {
		log.Fatal("error ResolveUDPAddr:", err)
		return
	}
	//Create the connection
	udpConn, err := net.ListenUDP(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}

	var buf [1048]byte

	//receive loop
	var lastKeyup uint32
	var startTime time.Time
	var callsign string
	var loss string = "0.00%"
	var lastSeq uint32

	transmitEnable := true
	for {
		n, err := udpConn.Read(buf[0:])
		if err != nil {
			log.Fatal("Error Reading: ", err)
		}
		if n < 4 {
			log.Println("bad packet length:", n)
			continue
		}
		//log.Println(hex.EncodeToString(buf[0:n]))
		firstFour := string(buf[0:4])
		//log.Println(firstFour)
		if firstFour != "USRP" {
			log.Println("Not a USRP packet")
			continue
		}
		seq := binary.BigEndian.Uint32(buf[4:8])
		//memory
		keyup := binary.BigEndian.Uint32(buf[12:16])
		//talkgroup
		usrpType := binary.LittleEndian.Uint32(buf[20:24])
		//mpxid
		//reserved
		audio := buf[32:]
		//binary.BigEndian.Uint16(buf)
		switch usrpType {
		case USRP_TYPE_VOICE:
			//log.Println("voice")
			//TODO do something with the audio
			if keyup != lastKeyup {
				startTime = time.Now()
			}
			if keyup == 0 {
				logEndTransmission(callsign, startTime, loss)
				transmitEnable = true
			}
			lastKeyup = keyup
		case USRP_TYPE_TEXT:
			log.Println("USRP TEXT", int(audio[0]))
			if audio[0] == TLV_TAG_SET_INFO {
				if !transmitEnable {
					logEndTransmission(callsign, startTime, loss)
				}
				callsign = findCall(audio[14:50])
				log.Println("Begin TX:", callsign)
				transmitEnable = false
			}
		case USRP_TYPE_PING:
			if (lastSeq + 1) == seq {
				log.Println("missed EOT")
				logEndTransmission(callsign, startTime, loss)
			}
			lastSeq = seq
		default:
			log.Println("unhandled USRP type:", usrpType)
		}

	}
}
