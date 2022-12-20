//go:build linux && amd64

package gohksdk

/*
   #cgo CFLAGS: -I../hksdk/linux/include/
   #cgo LDFLAGS: -L../hksdk/linux/lib/ -lHCCore -lhcnetsdk

   #include "hksdk.h"
*/
import "C"
import (
	"fmt"
)

func CapturePicture(channel, port int, ip, username, passwd, saveFile string) int {
	ret := C.NET_DVR_Init()
	if int(ret) != 1 {
		fmt.Printf("NET_DVR_Init failed error code = %v\n", C.NET_DVR_GetLastError())
		return -1
	}
	defer C.NET_DVR_Cleanup()

	var deviceInfoTmp C.NET_DVR_DEVICEINFO_V30

	lLoginID := C.NET_DVR_Login_V30(C.CString(ip), C.ushort(port), C.CString(username), C.CString(passwd), (C.LPNET_DVR_DEVICEINFO_V30)(&deviceInfoTmp))
	if lLoginID == -1 {
		fmt.Printf("Login to Device failed!\r\n")
		return -2
	}

	//组建jpg结构
	var JpgPara C.NET_DVR_JPEGPARA
	JpgPara.wPicSize = 0
	JpgPara.wPicQuality = 0

	if C.NET_DVR_CaptureJPEGPicture(lLoginID, C.int(channel), (C.LPNET_DVR_JPEGPARA)(&JpgPara), C.CString(saveFile)) == 0 {
		fmt.Printf("抓图失败，错误代码%v\n", C.NET_DVR_GetLastError())
		return -3
	}

	fmt.Printf("抓图成功!\n")

	if C.NET_DVR_Logout_V30(lLoginID) == 0 {
		fmt.Printf("Logout failed!\n")
	}
	lLoginID = -1

	return 0
}
