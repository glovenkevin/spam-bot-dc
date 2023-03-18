package helper

import "encoding/base64"

func ToBase64(bb []byte) string {
	if len(bb) == 0 {
		return ""
	}

	return base64.StdEncoding.EncodeToString(bb)
}
