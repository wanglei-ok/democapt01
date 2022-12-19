package main

import "democapt01/gohksdk"

func main() {
	gohksdk.CapturePicture(1, 8000, "192.168.112.102", "admin", "sg123456", "sss.jpg")
}
