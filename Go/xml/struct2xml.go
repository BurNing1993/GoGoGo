package main

import(
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct{
	XMLName xml.Name `xml:"serverd"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
}

type server struct{
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main(){
	v:=&Servers{Version:"1"}
	v.Svs=append(v.Svs,server{"Shanghai_VPN","127.0.0.1"})
	v.Svs=append(v.Svs,server{"Beijing_VPN","127.0.0.2"})
	output,err:=xml.MarshalIndent(v," ","   ")
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}