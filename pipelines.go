package gopipedrive

import (
	"context"
	"fmt"
	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
	"net/http"
)

func (c *Client) GetPipelines(ctx context.Context) ([]models.Pipeline, error) {
	pipelineEndpoint := fmt.Sprintf(utils.PipelinesEndpoint)
	url := c.APIBase + pipelineEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var personResponse models.PipelinesResponse
	err = c.SendWithAccessToken(req, &personResponse)
	if err != nil {
		return nil, err
	}

	return personResponse.Data, nil
}

func (c *Client) GetPipeline(ctx context.Context, id int) (*models.Pipeline, error) {
	pipelineEndpoint := fmt.Sprintf(utils.PipelineEndpoint, id)
	url := c.APIBase + pipelineEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var personResponse models.PipelineResponse
	err = c.SendWithAccessToken(req, &personResponse)
	if err != nil {
		return nil, err
	}

	return &personResponse.Data, nil
}

func (c *Client) GetStages(ctx context.Context) ([]models.Stage, error) {
	url := c.APIBase + utils.StagesEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var personResponse models.StagesResponse
	err = c.SendWithAccessToken(req, &personResponse)
	if err != nil {
		return nil, err
	}

	return personResponse.Data, nil
}

func (c *Client) GetStagesFromPipeline(ctx context.Context, pipelineID int) ([]models.Stage, error) {
	pipelineEndpoint := fmt.Sprintf(utils.StagesFromPipelineEndpoint, pipelineID)
	url := c.APIBase + pipelineEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var personResponse models.StagesResponse
	err = c.SendWithAccessToken(req, &personResponse)
	if err != nil {
		return nil, err
	}

	return personResponse.Data, nil
}
