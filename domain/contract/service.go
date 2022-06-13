package contract

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/kcalixto/penny-api/infra/context"
)

type InstanceService interface {
	StartEC2(context context.Context, instanceID string) ([]*ec2.InstanceStateChange, error)
	StopEC2(context context.Context, instanceID string) ([]*ec2.InstanceStateChange, error)
}
