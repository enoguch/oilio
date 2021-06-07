package main

import (
	"fmt"
	"path/filepath"
)

func helpMessage(originalProgramname string) string {
	programName := filepath.Base(originalProgramname)
	return fmt.Sprintf(`%s [OPTIONS] [FILEs...|DIRs...]
	OPTIONS
		-b, --byte			バイト数表示
		-c, --character		文字数表示
		-l, --line			行数表示
		-n, --no-ignore		ignoreファイルを参照しない
		-h, --help			ヘルプ表示

	`, programName)
}

func main() {
	fmt.Println(helpMessage("oilio"))
}
