package main

import ( 
	"fmt"
	"bufio"
	"os"
)

func main() {
	var score, sum int
	reader := bufio.NewReader(os.Stdin)

	for i:=0; i<5; i++ {
		fmt.Fscanln(reader, &score)
		if score < 40 {
			score = 40
		}
		sum += score
	}

	fmt.Println(sum/5)
}