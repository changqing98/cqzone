package snowflake

import (
    "fmt"
    "testing"
)

func TestGenerate(t *testing.T) {
    id := Generate()
    fmt.Print(id)
}
