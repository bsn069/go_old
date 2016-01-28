package bsn_msg

// import (
// 	// "fmt"
// 	// "github.com/bsn069/go/bsn_common"
// 	// "runtime"
// 	// "time"
// 	"net"
// )

// type sNetMsg struct {
// 	m_conn net.Conn
// }

// func (this *sNetMsg) Write(u16Type uint16, byMsg []byte) (err error) {
// 	var u16Len uint16 = uint16(len(byMsg))
// 	msgHeader := NewMsgHeader(u16Type, u16Len)
// 	byDatas := append(msgHeader.Serialize(), byMsg)
// 	err = conn.Write(byDatas)
// }

// func (this *sNetMsg) Read() (u16Type uint16, byMsgBody []byte, err error) {
// 	byMsgHeader := make([]byte, CMsgHeader_Size)
// 	_, err = conn.Read(byMsgHeader)
// 	if err != nil {
// 		return
// 	}
// 	msgHeader := NewMsgHeader(0, 0)
// 	msgHeader.DeSerialize(byMsgHeader)

// 	u16Len := msgHeader.Len()
// 	if u16Len > 0 {
// 		byMsgBody = make([]byte, u16Len)
// 		_, err = conn.Read(byMsgBody)
// 		if err != nil {
// 			return
// 		}
// 	}

// 	return msgHeader.Type(), byMsgBody, nil
// }
