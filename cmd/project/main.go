package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const round = 5

func main() {
	correct := 0

	for i := 0; i < round; i++ {
		a := random(11)
		b := random(11)

		fmt.Printf("%d + %d = \n", a, b)
		scan := bufio.NewScanner(os.Stdin)

		sum := a + b
		scan.Scan()
		strIn := scan.Text()
		in, err := strconv.Atoi(strIn)
		if err == nil && in == sum {
			correct++
		}
	}
	fmt.Printf("正解数 %d/%d", correct, round)
}

func random(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
