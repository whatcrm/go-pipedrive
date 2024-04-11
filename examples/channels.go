package main

import (
	"context"
	"fmt"

	gopipedrive "github.com/whatcrm/go-pipedrive"
	"github.com/whatcrm/go-pipedrive/models"
)

func main() {
	clientID := "7487ed88f96ef1a0"
	clientSecret := "1ad34c4b67254a453facf83618f50460f2d27788"
	redirectURI := "https://dev.pipedrive.whatcrm.net/oauth"

	domain := "whatcrm-sandbox"
	api := "v1u:AQIBAHj-LzTNK2yuuuaLqifzhWb9crUNKTpk4FlQ9rjnXqp_6AE-mSRsNEc-qt-JZF1QvUyHAAAAfjB8BgkqhkiG9w0BBwagbzBtAgEAMGgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMGqcEIUgSSD15bcXPAgEQgDtLA6W-OJ_ytVMrWzkVH7tyBFy3C5ZLb1hgRXukifjFKHrPSN0XT5mLYmrNgQAPNro5MhWeqqjIaVI6zA:kGaKgy9OxReNwkDFcrsxprstSr_3-flRwvdfHrxq4QRgcWd7kFWy6OojIjevSMfvDjSIfVh7udOrfq_YnUO6ymxvAu1mkSZDmrZhwdB2J0UtHJUCSMNfzgfRS-M9xoVjrcFUBUiIZAxCDp9mjzttuNRb5rJ1noWO3ClgoqY8LAai_sQCvMt2zaKCkXOWrKhehYV9SdDpleFyzYL6cnVuzUJzmtrxiGpo7YfZVwOGQ1FAN5_q49Lt64YfoXvvVPj1ua4f4yynjuPQS5aVFJXOdY7dregyOS1Y_-GB20PaNlS0TMQiB2U066vS55h1TMbcN-4IhN0gOiAODZvcbAKxciNxKRlbAeQUvnxL1YQhJtEQJXz9LsJCoy5PHYQqw6XY7aIOFZUfji8BIRzsUvRBB-VuOFv8Wg"
	client, err := gopipedrive.NewClient(domain, clientID, clientSecret, redirectURI)
	if err != nil {
		fmt.Println("error: ", err)
	}

	client.Token = api

	channel := models.ChannelRequest{
		Name: "name",
		ProviderChannelID: "1",
		AvatarURL: "1",
		TemplateSupport: false,
		ProviderType: "other",
	}

	c, err := client.AddChannel(context.Background(), channel)
	fmt.Println("res: ", c, " err: ", err)
}
