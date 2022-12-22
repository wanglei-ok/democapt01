package main

import (
	"democapt01/gohksdk"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/rs/xid"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	tempPath = "temp/"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(cors())
	router.POST("/hls/:stream", HlsPostHandler)
	router.DELETE("/hls/:stream", HlsDeleteHandler)
	router.POST("/rtsp2hls/:stream", Rtsp2HlsPostHandler)
	router.DELETE("/public/:file", DeleteFile)
	router.StaticFS("/public", http.Dir("temp"))
	router.POST("/combine", Combine)
	router.GET("/captureJpeg", CaptureJpeg)
	router.POST("/getfilebytime", GetFileByTimePostHandler)
	router.GET("/getfilebytime", GetFileByTimeGetHandler)
	router.POST("/upload", UploadFileHandle)
	router.Run(":10086")
}

func Rtsp2HlsPostHandler(c *gin.Context) {
	url, _ := c.GetPostForm("url")
	h264, _ := c.GetPostForm("h264")
	stream := c.Param("stream")

	_, found := Cache.Get(stream)
	if found {
		c.JSON(400, gin.H{"code": 400, "desc": "stream already exist"})
		return
	}
	if !IsExist(tempPath + stream) {
		os.Mkdir(tempPath+stream, 777)

		args := []string{"-y", "-rtsp_transport", "tcp", "-i", url}
		if h264 != "" {
			args = append(args, "-force_key_frames", "expr:gte(t,n_forced*2)", "-c:a", "copy", "-c:v", "libx264")
		} else {
			args = append(args, "-c", "copy")
		}
		args = append(args, "-f", "hls", "-hls_flags", "delete_segments", tempPath+stream+"/file.m3u8")

		cmd := exec.Command("ffmpeg", args...)
		err := cmd.Start()
		if err != nil {
			c.JSON(400, gin.H{"code": 400, "desc": err.Error()})
			return
		}
		Cache.Set(stream, cmd, cache.NoExpiration)
	}

	c.JSON(200, gin.H{"code": 200, "desc": "rtsp2hls success"})
}

func HlsDeleteHandler(c *gin.Context) {
	stream := c.Param("stream")
	v, found := Cache.Get(stream)
	if found {
		Cache.Delete(stream)
		cmd := v.(*exec.Cmd)
		cmd.Process.Kill()
		cmd.Wait()
	}

	if !IsExist(tempPath + stream) {
		c.JSON(400, gin.H{"code": 400, "desc": "hls stream not exists"})
		return
	}
	err := os.RemoveAll(tempPath + stream)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "desc": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "desc": "delete hls success"})
}

func HlsPostHandler(c *gin.Context) {
	file, _ := c.GetPostForm("file")
	h264, _ := c.GetPostForm("h264")
	stream := c.Param("stream")

	if !IsExist(tempPath + stream) {
		os.Mkdir(tempPath+stream, 777)

		args := []string{"-y", "-i", tempPath + file}
		if h264 != "" {
			args = append(args, "-force_key_frames", "expr:gte(t,n_forced*2)", "-c:a", "copy", "-c:v", "libx264")
		} else {
			args = append(args, "-c", "copy")
		}
		args = append(args, "-f", "hls", "-hls_list_size", "30000", tempPath+stream+"/file.m3u8")

		cmd := exec.Command("ffmpeg", args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(400, gin.H{"code": 400, "desc": err.Error(), "output": string(out)})
			return
		}
	}

	c.JSON(200, gin.H{"code": 200, "desc": "convert hls success"})

}

func IsExist(s string) bool {
	_, err := os.Stat(s)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func DeleteFile(c *gin.Context) {
	file := c.Param("file")
	err := os.Remove(tempPath + file)
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "desc": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "desc": "delete success"})
}

func Combine(c *gin.Context) {
	video, _ := c.GetPostForm("video")
	audio, _ := c.GetPostForm("audio")

	videoExt := filepath.Ext(video)
	audioExt := filepath.Ext(audio)
	videoPrefix := video[0 : len(video)-len(videoExt)]
	audioPrefix := audio[0 : len(audio)-len(audioExt)]
	outputFile := fmt.Sprintf("%s_%s.mp4", videoPrefix, audioPrefix)
	cmd := exec.Command("ffmpeg", "-y", "-i", tempPath+video, "-i", tempPath+audio, "-vcodec", "copy",
		"-acodec", "aac", "-map", "0:v:0", "-map", "1:a:0", tempPath+outputFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(400, gin.H{"code": 400, "desc": err.Error(), "output": string(out)})
		return
	}
	c.JSON(200, gin.H{"code": 200, "desc": "combine success", "file": outputFile})
}

func UploadFileHandle(c *gin.Context) {
	file, _ := c.FormFile("uploadfile")
	name := file.Filename

	if len([]byte(name)) == 0 {
		c.JSON(400, gin.H{"code": 400, "desc": "filename invalid"})
		return
	}

	if file != nil {
		if err := c.SaveUploadedFile(file, "temp/"+name); err != nil {
			c.JSON(500, gin.H{"code": 500, "desc": "save file error"})
			return
		}
	} else {
		c.JSON(400, gin.H{"code": 400, "desc": "the uploadfile invalid"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "desc": "upload success"})
}

func GetFileByTimeGetHandler(c *gin.Context) {
	file, _ := c.GetQuery("file")
	v, found := Cache.Get("temp/" + file + ".tmp")
	if !found {
		c.JSON(200, gin.H{"code": 200, "desc": "Not found task to downloading"})
		return
	}
	myMap := v.(map[string]any)
	c.JSON(200, gin.H{"code": 200, "desc": myMap["message"], "nPos": myMap["nPost"]})
}

func GetFileByTimePostHandler(c *gin.Context) {
	channel := 1
	port := 8000
	if v, ok := c.GetPostForm("channel"); ok {
		channel, _ = strconv.Atoi(v)
	}

	if v, ok := c.GetPostForm("port"); ok {
		port, _ = strconv.Atoi(v)
	}
	ip, _ := c.GetPostForm("ip")
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")

	timeCond, _ := c.GetPostForm("time")
	if len(timeCond) != 28 {
		c.JSON(400, gin.H{"code": 400, "desc": "time param invalid"})
		return
	}

	fileName := fmt.Sprintf("%s_%d_%v.mp4", strings.Replace(ip, ".", "_", -1), channel, timeCond)
	saveFile := fmt.Sprintf("temp/%s", fileName)

	if IsExist(saveFile) {
		c.JSON(400, gin.H{"code": 400, "desc": "file exists"})
		return
	}
	go func() {
		ret := gohksdk.GetFileByTime(channel, port, ip, username, password, saveFile+".tmp", timeCond, func(s string, i int, m string) {
			Cache.Set(s, map[string]any{"nPost": i, "message": m}, cache.DefaultExpiration)
		})
		if ret == 0 {
			filepath.Walk("temp/", func(path string, info os.FileInfo, err error) error {
				if strings.Contains(path, fileName) {
					os.Rename(path, strings.Replace(path, ".tmp", "", -1))
				}
				return nil
			})

		} else {
			filepath.Walk("temp/", func(path string, info os.FileInfo, err error) error {
				if strings.Contains(path, fileName) {
					os.Remove(path)
				}
				return nil
			})

		}
	}()

	c.JSON(200, gin.H{"code": 200, "desc": "succeed", "file": strings.Replace(saveFile, "temp/", "", -1)})
	//defer os.Remove(saveFile)
	//c.File(saveFile)
}

func CaptureJpeg(c *gin.Context) {

	channel := 1
	port := 8000
	if v, ok := c.GetQuery("channel"); ok {
		channel, _ = strconv.Atoi(v)
	}

	if v, ok := c.GetQuery("port"); ok {
		port, _ = strconv.Atoi(v)
	}
	ip, _ := c.GetQuery("ip")
	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")

	id := xid.New()
	saveFile := fmt.Sprintf("temp/%v.jpg", id.String())
	ret := gohksdk.CaptureJpeg(channel, port, ip, username, password, saveFile)
	if ret != 0 {
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}

	//defer os.Remove(saveFile)
	c.File(saveFile)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
