package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode/utf8"

	flag "github.com/spf13/pflag"
)

type options struct {
	bytet     bool
	character bool
	word      bool
	line      bool
	help      bool
	args      []string
}

type counts struct {
	bytes int
	chars int
	words int
	lines int
}

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

func parseArgs(args []string) (*options, error) {
	opts := &options{}
	flags := flag.NewFlagSet("oilio", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.bytet, "byte", "b", false, "バイト数")
	flags.BoolVarP(&opts.character, "character", "c", false, "文字数")
	flags.BoolVarP(&opts.word, "word", "w", false, "単語数")
	flags.BoolVarP(&opts.line, "line", "l", false, "行数")
	flags.BoolVarP(&opts.help, "help", "h", false, "ヘルプ表示")

	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	if !opts.bytet && !opts.character && !opts.word && !opts.line {
		opts.bytet = true
		opts.character = true
		opts.word = true
		opts.line = true
	}
	opts.args = flags.Args()[1:]

	return opts, nil
}

func printer(o *options, c counts, filename string) {
	if o.line {
		fmt.Printf("line: %5d   ", c.lines)
	}
	if o.bytet {
		fmt.Printf("byte: %5d   ", c.bytes)
	}
	if o.character {
		fmt.Printf("char: %5d   ", c.chars)
	}
	if o.word {
		fmt.Printf("word: %5d   ", c.words)
	}
	fmt.Println(filename)

}

func Count(o *options, filenames []string) {
	var c counts
	for _, filename := range filenames {
		fileInfo, _ := os.Stat(filename)
		if fileInfo.IsDir() {
			/*
				root := fileInfo.Name()
				fmt.Println(filename)
			*/
			rootFiles, _ := ioutil.ReadDir(filename)
			for _, rootFile := range rootFiles {
				if rootFile.IsDir() {
					continue
				}
				fullPath := filepath.Join(filename, rootFile.Name())
				rfp, err := os.Open(fullPath)
				if err != nil {
					continue
				}
				r, _ := ioutil.ReadAll(rfp)
				c.lines = bytes.Count(r, []byte{'\n'})
				c.bytes = len(r)
				c.chars = utf8.RuneCountInString(string(r))
				c.words = len(bytes.Fields(r))

				printer(o, c, fullPath)
			}

			continue
		}

		fp, err := os.Open(filename)
		if err != nil {
			continue
		}
		a, _ := ioutil.ReadAll(fp)
		c.lines = bytes.Count(a, []byte{'\n'})
		c.bytes = len(a)
		c.chars = utf8.RuneCountInString(string(a))
		c.words = len(bytes.Fields(a))

		printer(o, c, filename)

	}
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return 0
	}
	Count(opts, opts.args)

	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
