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
HWND hwndSettings;
HWND hwndSettings_button;
HWND hwndSettings_listbox;
HWND hwndSettings_listbox2;
HWND hwndSettings_listbox3;
HWND hwndSettings_listbox4;
HWND hwndSettings_static;
HINSTANCE hinst;
HINSTANCE hinstSettings;

POINT sclDataL[100];
POINT sclDataR[100];
POINT scrDataL[100];
POINT scrDataR[100];

char *tagstring;
char *portString;

char asahikaseil[50];
char asahikaseir[50];
char scrl[50];
char scrr[50];

#define WIDTH 500
#define HEIGHT 500

#define TIMER_ID 1
#define BUTTON_ID 0
#define LIST_ID 2
#define LIST_ID2 3
#define LIST_ID3 4
#define LIST_ID4 5
#define STATIC_ID 6

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
LRESULT CALLBACK WinProcSettings(HWND hwnd,UINT msg,WPARAM wp,LPARAM lp)
{
    char buf[50];
    int index, index2, index3, index4;

    switch(msg){
    case WM_DESTROY:
        PostQuitMessage(0);
        return 0;
    case WM_COMMAND:
        switch(LOWORD(wp)){
        case BUTTON_ID:
            memset(buf, 0, sizeof(buf));
            index=SendMessage(hwndSettings_listbox, LB_GETCURSEL, 0, 0);
            index2=SendMessage(hwndSettings_listbox2, LB_GETCURSEL, 0, 0);
            index3=SendMessage(hwndSettings_listbox3, LB_GETCURSEL, 0, 0);
            index4=SendMessage(hwndSettings_listbox4, LB_GETCURSEL, 0, 0);
            if (index == index2 || index3 == index4 ) {
                MSG("You chose the same port!");
            } else {
                SendMessage(hwndSettings_listbox, LB_GETTEXT, index, (WPARAM)buf);
                strcpy(asahikaseil, buf);
                SendMessage(hwndSettings_listbox2, LB_GETTEXT, index2, (WPARAM)buf);
                strcpy(asahikaseir, buf);
                SendMessage(hwndSettings_listbox3, LB_GETTEXT, index3, (WPARAM)buf);
                strcpy(scrl, buf);
                SendMessage(hwndSettings_listbox4, LB_GETTEXT, index4, (WPARAM)buf);
                strcpy(scrr, buf);
            }
            return 0;
        }
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

void AddListbox(char *port) {
    if (strstr(port, "Bth")) {
        SendMessage(hwndSettings_listbox, LB_ADDSTRING, 0, (WPARAM)port);
        SendMessage(hwndSettings_listbox2, LB_ADDSTRING, 0, (WPARAM)port);
    } else if (strstr(port, "V")) {
        SendMessage(hwndSettings_listbox3, LB_ADDSTRING, 0, (WPARAM)port);
        SendMessage(hwndSettings_listbox4, LB_ADDSTRING, 0, (WPARAM)port);
    } else {
        printf("Error!!!");
    }
}

int WINAPI WinMainSettings(HINSTANCE hInstance,HINSTANCE hPrevInstance,LPSTR lpCmdLine,int nShowCmd)
{
    MSG msg;
    WNDCLASS wcSettings;

    wcSettings.style = CS_HREDRAW | CS_VREDRAW;
    wcSettings.lpfnWndProc=WinProcSettings;
    wcSettings.cbClsExtra=wcSettings.cbWndExtra=0;
    wcSettings.hInstance=hInstance;
    wcSettings.hCursor=wcSettings.hIcon=NULL;
    wcSettings.hbrBackground=(HBRUSH)COLOR_WINDOW;
    wcSettings.lpszClassName="testSettings";
    wcSettings.lpszMenuName=NULL;

    if(!RegisterClass(&wcSettings)){
        MSG("WTF!2");
        return -1;
    }

    hwndSettings=CreateWindowA("testSettings","SettingsWindow",WS_VISIBLE | WS_SYSMENU | WS_CAPTION | WS_MINIMIZEBOX,0,0,300,350,NULL,NULL,hinstSettings,NULL);
    hwndSettings_button=CreateWindowA("button","Confirm",WS_VISIBLE | WS_CHILD | BS_PUSHBUTTON,100,300,100,20,hwndSettings,(HMENU)BUTTON_ID,hInstance,NULL);
    hwndSettings_listbox=CreateWindowA("listbox", NULL,WS_VISIBLE | WS_CHILD | LBS_SORT | LBS_NOTIFY | WS_VSCROLL, 0,0,150,150,hwndSettings,(HMENU)LIST_ID,hInstance,NULL);
    hwndSettings_listbox2=CreateWindowA("listbox", NULL,WS_VISIBLE | WS_CHILD | LBS_SORT | LBS_NOTIFY | WS_VSCROLL, 150,0,150,150,hwndSettings,(HMENU)LIST_ID2,hInstance,NULL);
    hwndSettings_listbox3=CreateWindowA("listbox", NULL,WS_VISIBLE | WS_CHILD | LBS_SORT | LBS_NOTIFY | WS_VSCROLL, 0,150,150,150,hwndSettings,(HMENU)LIST_ID3,hInstance,NULL);
    hwndSettings_listbox4=CreateWindowA("listbox", NULL,WS_VISIBLE | WS_CHILD | LBS_SORT | LBS_NOTIFY | WS_VSCROLL, 150,150,150,150,hwndSettings,(HMENU)LIST_ID4,hInstance,NULL);
    hwndSettings_static=CreateWindowA("static", "test", WS_VISIBLE | WS_CHILD, 0,300,100,50,hwndSettings,(HMENU)STATIC_ID,hInstance,NULL);

    if(hwndSettings==NULL){
        MSG("Window Failed.");
    }
    if(hwndSettings_button==NULL){
        MSG("Button Failed.");
    }

    int check;

    while(check=GetMessage(&msg,NULL,0,0)){
        if(check==-1){
            break;
        }
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }

    UnregisterClass("testSettings",hinstSettings);

    return 0;

}
*/
import "C"

import "github.com/IzumiSagiri/serial"

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

func ConvertToString(wtf [50]C.char) string {
	var charray []byte
	for i := range wtf {
		if wtf[i] != 0 {
			charray = append(charray, byte(wtf[i]))
		}
	}
	return string(charray)
}

func main() {
	var m = make(map[string]string)
	serial.GetAvailableComms(m)
	tempCh := make(chan bool)
	go func() {
		C.WinMainSettings(C.hinstSettings, nil, nil, 5)
		tempCh <- true
	}()
	time.Sleep(100 * time.Millisecond)
	for k, v := range m {
		C.AddListbox(C.CString(fmt.Sprint(k + " : " + v)))
	}
	<-tempCh
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
	scrlTempString := ConvertToString(C.scrl)
	if scrlTempString == "" {
		scrl.Accessable = false
	} else {
		scrl.Accessable = true
		scrl.Port.Name = "//./" + scrlTempString[strings.Index(scrlTempString, "COM"):]
	}
	scrl.Port.Baudrate = 19200
	scrrTempString := ConvertToString(C.scrr)
	if scrrTempString == "" {
		scrr.Accessable = false
	} else {
		scrr.Accessable = true
		scrr.Port.Name = "//./" + scrrTempString[strings.Index(scrrTempString, "COM"):]
	}
	scrr.Port.Baudrate = 19200
	asahikaseilTempString := ConvertToString(C.asahikaseil)
	if asahikaseilTempString == "" {
		asahikaseil.Accessable = false
	} else {
		asahikaseil.Accessable = true
		asahikaseil.Port.Name = "//./" + asahikaseilTempString[strings.Index(asahikaseilTempString, "COM"):]
	}
	asahikaseil.Port.Baudrate = 115200
	asahikaseirTempString := ConvertToString(C.asahikaseir)
	if asahikaseirTempString == "" {
		asahikaseir.Accessable = false
	} else {
		asahikaseir.Accessable = true
		asahikaseir.Port.Name = "//./" + asahikaseirTempString[strings.Index(asahikaseirTempString, "COM"):]
	}
	asahikaseir.Port.Baudrate = 115200
	if scrl.Accessable {
		temp := scrl.Port.Open()
		if !temp {
			log.Print("Open scrl port failure.")
			scrl.Accessable = false
		} else {
			scrl.Accessable = true
		}
	}
	if scrr.Accessable {
		temp := scrr.Port.Open()
		if !temp {
			log.Print("Open scrr port failure.")
			scrr.Accessable = false
		} else {
			scrr.Accessable = true
		}
	}
	if asahikaseil.Accessable {
		temp := asahikaseil.Port.Open()
		if !temp {
			log.Print("Open asahikaseil port failure.")
			asahikaseil.Accessable = false
		} else {
			asahikaseil.Accessable = true
		}
	}
	if asahikaseir.Accessable {
		temp := asahikaseir.Port.Open()
		if !temp {
			log.Print("Open asahikaseir port failure.")
			asahikaseir.Accessable = false
		} else {
			asahikaseir.Accessable = true
		}
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
				log.Println("Process Finished.")
				break
			}
			serial.SoundController(status)
			status.Ticked = false
		}
		if scrl.Accessable {
			scrl.Cond.L.Lock()
			for !scrl.Changed {
				scrl.Cond.Wait()
			}
			scrl.Changed = false
			scrl.Cond.L.Unlock()
		}
		if scrr.Accessable {
			scrr.Cond.L.Lock()
			for !scrr.Changed {
				scrr.Cond.Wait()
			}
			scrr.Changed = false
			scrr.Cond.L.Unlock()
		}
		if asahikaseil.Accessable {
			asahikaseil.Cond.L.Lock()
			for !asahikaseil.Changed {
				asahikaseil.Cond.Wait()
			}
			asahikaseil.Changed = false
			asahikaseil.Cond.L.Unlock()
		}
		if asahikaseir.Accessable {
			asahikaseir.Cond.L.Lock()
			for !asahikaseir.Changed {
				asahikaseir.Cond.Wait()
			}
			asahikaseir.Changed = false
			asahikaseir.Cond.L.Unlock()
		}

		scrl.Mux.Lock()
		scrr.Mux.Lock()
		asahikaseil.Mux.Lock()
		asahikaseir.Mux.Lock()

		if asahikaseil.Accessable {
			asahikaseil.SCL = serial.CalculateFromEADs(asahikaseil.Ead1_h, asahikaseil.Ead1_l, asahikaseil.Ead2_h, asahikaseil.Ead2_l)
			asahikaseil.Level = serial.CalculateLevel(asahikaseil.Alg11, asahikaseil.Alg12, asahikaseil.Alg13, asahikaseil.Alg14)
		}
		if asahikaseir.Accessable {
			asahikaseir.SCL = serial.CalculateFromEADs(asahikaseir.Ead1_h, asahikaseir.Ead1_l, asahikaseir.Ead2_h, asahikaseir.Ead2_l)
			asahikaseir.Level = serial.CalculateLevel(asahikaseir.Alg11, asahikaseir.Alg12, asahikaseir.Alg13, asahikaseir.Alg14)
		}

		C.tagstring = C.CString(status.Data[status.TagCounter].Tag)
		renewData(C.HEIGHT/2-int(asahikaseil.SCL*10), C.HEIGHT/2-int(asahikaseir.SCL*10), C.HEIGHT-int(scrl.SCR), C.HEIGHT-int(scrr.SCR))
		C.renewDrawing(C.hwnd)

		serial.CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)

		file.OutputData(status, scrl, scrr, asahikaseil, asahikaseir)

		fmt.Print(status.TagCounter)

		scrl.Mux.Unlock()
		scrr.Mux.Unlock()
		asahikaseil.Mux.Unlock()
		asahikaseir.Mux.Unlock()
	}
	log.Println("Output File.")
	file.OutputAverage(scrl, scrr, asahikaseil, asahikaseir)
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
	if data < 0 {
		data = 0
	}
	if data2 < 0 {
		data2 = 0
	}
	if data3 < 250 {
		data3 = 250
	}
	if data4 < 250 {
		data4 = 250
	}
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
