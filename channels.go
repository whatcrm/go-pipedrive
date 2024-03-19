package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

func (c *Client) AddChannel(ctx context.Context, channel models.ChannelRequest) error {
	url := c.APIBase + utils.ChannelEndPoint

	requestBodyBytes, err := json.Marshal(channel)
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
