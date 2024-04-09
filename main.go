package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano()) //  设置随机数种子
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin) //把一个文件转换成一个 reader 变量
	//fmt.Println(reader)
	for {
		input, err := reader.ReadString('\n') //读取一行
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n") //  去掉换行符

		guess, err := strconv.Atoi(input) //  将字符串转换成数字
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
