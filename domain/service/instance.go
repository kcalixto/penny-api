package service

import (
	"github.com/kcalixto/penny-api/constants"
	"github.com/kcalixto/penny-api/domain/contract"
	"github.com/kcalixto/penny-api/infra/context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type instanceService struct{}

func NewInstanceService() contract.InstanceService {
	return instanceService{}
}

func (s instanceService) StartEC2(context context.Context, instanceID string) (output []*ec2.InstanceStateChange, err error) {
	svc := ec2.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(constants.REGION),
	})))

	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
		DryRun: aws.Bool(true),
	}
	result, err := svc.StartInstances(input)
	awsErr, ok := err.(awserr.Error)

	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.StartInstances(input)
		if err != nil {
			return output, err
		}
	} else { // This could be due to a lack of permissions
		return output, err
	}

	return *&result.StartingInstances, nil
}

func (s instanceService) StopEC2(context context.Context, instanceID string) (output []*ec2.InstanceStateChange, err error) {
	svc := ec2.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(constants.REGION),
	})))

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
		DryRun: aws.Bool(true),
	}
	result, err := svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.StopInstances(input)
		if err != nil {
			return output, err
		}
	} else {
		return output, err
	}

	return result.StoppingInstances, nil
}
