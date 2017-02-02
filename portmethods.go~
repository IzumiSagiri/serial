package serialportmethods

/*
#cgo LDFLAGS: -lwinmm
#include <stdlib.h>
#include <stdio.h>
#include <Windows.h>
#pragma comment(lib, "winmm.lib")
#pragma comment(lib, "kernel32.lib")

uintptr_t _openAPort(char commName[])
{
	return (uintptr_t)CreateFile(commName, GENERIC_READ, 0, NULL, OPEN_EXISTING, 0, NULL);
}

BOOL closeAPort(uintptr_t handle)
{
    return CloseHandle((HANDLE)handle);
}

BOOL setCommState(uintptr_t handle, unsigned int baudrate)
{
	DCB dcb;
	BOOL fSuccess;
	SecureZeroMemory(&dcb, sizeof(DCB));
	dcb.DCBlength = sizeof(DCB);
    switch (baudrate) {
    case 19200:
        dcb.BaudRate = CBR_19200;
        break;
    case 115200:
        dcb.BaudRate = CBR_115200;
        break;
    default:
        printf("Error");
    }
	dcb.ByteSize = 8;
	dcb.Parity = NOPARITY;
	dcb.StopBits = ONESTOPBIT;

	fSuccess = SetCommState((HANDLE)handle, &dcb);
	return fSuccess;
}

BOOL setCommTimeouts(uintptr_t handle)
{
    COMMTIMEOUTS Cptimeouts;
    BOOL fSuccess;
    Cptimeouts.ReadIntervalTimeout = 100;
    Cptimeouts.ReadTotalTimeoutMultiplier = 100;
    Cptimeouts.ReadTotalTimeoutConstant = 100;
    Cptimeouts.WriteTotalTimeoutMultiplier = 100;
    Cptimeouts.WriteTotalTimeoutConstant = 100;

    fSuccess = SetCommTimeouts((HANDLE)handle, &Cptimeouts);
    return fSuccess;
}

unsigned char readPort(uintptr_t handle)
{
	unsigned char buf[1];
	int n=0;
	ReadFile((HANDLE)handle, buf, 1, (LPDWORD)((void *)&n), NULL);
	return(buf[0]);
}
*/
import "C"

import "fmt"
import "strconv"
import "time"

type Port struct {
	Name     string
	Baudrate uint32
	Handle   C.uintptr_t
}

func (port *Port) Open() bool {
	time.Sleep(time.Second)
	port.Handle = C._openAPort(C.CString(fmt.Sprint("//./", port.Name)))
	time.Sleep(time.Second)
	fSuccess := C.setCommState(port.Handle, C.uint(port.Baudrate))
	if fSuccess == 0 {
		return false
	} else {
		fSuccess2 := C.setCommTimeouts(port.Handle)
		if fSuccess2 == 0 {
			return false
		} else {
			return true
		}
	}
	return true
}

func (port Port) Read() byte {
	temp := byte(C.readPort(port.Handle))
	return temp
}

func (port Port) Close() {
	temp := C.closeAPort(port.Handle)
	if temp == 0 {
		fmt.Println("Close Failure.")
	}
}

func (port Port) ReadATwoAsciiData() (int, bool) {
	s, err := strconv.ParseInt(string(port.Read())+string(port.Read()), 16, 32)
	if err != nil {
		return 0, false
	}
	return int(s), true
}

func (port Port) GetSCR() byte {
	return port.Read()
}

func (port Port) GetAsahikasei() (byte, byte, byte, byte, int, int, int, int, int) {
	var fSuccess int
	for port.Read() != '\r' {
	}
	if port.Read() != '\n' {
		fSuccess = 1
	}
	ead1_h, flag1 := port.ReadATwoAsciiData()
	ead1_l, flag2 := port.ReadATwoAsciiData()
	ead2_h, flag3 := port.ReadATwoAsciiData()
	ead2_l, flag4 := port.ReadATwoAsciiData()
	if !flag1 || !flag2 || !flag3 || !flag4 {
		fSuccess = 2
	}
	for i := 0; i < 24; i++ {
		temp := port.Read()
		if temp == '\r' || temp == '\n' {
			fSuccess = 3
		}
	}
	alg11, flag5 := port.ReadATwoAsciiData()
	alg12, flag6 := port.ReadATwoAsciiData()
	alg13, flag7 := port.ReadATwoAsciiData()
	alg14, flag8 := port.ReadATwoAsciiData()
	if !flag5 || !flag6 || !flag7 || !flag8 {
		fSuccess = 4
	}
	if alg11 > 4 || alg11 < -4 || alg12 > 4 || alg12 < -4 || alg13 > 4 || alg13 < -4 || alg14 > 4 || alg14 < -4 {
		fSuccess = 5
	}
	if fSuccess != 0 {
		return 0, 0, 0, 0, 0, 0, 0, 0, fSuccess
	} else {
		return byte(ead1_h), byte(ead1_l), byte(ead2_h), byte(ead2_l), alg11, alg12, alg13, alg14, 0
	}
}
