package main

import (
	"bufio"
	"flag"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var listUrl = "https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_best.txt"

var tmpl = `bt-tracker={{range $k, $v := .}}{{if eq $k 0}}{{$v}}{{else}},{{$v}}{{end}}{{end}}
enable-dht=true
bt-enable-lpd=true
enable-peer-exchange=true
`

func downloadList() ([]string, error) {
	resp, err := http.Get(listUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rd := bufio.NewReader(resp.Body)
	res := []string{}
	for {
		line1, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		//fmt.Println("url:", string(line1), len(line1))
		if len(line1) > 0 {
			res = append(res, string(line1))
		}
	}
	return res, nil
}

func writeConf(data []string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	confPath := filepath.Join(home, ".aria2", "aria2.conf")
	fp, err := os.Create(confPath)
	if err != nil {
		return err
	}
	defer fp.Close()
	t := template.New("")
	t, err = t.Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(fp, data)
}

func main() {
	flag.Parse()
	if len(flag.Arg(0)) > 0 {
		listUrl = flag.Arg(0)
	}
	list1, err := downloadList()
	if err != nil {
		panic(err.Error())
	}
	err = writeConf(list1)
	if err != nil {
		panic(err.Error())
	}
}
