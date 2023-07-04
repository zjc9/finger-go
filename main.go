package main

import (
	"G2/cmd"
	"github.com/kataras/golog"
	"time"
)

var log = golog.Child("[main]")

func main() {
	start := time.Now()
	cmd.Init()
	t := time.Now().Sub(start)
	log.Infof("[*] 扫描结束,耗时: %s\n", t)
	//get, err := http.Get(`https://bbs.colg.cn/`)
	//if err != nil {
	//	log.Error(err)
	//
	//}
	//body := get.Body
	////for k, v := range body {
	////	fmt.Println(k, v)
	////}
	//buf := new(bytes.Buffer)
	//_, err = buf.ReadFrom(body)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(buf.String())

}
