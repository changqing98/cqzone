package sha

import (
    "testing"
)

func TestEncrypt(t *testing.T) {
    content := "test_password"
    result := EncryptWithSalt(content)
    if 57 != len(result){
        t.Errorf("Wrong encrypted content length")
    }

    if !Verify(result, content) {
        t.Errorf("Wrong encrypted result, expected = %s", result)
    }
}
