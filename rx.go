package usrp

import (
	"encoding/binary"
	"log"
	"net"
	"strings"
)

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
	return strings.Trim(string(b[0:findZero]), " ")
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
	//var lastKeyup uint32
	var callsign string
	var lastSeq uint32

	var t Transmission // holds the transmission stuff while active
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
		audio := buf[32:n]
		//binary.BigEndian.Uint16(buf)
		switch usrpType {
		case USRP_TYPE_VOICE:
			//log.Println("voice")
			t.Audio.Write(audio) //append audio to the transmission buffer
			//if keyup != lastKeyup {
			// do nothing for right now
			//}
			if keyup == 0 {
				t.EndTransmission()
			}
			//lastKeyup = keyup
		case USRP_TYPE_TEXT:
			//log.Println("USRP TEXT", int(audio[0]))
			if audio[0] == TLV_TAG_SET_INFO {
				// disabled to see if this matters
				//if !transmitEnable {
				//	t.EndTransmission()
				//}
				callsign = findCall(audio[14:50])
				t = NewTransmission(config.Group, callsign) //create a new transmission
				//transmitEnable = false
			}
		case USRP_TYPE_PING:
			if (lastSeq + 1) == seq {
				log.Println("missed EOT")
				t.EndTransmission()
			}
			lastSeq = seq
		default:
			log.Println("unhandled USRP type:", usrpType)
		}

	}
}
