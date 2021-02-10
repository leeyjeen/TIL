// https://www.acmicpc.net/problem/1049
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

	var minPackage, minSingle = 1000, 1000
	for i := 0; i < m; i++ {
		var packagePrice, singlePrice int
		fmt.Fscanln(reader, &packagePrice, &singlePrice)
		if minPackage > packagePrice {
			minPackage = packagePrice
		}
		if minSingle > singlePrice {
			minSingle = singlePrice
		}
	}
	var money int
	if minSingle*6 > minPackage {
		money += (n / 6) * minPackage
		n %= 6
		if n > 0 {
			if minSingle*n > minPackage {
				money += minPackage
			} else {
				money += minSingle * n
			}
		}
	} else {
		money += n * minSingle
	}

	fmt.Fprintln(writer, money)
}
