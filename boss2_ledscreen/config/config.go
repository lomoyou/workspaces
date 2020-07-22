package config

import (
	"boss2_ledscreen/log"
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"os"
)
type LEDControler struct {
	ControlIP string `json:"controlIP"`
	CounterName string `json:"counterName"`
	ControlPort int `json:"controlPort"`
	Ledwidth uint8 `json:"ledwidth"`
	Ledheight uint8 `json:"ledheight"`
	Texthead string `json:"texthead"`
	Dismode string `json:"leddismode"`
	TextColor string `json:"textColor"`
	Timeinterval int `json:"timeinterval"`
}


type Config struct{
	CtInfoA LEDControler `json:"ledControlerA"`
	CtInfoB LEDControler `json:"ledControlerB"`
	TurnLEDControler bool `json:"turnledControlB"`
	WebAddr string  `json:"webAddr"`
}


var Syscfg Config

func init() {
	var filename string
	flag.StringVar(&filename, "f", "./config.json", "默认配置文件为./config.json")
	flag.Parse()

	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		log.Errorf("open file config.json err:%v", err)
	}
	//读取文件内容
	reader := bufio.NewReader(file)
	data := make([]byte, 1024)
	n, err := reader.Read(data)
	if err != nil {
		if err == io.EOF {
			//关闭文件
			err = file.Close()
			if err != nil {
				log.Errorf("config.cfg close err:%v", err)
			}
		}
	}

	if (n > 0) {

		err = json.Unmarshal(data[:n], &Syscfg)
		if err != nil {
			log.Errorf("config unmarshal err:%v", err)
		}else {
			log.Infof("config:%v",Syscfg)
		}

	}
}

