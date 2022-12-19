package main

/*
#cgo CFLAGS: -I./hksdk/include/
#cgo LDFLAGS: -L./hksdk/lib/ -lHCCore -lHCNetSDK

#include "GoWin.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	ret := C.NET_DVR_Init()
	if int(ret) != 1 {
		fmt.Printf("NET_DVR_Init failed error code = %v\n", C.NET_DVR_GetLastError())
		return
	}
	defer C.NET_DVR_Cleanup()

	var deviceInfoTmp C.NET_DVR_DEVICEINFO_V30

	lLoginID := C.NET_DVR_Login_V30(C.CString("192.168.112.102"), 8000, C.CString("admin"), C.CString("sg123456"), (C.LPNET_DVR_DEVICEINFO_V30)(&deviceInfoTmp))
	if lLoginID == -1 {
		fmt.Printf("Login to Device failed!\r\n")
		return
	}

	//组建jpg结构
	var JpgPara C.NET_DVR_JPEGPARA
	JpgPara.wPicSize = 0
	JpgPara.wPicQuality = 0

	if C.NET_DVR_CaptureJPEGPicture(lLoginID, 1, (C.LPNET_DVR_JPEGPARA)(&JpgPara), C.CString("sss.jpg")) == 0 {
		fmt.Printf("抓图失败，错误代码%v\n", C.NET_DVR_GetLastError())
		return
	}

	fmt.Printf("抓图成功!\n")

	if C.NET_DVR_Logout_V30(lLoginID) == 0 {
		fmt.Printf("Logout failed!\n")
	}
	lLoginID = -1

}
