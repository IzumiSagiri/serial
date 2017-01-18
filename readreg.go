package serialportmethods

/*
#cgo LDFLAGS: -ladvapi32
#include <stdio.h>
#include <windows.h>
#pragma comment(lib, "advapi32.lib")

UCHAR strSerialList[256][25];
UCHAR *test = "test";

void readCommPorts()
{
    int i = 0;
    CHAR Name[25];
    UCHAR szPortName[25];
    LONG Status;
    DWORD dwIndex = 0;
    DWORD dwName;
    DWORD dwSizeofPortName;
    DWORD Type;
    HKEY hKey;
    LPCTSTR data_Set = "HARDWARE\\DEVICEMAP\\SERIALCOMM\\";
    dwName = sizeof(Name);
    dwSizeofPortName = sizeof(szPortName);
    long ret0 = RegOpenKeyEx(HKEY_LOCAL_MACHINE, data_Set, 0, KEY_READ, &hKey);
    if(ret0 == ERROR_SUCCESS)
    {
        do
        {
            Status = RegEnumValue(hKey, dwIndex++, Name, &dwName, NULL, &Type, szPortName, &dwSizeofPortName);
            if((Status == ERROR_SUCCESS) || (Status == ERROR_MORE_DATA))
            {
                strcpy(strSerialList[i], szPortName);
                i++;
            }
            dwName = sizeof(Name);
            dwSizeofPortName = sizeof(szPortName);
        } while((Status == ERROR_SUCCESS) || (Status == ERROR_MORE_DATA));
        RegCloseKey(hKey);
    }
}
*/
import "C"
import "fmt"

func GetAvailableComms() []string {
	var commsAvailable []string
	C.readCommPorts()
	for i := range C.strSerialList {
		if C.strSerialList[i][0] != 0 {
			commsAvailable = append(commsAvailable, convertToString(C.strSerialList[i]))
		}
	}
	fmt.Println(commsAvailable)
	return commsAvailable
}

func convertToString(wtf [25]C.UCHAR) string {
	var charray []byte
	for i := range wtf {
		if wtf[i] != 0 {
			charray = append(charray, byte(wtf[i]))
		}
	}
	return string(charray)
}
