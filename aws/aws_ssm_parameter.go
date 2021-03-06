// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmParameter(client *Client) ([]Resource, error) {
	req := client.Ssmconn.DescribeParametersRequest(&ssm.DescribeParametersInput{})

	var result []Resource

	p := ssm.NewDescribeParametersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Parameters {

			result = append(result, Resource{
				Type:      "aws_ssm_parameter",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
