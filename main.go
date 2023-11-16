package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	// 임시 디렉토리 생성
	tmpDir, err := os.MkdirTemp("", "named_pipe")
	if err != nil {
		fmt.Printf("Error creating temporary directory: %s\n", err)
		os.Exit(1)
	}

	// Named pipe의 경로 생성
	namedPipe := filepath.Join(tmpDir, "stdout")

	// Named pipe 생성
	err = syscall.Mkfifo(namedPipe, 0666)
	if err != nil {
		fmt.Printf("Error creating named pipe: %s\n", err)
		os.Remove(tmpDir) // 에러 발생 시 임시 디렉토리 정리
		os.Exit(1)
	}

	fmt.Println("Named pipe created at:", namedPipe)

	// 별도의 고루틴에서 Named pipe로 데이터 쓰기
	go func() {
		// 파일 열기
		file, err := os.OpenFile(namedPipe, os.O_WRONLY, os.ModeNamedPipe)
		if err != nil {
			fmt.Printf("Error opening named pipe for writing: %s\n", err)
			return
		}
		defer file.Close()

		// 데이터 쓰기
		_, err = file.WriteString("Hello, named pipe!\n")
		if err != nil {
			fmt.Printf("Error writing to named pipe: %s\n", err)
			return
		}
	}()

	// Named pipe에서 데이터 읽기
	file, err := os.OpenFile(namedPipe, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		fmt.Printf("Error opening named pipe for reading: %s\n", err)
		os.Remove(tmpDir)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("Received: %s", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from named pipe: %s\n", err)
	}

	// 사용 후, 임시 디렉토리와 named pipe 정리
	os.Remove(namedPipe)
	os.Remove(tmpDir)
}
