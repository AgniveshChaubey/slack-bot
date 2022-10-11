package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"

)

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xapp-1-A045CQ80JNR-4209484681984-7fc0d9a6290ff9bfede7f659f9996be90dd0d8856b611e891699fb877ed544d1")
	os.Setenv("CHANNEL_ID", "C0457BN7Z1U")
	api:= slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"hva.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %S\n", file.Name, file.URL)
	}
}