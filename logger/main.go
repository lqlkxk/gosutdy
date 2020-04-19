package main

import "github.com/lqlkxk/logger/mylog"

func main() {
	log := mylog.GetLogger("info")
	log.SetFilePath("D:/logs")
	log.SetMaxSiza(100)
	//mylog.SetFilePath()
	log.Info("[测试输出]:%d", 111, 222)
	log.Debug("测试")
}
