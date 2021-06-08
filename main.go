package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func helpMessage(originalProgramname string) string {
	programName := filepath.Base(originalProgramname)
	return fmt.Sprintf(`%s [OPTIONS] [FILEs...|DIRs...]
	OPTIONS
		-b, --byte		バイト数表示
		-c, --character		文字数表示
		-l, --line		行数表示
		-n, --no-ignore		ignoreファイルを参照しない
		-h, --help		ヘルプ表示

	ARGUMENTS
		FILEs...		ファイル名の指定．zip/tar/tar.gz/tar.bz2/jar/warファイルも可
		DIRs...			ディレクトリの中のファイルをカウントする

	`, programName)
}

/*
var (
	byteFlag   int
	charFlag   int
	lineFlag   int
	ignoreFlag bool
	helpFlag   string
)

func options() {
	flag.IntVar(&byteFlag, "b", 0, "バイト数表示")
	flag.IntVar(&byteFlag, "byte", 0, "バイト数表示")
	flag.IntVar(&charFlag, "c", 0, "文字数表示")
	flag.IntVar(&charFlag, "character", 0, "文字数表示")
	flag.IntVar(&lineFlag, "l", 0, "行数表示")
	flag.IntVar(&lineFlag, "line", 0, "行数表示")
	flag.BoolVar(&ignoreFlag, "n", false, "ignoreファイルを参照しない")
	flag.BoolVar(&ignoreFlag, "no-ignore", false, "ignoreファイルを参照しない")

	flag.Parse()
}

*/

func goMain(args []string) int {
	fmt.Println(helpMessage(args[0]))
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
