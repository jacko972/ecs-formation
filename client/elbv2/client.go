package elbv2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"

	"github.com/stormcat24/ecs-formation/client/util"
)

type Client interface {
	DescribeLoadBalancers(names []string) (*elbv2.DescribeLoadBalancersOutput, error)
	CreateRule(params *elbv2.CreateRuleInput) ([]*elbv2.Rule, error)
	DeleteRule(ruleArn string) error
	DescribeRule(params *elbv2.DescribeRulesInput) ([]*elbv2.Rule, error)
	ModifyRule(params *elbv2.ModifyRuleInput) ([]*elbv2.Rule, error)
	CreateTargetGroup(params *elbv2.CreateTargetGroupInput) ([]*elbv2.TargetGroup, error)
	DeleteTargetGroup(targetGroupArn string) error
	DescribeTargetGroup(params *elbv2.DescribeTargetGroupsInput) ([]*elbv2.TargetGroup, error)
	ModifyTargetGroup(params *elbv2.ModifyTargetGroupInput) ([]*elbv2.TargetGroup, error)
}

type DefaultClient struct {
	service *elbv2.ELBV2
}

func (c DefaultClient) DescribeLoadBalancers(names []string) (*elbv2.DescribeLoadBalancersOutput, error) {

	params := elbv2.DescribeLoadBalancersInput{
		LoadBalancerArns: aws.StringSlice(names),
	}

	result, err := c.service.DescribeLoadBalancers(&params)
	if util.IsRateExceeded(err) {
		return c.DescribeLoadBalancers(names)
	}

	return result, err
}

func (c DefaultClient) CreateRule(params *elbv2.CreateRuleInput) ([]*elbv2.Rule, error) {

	result, err := c.service.CreateRule(params)
	if util.IsRateExceeded(err) {
		return c.CreateRule(params)
	}

	return result.Rules, err
}

func (c DefaultClient) DeleteRule(ruleArn string) error {

	params := elbv2.DeleteRuleInput{
		RuleArn: aws.String(ruleArn),
	}

	_, err := c.service.DeleteRule(&params)
	if util.IsRateExceeded(err) {
		return c.DeleteRule(ruleArn)
	}

	return err
}

func (c DefaultClient) DescribeRule(params *elbv2.DescribeRulesInput) ([]*elbv2.Rule, error) {

	result, err := c.service.DescribeRules(params)
	if util.IsRateExceeded(err) {
		return c.DescribeRule(params)
	}

	return result.Rules, err
}

func (c DefaultClient) ModifyRule(params *elbv2.ModifyRuleInput) ([]*elbv2.Rule, error) {

	result, err := c.service.ModifyRule(params)
	if util.IsRateExceeded(err) {
		return c.ModifyRule(params)
	}

	return result.Rules, err
}

func (c DefaultClient) CreateTargetGroup(params *elbv2.CreateTargetGroupInput) ([]*elbv2.TargetGroup, error) {

	result, err := c.service.CreateTargetGroup(params)
	if util.IsRateExceeded(err) {
		return c.CreateTargetGroup(params)
	}

	return result.TargetGroups, err
}

func (c DefaultClient) DeleteTargetGroup(targetGroupArn string) error {

	params := elbv2.DeleteTargetGroupInput{
		TargetGroupArn: aws.String(targetGroupArn),
	}

	_, err := c.service.DeleteTargetGroup(&params)
	if util.IsRateExceeded(err) {
		return c.DeleteTargetGroup(targetGroupArn)
	}

	return err
}

func (c DefaultClient) DescribeTargetGroup(params *elbv2.DescribeTargetGroupsInput) ([]*elbv2.TargetGroup, error) {

	result, err := c.service.DescribeTargetGroups(params)
	if util.IsRateExceeded(err) {
		return c.DescribeTargetGroup(params)
	}

	return result.TargetGroups, err
}

func (c DefaultClient) ModifyTargetGroup(params *elbv2.ModifyTargetGroupInput) ([]*elbv2.TargetGroup, error) {

	result, err := c.service.ModifyTargetGroup(params)
	if util.IsRateExceeded(err) {
		return c.ModifyTargetGroup(params)
	}

	return result.TargetGroups, err
}
