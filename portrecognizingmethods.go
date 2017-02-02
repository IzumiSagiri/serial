package serial

import "fmt"
import "time"

func RecognizeAsahikaseiAndSCRPorts(allPorts []Port, asahikaseiPorts, SCRPorts []Port) ([]Port, []Port) {
	var leftPorts, leftLeftPorts []Port
	asahikaseiPorts, leftPorts = RecognizePortsbyaMethod(allPorts, asahikaseiPorts, leftPorts, checkIfPortIsAsahikasei)
	SCRPorts, _ = RecognizePortsbyaMethod(leftPorts, SCRPorts, leftLeftPorts, checkIfPortIsSCR)
	return asahikaseiPorts, SCRPorts
}

func RecognizePortsbyaMethod(bufferIn []Port, bufferRecognized, bufferLeft []Port, methodUsed func(Port) bool) ([]Port, []Port) {
	for _, v := range bufferIn {
		if methodUsed(v) {
			bufferRecognized = append(bufferRecognized, v)
		} else {
			bufferLeft = append(bufferLeft, v)
		}
	}
	return bufferRecognized, bufferLeft
}

func checkIfPortIsAsahikasei(port Port) bool {
	fmt.Println("Check Asahikasei", port.Name)
	c := make(chan bool)
	go func() {
		for i := 0; i <= 100; i++ {
			temptemp := port.Read()
			fmt.Print(temptemp)
			if temptemp == '\r' {
				if temptemptemp := port.Read(); temptemptemp == '\n' {
					c <- true
				} else {
					c <- false
				}
			}
		}
	}()
	select {
	case temp := <-c:
		return temp
	case <-time.After(500 * time.Millisecond):
		return false
	}
	return false
}

func checkIfPortIsSCR(port Port) bool {
	fmt.Println("Check SCR", port.Name)
	var receivedBuffer []byte
	c := make(chan byte)
	for i := 0; i < 5; i++ {
		go func() { c <- port.Read() }()
		select {
		case temptemp := <-c:
			fmt.Print(temptemp)
			receivedBuffer = append(receivedBuffer, temptemp)
		case <-time.After(500 * time.Millisecond):
			return false
		}
	}
	var tempBool bool
	tempBool = true
	for i := 0; i < 4; i++ {
		temp := int(receivedBuffer[i+1]) - int(receivedBuffer[i])
		if temp < -50 && temp > 50 {
			tempBool = false
		}
	}
	if receivedBuffer[0] == receivedBuffer[1] && receivedBuffer[0] == receivedBuffer[2] && receivedBuffer[0] == receivedBuffer[3] && receivedBuffer[0] == receivedBuffer[4] && receivedBuffer[0] != 0 {
		tempBool = false
	}
	return tempBool
}

func CheckIfPortIsTheSame(port1, port2 Port) bool {
	var temp1 [5]byte
	var temp2 [5]byte
	for i := 0; i < 5; i++ {
		temp1[i] = port1.Read()
		temp2[i] = port2.Read()
	}
	tempBool := true
	for i := 0; i < 5; i++ {
		if temp1[i] != temp2[i] {
			tempBool = false
		}
	}
	return tempBool
}
