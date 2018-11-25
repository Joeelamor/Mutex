package main

import (
	"Mutex/util"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) (int, int, int, int, []util.HostInfo) {
	var (
		NodeNum       int
		InterReqDelay int
		CsExecTime    int
		ReqNum        int
		HostList      []util.HostInfo
	)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	NodeNum, _ = strconv.Atoi(words[0])
	InterReqDelay, _ = strconv.Atoi(words[1])
	CsExecTime, _ = strconv.Atoi(words[2])
	ReqNum, _ = strconv.Atoi(words[3])

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) == 0 {
			continue
		}
		id, _ := strconv.ParseInt(words[0], 10, 32)
		hostname := words[1]
		port := words[2]

		hostInfo := util.HostInfo{
			Id:       int32(id),
			HostName: hostname,
			Port:     port,
		}
		HostList = append(HostList, hostInfo)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return NodeNum, InterReqDelay, CsExecTime, ReqNum, HostList
}

