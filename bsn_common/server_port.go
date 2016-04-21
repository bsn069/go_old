package bsn_common

import (
// "bsn/bsn_common"
// "time"
)

// Gate 		2xxxx xxxx(gate id)
func GatePort(id uint32) uint16 {
	return 20000 + uint16(id)
}

// GateConfig 	300xx xx(id)
func GateConfigPort(id uint32) uint16 {
	return 30000 + uint16(id%100)
}

// Server 4aabb aa(server type) bb(server id)
func ServerPort(serverType, id uint32) uint16 {
	return uint16(40000 + (serverType%100)*100 + id%100)
}
