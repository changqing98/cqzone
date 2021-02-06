package snowflake

import "github.com/bwmarrin/snowflake"

func Generate() int64 {
    node, err := snowflake.NewNode(1)
    if err != nil {
        panic(err)
    }
    id := node.Generate()
    return id.Int64()
}
