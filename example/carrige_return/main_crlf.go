// ファイル型をcrlfにしてみる
package main

import "fmt"

func main() {
	// こっちでは\r\nは改行認識 cat -eで出力しても\rは入っている
	fmt.Print("Hello\r\n")
	// cat -e だと末尾の\rが削除されているのが確認できる
	msg := `Hello
`
	fmt.Print(msg)
}
