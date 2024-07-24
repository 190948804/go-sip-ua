package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/cloudwebrtc/go-sip-ua/pkg/account"
	"github.com/cloudwebrtc/go-sip-ua/pkg/media/rtp"
	"github.com/cloudwebrtc/go-sip-ua/pkg/session"
	"github.com/cloudwebrtc/go-sip-ua/pkg/stack"
	"github.com/cloudwebrtc/go-sip-ua/pkg/ua"
	"github.com/cloudwebrtc/go-sip-ua/pkg/utils"
	"github.com/ghettovoice/gosip/log"
	"github.com/ghettovoice/gosip/sip"
	"github.com/ghettovoice/gosip/sip/parser"
	"github.com/ghettovoice/gosip/util"
)

var (
	logger      log.Logger
	udp         *rtp.RtpUDPStream
	tcp         *RtpTCPStream
	stackAgent  *stack.SipStack
	userAgent   *ua.UserAgent
	keepaliveSN int
	profile     *account.Profile
	recipient   sip.SipUri

	rport int
	lport int
)

func init() {
	logger = utils.NewLogrusLogger(log.InfoLevel, "Client", nil)
	lport = 30000
}

func createUdp() *rtp.RtpUDPStream {
	udp = rtp.NewRtpUDPStream("127.0.0.1", rtp.DefaultPortMin, rtp.DefaultPortMax, func(data []byte, raddr net.Addr) {
		logger.Infof("Rtp recevied: %v, laddr %s : raddr %s", len(data), udp.LocalAddr().String(), raddr)
		dest, _ := net.ResolveUDPAddr(raddr.Network(), raddr.String())
		logger.Infof("Echo rtp to %v", raddr)
		udp.Send(data, dest)
	})
	udp.Log().SetLevel(uint32(log.ErrorLevel))

	go udp.Read()
	return udp
}

func createTcp(port int, raddr *net.TCPAddr) *RtpTCPStream {
	tcp = NewRtpTCPStream("192.168.1.125", port, raddr, func(data []byte) {
		logger.Infof("Rtp recevied: %v, laddr %s : raddr %s", len(data), udp.LocalAddr().String(), udp.RemoteAddr().String())
	})
	tcp.Log().SetLevel(uint32(log.InfoLevel))

	go tcp.Read()
	return tcp
}

func main() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	stackAgent = stack.NewSipStack(&stack.SipStackConfig{
		UserAgent:  "Go Sip Client/example-client",
		Extensions: []string{"replaces", "outbound"},
		Dns:        "8.8.8.8"})

	stackAgent.OnRequest(sip.MESSAGE, handleStatus)
	stackAgent.OnRequest(sip.SUBSCRIBE, handleStatus)

	listen := "0.0.0.0:5080"
	sipServer := "192.168.1.125:5060"
	sipServerID := "34020000002000000001"
	sipDeviceID := "34020000001190000001"

	logger.Infof("Listen => %s", listen)
	if err := stackAgent.Listen("udp", listen); err != nil {
		logger.Panic(err)
	}

	userAgent = ua.NewUserAgent(&ua.UserAgentConfig{
		SipStack: stackAgent,
	})

	userAgent.InviteStateHandler = func(sess *session.Session, req *sip.Request, resp *sip.Response, state session.Status) {
		logger.Infof("InviteStateHandler: state => %v, type => %s", state, sess.Direction())
		//回复101
		//回复200
		switch state {
		case session.InviteReceived:
			bodys := strings.Split(sess.Request().Body(), "\r")
			ip := ""
			for _, v := range bodys {
				v = strings.ReplaceAll(v, "\n", "")
				if strings.HasPrefix(v, "c=") {
					ip = strings.Split(v, " ")[2]
				}
				if strings.HasPrefix(v, "m=") {
					ports := strings.Split(v, " ")[1]
					rport, _ = strconv.Atoi(ports)
				}
			}

			sdp := BuildLocalSdp("192.168.1.125", lport)
			sess.ProvideAnswer(sdp)
			sess.Accept(200)

			fmt.Printf("%s, %d\n", ip, rport)
		case session.Confirmed:
			go SendStream("192.168.1.125", rport, lport)
		case session.Canceled:
			fallthrough
		case session.Failure:
			fallthrough
		case session.Terminated:
			tcp.Close()
		}
	}

	userAgent.RegisterStateHandler = func(state account.RegisterState) {
		logger.Infof("RegisterStateHandler: user => %s, state => %v, expires => %v", state.Account.AuthInfo.AuthUser, state.StatusCode, state.Expiration)
		if state.StatusCode == 200 {
			SendDevice()
			go Keepalive(sipDeviceID, recipient)
		}
	}

	//本机信息
	uri, err := parser.ParseUri("sip:" + sipDeviceID + "@127.0.0.1:5080")
	if err != nil {
		logger.Error(err)
	}

	profile = account.NewProfile(uri.Clone(), "goSIP/example-client",
		&account.AuthInfo{
			AuthUser: "34020000001190000001",
			Password: "100",
			Realm:    "",
		},
		1800,
		stackAgent,
	)

	//客户端
	recipient, err = parser.ParseSipUri("sip:" + sipServerID + "@" + sipServer + ";transport=udp")
	if err != nil {
		logger.Error(err)
	}

	userAgent.SendRegister(profile, recipient, profile.Expires, nil)
	time.Sleep(time.Second * 3)

	//udp = createUdp()
	//udpLaddr := udp.LocalAddr()
	//sdp := mock.BuildLocalSdp(udpLaddr.IP.String(), udpLaddr.Port)
	//go userAgent.Invite(profile, uri, recipient, &sdp)

	logger.Infof("go sip client start over")

	<-stop
	logger.Infof("go sip stop ...")
	//register.SendRegister(0)
	userAgent.Shutdown()
}

func SendStream(rip string, rport, lport int) {
	logger.Infof("SendStream begin...")
	time.Sleep(time.Second * 3)
	rAddr := &net.TCPAddr{IP: net.ParseIP(rip), Port: rport}
	tcp = createTcp(lport, rAddr)
}

func Keepalive(deviceID string, recipient sip.SipUri) {
	logger.Info("start Keepalive")
	str := "<?xml version=\"1.0\" encoding=\"GB2312\" standalone=\"yes\" ?><Notify><CmdType>Keepalive</CmdType><SN>%d</SN><DeviceID>%s</DeviceID><Status>OK</Status></Notify>"
	for {
		time.Sleep(time.Second * 20)

		keepaliveSN++
		msg := fmt.Sprintf(str, keepaliveSN, deviceID)
		logger.Info("Keepalive: body => %s", msg)

		builder := sip.NewRequestBuilder()
		from := &sip.Address{
			DisplayName: sip.String{Str: profile.DisplayName},
			Uri:         profile.URI,
			Params:      sip.NewParams().Add("tag", sip.String{Str: util.RandString(8)}),
		}
		contact := profile.Contact()

		to := &sip.Address{
			Uri: profile.URI,
		}

		builder.SetMethod(sip.MESSAGE)
		builder.SetFrom(from)
		builder.SetTo(to)
		builder.SetContact(contact)
		builder.SetRecipient(recipient.Clone())
		builder.SetBody(msg)

		contentType := sip.ContentType("Application/MANSCDP+xml")
		builder.SetContentType(&contentType)

		req, err := builder.Build()
		if err != nil {
			logger.Error(err)
			continue
		}
		stackAgent.Request(req)
	}
}

func handleStatus(request sip.Request, tx sip.ServerTransaction) {
	logger.Infof("handleStatus: Request => %s, body => %s", request.Short(), request.Body())
	response := sip.NewResponseFromRequest(request.MessageID(), request, 200, "OK", "")

	if viaHop, ok := request.ViaHop(); ok {
		var (
			host string
			port sip.Port
		)
		host = viaHop.Host
		if viaHop.Params != nil {
			if received, ok := viaHop.Params.Get("received"); ok && received.String() != "" {
				host = received.String()
			}
			if viaHop.Port != nil {
				port = *viaHop.Port
			} else if rport, ok := viaHop.Params.Get("rport"); ok && rport != nil && rport.String() != "" {
				if p, err := strconv.Atoi(rport.String()); err == nil {
					port = sip.Port(uint16(p))
				}
			} else {
				port = sip.DefaultPort(request.Transport())
			}
		}

		dest := fmt.Sprintf("%v:%v", host, port)
		response.SetDestination(dest)
	}

	tx.Respond(response)
}

func SendDevice() {
	keepaliveSN++
	sipDeviceID := "34020000001190000001"
	msg := GetResponseXmlStr(keepaliveSN, sipDeviceID, 1, sipDeviceID, "摄像机1", "video-2023", "whxd")
	builder := sip.NewRequestBuilder()
	from := &sip.Address{
		DisplayName: sip.String{Str: profile.DisplayName},
		Uri:         profile.URI,
		Params:      sip.NewParams().Add("tag", sip.String{Str: util.RandString(8)}),
	}
	contact := profile.Contact()

	to := &sip.Address{
		Uri: profile.URI,
	}

	builder.SetMethod(sip.MESSAGE)
	builder.SetFrom(from)
	builder.SetTo(to)
	builder.SetContact(contact)
	builder.SetRecipient(recipient.Clone())
	builder.SetBody(msg)

	contentType := sip.ContentType("Application/MANSCDP + xml")
	builder.SetContentType(&contentType)

	req, err := builder.Build()
	if err != nil {
		logger.Error(err)
	}
	stackAgent.Request(req)
}
