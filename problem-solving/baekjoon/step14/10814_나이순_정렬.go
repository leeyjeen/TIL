package main
 
import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

type userInfo struct {
	age int
	name string
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var users = make([]userInfo, n)
	for i:=0; i<n; i++ {
		fmt.Fscanln(reader, &users[i].age, &users[i].name)
	}

	sort.SliceStable(users, func(i,j int) bool {
		return users[i].age < users[j].age
	})

	for i:=0; i<n;i++ {
		fmt.Fprintln(writer, users[i].age, users[i].name)
	}
}