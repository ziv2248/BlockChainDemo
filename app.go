package main

import (
	"bufio"
	"fmt"
	app "myBlockChain/internal"
	"os"
)

func Help() {
	fmt.Println("===============")
	fmt.Println("輸入1 新建區塊")
	fmt.Println("輸入2 查詢區塊鏈")
	fmt.Println("輸入3 離開")
	fmt.Println("===============")
}

func main() {
	fmt.Println("=====歡迎使用Ziv區塊鏈=====")
	var (
		op string
	)

	NewBlockchain := app.CreateBlockchain()
	for {
		fmt.Println("\n請輸入 h, 1, 2, 3 進行操作")
		fmt.Scanln(&op)
		if op == "h" {
			fmt.Println("\nHelp: ")
			Help()
		} else if op == "1" {
			fmt.Println("請輸入數據:")
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			NewBlockchain.AddBlock(data)
			fmt.Println("***新區塊加入區塊鏈***")
		} else if op == "2" {
			for i, block := range NewBlockchain.Blocks {
				fmt.Printf("前一區塊hash: %x\n", block.PrevBlockHash)
				fmt.Printf("當前區塊hash: %x\n", block.Hash)
				fmt.Printf("值: %s\n", block.Data)
				if i != 0 {
					pow := app.InitProofOfWork(block)
					fmt.Println("驗證:", pow.Validate())
				}
				fmt.Println()
			}
		} else if op == "3" {
			fmt.Println("=====感謝您使用Ziv區塊鏈=====")
			break
		} else {
			fmt.Println("輸入必須為 h, 1, 2, 3 其中之一")
		}
	}
}
