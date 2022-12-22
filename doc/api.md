# The democapt API

  * [基础功能](#基础功能)
    * [实时抓图](#实时抓图)
    * [下载录像](#下载录像)
    * [查询下载进度](#查询下载进度)
    * [合并音视频](#合并音视频)
  * [文件管理](#文件管理)
    * [文件列表](#文件列表)
    * [文件上传](#文件上传)
    * [文件删除](#文件删除)
  * [HLS转换](#HLS)
    * [RTSP拉流转换为HLS](#RTSP2HLS)
    * [MP4文件转换为HLS](#MP42HLS)
    * [删除HLS](#delete-hls)
  * [Video endpoints](#Video-endpoints)
    * [HLS](#HLS-endpoint)
  



## 基础功能
### 实时抓图
#### Request
`GET /captureJpeg`

```base
wget http://localhost:10086/captureJpeg?channel=1&port=8000&ip=192.168.112.102&username=admin&password=a1234
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| channel | 通道号 | 33 |
| ip | IP地址 | 192.168.112.103 |
| port | 端口号 | 8000 |
| username | 注册账号 | admin |
| password | 密码 | a1234 |


#### Response
JPEG FILE

### 下载录像
#### Request
`POST /getfilebytime`

```base
curl -X "POST" -d 'channel=33&ip=192.168.112.103&username=admin&password=a1234&time=2022121218011120221212181122' http://localhost:10086/getfilebytime
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| channel | NVR通道号 | 33 |
| ip | IP地址 | 192.168.112.103 |
| port | 端口号 | 8000 |
| username | 注册账号 | admin |
| password | 密码 | a1234 |
| time | 录像起止时间，精确到秒，14字符开始时间+14字符截止时间，时间格式YYYYMMDDhhmmss | 2022121218011120221212181122 |


#### Response
```json
{
	"code": 200,
	"desc": "succeed",
	"file": "out.mp4"
}
```

```json
{
	"code": 400,
	"desc": "time param invalid"
}
```

```json
{
	"code": 400,
	"desc": "file exists"
}
```


#### Return

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| code | 返回代码。200 成功 400 失败 | 200 |
| desc | 描述信息。succeed调用成功，表示正确处理了请求，可使用文件名查询下载进度 | succeed  |
| file | 文件名用来查询下载进度或使用HTTP GET /public/<文件名>下载文件 | out.mp4 |



### 查询下载进度
#### Request
`GET /getfilebytime`

```base
curl "http://localhost:10086/getfilebytime?file=192_168_112_103_33_2022121218011120221212181122.mp4"
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| file | 指定要查询的文件，这个文件名是由下载请求的应答中返回地 | 192_168_112_103_33_2022121218011120221212181122.mp4 |


#### Response
```json
{
	"code": 200,
	"desc": "Be downloading...",
	"nPos": 23
}
```
```json
{
	"code": 200,
	"desc": "Not found task to downloading"
}
```
```json
{
	"code": 200,
	"desc": "Login to Device failed!",
	"nPos": 0
}
```


#### Return

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| code | 返回代码。200 成功| 200 |
| desc | 描述信息。 | Login to Device failed!  |
| nPos | 下载进度，百分比 | 0-100 |



### 合并音视频
#### Request
`POST /combine`

```base
curl -X "POST" -d 'video=192_168_112_103_33_2022121218011120221212181122.mp4&audio=5.wma' http://localhost:10086/combine
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| video | 指定视频文件 | 192_168_112_103_33_2022121218011120221212181122.mp4 |
| audio | 指定音频文件 | 5.wma |


#### Response
```json
{
	"code": 200,
	"desc": "combine success",
	"file": "192_168_112_103_33_2022121218011120221212181122_5.mp4"
}
```

```json
{
	"code": 400,
	"desc": "exit status 1",
	"output": "ffmpeg version 4.2.7-0ubuntu0.1 Copyright (c) 2000-2022 the FFmpeg developers\n  built with gcc 9 (Ubuntu 9.4.0-1ubuntu1~20.04.1)\n  configuration: --prefix=/usr --extra-version=0ubuntu0.1 --toolchain=hardened --libdir=/usr/lib/x86_64-linux-gnu --incdir=/usr/include/x86_64-linux-gnu --arch=amd64 --enable-gpl --disable-stripping --enable-avresample --disable-filter=resample --enable-avisynth --enable-gnutls --enable-ladspa --enable-libaom --enable-libass --enable-libbluray --enable-libbs2b --enable-libcaca --enable-libcdio --enable-libcodec2 --enable-libflite --enable-libfontconfig --enable-libfreetype --enable-libfribidi --enable-libgme --enable-libgsm --enable-libjack --enable-libmp3lame --enable-libmysofa --enable-libopenjpeg --enable-libopenmpt --enable-libopus --enable-libpulse --enable-librsvg --enable-librubberband --enable-libshine --enable-libsnappy --enable-libsoxr --enable-libspeex --enable-libssh --enable-libtheora --enable-libtwolame --enable-libvidstab --enable-libvorbis --enable-libvpx --enable-libwavpack --enable-libwebp --enable-libx265 --enable-libxml2 --enable-libxvid --enable-libzmq --enable-libzvbi --enable-lv2 --enable-omx --enable-openal --enable-opencl --enable-opengl --enable-sdl2 --enable-libdc1394 --enable-libdrm --enable-libiec61883 --enable-nvenc --enable-chromaprint --enable-frei0r --enable-libx264 --enable-shared\n  libavutil      56. 31.100 / 56. 31.100\n  libavcodec     58. 54.100 / 58. 54.100\n  libavformat    58. 29.100 / 58. 29.100\n  libavdevice    58.  8.100 / 58.  8.100\n  libavfilter     7. 57.100 /  7. 57.100\n  libavresample   4.  0.  0 /  4.  0.  0\n  libswscale      5.  5.100 /  5.  5.100\n  libswresample   3.  5.100 /  3.  5.100\n  libpostproc    55.  5.100 / 55.  5.100\nInput #0, mpeg, from 'temp/192_168_112_103_33_2022121218011120221212181122.mp4':\n  Duration: 00:10:10.92, start: 4752.810000, bitrate: 4025 kb/s\n    Stream #0:0[0x1e0]: Video: hevc (Main), yuvj420p(pc, bt709), 1920x1080, 25 fps, 25 tbr, 90k tbn\ntemp/5.wma: No such file or directory\n"
}
```

#### Return

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| code | 返回代码。200 成功| 200 |
| desc | 描述信息。 | combine success  |
| file | 输出文件文件名 | 192_168_112_103_33_2022121218011120221212181122_5.mp4 |
| output | 合并失败时，输出导致失败的信息|  |


## 文件管理
### 文件列表
#### Request
`GET /public`

浏览器打开 http://localhost:10086/public

### 文件上传
#### Request
`POST /upload`

```bash
curl -F 'uploadfile=@15.wma' http://localhost:10086/upload
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| uploadfile | 上传文件数据 |  |


#### Response
```json
{
	"code": 200,
	"desc": "upload success"
}
```
```json
{
	"code": 500,
	"desc": "save file error"
}
```
```json
{
	"code": 400,
	"desc": "the uploadfile invalid"
}
```

#### Return

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| code | 返回代码。200 成功,400失败，500服务器错误| 200 |
| desc | 描述信息。 | upload success  |

### 文件删除
#### Request
`DELETE /public/:file`

```base
curl -X "DELETE" http://localhost:10086/public/192_168_112_103_33_2022121218011120221212181122.mp4
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| :file | 指定要删除的文件名 | 192_168_112_103_33_2022121218011120221212181122.mp4 |


#### Response
```json
{
	"code": 200,
	"desc": "delete success"
}
```
```json
{
	"code": 400,
	"desc": ""
}
```

#### Return

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| code | 返回代码。200 成功| 200 |
| desc | 描述信息。 | delete success  |

## <a id='HLS'>HLS转换</a>
### <a id='RTSP2HLS'>RTSP拉流转换为HLS</a>

#### Request
`POST /rtsp2hls/:stream`

```base
curl -X "POST" -d 'url=rtsp://admin:a1234@192.168.112.102:554/h264/ch1/main/av_stream' http://localhost:10086/rtsp2hls/stream123
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| :stream | 自定义流ID | stream123 |
| url | 拉流地址URL | rtsp://admin:a1234@192.168.112.102:554/h264/ch1/main/av_stream |
| h264 | h264转码开关，h264=1开启转码，如果浏览器无法播放可尝试转码 | 1 |


#### Response
```json
{
	"code": 200,
	"desc": "rtsp2hls success"
}
```
```json
{
	"code": 400,
	"desc": "stream already exist"
}
```

#### Return

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| code | 返回代码。200 成功,400 失败| 400 |
| desc | 描述信息。 | stream already exist |


### <a id='MP42HLS'>MP4文件转换为HLS</a>

#### Request
`POST /hls/:stream`

```base
curl -X "POST" -d 'file=out.mp4' http://localhost:10086/hls/stream456
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| :stream | 自定义流ID | stream456 |
| file | 需要转换的文件，文件必须在服务器上，可以是录像下载、合并音视频输出的文件，和可以是上传的文件 | out.mp4 |
| h264 | h264转码开关，h264=1开启转码，如果浏览器无法播放可尝试转码 | 1 |


#### Response
```json
{
	"code": 200,
	"desc": "convert hls success"
}
```
```json
{
	"code": 400,
	"desc": "stream already exist"
}
```

#### Return

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| code | 返回代码。200 成功,400 失败| 400 |
| desc | 描述信息。 | stream already exist |

### <a id='delete-hls'>删除HLS</a>
#### Request
`DELETE /hls/:stream`

```base
curl -X "DELETE" http://localhost:10086/hls/stream123
```

#### Param

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| :stream | 自定义流ID | stream123 |


#### Response
```json
{
	"code": 200,
	"desc": "delete hls success"
}
```
```json
{
	"code": 400,
	"desc": "hls stream not exists"
}
```

#### Return

| 名称 | 说明 | 示例 |
| --- | --- | --- |
| code | 返回代码。200 成功,400 失败| 400 |
| desc | 描述信息。 | hls stream not exists |


## <a id='Video-endpoints'>Video endpoints</a>

### <a id='HLS-endpoint'>HLS</a>

`GET /public/{STREAM_ID}/file.m3u8`

```bash
curl http://localhost:10086/public/{STREAM_ID}/file.m3u8
```

```bash
ffplay http://localhost:10086/public/{STREAM_ID}/file.m3u8
```
