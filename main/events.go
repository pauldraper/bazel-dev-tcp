package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type BuildStatus int

const (
	BuildStatusNone BuildStatus = iota
	BuildStatusBuilding
	BuildStatusBuilt
	BuildStatusFailed
)

func ReadBuildStatus(reader *bufio.Reader) (BuildStatus, error) {
	var status BuildStatus
	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return BuildStatusNone, err
		}

		var d map[string]interface{}
		err = json.Unmarshal([]byte(text), &d)

		switch d["type"] {
		case "GRAPH_CHANGE":
			status = BuildStatusBuilding
		case "SOURCE_CHANGE":
			status = BuildStatusBuilding
		case "RUN_DONE":
			status = BuildStatusBuilt
		case "RUN_FAILED":
			status = BuildStatusFailed
		}
	}
	return status, nil
}

func UpdateBuildStatus(reader *bufio.Reader, status BuildStatus) (BuildStatus, error) {
	newStatus, err := ReadBuildStatus(reader)
	if err != nil {
		return status, err
	}

	if newStatus != BuildStatusNone {
		status = newStatus
	}
	return status, nil
}

/*
func StdinMessage(ch chan struct{}) {
	reader := bufio.NewReader(os.Stdin)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		ch <- struct{}{}
	}
}
*/
