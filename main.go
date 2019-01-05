package asura

import (
    "time"
)

var CurrentDate string
var CurrentTime string

func init() {
    CurrentDate = Date("2006-01-02")
    CurrentTime = Date("2006-01-02 15:04:05")
}
func Date(format string) string {
    if 0 == len(format) {
        format = "2006-01-02 15:04:05"
    }
    return time.Now().Format(format)
}
