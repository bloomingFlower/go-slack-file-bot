package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	filePath := ""
	if len(os.Args) > 1 {
		filePath = os.Args[1]
		//fmt.Println("첫 번째 인자:", filePath)
	} else {
		fmt.Println("파일경로가 제공되지 않았습니다.")
		return
	}

	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")
	api := slack.New(slackBotToken)
	channelArr := []string{channelId}
	fileArr := []string{filePath}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)

		//fileInfo, _, _, err := api.GetFileInfo(file.ID, 1, 1)
		//if err != nil {
		//	fmt.Printf("Error retrieving file info: %s\n", err)
		//	return
		//}
		//
		//fmt.Printf("Name: %s, URL: %s\n", fileInfo.Name, fileInfo.URLPrivate)

	}
}
