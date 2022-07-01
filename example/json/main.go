package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    type Hoge struct {
        Field1 string `json:"field1"`

        // "omitempty" オプションは、フィールドの値が空の値の場合
        // エンコーディングするときに省略される
        Field2 string `json:"field2,omitempty"`

        // "-" は、常にフィールドは省略される
        Field3 string `json:"-"`
    }

    jsonBlob := []byte(`[
    {"field1": "someting1", "field2": "something2", "field3": "something3"},
    {"field1": null,        "field2": null,         "field3": null},
    {"field1": "hello",        "field2": null,         "field3": 999}
]`)

    var hoges []Hoge
    err := json.Unmarshal(jsonBlob, &hoges)
    if err != nil {
        fmt.Println("error:", err)
        return
    }

    // デコード時には`omitempty`は問題なくデコードされ、
    // `"-"`タグは省略される（フィールドの値は初期値）ようです。
    fmt.Printf("%+v\n", hoges)
	// fmt.Println(hoges[0].Field3)
    //=> [{Field1:someting1 Field2:something2 Field3:} {Field1: Field2: Field3:}]
}
