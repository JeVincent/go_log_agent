package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"kafka"
	"tailf"
)

func main()  {
	filename := "/Users/xujianhao/go/go_log_agent/conf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf fail, err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}

	logs.Debug("load conf succ, config %v", appConfig)

	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}
	logs.Debug("initialize all succ")
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}

	logs.Info("program exited")
}
