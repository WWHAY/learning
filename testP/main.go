package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func myFunc9() {
	filePath := "./err.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	write := bufio.NewWriter(file)

	filename := "./data.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		appInfo := &struct {
			AppName  string `json:"AppName"`
			NebulaID string `json:"NebulaID"`
		}{}

		appInfo.AppName = "vgroup"
		appInfo.NebulaID = line

		url := "http://dev-api-nebula.weizhipin.com/v1/nebula/room_manager/delete_nebula"

		buf, err := json.Marshal(appInfo)
		if err != nil {
			fmt.Println(err)
			write.WriteString(line + "\n")
			write.Flush()
			continue
		}
		resp, err := http.Post(url, "", buf)
		if err != nil {
			fmt.Println(err)
			write.WriteString(line + "\n")
			write.Flush()
			continue
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			write.WriteString(line + "\n")
			write.Flush()
			continue
		}
		deleteTIMGroupResp := HTTPResult{}
		err = json.Unmarshal(respBody, &deleteTIMGroupResp)
		if err != nil {
			fmt.Println(err)
			write.WriteString(line + "\n")
			write.Flush()
			continue
		}
		if deleteTIMGroupResp.StatusCode != 0 {
			fmt.Println(err)
			write.WriteString(line + "\n")
			write.Flush()
			continue
		}

		time.Sleep(100 * time.Millisecond)
		if i%100 == 0 {
			fmt.Printf("success record %d\n", i)
		}
	}
}
