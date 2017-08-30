package main

import (
	"encoding/json"
	//"time"

	log "github.com/thinkboy/log4go"
	"golang.org/x/net/websocket"
	"fmt"
)

func initWebsocket() {
	origin := "http://" + Conf.WebsocketAddr + "/sub"
	url := "ws://" + Conf.WebsocketAddr + "/sub"
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Error("websocket.Dial(\"%s\") error(%v)", Conf.WebsocketAddr, err)
		return
	}
	proto := new(Proto)
	proto.Ver = 1
	// auth
	// test handshake timeout
	// time.Sleep(time.Second * 31)
	proto.Operation = OP_AUTH
	seqId := int32(0)
	proto.SeqId = seqId
	proto.Body = []byte("{\"uid\":12345,\"roomid\":3001}")
	if err = websocketWriteProto(conn, proto); err != nil {
		log.Error("websocketWriteProto() error(%v)", err)
		return
	}

	var props []Proto
	if err = websocketReadProto(conn, props); err != nil {
		log.Error("websocketReadProto() error(%v)", err)
		return
	}
	log.Debug("auth ok, proto: %v", proto.Body)
	seqId++
	proto1 := new(Proto)
	proto1.Operation = OP_HEARTBEAT
	proto1.SeqId = seqId
	proto1.Body = nil
	if err = websocketWriteProto(conn, proto1); err != nil {
		log.Error("tcpWriteProto() error(%v)", err)
		return
	}
	p,err := readWs(conn)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(p[0].Body))
	seqId++
	proto1.Operation = OP_UNREAD_NUM
	proto1.SeqId = seqId
	proto1.Body = []byte("{}")
	if err = websocketWriteProto(conn, proto1); err != nil {
		log.Error("tcpWriteProto() error(%v)", err)
		return
	}
	p,err = readWs(conn)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(p[0].Operation))
	fmt.Println(string(p[0].Body))
	select {

	}
}

func websocketReadProto(conn *websocket.Conn, p []Proto) error {

	return websocket.JSON.Receive(conn,&p)
}

func readWs(conn *websocket.Conn)([]Proto,error){
	var p []Proto
	err := websocket.JSON.Receive(conn,&p)
	return p,err
}

func websocketWriteProto(conn *websocket.Conn, p *Proto) error {
	if p.Body == nil {
		p.Body = []byte("{}")
	}
	msg, _ := json.Marshal(p)
	log.Debug("%s", string(msg))
	return websocket.JSON.Send(conn, p)
}
