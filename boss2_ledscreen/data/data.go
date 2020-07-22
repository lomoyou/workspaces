package data

import (
	"boss2_ledscreen/config"
	"boss2_ledscreen/crc"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

const (
	VAR_STATIC_DISPLAY = 0x01
	VAR_MOVELIFT_DISPLAY = 0x03
)
/*-----------------------------------------------------------------------------
Function	: PackingFrame
Description	: 数据封帧
Input		: indata - 封帧前的数据
			  datasize  - 封帧前的数据长度
Return		: []byte 封帧后数据的
			: int 封帧后的数据长度
Note		: None
------------------------------------------------------------------------------*/
func ParkingFrame(indata []byte, datasize int)  ([]byte, int) {
	var frame []byte
	var outdata []byte
	frame = make([]byte, 8)
	for i :=0; i < len(frame) + datasize; i++ {
		if i< len(frame){
			outdata = append(outdata, 0xA5)
		}else {
			outdata = append(outdata, indata[i-len(frame)])
		}
	}
	outdata = append(outdata, 0x5A)
	return outdata, len(frame) + datasize
}

/*-----------------------------------------------------------------------------
Function	: ReturnData
Description	: 反转义函数
Input		: data -反转义前的数据
			  datalen - 反转义前的数据长度
Output		: None
Return		: []byte 反转义后数据
			: int 反转义后的数据长度
Note		: None
------------------------------------------------------------------------------*/
func ReturnData(indata []byte, datalen int) ([]byte, int) {
	var j int = 0
	var outdata []byte
	outdata = make([]byte, datalen)
	for i := 0; i < datalen; i++ {
		if indata[i] == 0xA6 && indata[i+1] == 0x02 {
			outdata[j] = 0xA5
			j++
		}else if indata[i] == 0xA6 && indata[i+1] == 0x01 {
			outdata[j] = 0xA6
			j++
		}else if indata[i] == 0x5B && indata[i+1] == 0x02 {
			outdata[j] = 0x5A
			j++
		}else if indata[i] == 0xA6 && indata[i+1] == 0x01 {
			outdata[j] = 0x5B
			j++
		}else {
			outdata[j] = indata[i]
			j++
		}
	}
	return outdata, j
}

/*-----------------------------------------------------------------------------
Function	: TurnData
Description	: Phy1层数据字符转义
Input		: indata -转义前的数据
			  datalen - 转义前的数据长度
Output		: None
Return		: []byte 转义后的数据
			： int 转义后数据长度
Note		: None
------------------------------------------------------------------------------*/
func TurnDate(indata []byte , datalen int) ([]byte, int) {
	var outdata []byte
	var i = 0

	for i = 0; i < datalen; i++ {
		if indata[i] == 0xA5 {
			//outdata[j] = 0xA6
			//outdata[j+1] = 0x02
			//j=j+2
			outdata = append(outdata, 0xA6)
			outdata = append(outdata, 0x02)
		}else if indata[i] == 0xA6 {
			//outdata[j] = 0xA6
			//outdata[j+1] = 0x01
			//j=j+2
			outdata = append(outdata, 0xA6)
			outdata = append(outdata, 0x01)
		} else if indata[i] == 0x5A {
			//outdata[j] = 0x5B
			//outdata[j+1] = 0x02
			//j=j+2
			outdata = append(outdata, 0x5B)
			outdata = append(outdata, 0x02)
		} else if indata[i] == 0x5B {
			//outdata[j] = 0x5B
			//outdata[j+1] = 0x01
			//j=j+2
			outdata = append(outdata, 0x5B)
			outdata = append(outdata, 0x01)
		}else{
			//fmt.Printf("j=%v, i=%v\n",j,i )
			//outdata[j] = indata[i]
			//j++
			outdata = append(outdata, indata[i])
		}
	}

	return outdata, len(outdata)
}

/*-----------------------------------------------------------------------------
Function	: DataPhyOne
Description	: phy1层数据组装
Input		: packetheaderdata -包头数据
			  datalen1 -包头数据长度
              datafiled - 数据域数据
              datalen2 - 数据域数据长度
Output		: dataphy1data -封包的数据（含包校验）
Return		: 封包后的数据长度
Note		: None
------------------------------------------------------------------------------*/
func DataPhyOne(headdata []byte, headlen int, fileddata []byte, filedlen int) ([]byte, int) {
	var outdata []byte
	for i :=0; i < headlen; i ++ {
		outdata = append(outdata, headdata[i])
	}

	for j :=0; j < filedlen; j++ {
		outdata = append(outdata, fileddata[j])
	}

	crcdata := crc.CheckCRC(outdata)
	tmep := crc.Uint16ToBytes(crcdata)
	for _,v := range tmep {
		outdata = append(outdata, v)
	}

	return outdata, headlen+filedlen+2
}

/*-----------------------------------------------------------------------------
Function	: MakePackHerDate
Description	: 包头组装
Input		: DataLen - 数据位长度
Output		: None
Return		: []byte 包头数据
			: int 包头数据长度
Note		: None
------------------------------------------------------------------------------*/
func MakeHeadDate(datalen uint16) ([]byte, int) {
	var outdata []byte
	outdata = make([]byte, 12)

	outdata[0] = 0x01
	outdata[1] = 0x00
	outdata[2] = 0x00
	outdata[3] = 0x80
	outdata[4] = 0x00
	outdata[5] = 0x00
	outdata[6] = 0x00
	outdata[7] = 0x00
	outdata[8] = 0x00
	outdata[9] = 0x01
	outdata[10] = 0x61
	outdata[11] = 0x02
	tmep := crc.Uint16ToBytes(datalen)

	for _, v := range tmep{
		outdata = append(outdata, v)
	}

	return outdata, 14
}

/*-----------------------------------------------------------------------------
Function	: PackData
Description : 数据封包
Input		: datalen - 数据长度
			  indata - 源数据
Output		: None
Return		: []byte 包头数据
			: int 包头数据长度
Note		: 数据长度和源数据均为requestmessagepacking函数组装之后的数据
------------------------------------------------------------------------------*/
func PacKData(datalen int, indata []byte) ([]byte,int) {
	//组装要发送的数据
	dmdata, dmlen := MakeHeadDate(uint16(datalen))
	dDdata, dDlen := DataPhyOne(dmdata, dmlen, indata, datalen)
	dTdata, dTlen := TurnDate(dDdata, dDlen)
	dPdata, dPlen := ParkingFrame(dTdata, dTlen)

	return dPdata, dPlen
}

/*-----------------------------------------------------------------------------
Function	: AreaDataFormat
Description : 区域文本数据显示
Input		: runmode - 运行模式
			  vate - 显示速度
			  datalen - 文本数据长度
			  indata - 文本数据
Output		: None
Return		: []byte 包头数据
			: int 包头数据长度
Note		: 该函数对要在区域内显示的数据进行组装,单行显示
------------------------------------------------------------------------------*/
func AreaDataFormat(width byte, height byte, dismode byte, datalen int32, indata []byte ) ([]byte, int) {
	var outdata []byte
	outdata = make([]byte, 23)
	outdata[5] = width/8    //区域宽度，默认以字节(8 个像素点)为单位 高字节最高位为 1 时，表示以像素点为单位
	outdata[6] = 0x00
	outdata[7] = height 	//区域高度，默认以像素点为单位
	outdata[8] = 0x00	//区域高度，以像素点为单位
	outdata[9] = 0x00 	//动态区域编号
	outdata[10] = 0x00	//行间距
	outdata[11] = 0x00  //运行模式
	outdata[12] = 0x01
	outdata[16] = 0x60   //对齐方式
	outdata[17] = 0x01   //单行显示
	outdata[18] = 0x01   //手动换行
	outdata[19] = dismode  //显示方式
	outdata[20] = 0x00
	outdata[21] = 0x18  //滚动速度
	outdata[22] = 0x00

	temp := crc.Int32ToBytes(datalen)
	for _, v := range temp {
		outdata = append(outdata, v)
	}

	for _, val := range indata {
		outdata = append(outdata, val)
	}

	return outdata, 27+int(datalen)
}

/*-----------------------------------------------------------------------------
Function	: SendShowData
Description : 发送显示数据
Input		: datalen - 数据长度
			  indata - 文本数据
Output		: None
Return		: []byte 返回数据
			: int 返回数据长度
Note		: None
------------------------------------------------------------------------------*/
func SendShowData(datalen uint16, indata []byte) ([]byte, int) {
	outdata := make([]byte, 8)
	outdata[0] = 0xA3
	outdata[1] = 0x06
	outdata[2] = 0x01
	outdata[3] = 0x01
	outdata[4] = 0x00
	outdata[5] = 0x01 //删除区域个数
	outdata[6] = 0x00 //删除区域的ID号
	outdata[7] = 0x01 //更新区域的个数

	temp := crc.Uint16ToBytes(datalen)
	for _, v := range temp {
		outdata = append(outdata, v)
	}

	for _, val := range indata {
		outdata = append(outdata, val)
	}

	return outdata, 10+int(datalen)
}

/*-----------------------------------------------------------------------------
Function	: TextDataPack
Description : 文本数据封包
Input		: colour - 文本颜色
			  intdata - 文本数据
Output		: None
Return		: []byte 返回数据
			: int 返回数据长度
Note		: None
------------------------------------------------------------------------------*/
func TextDataPack(color string, indata []byte, dismode string, width byte, height byte) ([]byte, int) {
	rdata := `\FO100\C1`
	gdata := `\FO100\C2`
	ydata := `\FO100\C3`
	var outdata []byte

	if color == "red" {
		for _, v := range []byte(rdata) {
			outdata = append(outdata, v)
		}
	}else if color == "green" {
		for _, v := range []byte(gdata) {
			outdata = append(outdata, v)
		}
	}else if color == "yellow" {
		for _, v := range []byte(ydata) {
			outdata = append(outdata, v)
		}
	}else {
		for _, v := range []byte(rdata) {
			outdata = append(outdata, v)
		}
	}

	for _, v := range (indata) {
		outdata = append(outdata, v)
	}
	var mode byte
	if dismode == "static" {
		mode = VAR_STATIC_DISPLAY
	}else if dismode == "move" {
		mode = VAR_MOVELIFT_DISPLAY
	}

	buf, ret := AreaDataFormat(  width, height,  mode, int32(len(outdata)), outdata)
	return buf, ret
}

func Infotodata(indata []byte, sysled config.LEDControler ) ([]byte,int) {

	Tdata, ret := TextDataPack(sysled.TextColor, indata, sysled.Dismode, sysled.Ledwidth, sysled.Ledheight)
	if ret > 0 {
		Sdata, ret := SendShowData(uint16(ret), Tdata)
		if ret > 0 {
			Pdata, ret := PacKData(ret, Sdata)
			if ret > 0 {
				return Pdata, ret
			}
		}
	}
	return nil, 0
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}