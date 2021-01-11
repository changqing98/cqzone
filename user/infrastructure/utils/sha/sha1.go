package sha

import (
    "crypto/sha1"
    "encoding/hex"
    "io"
)

func /**/Encrypt(content string) string {
    h := sha1.New()
    io.WriteString(h, content)
    return hex.EncodeToString(h.Sum(nil))
}
