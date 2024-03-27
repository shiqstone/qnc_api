package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var sender string // Must be a verified email address or domain

type SesHandler struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewSesService(ctx context.Context, c *app.RequestContext) *SesHandler {
	return &SesHandler{ctx: ctx, c: c}
}

func SendEmail(recipient string, subject string, content string, html string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))),
	)
	if err != nil {
		hlog.Error(err)
		return err
	}

	client := ses.NewFromConfig(cfg)

	htmlBody := "<h1>" + subject + "</h1><p>" + content + "<p>"
	if len(html) == 0 {
		htmlBody = html
	}
	textBody := content
	charset := "UTF-8"

	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{recipient},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String(charset),
					Data:    aws.String(htmlBody),
				},
				Text: &types.Content{
					Charset: aws.String(charset),
					Data:    aws.String(textBody),
				},
			},
			Subject: &types.Content{
				Charset: aws.String(charset),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	output, err := client.SendEmail(context.TODO(), input)
	if err != nil {
		hlog.Error(err)
		return err
	}

	hlog.Infof("Email Sent to address: %s\n", recipient)
	hlog.Debugf("Message ID: %s\n", *output.MessageId)
	return nil
}
