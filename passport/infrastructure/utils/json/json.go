package json

import "encoding/json"

func ToJsonString(obj interface{}) string {
    result, err := json.Marshal(obj)
    if err != nil {
        panic(err)
    }
    return string(result)
}
