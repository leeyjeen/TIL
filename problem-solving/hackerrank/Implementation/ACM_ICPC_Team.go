// https://www.hackerrank.com/challenges/acm-icpc-team/problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	var topics []string
	for i := 0; i < n; i++ {
		var topic string
		fmt.Fscanln(reader, &topic)
		topics = append(topics, topic)
	}
	maxTopic, maxCount := acmTeam(topics)
	fmt.Fprintln(writer, maxTopic)
	fmt.Fprintln(writer, maxCount)
}

func acmTeam(topics []string) (maxTopic, maxCount int) {
	for i := 0; i < len(topics)-1; i++ {
		topicI := topics[i]
		for j := i + 1; j < len(topics); j++ {
			var topicCount int
			topicJ := topics[j]
			for k := range topicI {
				if topicI[k] == 49 || topicJ[k] == 49 {
					topicCount++
				}
			}
			if maxTopic == topicCount {
				maxCount++
			} else if maxTopic < topicCount {
				maxTopic = topicCount
				maxCount = 1
			}
		}
	}
	return maxTopic, maxCount
}
