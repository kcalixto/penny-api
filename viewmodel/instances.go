package viewmodel

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceArray struct {
	Reservations []Reservation `json:"Reservations"`
}

type Reservation struct {
	Instances []Instance
}

type Instance struct {
	ID             string
	InstanceType   string
	SecurityGroups []SecurityGroup
	SubnetID       string
	Network        Network
	State          string
	Tags           []Tag
	VpcId          string
}

type Network struct {
	PublicDnsName   string
	PublicIpAddress string
}

type SecurityGroup struct {
	GroupID   string
	GroupName string
}

type Tag struct {
	Key   string
	Value string
}

func ParseDescribeInstancesOutput(instance *ec2.DescribeInstancesOutput) []Instance {
	var parsedInstances []Instance

	//Reading reservations
	for _, reservations := range instance.Reservations {
		for _, instance := range reservations.Instances {
			parsedInstances = append(parsedInstances, parseInstance(instance))
		}
	}

	return parsedInstances
}

func parseInstance(ec2 *ec2.Instance) Instance {
	return Instance{
		ID:             *ec2.InstanceId,
		InstanceType:   *ec2.InstanceType,
		SecurityGroups: parseSecurityGroups(ec2.SecurityGroups),
		SubnetID:       *ec2.SubnetId,
		Network:        parseNetwork(ec2),
		State:          *ec2.State.Name,
		Tags:           parseTags(ec2.Tags),
		VpcId:          *ec2.VpcId,
	}
}

func parseSecurityGroups(securityGroups []*ec2.GroupIdentifier) (output []SecurityGroup) {
	for _, sg := range securityGroups {
		output = append(output, SecurityGroup{
			GroupID:   *sg.GroupId,
			GroupName: *sg.GroupName,
		})
	}

	return output
}

func parseNetwork(ec2 *ec2.Instance) (output Network) {
	if ec2.PublicDnsName != nil {
		output.PublicDnsName = *ec2.PublicDnsName
	}
	if ec2.PublicIpAddress != nil {
		output.PublicIpAddress = *ec2.PublicIpAddress
	}

	return output
}

func parseTags(tags []*ec2.Tag) (output []Tag) {
	for _, tag := range tags {
		output = append(output, Tag{
			Key:   *tag.Key,
			Value: *tag.Value,
		})
	}

	return output
}
