package lox

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/scanner"
)

var STRIDE int = 1024
var HadError = false

type GenericList [T comparable] struct {
	Data []T
}

func (l *GenericList[T]) Append(value T){
	l.Data = append(l.Data, value)
}

func NewGenericList[T comparable]() *GenericList[T]{
	return &GenericList[T]{
		Data: []T{},
	}
}

func Report(line int, where string, message string){
	HadError = true
	l := log.New(os.Stderr, "", 0)
	l.Println("[line ", line, "] Error", where, ": ", message)
}

func Error(line int, message string){
	Report(line, "", message)
}

func Run(source string) {
	var s scanner.Scanner
	s.Init(strings.NewReader(source))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan(){
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

	if HadError {
		os.Exit(65)
	}
}

func RunPrompt() {

	s := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Lox interpretter\nCtrl+D or Ctrl+Z to end")

	fmt.Print("> ")
	for s.Scan(){
		line := s.Text()
		Run(line)
		fmt.Print("> ")
		
		HadError = false
	}

	if err := s.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func RunFile() {
	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 0, STRIDE)
	str := ""

	for {
		n, err := io.ReadFull(reader, buf[:cap(buf)])
		buf = buf[:n]
		if err != nil {
			if err == io.EOF {
				break // expected
			}
			if err != io.ErrUnexpectedEOF {
				fmt.Fprintln(os.Stderr, err)
				break
			}
			str = str + string(buf)
		}
	}
	Run(str)
}

func main() {
	if len(os.Args) == 1 {
		RunPrompt()
	} else if len(os.Args) == 2 {
		RunFile()
	} else {
		fmt.Println("usage: jlox filename.lox")
	}
}
