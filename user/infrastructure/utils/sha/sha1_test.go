package sha

import (
    "fmt"
    "testing"
)

func TestEncrypt(t *testing.T) {
    content := "test_password"
    result := Encrypt(content)
    fmt.Printf(result)
}
