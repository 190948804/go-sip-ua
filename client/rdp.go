package main

import (
	"errors"
	"net"
	"time"

	"github.com/190948804/go-sip-ua/pkg/utils"
	"github.com/ghettovoice/gosip/log"
	"github.com/pixelbender/go-sdp/sdp"
)

var (
	ErrPort = errors.New("invalid port")
)

const (
	DefaultPortMin = 30000
	DefaultPortMax = 65530
)

type RtpTCPStream struct {
	conn     *net.TCPConn
	stop     bool
	onPacket func(pkt []byte)
	laddr    *net.TCPAddr
	raddr    *net.TCPAddr
	logger   log.Logger
}

func NewRtpTCPStream(bind string, port int, raddr *net.TCPAddr, callback func(pkt []byte)) *RtpTCPStream {

	logger := utils.NewLogrusLogger(log.DebugLevel, "Media", nil)
	lAddr := &net.TCPAddr{IP: net.ParseIP(bind), Port: port}
	conn, err := net.DialTCP("tcp", lAddr, raddr)
	if err != nil {
		logger.Errorf("ListenUDP: err => %v", err)
		return nil
	}

	return &RtpTCPStream{
		conn:     conn,
		stop:     false,
		onPacket: callback,
		laddr:    lAddr,
		logger:   logger,
	}
}

func (r *RtpTCPStream) Log() log.Logger {
	return r.logger
}

func (r *RtpTCPStream) RemoteAddr() *net.TCPAddr {
	return r.raddr
}

func (r *RtpTCPStream) LocalAddr() *net.TCPAddr {
	return r.laddr
}

func (r *RtpTCPStream) Close() {
	r.stop = true
	r.conn.Close()
}

func (r *RtpTCPStream) Send(pkt []byte) (int, error) {
	r.Log().Debugf("Send to %v, length %d", r.raddr.String(), len(pkt))
	return r.conn.Write(pkt)
}

func (r *RtpTCPStream) Read() {
	r.Log().Infof("Read")

	buf := make([]byte, 1500)
	for {
		if r.stop {
			r.Log().Infof("Terminate: stop rtp conn now!")
			return
		}
		n, err := r.conn.Read(buf)
		if err != nil {
			r.Log().Warnf("RTP Conn [%v] refused, err: %v, stop now!", r.raddr, err)
			return
		}
		r.Log().Tracef("Read rtp from: %v, length: %d", r.raddr.String(), n)
		if !r.stop {
			r.onPacket(buf[0:n])
		}
	}
}

func BuildLocalSdp(host string, port int) string {
	sdp := &sdp.Session{
		Origin: &sdp.Origin{
			Username:       "-",
			Address:        host,
			SessionID:      time.Now().UnixNano() / 1e6,
			SessionVersion: time.Now().UnixNano() / 1e6,
		},
		Timing: &sdp.Timing{Start: time.Time{}, Stop: time.Time{}},
		//Name: "Example",
		Connection: &sdp.Connection{
			Address: host,
		},
		//Bandwidth: []*sdp.Bandwidth{{Type: "AS", Value: 117}},
		Media: []*sdp.Media{
			{
				//Bandwidth: []*sdp.Bandwidth{{Type: "TIAS", Value: 96000}},
				Connection: []*sdp.Connection{{Address: host}},
				Mode:       sdp.SendRecv,
				Type:       "video",
				Port:       port,
				Proto:      "TCP/RTP/AVP",
				Format: []*sdp.Format{
					{Payload: 0, Name: "PCMU", ClockRate: 8000},
					{Payload: 8, Name: "PCMA", ClockRate: 8000},
					//{Payload: 18, Name: "G729", ClockRate: 8000, Params: []string{"annexb=yes"}},
					{Payload: 106, Name: "telephone-event", ClockRate: 8000, Params: []string{"0-16"}},
				},
			},
		},
	}
	return sdp.String()
}
