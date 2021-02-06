package string

import (
    "math/rand"
    "time"
)

func IsEmpty(str string) bool {
    return len(str) == 0
}

var digitalRunes = []rune("0123456789")

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomIntString(n int) string {
    rand.Seed(time.Now().Unix())
    b := make([]rune, n)
    for i := range b {
        b[i] = digitalRunes[rand.Intn(len(digitalRunes))]
    }
    return string(b)
}

func RandomString(n int) string {
    rand.Seed(time.Now().Unix())
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
