//go:build windows && amd64

package gohksdk

/*
   #cgo CFLAGS: -I../hksdk/win/include/
   #cgo LDFLAGS: -L../hksdk/win/lib/ -lHCCore -lhcnetsdk

   #include "hksdk.h"
*/
import "C"
import (
	"fmt"
	"strconv"
	"time"
)

func CaptureJpeg(channel, port int, ip, username, passwd, saveFile string) int {
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
	JpgPara.wPicSize = 0xff
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

func GetFileByTime(channel, port int, ip, username, passwd, saveFile, timeCond string) int {
	if len(timeCond) != 28 {
		return -3
	}

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
	defer C.NET_DVR_Logout_V30(lLoginID)

	var struDownloadCond C.NET_DVR_PLAYCOND
	struDownloadCond.dwChannel = C.uint(channel)

	var t int
	t, _ = strconv.Atoi(timeCond[0:4])
	struDownloadCond.struStartTime.dwYear = C.uint(t)
	t, _ = strconv.Atoi(timeCond[4:6])
	struDownloadCond.struStartTime.dwMonth = C.uint(t)
	t, _ = strconv.Atoi(timeCond[6:8])
	struDownloadCond.struStartTime.dwDay = C.uint(t)
	t, _ = strconv.Atoi(timeCond[8:10])
	struDownloadCond.struStartTime.dwHour = C.uint(t)
	t, _ = strconv.Atoi(timeCond[10:12])
	struDownloadCond.struStartTime.dwMinute = C.uint(t)
	t, _ = strconv.Atoi(timeCond[12:14])
	struDownloadCond.struStartTime.dwSecond = C.uint(t)
	t, _ = strconv.Atoi(timeCond[14:18])
	struDownloadCond.struStopTime.dwYear = C.uint(t)
	t, _ = strconv.Atoi(timeCond[18:20])
	struDownloadCond.struStopTime.dwMonth = C.uint(t)
	t, _ = strconv.Atoi(timeCond[20:22])
	struDownloadCond.struStopTime.dwDay = C.uint(t)
	t, _ = strconv.Atoi(timeCond[22:24])
	struDownloadCond.struStopTime.dwHour = C.uint(t)
	t, _ = strconv.Atoi(timeCond[24:26])
	struDownloadCond.struStopTime.dwMinute = C.uint(t)
	t, _ = strconv.Atoi(timeCond[26:28])
	struDownloadCond.struStopTime.dwSecond = C.uint(t)

	hPlayback := C.NET_DVR_GetFileByTime_V40(lLoginID, C.CString(saveFile), &struDownloadCond)
	if hPlayback < 0 {
		fmt.Printf("NET_DVR_GetFileByTime_V40 fail,last error %v\n", C.NET_DVR_GetLastError())
		return -4
	}

	//---------------------------------------
	//开始下载
	if C.NET_DVR_PlayBackControl_V40(hPlayback, C.NET_DVR_PLAYSTART, C.LPVOID(C.NULL), C.uint(0), C.LPVOID(C.NULL), (*C.uint)(C.NULL)) == 0 {
		fmt.Printf("Play back control failed [%v]\n", C.NET_DVR_GetLastError())
		return -5
	}

	nPos := 0
	for nPos = 0; nPos < 100 && nPos >= 0; nPos = int(C.NET_DVR_GetDownloadPos(hPlayback)) {
		fmt.Printf("Be downloading... %d %%\n", nPos)
		time.Sleep(1 * time.Second)
	}
	if C.NET_DVR_StopGetFile(hPlayback) == 0 {
		fmt.Printf("failed to stop get file [%v]\n", C.NET_DVR_GetLastError())
		return -6
	}
	if nPos < 0 || nPos > 100 {
		fmt.Printf("download err [%v]\n", C.NET_DVR_GetLastError())
		return -7
	}
	fmt.Printf("Be downloading... %d %%\n", nPos)

	return 0
}
