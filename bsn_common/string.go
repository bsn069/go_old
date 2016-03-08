package bsn_common

import (
	// "encoding/binary"
	"strings"
)

func StringIsUpper(this *string) bool {
	strUpper := strings.ToUpper(*this)
	return 0 == strings.Compare(strUpper, *this)
}
