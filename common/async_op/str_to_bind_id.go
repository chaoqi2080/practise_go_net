package async_op

import "hash/crc32"

func Str2BindId(strVal string) int {
	v := int(crc32.ChecksumIEEE([]byte(strVal)))

	if v >= 0 {
		return v
	} else {
		return -v
	}
}
