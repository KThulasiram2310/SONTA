package main

import (
    "github.com/thanhpk/randstr"
    "fmt"
)

func main() {

    StringToEncode := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    Encoding := base64.StdEncoding.EncodeToString([]byte(StringToEncode))
    fmt.Println(Encoding)                                        
}