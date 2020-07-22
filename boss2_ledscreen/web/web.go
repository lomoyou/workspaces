package web

import (
	"boss2_ledscreen/config"
	"boss2_ledscreen/log"
	"bytes"
	_ "context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "strconv"
	"time"
)

var DataInchanA = make(chan int, 1)
var DataInchanB = make(chan int, 1)

const (
	userinfo = "subin"
	methodinfo = "getemptyspace"
	versioninfo = "1.0"
	key = "866585558226OPKE8EFE9873"
)


func (post postinfo)Signinfo() string{
	str := fmt.Sprintf("data=%s&method=%s&requesttoken=%s&timestamp=%s&user=%s&version=%s",
		post.Data,
		post.Method,
		post.Requesttoken,
		post.Timestamp,
		post.User,
		post.Version)

	date := time.Now().Format("2006-01-02")

	buf :=fmt.Sprintf("&%v&key=%v", date, key)
	return EncodeMD5(str+buf)
}

func EncodeMD5(value string) string{
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func HttpPost(url string, data []byte, contentype string) (string, error){
	clinet := &http.Client{Timeout: 5*time.Second}
	resp, err := clinet.Post(url, contentype, bytes.NewBuffer(data))
	if err != nil {
		log.Errorf("ledscreen post failed,err:%v",err)
		return "", err
	}
	defer resp.Body.Close()

	result, _ :=ioutil.ReadAll(resp.Body)
	return string(result), nil

	//jsonstr, err := json.Marshal(data)
	//if err != nil {
	//	log.Errorf("json.Marshal failed, err:%v", err)
	//	return "", err
	//}
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
	//if err != nil {
	//	log.Errorf("http.NewRequest failed, err:%v", err)
	//	return "", err
	//}
	//req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//clienttcpa := http.Client{}
	//resp, err := clienttcpa.Do(req.WithContext(context.TODO()))
	//if err != nil {
	//	log.Errorf("clienttcpa.Do failed, err:%v", err)
	//	return "", nil
	//}
	//log.Errorf("status_code:%v", resp.StatusCode)
	//log.Errorf("status:%v", resp.Status)
	//respbytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Errorf("ioutil.ReadAll failed, err:%v", err)
	//	return "", nil
	//}
	//return string(respbytes), nil
}


func Parseresult(indata string) (int, error) {
	reqInfo := new(resultinfo)
	err := json.Unmarshal([]byte(indata), reqInfo)
	if err != nil {
		log.Errorf("postinfo json.Umarshal failed, err:%v", err)
		return 0, err
	}

	retdata := new(carinfo)
	if reqInfo.Code == 0 && reqInfo.Success {
		err = json.Unmarshal([]byte(reqInfo.Data), retdata)
		if err != nil {
			log.Errorf("carinfo json.Umarshal failed, err:%v", err)
			return 0, err
		}
	}

	return retdata.Empty_num, nil
}

func QuerySpaces(counterName string, Inchan chan int, interval int ) {
	for {
		reqpost := new(postinfo)
		reqpost.User = userinfo
		reqpost.Version = versioninfo
		reqpost.Method = methodinfo
		timedate := time.Now().UnixNano()
		timebuf := fmt.Sprintf("%v", timedate/1e6)
		reqpost.Timestamp = timebuf
		reqpost.Data = fmt.Sprintf(`{"parkcode":"", "counterName":"%v"}`, counterName);
		reqpost.Requesttoken = generateSubId()
		signinfo := reqpost.Signinfo()
		reqpost.Sign = signinfo
		jsonstr,err := json.Marshal(reqpost)
		if err != nil {
			log.Warnf("json.marshal failed, err:%v", err)
		}
		//log.Infof("Post request info:%v", string(jsonstr))

		responseinfo, err := HttpPost(config.Syscfg.WebAddr, jsonstr, "application/json")
		if err != nil {
			log.Errorf("HttpPost failed, err:%v", err)
		}
		//log.Infof("Post response info:%v",responseinfo)
		num, err := Parseresult(responseinfo)
		if err != nil {
			log.Errorf("Parseresult failed, err:%v", err)
		}else{
			go func() {
				Inchan <- num
			}()

		}

		time.Sleep(time.Duration(interval)*time.Minute)
	}
}
func generateSubId() string{
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-")
	b :=make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
