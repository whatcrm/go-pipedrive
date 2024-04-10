package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

func (c *Client) AddChannel(ctx context.Context, channelReq models.ChannelRequest) (channel models.Channel, err error) {
	url := c.APIBase + utils.ChannelEndPoint
	fmt.Println(url)

	requestBodyBytes, err := json.Marshal(channelReq)
	if err != nil {
		return channel, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return channel, err
	}

	req.Header.Set("Content-Type", "application/json")

	response := &models.ChannelReponse{}
	
	err = c.SendWithAccessToken(req, response)

	if response.Data.ID != "" {
		channel = response.Data
	}
	
	return channel, err
}

func (c *Client) ReceiveMessage(ctx context.Context, messageRequest models.MessageRequest) error {
	url :=c.APIBase + utils.ReceiveMessageEndPoint

	requestBodyBytes, err := json.Marshal(messageRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	return c.Send(req, nil)
}

func (c *Client) DeleteChannel(ctx context.Context, channelID string) error {
	url := c.APIBase + utils.ChannelEndPoint + channelID

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) DeleteConversation(ctx context.Context, channelID, conversationID string) error {
	url := c.APIBase + utils.ChannelEndPoint + channelID + utils.ConversationEndPoint + conversationID

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}
