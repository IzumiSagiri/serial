package main

/*
#cgo LDFLAGS: -lgdi32
#include <Windows.h>
#include <stdio.h>
#include <stdlib.h>
#pragma comment(lib, "gdi32.lib")

#define MSG(m) {\
    MessageBoxA(NULL,m,NULL,MB_OK);}

HWND hwnd;
HINSTANCE hinst;

POINT sclDataL[100];
POINT sclDataR[100];
POINT scrDataL[100];
POINT scrDataR[100];

char *tagstring;

#define WIDTH 500
#define HEIGHT 500

#define TIMER_ID 1

int renewDrawing(HWND hwnd)
{
    int i;
    HDC hdc;
    hdc=GetDC(hwnd);

    SelectObject(hdc,GetStockObject(BLACK_BRUSH));

    SelectObject(hdc,GetStockObject(WHITE_PEN));

    Rectangle(hdc,0,0,WIDTH,HEIGHT);

    MoveToEx(hdc,0,HEIGHT/2,(LPPOINT)NULL);
    LineTo(hdc,WIDTH,HEIGHT/2);
    MoveToEx(hdc,WIDTH/2,0,(LPPOINT)NULL);
    LineTo(hdc,WIDTH/2,HEIGHT);

    TextOut(hdc,0,0,tagstring,sizeof(tagstring));

    MoveToEx(hdc,sclDataL[0].x,sclDataL[0].y,(LPPOINT)NULL);
    for(i=1;i<100;i++)
    {
        LineTo(hdc,sclDataL[i].x,sclDataL[i].y);
    }

    MoveToEx(hdc,sclDataR[0].x,sclDataR[0].y,(LPPOINT)NULL);
    for(i=1;i<100;i++)
    {
        LineTo(hdc,sclDataR[i].x,sclDataR[i].y);
    }

    MoveToEx(hdc,scrDataL[0].x,scrDataL[0].y,(LPPOINT)NULL);
    for(i=1;i<100;i++)
    {
        LineTo(hdc,scrDataL[i].x,scrDataL[i].y);
    }

    MoveToEx(hdc,scrDataR[0].x,scrDataR[0].y,(LPPOINT)NULL);
    for(i=1;i<100;i++)
    {
        LineTo(hdc,scrDataR[i].x,scrDataR[i].y);
    }

    ReleaseDC(hwnd,hdc);

    return 0;
}
LRESULT CALLBACK WinProc(HWND hwnd,UINT msg,WPARAM wp,LPARAM lp)
{
    HDC hdc;
    PAINTSTRUCT ps;
    HPEN pen;
    static HBRUSH brush,brush2;

    switch(msg){
    case WM_DESTROY:
        PostQuitMessage(0);
        return 0;
    case WM_PAINT:
        hdc=GetDC(hwnd);

        SelectObject(hdc,GetStockObject(WHITE_PEN));
        MoveToEx(hdc,0,HEIGHT/2,(LPPOINT)NULL);
        LineTo(hdc,WIDTH,HEIGHT/2);
        MoveToEx(hdc,WIDTH/2,0,(LPPOINT)NULL);
        LineTo(hdc,WIDTH/2,HEIGHT);

        ReleaseDC(hwnd,hdc);

        return 0;
    }
    return DefWindowProc(hwnd,msg,wp,lp);
}

int WINAPI WinMain(HINSTANCE hInstance,HINSTANCE hPrevInstance,LPSTR lpCmdLine,int nShowCmd)
{
    MSG msg;
    WNDCLASS wc;

    wc.style = CS_HREDRAW | CS_VREDRAW;
    wc.lpfnWndProc=WinProc;
    wc.cbClsExtra=wc.cbWndExtra=0;
    wc.hInstance=hInstance;
    wc.hCursor=wc.hIcon=NULL;
    wc.hbrBackground=(HBRUSH)GetStockObject(BLACK_BRUSH);
    wc.lpszClassName="test";
    wc.lpszMenuName=NULL;

    if(!RegisterClass(&wc)){
        MSG("WTF!");
        return -1;
    }

    hwnd=CreateWindowA("test","testWindow",WS_VISIBLE | WS_SYSMENU | WS_CAPTION | WS_MINIMIZEBOX,0,0,WIDTH,HEIGHT,NULL,NULL,hinst,NULL);

    if(hwnd==NULL){
        MSG("Window Failed.");
    }


    int check;

    while(check=GetMessage(&msg,NULL,0,0)){
        if(check==-1){
            break;
        }
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }

    UnregisterClass("test",hinst);

    return 0;

}
*/
import "C"

import "github.com/IzumiSagiri/serial"

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	tempCh := make(chan bool)
	go func() {
		C.WinMain(C.hinst, nil, nil, 5)
		tempCh <- true
	}()
	go everything()
	<-tempCh
}

func everything() {
	//Port Initialize.
	var scrl, scrr serial.SCR
	var asahikaseil, asahikaseir serial.Asahikasei
	scrl.Cond = *sync.NewCond(new(sync.Mutex))
	scrr.Cond = *sync.NewCond(new(sync.Mutex))
	asahikaseil.Cond = *sync.NewCond(new(sync.Mutex))
	asahikaseir.Cond = *sync.NewCond(new(sync.Mutex))
	scrl.Port.Name = "//./COM10"
	scrl.Port.Baudrate = 19200
	scrr.Port.Name = "//./COM13"
	scrr.Port.Baudrate = 19200
	asahikaseil.Port.Name = "//./COM12"
	asahikaseil.Port.Baudrate = 115200
	asahikaseir.Port.Name = "//./COM14"
	asahikaseir.Port.Baudrate = 115200
	temp := scrl.Port.Open()
	if !temp {
		log.Fatal("Open scrl port failure.")
	}
	temp = scrr.Port.Open()
	if !temp {
		log.Fatal("Open scrr port failure.")
	}
	temp = asahikaseil.Port.Open()
	if !temp {
		log.Fatal("Open asahikaseil port failure.")
	}
	temp = asahikaseir.Port.Open()
	if !temp {
		log.Fatal("Open asahikaseir port failure.")
	}

	//StatusController Start.
	var status serial.Status
	status.Start()
	serial.SoundController(status)

	//File Initialize.
	var file serial.FileController
	file.Name = fmt.Sprintf("%v_%v_%v_%v_%v_%v.csv", status.TimeStarted.Year(), status.TimeStarted.Month(), status.TimeStarted.Day(), status.TimeStarted.Hour(), status.TimeStarted.Minute(), status.TimeStarted.Second())
	file.Open()
	file.Write("Duration, Tag, SCR left, SCR right, SCL left, SCL right, SCR Level left, SCR Level right\n")

	//Drawing Initialize
	initializeData()

	//DataReading Start.
	scrl.StartFlushing()
	scrr.StartFlushing()
	asahikaseil.StartFlushing()
	asahikaseir.StartFlushing()

	//Main Loop.
	for status.TagCounter <= 30 {
		status.Set()
		if status.Ticked {
			if status.TagCounter >= 31 {
				break
			}
			serial.SoundController(status)
			status.Ticked = false
		}
		scrl.Cond.L.Lock()
		for !scrl.Changed {
			scrl.Cond.Wait()
		}
		scrl.Changed = false
		scrl.Cond.L.Unlock()
		scrr.Cond.L.Lock()
		for !scrr.Changed {
			scrr.Cond.Wait()
		}
		scrr.Changed = false
		scrr.Cond.L.Unlock()
		asahikaseil.Cond.L.Lock()
		for !asahikaseil.Changed {
			asahikaseil.Cond.Wait()
		}
		asahikaseil.Changed = false
		asahikaseil.Cond.L.Unlock()
		asahikaseir.Cond.L.Lock()
		for !asahikaseir.Changed {
			asahikaseir.Cond.Wait()
		}
		asahikaseir.Changed = false
		asahikaseir.Cond.L.Unlock()

		scrl.Mux.Lock()
		scrr.Mux.Lock()
		asahikaseil.Mux.Lock()
		asahikaseir.Mux.Lock()

		asahikaseil.SCL = serial.CalculateFromEADs(asahikaseil.Ead1_h, asahikaseil.Ead1_l, asahikaseil.Ead2_h, asahikaseil.Ead2_l)
		asahikaseil.Level = serial.CalculateLevel(asahikaseil.Alg11, asahikaseil.Alg12, asahikaseil.Alg13, asahikaseil.Alg14)
		asahikaseir.SCL = serial.CalculateFromEADs(asahikaseir.Ead1_h, asahikaseir.Ead1_l, asahikaseir.Ead2_h, asahikaseir.Ead2_l)
		asahikaseir.Level = serial.CalculateLevel(asahikaseir.Alg11, asahikaseir.Alg12, asahikaseir.Alg13, asahikaseir.Alg14)

		C.tagstring = C.CString(status.Data[status.TagCounter].Tag)
		renewData(C.HEIGHT/2-int(asahikaseil.SCL*100), C.HEIGHT/2-int(asahikaseir.SCL*100), C.HEIGHT-int(scrl.SCR), C.HEIGHT-int(scrr.SCR))
		C.renewDrawing(C.hwnd)

		switch status.TagCounter {
		default:
		case 0:
		case 1:
			scrl.SCRCalm += int(scrl.SCR)
			scrr.SCRCalm += int(scrr.SCR)
			asahikaseil.SCLSum07 += asahikaseil.SCL
			asahikaseil.SCLCounterSum07++
			asahikaseil.SCLSum01 += asahikaseil.SCL
			asahikaseil.SCLCounterSum01++
			asahikaseir.SCLSum07 += asahikaseir.SCL
			asahikaseir.SCLCounterSum07++
			asahikaseir.SCLSum01 += asahikaseir.SCL
			asahikaseir.SCLCounterSum01++
		case 2:
			scrl.SCRCalm += int(scrl.SCR)
			scrr.SCRCalm += int(scrr.SCR)
			asahikaseil.SCLSum07 += asahikaseil.SCL
			asahikaseil.SCLCounterSum07++
			asahikaseil.SCLSum12 += asahikaseil.SCL
			asahikaseil.SCLCounterSum12++
			asahikaseir.SCLSum07 += asahikaseir.SCL
			asahikaseir.SCLCounterSum07++
			asahikaseir.SCLSum12 += asahikaseir.SCL
			asahikaseir.SCLCounterSum12++
		case 3:
			scrl.SCRCalm += int(scrl.SCR)
			scrr.SCRCalm += int(scrr.SCR)
			asahikaseil.SCLSum07 += asahikaseil.SCL
			asahikaseil.SCLCounterSum07++
			asahikaseil.SCLSum23 += asahikaseil.SCL
			asahikaseil.SCLCounterSum23++
			asahikaseil.SCLSum27 += asahikaseil.SCL
			asahikaseil.SCLCounterSum27++
			asahikaseir.SCLSum07 += asahikaseir.SCL
			asahikaseir.SCLCounterSum07++
			asahikaseir.SCLSum23 += asahikaseir.SCL
			asahikaseir.SCLCounterSum23++
			asahikaseir.SCLSum27 += asahikaseir.SCL
			asahikaseir.SCLCounterSum27++
		case 4:
			scrl.SCRCalm += int(scrl.SCR)
			scrr.SCRCalm += int(scrr.SCR)
			asahikaseil.SCLSum07 += asahikaseil.SCL
			asahikaseil.SCLCounterSum07++
			asahikaseil.SCLSum34 += asahikaseil.SCL
			asahikaseil.SCLCounterSum34++
			asahikaseil.SCLSum27 += asahikaseil.SCL
			asahikaseil.SCLCounterSum27++
			asahikaseir.SCLSum07 += asahikaseir.SCL
			asahikaseir.SCLCounterSum07++
			asahikaseir.SCLSum34 += asahikaseir.SCL
			asahikaseir.SCLCounterSum34++
			asahikaseir.SCLSum27 += asahikaseir.SCL
			asahikaseir.SCLCounterSum27++
		case 5:
			scrl.SCRCalm += int(scrl.SCR)
			scrr.SCRCalm += int(scrr.SCR)
			asahikaseil.SCLSum07 += asahikaseil.SCL
			asahikaseil.SCLCounterSum07++
			asahikaseil.SCLSum45 += asahikaseil.SCL
			asahikaseil.SCLCounterSum45++
			asahikaseil.SCLSum27 += asahikaseil.SCL
			asahikaseil.SCLCounterSum27++
			asahikaseir.SCLSum07 += asahikaseir.SCL
			asahikaseir.SCLCounterSum07++
			asahikaseir.SCLSum45 += asahikaseir.SCL
			asahikaseir.SCLCounterSum45++
			asahikaseir.SCLSum27 += asahikaseir.SCL
			asahikaseir.SCLCounterSum27++
		case 6:
			scrl.SCRCalm += int(scrl.SCR)
			scrr.SCRCalm += int(scrr.SCR)
			asahikaseil.SCLSum07 += asahikaseil.SCL
			asahikaseil.SCLCounterSum07++
			asahikaseil.SCLSum56 += asahikaseil.SCL
			asahikaseil.SCLCounterSum56++
			asahikaseil.SCLSum27 += asahikaseil.SCL
			asahikaseil.SCLCounterSum27++
			asahikaseir.SCLSum07 += asahikaseir.SCL
			asahikaseir.SCLCounterSum07++
			asahikaseir.SCLSum56 += asahikaseir.SCL
			asahikaseir.SCLCounterSum56++
			asahikaseir.SCLSum27 += asahikaseir.SCL
			asahikaseir.SCLCounterSum27++
		case 7:
			scrl.SCRCalm += int(scrl.SCR)
			scrr.SCRCalm += int(scrr.SCR)
			asahikaseil.SCLSum07 += asahikaseil.SCL
			asahikaseil.SCLCounterSum07++
			asahikaseil.SCLSum67 += asahikaseil.SCL
			asahikaseil.SCLCounterSum67++
			asahikaseil.SCLSum27 += asahikaseil.SCL
			asahikaseil.SCLCounterSum27++
			asahikaseir.SCLSum07 += asahikaseir.SCL
			asahikaseir.SCLCounterSum07++
			asahikaseir.SCLSum67 += asahikaseir.SCL
			asahikaseir.SCLCounterSum67++
			asahikaseir.SCLSum27 += asahikaseir.SCL
			asahikaseir.SCLCounterSum27++
		case 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29:
			scrl.SCRSound += int(scrl.SCR)
			scrr.SCRSound += int(scrr.SCR)
			asahikaseil.SCLSumSound += asahikaseil.SCL
			asahikaseil.SCLCounterSumSound++
			asahikaseir.SCLSumSound += asahikaseir.SCL
			asahikaseir.SCLCounterSumSound++
		}

		file.OutputData(status, scrl, scrr, asahikaseil, asahikaseil)

		fmt.Print(status.TagCounter)

		scrl.Mux.Unlock()
		scrr.Mux.Unlock()
		asahikaseil.Mux.Unlock()
		asahikaseir.Mux.Unlock()
	}
	file.OutputSummary(asahikaseir.SCLSum07/float64(asahikaseir.SCLCounterSum07), asahikaseil.SCLSum07/float64(asahikaseil.SCLCounterSum07), (asahikaseir.SCLSum07/float64(asahikaseir.SCLCounterSum07)+asahikaseil.SCLSum07/float64(asahikaseil.SCLCounterSum07))/2, asahikaseir.SCLSum27/float64(asahikaseir.SCLCounterSum27), asahikaseil.SCLSum27/float64(asahikaseil.SCLCounterSum27), (asahikaseir.SCLSum27/float64(asahikaseir.SCLCounterSum27)+asahikaseil.SCLSum27/float64(asahikaseil.SCLCounterSum27))/2, asahikaseir.SCLSum01/float64(asahikaseir.SCLCounterSum01), asahikaseil.SCLSum01/float64(asahikaseil.SCLCounterSum01), (asahikaseir.SCLSum01/float64(asahikaseir.SCLCounterSum01)+asahikaseil.SCLSum01/float64(asahikaseil.SCLCounterSum01))/2, asahikaseir.SCLSum12/float64(asahikaseir.SCLCounterSum12), asahikaseil.SCLSum12/float64(asahikaseil.SCLCounterSum12), (asahikaseir.SCLSum12/float64(asahikaseir.SCLCounterSum12)+asahikaseil.SCLSum12/float64(asahikaseil.SCLCounterSum12))/2, asahikaseir.SCLSum23/float64(asahikaseir.SCLCounterSum23), asahikaseil.SCLSum23/float64(asahikaseil.SCLCounterSum23), (asahikaseir.SCLSum23/float64(asahikaseir.SCLCounterSum23)+asahikaseil.SCLSum23/float64(asahikaseil.SCLCounterSum23))/2, asahikaseir.SCLSum34/float64(asahikaseir.SCLCounterSum34), asahikaseil.SCLSum34/float64(asahikaseil.SCLCounterSum34), (asahikaseir.SCLSum34/float64(asahikaseir.SCLCounterSum34)+asahikaseil.SCLSum34/float64(asahikaseil.SCLCounterSum34))/2, asahikaseir.SCLSum45/float64(asahikaseir.SCLCounterSum45), asahikaseil.SCLSum45/float64(asahikaseil.SCLCounterSum45), (asahikaseir.SCLSum45/float64(asahikaseir.SCLCounterSum45)+asahikaseil.SCLSum45/float64(asahikaseil.SCLCounterSum56))/2, asahikaseir.SCLSum56/float64(asahikaseir.SCLCounterSum56), asahikaseil.SCLSum56/float64(asahikaseil.SCLCounterSum56), (asahikaseir.SCLSum56/float64(asahikaseir.SCLCounterSum56)+asahikaseil.SCLSum56/float64(asahikaseil.SCLCounterSum56))/2, asahikaseir.SCLSum67/float64(asahikaseir.SCLCounterSum67), asahikaseil.SCLSum67/float64(asahikaseil.SCLCounterSum67), (asahikaseir.SCLSum67/float64(asahikaseir.SCLCounterSum67)+asahikaseil.SCLSum67/float64(asahikaseil.SCLCounterSum67))/2, asahikaseir.SCLSumSound/float64(asahikaseir.SCLCounterSumSound), asahikaseil.SCLSumSound/float64(asahikaseil.SCLCounterSumSound), (asahikaseir.SCLSumSound/float64(asahikaseir.SCLCounterSumSound)+asahikaseil.SCLSumSound/float64(asahikaseil.SCLCounterSumSound))/2, scrr.SCRCalm, scrl.SCRCalm, scrr.SCRCalm+scrl.SCRCalm, scrr.SCRSound, scrl.SCRSound, scrr.SCRSound+scrl.SCRSound)
}

func initializeData() error {
	dx := float64(C.WIDTH) / 2 / 99
	for i := 0; i < 100; i++ {
		C.scrDataL[i].x = C.LONG(int64(float64(i) * dx))
		C.scrDataL[i].y = C.LONG(C.HEIGHT)
		C.scrDataR[i].x = C.LONG(int64(C.WIDTH/2 + float64(i)*dx))
		C.scrDataR[i].y = C.LONG(C.HEIGHT)
		C.sclDataL[i].x = C.LONG(int64(float64(i) * dx))
		C.sclDataL[i].y = C.LONG(C.HEIGHT / 2)
		C.sclDataR[i].x = C.LONG(int64(C.WIDTH/2 + float64(i)*dx))
		C.sclDataR[i].y = C.LONG(C.HEIGHT / 2)
	}
	return nil
}

func renewData(data, data2, data3, data4 int) {
	for i := 0; i < 99; i++ {
		C.sclDataL[i].y = C.sclDataL[i+1].y
		C.sclDataR[i].y = C.sclDataR[i+1].y
		C.scrDataL[i].y = C.scrDataL[i+1].y
		C.scrDataR[i].y = C.scrDataR[i+1].y
	}
	C.sclDataL[99].y = C.LONG(data)
	C.sclDataR[99].y = C.LONG(data2)
	C.scrDataL[99].y = C.LONG(data3)
	C.scrDataR[99].y = C.LONG(data4)
}
