package main

import (
	"bytes"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"

	"database/sql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ssmconfig "github.com/ianlopshire/go-ssm-config"
	_ "github.com/lib/pq"

	"github.com/terraform-and-serverless/app/pkg/pg"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type Config struct {
	RdsConnectionString 	string		`ssm:"rds-connection-string" required:"true"`
}

const ssmBasePath string = "/env/app"

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	var config Config

	log.Info("Gathering configuration from SSM")
	err := ssmconfig.Process(ssmBasePath, &config)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	log.Info("Opening database connection")
	db, err := sql.Open("postgres",config.RdsConnectionString)
	if err != nil {
		return Response{StatusCode: 500}, err
	}
	defer db.Close()

	log.Info("Pinging database")
	err = pg.Ping(db)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
