package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed templates/*
var templates embed.FS

func main() {
	// 定义命令行参数
	interactive := flag.Bool("i", false, "Interactive mode: ask for confirmation before creating each file")
	flag.Parse()

	// 要创建的文件及其内容
	files := []string{
		".gitignore",
		".dockerignore",
		"Makefile",
		"Dockerfile",
		"README.md",
	}
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// 创建文件
	for _, filename := range files {
		// 读取模板文件内容
		content, err := templates.ReadFile("templates/" + filename)
		if err != nil {
			fmt.Printf("Error reading template file %s: %s\n", filename, err)
			continue
		}

		filePath := filepath.Join(dir, filename)

		if *interactive {
			fmt.Printf("Do you want to create %s? (y/n): ", filename)
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)

			if answer != "y" && answer != "Y" {
				fmt.Printf("Skipping %s\n", filename)
				continue
			}
		}

		err = os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error creating file %s: %s\n", filename, err)
			continue
		}
		fmt.Printf("Created %s\n", filename)
	}
}
