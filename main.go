package main

import (
	"bufio"
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
func isDir(filename string) {
	filename :=
}
*/

func fileReadLine(fileName string) {
	line := 0
	fp, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line++
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Line:", line)
}

func goMain(args []string) int {
	if len(os.Args) == 1 {
		fmt.Println(helpMessage(args[0]))
		return 0
	} else if args[1] == "-h" || args[1] == "--help" {
		fmt.Println(helpMessage(args[0]))
		return 0
	}

	fileReadLine(args[1])
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
