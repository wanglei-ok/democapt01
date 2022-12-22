#ifndef _HKSDK_H
#define _HKSDK_H

#define BOOL int
typedef  unsigned int       DWORD;
typedef  unsigned short     WORD;
typedef  unsigned short     USHORT;
typedef  short              SHORT;
typedef  int                LONG;
typedef  unsigned char      BYTE;
typedef  unsigned int       UINT;
typedef  void*              LPVOID;
typedef  void*              HANDLE;
typedef  unsigned int*      LPDWORD; 
typedef  unsigned long long UINT64;
typedef  signed long long   INT64;

#define SERIALNO_LEN            48      //序列号长度
#define STREAM_ID_LEN   32

#define NET_DVR_PLAYSTART        1//开始播放
#define NET_DVR_PLAYSTOP        2//停止播放
#define NET_DVR_PLAYPAUSE        3//暂停播放
#define NET_DVR_PLAYRESTART        4//恢复播放

//NET_DVR_Login_V30()参数结构
typedef struct
{
    BYTE sSerialNumber[SERIALNO_LEN];  //序列号
    BYTE byAlarmInPortNum;                //报警输入个数
    BYTE byAlarmOutPortNum;                //报警输出个数
    BYTE byDiskNum;                    //硬盘个数
    BYTE byDVRType;                    //设备类型, 1:DVR 2:ATM DVR 3:DVS ......
    BYTE byChanNum;                    //模拟通道个数
    BYTE byStartChan;                    //起始通道号,例如DVS-1,DVR - 1
    BYTE byAudioChanNum;                //语音通道数
    BYTE byIPChanNum;                    //最大数字通道个数，低位  
    BYTE byZeroChanNum;            //零通道编码个数 //2010-01-16
    BYTE byMainProto;            //主码流传输协议类型 0-private, 1-rtsp,2-同时支持private和rtsp
    BYTE bySubProto;                //子码流传输协议类型0-private, 1-rtsp,2-同时支持private和rtsp
    BYTE bySupport;        //能力，位与结果为0表示不支持，1表示支持，
    //bySupport & 0x1, 表示是否支持智能搜索
    //bySupport & 0x2, 表示是否支持备份
    //bySupport & 0x4, 表示是否支持压缩参数能力获取
    //bySupport & 0x8, 表示是否支持多网卡
    //bySupport & 0x10, 表示支持远程SADP
    //bySupport & 0x20, 表示支持Raid卡功能
    //bySupport & 0x40, 表示支持IPSAN 目录查找
    //bySupport & 0x80, 表示支持rtp over rtsp
    BYTE bySupport1;        // 能力集扩充，位与结果为0表示不支持，1表示支持
    //bySupport1 & 0x1, 表示是否支持snmp v30
    //bySupport1 & 0x2, 支持区分回放和下载
    //bySupport1 & 0x4, 是否支持布防优先级    
    //bySupport1 & 0x8, 智能设备是否支持布防时间段扩展
    //bySupport1 & 0x10, 表示是否支持多磁盘数（超过33个）
    //bySupport1 & 0x20, 表示是否支持rtsp over http    
    //bySupport1 & 0x80, 表示是否支持车牌新报警信息2012-9-28, 且还表示是否支持NET_DVR_IPPARACFG_V40结构体
    BYTE bySupport2; /*能力，位与结果为0表示不支持，非0表示支持                            
                     bySupport2 & 0x1, 表示解码器是否支持通过URL取流解码
                     bySupport2 & 0x2,  表示支持FTPV40
                     bySupport2 & 0x4,  表示支持ANR
                     bySupport2 & 0x8,  表示支持CCD的通道参数配置
                     bySupport2 & 0x10,  表示支持布防报警回传信息（仅支持抓拍机报警 新老报警结构）
                     bySupport2 & 0x20,  表示是否支持单独获取设备状态子项
    bySupport2 & 0x40,  表示是否是码流加密设备*/
    WORD wDevType;              //设备型号
    BYTE bySupport3; //能力集扩展，位与结果为0表示不支持，1表示支持
    //bySupport3 & 0x1, 表示是否支持批量配置多码流参数
    // bySupport3 & 0x4 表示支持按组配置， 具体包含 通道图像参数、报警输入参数、IP报警输入、输出接入参数、
    // 用户参数、设备工作状态、JPEG抓图、定时和时间抓图、硬盘盘组管理 
    //bySupport3 & 0x8为1 表示支持使用TCP预览、UDP预览、多播预览中的"延时预览"字段来请求延时预览（后续都将使用这种方式请求延时预览）。而当bySupport3 & 0x8为0时，将使用 "私有延时预览"协议。
    //bySupport3 & 0x10 表示支持"获取报警主机主要状态（V40）"。
    //bySupport3 & 0x20 表示是否支持通过DDNS域名解析取流
    
    BYTE byMultiStreamProto;//是否支持多码流,按位表示,0-不支持,1-支持,bit1-码流3,bit2-码流4,bit7-主码流，bit-8子码流
    BYTE byStartDChan;        //起始数字通道号,0表示无效
    BYTE byStartDTalkChan;    //起始数字对讲通道号，区别于模拟对讲通道号，0表示无效
    BYTE byHighDChanNum;        //数字通道个数，高位
    BYTE bySupport4;        //能力集扩展，位与结果为0表示不支持，1表示支持
    //bySupport4 & 0x02 表示是否支持NetSDK透传接口（NET_DVR_STDXMLConfig）透传表单格式
    //bySupport4 & 0x4表示是否支持拼控统一接口
    //bySupport4 & 0x80 支持设备上传中心报警使能。表示判断调用接口是 NET_DVR_PDC_RULE_CFG_V42还是 NET_DVR_PDC_RULE_CFG_V41
    BYTE byLanguageType;// 支持语种能力,按位表示,每一位0-不支持,1-支持  
    //  byLanguageType 等于0 表示 老设备
    //  byLanguageType & 0x1表示支持中文
    //  byLanguageType & 0x2表示支持英文
    BYTE byVoiceInChanNum;   //音频输入通道数 
    BYTE byStartVoiceInChanNo; //音频输入起始通道号 0表示无效
    BYTE  bySupport5;  //按位表示,0-不支持,1-支持,bit0-支持多码流
    //bySupport5 &0x01表示支持wEventTypeEx ,兼容dwEventType 的事件类型（支持行为事件扩展）--先占住，防止冲突
    //bySupport5 &0x04表示是否支持使用扩展的场景模式接口
    /*
       bySupport5 &0x08 设备返回该值表示是否支持计划录像类型V40接口协议(DVR_SET_RECORDCFG_V40/ DVR_GET_RECORDCFG_V40)(在该协议中设备支持类型类型13的配置)
       之前的部分发布的设备，支持录像类型13，则配置录像类型13。如果不支持，统一转换成录像类型3兼容处理。SDK通过命令探测处理)
       bySupport5 &0x10 设备返回改值表示支持超过255个预置点
    */
    BYTE  bySupport6;   //能力，按位表示，0-不支持,1-支持
    //bySupport6 0x1  表示设备是否支持压缩 
    //bySupport6 0x2 表示是否支持流ID方式配置流来源扩展命令，DVR_SET_STREAM_SRC_INFO_V40
    //bySupport6 0x4 表示是否支持事件搜索V40接口
    //bySupport6 0x8 表示是否支持扩展智能侦测配置命令
    //bySupport6 0x40表示图片查询结果V40扩展
    BYTE  byMirrorChanNum;    //镜像通道个数，<录播主机中用于表示导播通道>
    WORD wStartMirrorChanNo;  //起始镜像通道号
    BYTE bySupport7;   //能力,按位表示,0-不支持,1-支持
    // bySupport7 & 0x1  表示设备是否支持 INTER_VCA_RULECFG_V42 扩展
    // bySupport7 & 0x2  表示设备是否支持 IPC HVT 模式扩展
    // bySupport7 & 0x04  表示设备是否支持 返回锁定时间
    // bySupport7 & 0x08  表示设置云台PTZ位置时，是否支持带通道号
    // bySupport7 & 0x10  表示设备是否支持双系统升级备份
    // bySupport7 & 0x20  表示设备是否支持 OSD字符叠加 V50
    // bySupport7 & 0x40  表示设备是否支持 主从跟踪（从摄像机）
    // bySupport7 & 0x80  表示设备是否支持 报文加密
    BYTE  byRes2;        //保留
}NET_DVR_DEVICEINFO_V30, *LPNET_DVR_DEVICEINFO_V30;

//图片质量
typedef struct 
{
/*注意：当图像压缩分辨率为VGA时，支持0=CIF, 1=QCIF, 2=D1抓图，
当分辨率为3=UXGA(1600x1200), 4=SVGA(800x600), 5=HD720p(1280x720),6=VGA,7=XVGA, 8=HD900p
    仅支持当前分辨率的抓图*/
    
    /* 可以通过能力集获取
       0-CIF，           1-QCIF，           2-D1，         3-UXGA(1600x1200), 4-SVGA(800x600),5-HD720p(1280x720)，
       6-VGA，           7-XVGA，           8-HD900p，     9-HD1080，     10-2560*1920，
       11-1600*304，     12-2048*1536，     13-2448*2048,  14-2448*1200， 15-2448*800，
       16-XGA(1024*768), 17-SXGA(1280*1024),18-WD1(960*576/960*480),      19-1080i,      20-576*576，     
       21-1536*1536,     22-1920*1920,      23-320*240,    24-720*720,    25-1024*768,
       26-1280*1280,     27-1600*600,       28-2048*768,   29-160*120,    55-3072*2048,
       64-3840*2160,     70-2560*1440,      75-336*256,
       78-384*256,         79-384*216,        80-320*256,    82-320*192,    83-512*384,
       127-480*272,      128-512*272,       161-288*320,   162-144*176,   163-480*640,
       164-240*320,      165-120*160,       166-576*720,   167-720*1280,  168-576*960,
       180-180*240,      181-360*480,       182-540*720,    183-720*960,  184-960*1280,
       185-1080*1440,      215-1080*720(占位，未测试),  216-360x640(占位，未测试),245-576*704(占位，未测试)
       500-384*288,
       0xff-Auto(使用当前码流分辨率)
    */
    WORD    wPicSize;            
    WORD    wPicQuality;            /* 图片质量系数 0-最好 1-较好 2-一般 */
}NET_DVR_JPEGPARA, *LPNET_DVR_JPEGPARA;

typedef struct
{
    DWORD dwYear;        //年
    DWORD dwMonth;        //月
    DWORD dwDay;        //日
    DWORD dwHour;        //时
    DWORD dwMinute;        //分
    DWORD dwSecond;        //秒
}NET_DVR_TIME, *LPNET_DVR_TIME;

typedef struct tagNET_DVR_PLAYCOND
{
    DWORD             dwChannel;
    NET_DVR_TIME     struStartTime;
    NET_DVR_TIME     struStopTime;
    BYTE             byDrawFrame;  //0:不抽帧，1：抽帧
    BYTE             byStreamType ; //码流类型，0-主码流 1-子码流 2-码流三
    BYTE             byStreamID[STREAM_ID_LEN];
    BYTE             byCourseFile;    //课程文件0-否，1-是
    BYTE             byDownload;    //是否下载 0-否，1-是
    BYTE             byOptimalStreamType;    //是否按最优码流类型回放 0-否，1-是（对于双码流设备，某一段时间内的录像文件与指定码流类型不同，则返回实际码流类型的录像）
    BYTE             byVODFileType; // 下载录像文件，文件格式 0-PS码流格式，1-3GP格式
    BYTE             byRes[26];    //保留
}NET_DVR_PLAYCOND, *LPNET_DVR_PLAYCOND;


BOOL __stdcall NET_DVR_Init();
BOOL __stdcall NET_DVR_Cleanup();
DWORD __stdcall NET_DVR_GetLastError();
LONG __stdcall NET_DVR_Login_V30(char *sDVRIP, WORD wDVRPort, char *sUserName, char *sPassword, LPNET_DVR_DEVICEINFO_V30 lpDeviceInfo);
BOOL __stdcall NET_DVR_Logout_V30(LONG lUserID);
BOOL __stdcall NET_DVR_CaptureJPEGPicture(LONG lUserID, LONG lChannel, LPNET_DVR_JPEGPARA lpJpegPara, char *sPicFileName);
LONG __stdcall NET_DVR_GetFileByTime_V40(LONG lUserID, char *sSavedFileName, LPNET_DVR_PLAYCOND  pDownloadCond);
BOOL __stdcall NET_DVR_PlayBackControl_V40(LONG lPlayHandle,DWORD dwControlCode, LPVOID lpInBuffer, DWORD dwInLen, LPVOID lpOutBuffer, DWORD *lpOutLen);
int __stdcall NET_DVR_GetDownloadPos(LONG lFileHandle);
BOOL __stdcall NET_DVR_StopGetFile(LONG lFileHandle);

#endif //_HKSDK_H
