package clienttcpb


import (
	"boss2_ledscreen/config"
	"boss2_ledscreen/data"
	"boss2_ledscreen/log"
	"boss2_ledscreen/web"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)


var conn net.Conn
func TcpClinet() {
	log.Info("正在连接计数器显示屏B")
	hostname:= fmt.Sprintf("%v:%v", config.Syscfg.CtInfoB.ControlIP, config.Syscfg.CtInfoB.ControlPort)
	for i := 0; i < 3; i++ {
		var err error
		conn, err = net.Dial("tcp", hostname )
		if err != nil {
			log.Errorf("显示屏B TCP connect failed,err:%v", err)
			log.Warn("显示屏B Reconnect after 10 seconds")
			if i == 2 {
				log.Errorf("显示屏B 3 failed connections, please check the hardware connection or network connection")
				os.Exit(1)
			}
			time.Sleep(time.Duration(10)*time.Second)
		}else {
			defer conn.Close()
			log.Infof("显示屏B Tcp connected server:%v", hostname)
			break
		}
	}

	go ReceiveDataFromServer()

	for {
		select {
		case remainspace, ok := <- web.DataInchanB:
			if ok {
				buf := fmt.Sprintf("%v%d", config.Syscfg.CtInfoB.Texthead, remainspace)
				log.Infof("显示屏B:%v", buf)
				strhead, err := data.Utf8ToGbk([]byte(buf))
				if err != nil {
					log.Errorf("UTF-8 -> GBK failed, err：%v", err)
				}

				rebuf, errnum := data.Infotodata(strhead, config.Syscfg.CtInfoB)
				if errnum < 0 {
					log.Error("显示屏B data.Infotodata failed")
				}else {

					_, err := conn.Write(rebuf)
					if err != nil {
						log.Errorf("显示屏B Data send failed,err:%v", err)
					}
					//log.Infof("Sent remainspace: %v", remainspace)
				}
			}

		}
	}

}



func ReceiveDataFromServer() {
	rdata := make([]byte, 1024)
	reader := bufio.NewReader(conn)
	for {
		_, err := reader.Read(rdata)
		if err != nil {
			if err == io.EOF {
				log.Errorf("显示屏B Disconnected from the tcp network,err:%v", err)
				conn.Close()
			}
		}
	}
}



