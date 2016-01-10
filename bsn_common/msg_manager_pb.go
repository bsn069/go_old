package bsn_common

import (
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

type IMsgManager_Pb interface {
	MsgManager_MakeMsg(u16MsgType uint16) (proto.Message, error)
	MsgManager_ProcMsg(u16MsgType uint16, pbMsg interface{}) error
}

type SMsgManager_Pb struct {
	m_callBack IMsgManager_Pb
	m_conn     net.Conn
}

func NewMsgManager(callBack IMsgManager_Pb) *SMsgManager_Pb {
	return &SMsgManager_Pb{m_callBack: callBack}
}

func (this *SMsgManager_Pb) Close() {
	if nil != this.m_conn {
		this.m_conn.Close()
		this.m_conn = nil
	}
}

func (this *SMsgManager_Pb) SetConn(conn net.Conn) {
	this.m_conn = conn
}

func (this *SMsgManager_Pb) Send(u16MsgType uint16, sMsg proto.Message) error {
	// log.Println("Send")

	data, err := proto.Marshal(sMsg)
	if err != nil {
		log.Println("Marshal error: ", err)
		return err
	}
	// log.Println("data: ", data)

	err = WriteMsg(this.m_conn, u16MsgType, data)
	if err != nil {
		log.Println("WriteMsg error: ", err)
		return err
	}

	return err
}

func (this *SMsgManager_Pb) Recv() error {
	// log.Println("Recv")

	err, u16MsgType, data := ReadMsg(this.m_conn)
	if err != nil {
		if err.Error() == "EOF" {
			log.Println("connect disconnect")
		} else {
			log.Println("ReadMsg error: ", err)
		}
		return err
	}
	// log.Println("data=", data)

	var msg proto.Message
	msg, err = this.m_callBack.MsgManager_MakeMsg(u16MsgType)
	if err != nil {
		log.Println("MsgManager_MakeMsg error: ", err)
		return err
	}

	err = proto.Unmarshal(data, msg)
	if err != nil {
		log.Println("unmarshaling error: ", err)
		return err
	}
	// log.Println("msg=", msg)

	err = this.m_callBack.MsgManager_ProcMsg(u16MsgType, msg)
	if err != nil {
		if err.Error() != "" {
			log.Println("MsgManager_ProcMsg error: ", err)
		} else {
			log.Println("disconnect connect")
		}
		return err
	}

	return err
}
