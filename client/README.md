### sip 通信

#### REGISTER 客户端注册
```
REGISTER sip:34020000002000000001@192.168.1.125:5060 SIP/2.0
Call-ID: b954e1e87c6d580eaec2c2755aba3e54
Contact: <sip:34020000001180000001@192.168.1.108:5060>
Content-Length: 0
CSeq: 1 REGISTER
Expires: 3600
From: <sip:34020000001180000001@192.168.1.108:5060>;tag=d28d156d2c72dcfb6a49ceb55f9627e8
Max-Forwards: 70
Route: <sip:34020000001180000001@192.168.1.125:5060;lr>
To: <sip:34020000001180000001@192.168.1.108:5060>
User-Agent: SIP UAS V3.0.0.1208500
Via: SIP/2.0/UDP 192.168.1.108:5060;rport;branch=z9hG4bK9ffcac0c7c50cdbc062009c67a39ecfb
```

```
Via: SIP/2.0/UDP 192.168.1.107:5060;rport;branch=z9hG4bK1821122895
From: <sip:34020000001320000001@192.168.1.107:5060>;tag=836318699
To: <sip:34020000001320000001@192.168.1.107:5060>
Call-ID: 1280868601@192.168.1.107
CSeq: 1 REGISTER
Contact: <sip:34020000001320000001@192.168.1.107:5060>
Max-Forwards: 70
User-Agent: DH SIP UAS V1.0
Expires: 3600
Content-Length: 0
```

```
REGISTER sip:34020000002000000001@192.168.1.125:5060 SIP/2.0
Via: SIP/2.0/UDP 192.100.100.125:5080;branch=z9hG4bK.8RxXCC3oFwdJZsrjpdmxbDvR60scNeer
CSeq: 1 REGISTER
From: <sip:34020000001190000001@127.0.0.1:5080>;tag=KR0lRo0j
To: <sip:34020000001190000001@127.0.0.1:5080>
Call-ID: xPCsP1arOS8HrdC60Xhho6CiQ2ayqts4
Contact: <sip:34020000001190000001@192.100.100.125:5080;transport=udp>;+sip.instance="<urn:uuid:c781fa55-35ba-11ee-895d-0826ae3b6a2b>"
Max-Forwards: 70
Content-Length: 0
Expires: 1800
Allow: INVITE, ACK, BYE, CANCEL, INFO, OPTIONS, UPDATE, MESSAGE, SUBSCRIBE
Supported: replaces, outbound
User-Agent: Go Sip Client/example-client
```

#### subscribe


#### message
```
<?xml version="1.0" encoding="GB2312" standalone="yes" ?><Response><CmdType>Catalog</CmdType><SN>1</SN><DeviceID>34020000001180000001</DeviceID><SumNum>4</SumNum><DeviceList Num="1"><Item><DeviceID>34020000001310000001</DeviceID><Name>´ó»ª1.107</Name><Manufacturer>General</Manufacturer><Model>IPC-HFW4150D-V2</Model><Owner>0</Owner><CivilCode>340200</CivilCode><Address>axy</Address><Parental>0</Parental><ParentID>34020000001180000001</ParentID><RegisterWay>1</RegisterWay><Secrecy>0</Secrecy><StreamNum>2</StreamNum><IPAddress>192.168.1.107</IPAddress><Port>80</Port><Password></Password><Status>ON</Status><Info><PTZType>3</PTZType><DownloadSpeed>1/2/4/8</DownloadSpeed></Info></Item></DeviceList></Response>
```

```
<?xml version="1.0" encoding="GB2312" standalone="yes" ?><Notify><CmdType>Alarm</CmdType><SN>9</SN><DeviceID>34020000001310000001</DeviceID><AlarmPriority>1</AlarmPriority><AlarmMethod>5</AlarmMethod><AlarmTime>2023-08-09T07:38:35</AlarmTime><AlarmDescription>ÊÓÆµ¶¯¼ì</AlarmDescription><AlarmInfo>11</AlarmInfo><Info><AlarmType>2</AlarmType><AlarmTypeParam><EventType>1</EventType></AlarmTypeParam></Info></Notify>
```

```
<?xml version="1.0" encoding="GB2312" ?>
<Response>
    <CmdType>Catalog</CmdType>
    <SN>1</SN>
    <DeviceID>34020000001320000001</DeviceID>
    <SumNum>1</SumNum>
    <DeviceList Num="1">
        <Item>
            <DeviceID>34020000001320000001</DeviceID>
            <Name>´ó»ª1.107</Name>
            <Manufacturer>DAHUA</Manufacturer>
            <Model>IPC-HFW4150D-V2</Model>
            <Owner>0</Owner>
            <CivilCode>6532</CivilCode>
            <Address>axy</Address>
            <Parental>0</Parental>
            <RegisterWay>1</RegisterWay>
            <Secrecy>0</Secrecy>
            <Status>ON</Status>
        </Item>
    </DeviceList>
</Response>
```


```
<?xml version="1.0" encoding="GB2312" standalone="yes" ?><Response><CmdType>Catalog</CmdType><SN>4</SN><DeviceID>34020000001180000001</DeviceID><SumNum>4</SumNum><DeviceList Num="1"><Item><DeviceID>34020000001310000004</DeviceID><Name>1.64 º£¿µÉãÏñ»ú</Name><Manufacturer>General</Manufacturer><Model>DS-2CD2725EFD-IS</Model><Owner>0</Owner><CivilCode>340200</CivilCode><Address>axy</Address><Parental>0</Parental><ParentID>34020000001180000001</ParentID><RegisterWay>1</RegisterWay><Secrecy>0</Secrecy><StreamNum>2</StreamNum><IPAddress>192.168.1.64</IPAddress><Port>80</Port><Password></Password><Status>ON</Status><Info><PTZType>3</PTZType><DownloadSpeed>1/2/4/8</DownloadSpeed></Info></Item></DeviceList></Response>
```

```
<?xml version="1.0" encoding="GB2312" standalone="yes" ?><Response><CmdType>Catalog</CmdType><SN>4</SN><DeviceID>34020000001180000001</DeviceID><SumNum>4</SumNum><DeviceList Num="1"><Item><DeviceID>34020000001310000002</DeviceID><Name>GB_Chn_002</Name><Manufacturer>General</Manufacturer><Model>DH-IVSS708-S1</Model><Owner>0</Owner><CivilCode>340200</CivilCode><Address>axy</Address><Parental>0</Parental><ParentID>34020000001180000001</ParentID><RegisterWay>1</RegisterWay><Secrecy>0</Secrecy><StreamNum>2</StreamNum><IPAddress>rtsp://admin@192.168.1.64:554/Streaming/Channels/101</IPAddress><Port>554</Port><Password></Password><Status>ON</Status><Info><PTZType>3</PTZType><DownloadSpeed>1/2/4/8</DownloadSpeed></Info></Item></DeviceList></Response>
```

```
<?xml version="1.0" encoding="GB2312" standalone="yes" ?><Response><CmdType>Catalog</CmdType><SN>4</SN><DeviceID>34020000001180000001</DeviceID><SumNum>4</SumNum><DeviceList Num="1"><Item><DeviceID>34020000001310000003</DeviceID><Name>1.64 º£¿µÉãÏñ»ú</Name><Manufacturer>General</Manufacturer><Model>DS-2CD2725EFD-IS</Model><Owner>0</Owner><CivilCode>340200</CivilCode><Address>axy</Address><Parental>0</Parental><ParentID>34020000001180000001</ParentID><RegisterWay>1</RegisterWay><Secrecy>0</Secrecy><StreamNum>2</StreamNum><IPAddress>192.168.1.64</IPAddress><Port>80</Port><Password></Password><Status>ON</Status><Info><PTZType>3</PTZType><DownloadSpeed>1/2/4/8</DownloadSpeed></Info></Item></DeviceList></Response>
```

#### Keepalive心跳
```
<?xml version="1.0" encoding="GB2312" standalone="yes" ?><Notify><CmdType>Keepalive</CmdType><SN>279</SN><DeviceID>34020000001180000001</DeviceID><Status>OK</Status></Notify>
<?xml version="1.0" encoding="GB2312" standalone="yes" ?><Notify><CmdType>Keepalive</CmdType><SN>280</SN><DeviceID>34020000001180000001</DeviceID><Status>OK</Status></Notify>
```

```
MESSAGE sip:34020000002000000001@192.168.1.125:5060 SIP/2.0
Via: SIP/2.0/UDP 192.168.1.107:5060;rport;branch=z9hG4bK990289307
From: <sip:34020000001320000001@192.168.1.107:5060>;tag=50882131
To: <sip:34020000002000000001@192.168.1.125:5060>
Call-ID: 432363298@192.168.1.107
CSeq: 20 MESSAGE
Max-Forwards: 70
User-Agent: DH SIP UAS V1.0
Content-Type: Application/MANSCDP+xml
Content-Length:   178

<?xml version="1.0" encoding="GB2312" ?>
<Notify>
    <CmdType>Keepalive</CmdType>
    <SN>4</SN>
    <DeviceID>34020000001320000001</DeviceID>
    <Status>OK</Status>
</Notify>
```

#### INVITE
```
server => client

INVITE sip:34020000001320000001@192.168.1.107:5060 SIP/2.0
Via: SIP/2.0/UDP 192.100.100.125:5060;branch=z9hG4bK.18pdKrzVPh65bxy4JpeBnpdDuMch5W8B
From: <sip:34020000002000000001@192.168.1.125:5060>;tag=454756032
To: <sip:34020000001320000001@192.168.1.107:5060>
Call-ID: 4547560325
User-Agent: Monibuca
CSeq: 3 INVITE
Max-Forwards: 70
Contact: <sip:34020000002000000001@192.168.1.125:5060>;tag=454756032
Content-Type: application/sdp
Content-Length: 232
Subject: 34020000001320000001:0200007686,34020000002000000001:0
Allow: INVITE, ACK, CANCEL, REGISTER, MESSAGE, NOTIFY, BYE

v=0
o=34020000001320000001 0 0 IN IP4 192.168.1.125
s=Play
u=34020000001320000001:0
c=IN IP4 192.168.1.125
t=0 0
m=video 58200 TCP/RTP/AVP 96
a=recvonly
a=rtpmap:96 PS/90000
y=0200007686
a=setup:passive
a=connection:new
```

```
client => server

SIP/2.0 100 Trying
Via: SIP/2.0/UDP 192.100.100.125:5060;branch=z9hG4bK.18pdKrzVPh65bxy4JpeBnpdDuMch5W8B;received=192.168.1.125
From: <sip:34020000002000000001@192.168.1.125:5060>;tag=454756032
To: <sip:34020000001320000001@192.168.1.107:5060>
Call-ID: 4547560325
CSeq: 3 INVITE
User-Agent: DH SIP UAS V1.0
Content-Length: 0

SIP/2.0 200 OK
Call-ID: 6110213552
Contact: <sip:34020000001180000001@192.168.1.108:5060>
Content-Length: 283
Content-Type: application/sdp
CSeq: 3 INVITE
From: <sip:34020000002000000001@192.168.1.125:5060>;tag=611021355
To: <sip:34020000001310000001@192.168.1.108:80>;tag=3e86e5287e6934107873f8a078dd3039
User-Agent: SIP UAS V3.0.0.1208500
Via: SIP/2.0/UDP 192.100.100.125:5060;received=192.168.1.125;rport=5060;branch=z9hG4bK.wAF3TBVMMrwiOOCPVxcc7D6ngsp4t9jK

v=0
o=34020000001310000001 0 0 IN IP4 192.168.1.108
s=Play
i=VCam Live Video
c=IN IP4 192.168.1.108
t=0 0
m=video 10032 TCP/RTP/AVP 96
a=sendonly
a=rtpmap:96 PS/90000
a=streamprofile:0
a=streamnumber:0
a=setup:active
a=connection:new
y=0200005549
f=v/0/0/0/0/0a/0/0/0


```


```
ACK sip:34020000001180000001@192.168.1.108:5060 SIP/2.0
Via: SIP/2.0/UDP 192.168.1.125:5060;branch=z9hG4bK.0HS0g8m0Dx4ZtMAw0w4lEIqX2EdxSFoX
Max-Forwards: 70
From: <sip:34020000002000000001@192.168.1.125:5060>;tag=460815743
To: <sip:34020000001310000003@192.168.1.108:80>;tag=957752ba5d739f2aba57fe6b4b383216
Call-ID: 4608157435
CSeq: 4 ACK
Content-Length: 0
User-Agent: GoSIP


```