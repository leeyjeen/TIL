// https://www.acmicpc.net/problem/16165
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	groupInfos := map[string]groupInfo{}
	memberInfos := map[string]string{}
	for i := 0; i < n; i++ {
		var groupName string
		fmt.Fscanln(reader, &groupName)
		var memberCount int
		fmt.Fscanln(reader, &memberCount)
		info := groupInfo{groupName: groupName}
		for j := 0; j < memberCount; j++ {
			var memberName string
			fmt.Fscanln(reader, &memberName)
			info.members = append(info.members, memberName)
			memberInfos[memberName] = groupName
		}
		groupInfos[groupName] = info
	}
	for i := 0; i < m; i++ {
		var quiz string
		fmt.Fscanln(reader, &quiz)
		var quizType int
		fmt.Fscanln(reader, &quizType)
		if quizType == 0 {
			sort.Strings(groupInfos[quiz].members)
			for _, v := range groupInfos[quiz].members {
				fmt.Fprintln(writer, v)
			}
		} else if quizType == 1 {
			fmt.Fprintln(writer, memberInfos[quiz])
		}
	}
}

type groupInfo struct {
	groupName string
	members   []string
}
