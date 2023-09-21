package main

import (
	"log"
	"os"

	"github.com/maskentir/qontalk/qontak"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	var (
		username     = os.Getenv("QONTAK_USERNAME")
		password     = os.Getenv("QONTAK_PASSWORD")
		grantType    = os.Getenv("QONTAK_GRANT_TYPE")
		clientID     = os.Getenv("QONTAK_CLIENT_ID")
		clientSecret = os.Getenv("QONTAK_CLIENT_SECRET")

		messageTemplateID = os.Getenv("QONTAK_MESSAGE_TEMPLATE_ID")
		channelIntegrationID = os.Getenv("QONTAK_CHANNEL_INTEGRATION_ID")
	)

	sdk := qontak.NewQontakSDKBuilder().WithClientCredentials(username, password, grantType, clientID, clientSecret).Build()

	if err := sdk.Authenticate(); err != nil {
		log.Println(err)	
	}

	message := qontak.DirectWhatsAppBroadcast{
		ToName: "Diky Yodihamzah",
		ToNumber: "6289684279559",
		MessageTemplateID: messageTemplateID,
		ChannelIntegrationID: channelIntegrationID,
		Language: map[string]string{
			"code": "id",
		},
		BodyParams: []qontak.KeyValueText{
			{
				Key: "1",
				Value: "full_name",
				ValueText: "Diky Yodihamzah",
			},
		},
	}

	if err := sdk.SendDirectWhatsAppBroadcast(message); err != nil {
		log.Println(err)
	}
}
