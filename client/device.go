package main

import (
	"encoding/xml"
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Response struct {
	XMLName    xml.Name `xml:"Response"`
	CmdType    string   `xml:"CmdType"`
	SN         int      `xml:"SN"`
	DeviceID   string   `xml:"DeviceID"`
	SumNum     int      `xml:"SumNum"`
	DeviceList DeviceList
}

type DeviceList struct {
	XMLName xml.Name `xml:"DeviceList"`
	Num     int      `xml:"Num,attr"`
	Device  Device
}

type Device struct {
	XMLName xml.Name `xml:"Item"`

	DeviceID     string
	Name         string
	Manufacturer string
	Model        string
	Owner        string
	CivilCode    string
	Address      string
	Parental     string
	RegisterWay  string
	Secrecy      string
	Status       string
}

func GetResponseXmlStr(sn int, deviceID string, sumNum int, childDeviceID, name, manufacturer, model string) string {
	decoder := simplifiedchinese.GBK.NewDecoder()
	nname, _ := decoder.String(name)
	cd := Device{
		DeviceID:     childDeviceID,
		Name:         nname,
		Manufacturer: manufacturer,
		Model:        model,
		Owner:        "0",
		CivilCode:    "111",
		Address:      "axy",
		Parental:     "0",
		RegisterWay:  "1",
		Secrecy:      "0",
		Status:       "ON",
	}

	r := &Response{
		CmdType:  "Catalog",
		SN:       sn,
		DeviceID: deviceID,
		SumNum:   sumNum,
	}
	r.DeviceList = DeviceList{
		Num:    1,
		Device: cd,
	}

	b, _ := xml.MarshalIndent(r, "", "") // 有缩进格式
	fmt.Printf("%v\n", string(b))
	return "<?xml version=\"1.0\" encoding=\"GB2312\" standalone=\"yes\" ?>" + string(b)
}
