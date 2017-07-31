package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	dir := "./trainData/"
	out, _ := exec.Command("ls", dir).Output()
	allFile := strings.Split(string(out), "\n")
	num := make(map[byte]([]string))
	for _, v := range allFile {
		if len(v) != 0 {
			num[v[0]] = append(num[v[0]], v)
		}
	}
	for k, _ := range num {
		err := exec.Command("mv", dir+string(k)+".jpg", dir+fmt.Sprintf("%s-%d.jpg", string(k), len(num[k]))).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
