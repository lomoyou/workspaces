package start

import (
	"boss2_ledscreen/log"
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
)

//通过更改注册表，实现开机自启动
func StartProgram() {
	key, exists, err := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		log.Errorf("registry.CreateKey failed,err:%v", err)
	}
	defer key.Close()
	if exists {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Errorf("Get FilePath failed, err:%v", err)
		}
		log.Infof("FilePath:%v", dir)
		key.SetStringValue("BossLedKey", dir + "\\boss2_ledscreen.exe")
	}


}