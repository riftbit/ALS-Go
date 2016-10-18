package main

import (
	"encoding/json"
	"runtime"
	"fmt"
)

func PrintObject(v interface{}) {
	res2B, _ := json.Marshal(v)
	fmt.Println(string(res2B))
}

func GetFuncName(level int) string {
	pc, _, _, _ := runtime.Caller(level)
	return runtime.FuncForPC(pc).Name()
}

func GetLineCall(level int) int {
	_, _, line, _ := runtime.Caller(level)
	return line
}

func GetFileCall(level int) string {
	_, file, _, _ := runtime.Caller(level)
	return file
}