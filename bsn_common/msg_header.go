package bsn_common

type SMsgHeader struct {
	m_u16Type uint16
	m_u16Len  uint16
}

const (
	cMsgHeader_Size uint16 = 4
)

func newMsgHeader(u16Type, u16Len uint16) *SMsgHeader {
	return &SMsgHeader{m_u16Type: u16Type, m_u16Len: u16Len}
}

func newMsgHeaderFromBytes(byData []byte) *SMsgHeader {
	return newMsgHeader(ReadUint16(byData), ReadUint16(byData[2:]))
}

func (this *SMsgHeader) bytes() []byte {
	var byRet = make([]byte, cMsgHeader_Size)
	WriteUint16(byRet, this.m_u16Type)
	WriteUint16(byRet[2:], this.m_u16Len)
	return byRet
}

func (this *SMsgHeader) getLen() uint16 {
	return this.m_u16Len
}

func (this *SMsgHeader) getType() uint16 {
	return this.m_u16Type
}
