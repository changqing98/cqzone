package sha

import (
    "crypto/sha1"
    "encoding/hex"
    StringUtils "github.com/changqing98/cqzone/user/infrastructure/utils/string"
    "io"
    "strings"
)

func EncryptWithSalt(content string) string {
    salt := StringUtils.RandomString(16)
    return encrypt(content, salt)
}

func encrypt(content string, salt string) string {
    h := sha1.New()
    appendedContent := salt + content
    _, _ = io.WriteString(h, appendedContent)
    return salt + "$" + hex.EncodeToString(h.Sum(nil))
}

func Verify(content string, source string) bool {
    salt := strings.Split(content, "$")[0]
    encryptedContent := encrypt(source, salt)
    return encryptedContent == content
}
