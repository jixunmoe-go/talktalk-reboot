package router

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func SHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	result := hex.EncodeToString(hash[:])
	return strings.ToLower(result)
}

func HashPassword(username string, password string, csrf csrfParams) string {
	sb := strings.Builder{}
	sb.WriteString(username)
	sb.WriteString(base64.StdEncoding.EncodeToString([]byte(SHA256(password))))
	sb.WriteString(csrf.Param)
	sb.WriteString(csrf.Token)
	return SHA256(sb.String())
}
