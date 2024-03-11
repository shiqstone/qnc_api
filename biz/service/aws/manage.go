package service

import (
	"context"
	"qnc/biz/mw/viper"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var (
	accessKey  string
	secretKey  string
	region     string
	instanceId string
)

type EC2Handler struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewEc2Service(ctx context.Context, c *app.RequestContext) *EC2Handler {
	return &EC2Handler{ctx: ctx, c: c}
}

func Init() {
	config := viper.Conf.Aws
	accessKey = config.AccessKey
	secretKey = config.SecretKey

	region = "ap-southeast-1"
	instanceId = "i-0dbf23e4a46fa2119"
}

func (s *EC2Handler) GetInstanceStatus() (status string, err error) {
	cfg, err := config.LoadDefaultConfig(s.ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))),
	)
	if err != nil {
		hlog.Error(err)
		return
	}
	ec2Client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeInstanceStatusInput{
		InstanceIds: []string{instanceId},
	}
	output, err := ec2Client.DescribeInstanceStatus(s.ctx, input)
	if err != nil {
		hlog.Error(err)
		return
	}
	hlog.Debug(output)
	status = "unknown"
	for _, instanceStatus := range output.InstanceStatuses {
		hlog.Debugf("%s: %s\n", *instanceStatus.InstanceId, instanceStatus.InstanceState.Name)
		if *instanceStatus.InstanceId == instanceId && instanceStatus.InstanceState.Name == "running" {
			status = (string)(instanceStatus.InstanceState.Name)
		}
	}
	return
}

func (s *EC2Handler) StartInstance() (err error) {
	cfg, err := config.LoadDefaultConfig(s.ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))),
	)
	if err != nil {
		hlog.Error(err)
		return
	}
	ec2Client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeInstanceStatusInput{
		InstanceIds: []string{instanceId},
	}
	output, err := ec2Client.DescribeInstanceStatus(s.ctx, input)
	if err != nil {
		hlog.Error(err)
		return
	}
	isRunning := false
	for _, instanceStatus := range output.InstanceStatuses {
		hlog.Debugf("%s: %s\n", *instanceStatus.InstanceId, instanceStatus.InstanceState.Name)
		if *instanceStatus.InstanceId == instanceId && instanceStatus.InstanceState.Name == "running" {
			isRunning = true
		}
	}
	if !isRunning {
		runInstance := &ec2.StartInstancesInput{
			InstanceIds: []string{instanceId},
		}
		hlog.Infof("Start %s\n", instanceId)
		if outputStart, errInstance := ec2Client.StartInstances(s.ctx, runInstance); errInstance != nil {
			return
		} else {
			hlog.Info(outputStart.StartingInstances)
		}
	} else {
		hlog.Infof("Skip starting %s", instanceId)
	}
	return
}

func (s *EC2Handler) StopInstance() (err error) {
	cfg, err := config.LoadDefaultConfig(s.ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))),
	)
	if err != nil {
		hlog.Error(err)
		return
	}
	ec2Client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeInstanceStatusInput{
		InstanceIds: []string{instanceId},
	}
	output, err := ec2Client.DescribeInstanceStatus(s.ctx, input)
	if err != nil {
		hlog.Error(err)
		return
	}
	isStop := false
	for _, instanceStatus := range output.InstanceStatuses {
		hlog.Debugf("%s: %s\n", *instanceStatus.InstanceId, instanceStatus.InstanceState.Name)
		if *instanceStatus.InstanceId == instanceId && instanceStatus.InstanceState.Name == "stop" {
			isStop = true
		}
	}
	if !isStop {
		stopInstance := &ec2.StopInstancesInput{
			InstanceIds: []string{instanceId},
		}
		hlog.Infof("Stop %s\n", instanceId)
		if outputStop, errInstance := ec2Client.StopInstances(s.ctx, stopInstance); errInstance != nil {
			return
		} else {
			hlog.Info(outputStop.StoppingInstances)
		}
	} else {
		hlog.Infof("Skip stop %s\n", instanceId)
	}
	return
}
