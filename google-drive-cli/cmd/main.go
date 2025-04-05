package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const (
	serviceAccountKeyPath = "/Users/takenariyamamoto/Devlopment/private/golang-practice/google-drive-cli/config/sample-pro-395702-a1810e69a578.json"
)

func main() {
	filePath := flag.String("file", "", "アップロードするファイルのパス")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("ファイルパスを指定してください。使用方法: google-drive-cli -file=<ファイルパス>")
	}

	ctx := context.Background()

	// Service Accountの認証情報を使用してDriveサービスを初期化
	service, err := drive.NewService(ctx, option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		log.Fatalf("Drive サービスの初期化に失敗しました: %v", err)
	}

	err = uploadFile(service, *filePath)
	if err != nil {
		log.Fatalf("ファイルのアップロードに失敗しました: %v", err)
	}

	fmt.Println("ファイルのアップロードが完了しました！")
}

func uploadFile(service *drive.Service, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ファイルを開けませんでした: %v", err)
	}
	defer file.Close()

	filename := filepath.Base(filePath)
	f := &drive.File{
		Name: filename,
		Parents: []string{
			"1xaE39-d5CUWLMd3EJza0sxyPjFrsuSJx",
		},
	}

	res, err := service.Files.Create(f).Media(file).Do()
	if err != nil {
		return fmt.Errorf("アップロードに失敗しました: %v", err)
	}

	fmt.Printf("ファイル '%s' をアップロードしました（ID: %s）\n", filename, res.Id)
	return nil
}
