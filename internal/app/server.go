// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

// ###############################$
// #########################&%$###$
// #################@%%##&|!|&####$
// ##############$|!!%$|!!!|@#####$
// ###########%!!!!!!!!;;!%#######$
// ########%!!!;;;!!;;;;!$########$
// #####&|!!;;;;;;;;;;!!||!$######$
// ###&|!!;;;;;;;;;;;!!!!!&#######$
// ##&!!;;;;;;;;;;;;;;;!|@########$
// ##%!;;;:'``':;;;;;;!%@#########$
// ##$!;;;:`..`:;;;;!!$###########$
// ##@|!;;;::::;;;!!|@############$
// ####%!!;;;;;;!!|&##############$
// ######&%!!!!|%@################$
// ###############################$
// LOGO IMAGE : https://avatars0.githubusercontent.com/u/48612456?s=200&v=4

package app

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/fast-mq/server/internal/conf"
	"github.com/higker/logker"
)

var (
	// data buffer
	buffers        [1024]byte
	errCreateServe = errors.New("create tcp server error")
	errAcceptConn  = errors.New("tcp accept error")
	errClientClose = "Close->|"
	errAddressMsg  = "Your ip address not at allow list."
)

// MQServer is message queue server.
type MQServer struct {
	// Bind IP
	Address string
	// Bind Port 0-65535
	Port uint16
	// Allow access IP
	AllowIP []string
	// Auth password
	Password string
}

// NewMQServer create server
func NewMQServer(option *conf.Option) *MQServer {
	return &MQServer{
		Address: option.Address,
		Port: func() uint16 {
			parseUint, err := strconv.ParseUint(option.Port, 10, 16)
			if err != nil {
				logker.Error("bind port error:", err)
				//os.Exit(1)
			}
			return uint16(parseUint)
		}(),
		AllowIP:  option.AllowIP,
		Password: option.Password,
	}
}

// Start START MqServer
func (mq *MQServer) Start() error {
	type this int
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", mq.Address, mq.Port))
	if err != nil {
		return errCreateServe
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			return errAcceptConn
		}
		go mq.checkClient(accept)
	}
}

func handleConn(con net.Conn) {
	for con != nil {
		n, _ := con.Read(buffers[:])
		msg := string(buffers[:n])
		logker.Warning("%v", msg)
		if msg == errClientClose {
			logker.Error("%s", "客户端连接已经关闭.")
			con.Close()
			return
		}
	}
}

// 验证客户端
func (ms *MQServer) checkClient(con net.Conn) {
	ip := strings.Split(con.RemoteAddr().String(), ":")[0]
	for _, allow := range ms.AllowIP {
		if allow == ip {
			go handleConn(con)
			return
		}
	}
	logker.Warning("%s", "There are client IP="+ip+" connection failed.")
	con.Write([]byte(errAddressMsg))
	con.Close()
}
