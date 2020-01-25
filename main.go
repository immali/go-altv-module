package main

import "C"
import goaltv "github.com/immali/go-altv"

var _m *goaltv.Module = goaltv.NewModule()

func main() {}

//export altMain
func altMain(core uintptr) bool {
	_m.SetCore(core)

	if _m.Start() {
		_m.LogInfo(">> GO Module Started")
		return true
	} else {
		return false
	}
}

//export GetSDKVersion
func GetSDKVersion() int {
	return _m.GetVersion()
}
