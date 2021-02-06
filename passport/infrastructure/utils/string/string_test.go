package string

import (
    "fmt"
    "testing"
)

func TestRandomString(t *testing.T) {
    fmt.Print(RandomString(5))
}

func TestRandomIntString(t *testing.T) {
    fmt.Println(RandomIntString(3))
}
