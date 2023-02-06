package tool

import "encoding/json"

func ToJsonString(v interface{}) string {
    if ret, err := json.Marshal(v); err != nil {
        return err.Error()
    } else {
        return string(ret)
    }
}
