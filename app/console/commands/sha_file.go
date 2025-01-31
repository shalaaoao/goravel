package commands

import (
	"crypto/sha256"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"io"
	"os"
)

type ShaFile struct {
}

// Signature The name and signature of the console command.
func (receiver *ShaFile) Signature() string {
	return "command:sha_file"
}

// Description The console command description.
func (receiver *ShaFile) Description() string {
	return "获取文件的hash值，带进度条显示"
}

// Extend The console command extend.
func (receiver *ShaFile) Extend() command.Extend {
	return command.Extend{
		Flags: []command.Flag{
			&command.StringFlag{
				Name:    "path",
				Value:   "",
				Aliases: []string{"p"},
				Usage:   "文件的路径",
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *ShaFile) Handle(ctx console.Context) error {

	filePath := ctx.Option("path")

	if filePath == "" {
		fmt.Println("请输入--path=文件路径 / -p=文件路径")
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()

	// 创建一个新的进度条
	bar := pb.Full.Start64(fileSize)

	// 创建一个新的代理 reader，它将报告读取的字节数
	reader := bar.NewProxyReader(file)

	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		fmt.Println(err)
		return nil
	}

	// 结束进度条
	bar.Finish()

	hashInBytes := hash.Sum(nil)[:]

	hashString := fmt.Sprintf("%x", hashInBytes)
	fmt.Println(hashString)

	return nil
}
