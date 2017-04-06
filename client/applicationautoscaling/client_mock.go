// Automatically generated by MockGen. DO NOT EDIT!
// Source: client/applicationautoscaling/client.go

package applicationautoscaling

import (
	applicationautoscaling "github.com/aws/aws-sdk-go/service/applicationautoscaling"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) DeleteScalingPolicy(params *applicationautoscaling.DeleteScalingPolicyInput) error {
	ret := _m.ctrl.Call(_m, "DeleteScalingPolicy", params)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) DeleteScalingPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteScalingPolicy", arg0)
}

func (_m *MockClient) DeregisterScalableTarget(params *applicationautoscaling.DeregisterScalableTargetInput) error {
	ret := _m.ctrl.Call(_m, "DeregisterScalableTarget", params)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) DeregisterScalableTarget(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterScalableTarget", arg0)
}

func (_m *MockClient) DescribeScalableTargets(params *applicationautoscaling.DescribeScalableTargetsInput) ([]*applicationautoscaling.ScalableTarget, error) {
	ret := _m.ctrl.Call(_m, "DescribeScalableTargets", params)
	ret0, _ := ret[0].([]*applicationautoscaling.ScalableTarget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) DescribeScalableTargets(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeScalableTargets", arg0)
}

func (_m *MockClient) DescribeScalingActivities(params *applicationautoscaling.DescribeScalingActivitiesInput) ([]*applicationautoscaling.ScalingActivity, error) {
	ret := _m.ctrl.Call(_m, "DescribeScalingActivities", params)
	ret0, _ := ret[0].([]*applicationautoscaling.ScalingActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) DescribeScalingActivities(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeScalingActivities", arg0)
}

func (_m *MockClient) DescribeScalingPolicies(params *applicationautoscaling.DescribeScalingPoliciesInput) ([]*applicationautoscaling.ScalingPolicy, error) {
	ret := _m.ctrl.Call(_m, "DescribeScalingPolicies", params)
	ret0, _ := ret[0].([]*applicationautoscaling.ScalingPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) DescribeScalingPolicies(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeScalingPolicies", arg0)
}

func (_m *MockClient) PutScalingPolicy(params *applicationautoscaling.PutScalingPolicyInput) (string, error) {
	ret := _m.ctrl.Call(_m, "PutScalingPolicy", params)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) PutScalingPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutScalingPolicy", arg0)
}

func (_m *MockClient) RegisterScalableTarget(params *applicationautoscaling.RegisterScalableTargetInput) error {
	ret := _m.ctrl.Call(_m, "RegisterScalableTarget", params)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) RegisterScalableTarget(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterScalableTarget", arg0)
}
