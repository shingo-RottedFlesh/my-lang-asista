// Command repl は Asista の対話実行環境（REPL）です。
// 第1章で lexer をつないで、入力をトークン列として表示するところから始めます。
package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "asista> "

func main() {
	fmt.Println("Asista REPL へようこそ。終了は Ctrl+D です。")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			fmt.Println()
			return
		}
		line := scanner.Text()
		// TODO(第1章): lexer で line をトークンに分解して表示する
		fmt.Println(line)
	}
}
