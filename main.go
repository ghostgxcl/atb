// Sean at Shanghai
// convert alipay bill to beancount version

package main


import (
	"log"
	"github.com/dilfish/tools"
	"flag"
)


var flagInput = flag.String("i", "", "input file name")
var flagOutput = flag.String("o", "output.bean", "output file name")
var flagConfig = flag.String("c", "config.conf", "config file name")

// Config set all default values
type Config struct {
	DefaultCurrency string `json:"defaultCurrency"`
	DefaultPlusAccount string `json:"defaultPlusAccount"`
	DefaultMinusAccount string `json:"defaultMinusAccount"`
	Title string `json:"title"`
	AccountList []AliBillAttr `json:"accountList"`
}

func main() {
	log.SetPrefix("AlipayBillToBean: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()
	if *flagInput == "" {
		flag.PrintDefaults()
		return
	}
	var conf Config
	err := tools.ReadConfig(*flagConfig, &conf)
	if err != nil {
		log.Println("read config error:", err)
		return
	}
	err = ReadAliBill(*flagInput, conf.AccountList)
	if err != nil {
		log.Println("read ali bill error:", err)
		return
	}
	err = WriteBean(&conf)
	if err != nil {
		log.Println("write bean file error:", err)
		return
	}
}
