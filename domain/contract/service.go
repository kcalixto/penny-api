package contract

import "github.com/kcalixto/penny-api/domain/context"

type InstanceService interface {
	StartEc2(context context.Context, instanceID string) error
	StopEc2(context context.Context, instanceID string) error
}
