package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jcbl1/myM2H/converter"
)

func main() {
	filename := os.Args[1]
	result, err := converter.MDtoHTML(filename)
	check(err)
	outfile := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		outfile += strconv.Itoa(rand.Int() % 10)
	}
	outfile += ".html"
	f, err := os.OpenFile("./"+outfile, os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	_, err = f.WriteString(result)
	check(err)
	f.Sync()
	fmt.Printf("转换成功，已写入./%s中\n", outfile)
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
