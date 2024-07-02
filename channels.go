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

// AddChannel creates a new channel with the provided channel request data.
func (c *Client) AddChannel(ctx context.Context, channelReq models.ChannelRequest) (channel models.Channel, err error) {
	url := c.APIBase + utils.ChannelEndPoint
	fmt.Println(url)

	requestBodyBytes, err := json.Marshal(channelReq)
	if err != nil {
		fmt.Println("marshal err: ", err)
		return channel, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Println("new req err: ", err)
		return channel, err
	}

	response := &models.ChannelReponse{}

	err = c.SendWithAccessToken(req, response)
	if err != nil {
		fmt.Println("send req err: ", err)
		return channel, err
	}

	if response.Data.ID != "" {
		channel = response.Data
	}

	return channel, err
}

// ReceiveMessage sends a message request to the specified endpoint.
func (c *Client) ReceiveMessage(ctx context.Context, messageRequest models.MessageRequest) error {
	url := c.APIBase + utils.ReceiveMessageEndPoint

	requestBodyBytes, err := json.Marshal(messageRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// DeleteChannel deletes a channel with the specified channel ID.
func (c *Client) DeleteChannel(ctx context.Context, channelID string) error {
	deleteChannelEndpoint := fmt.Sprintf(utils.DeleteChannelEndPoint, channelID)
	url := c.APIBase + deleteChannelEndpoint
	fmt.Println(url)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// DeleteConversation deletes a conversation within a specified channel using the channel ID and conversation ID.
func (c *Client) DeleteConversation(ctx context.Context, channelID, conversationID string) error {
	url := c.APIBase + utils.ChannelEndPoint + channelID + utils.ConversationEndPoint + conversationID

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
