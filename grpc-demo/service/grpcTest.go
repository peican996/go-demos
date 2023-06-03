package service

import (
	"context"
)

var TestService = &testService{}

type testService struct {
}

// ServiceMethodInvoke 定义方法
func (t *testService) ServiceMethodInvoke(_ context.Context, serviceParam *ServiceParam) (*ClientParam, error) {
	stock := t.getResult(serviceParam.GetServiceParam1(), serviceParam.GetServiceParam2())
	return &ClientParam{ClientParam1: stock}, nil
}

func (t *testService) mustEmbedUnimplementedTestServiceServer() {}

func (t *testService) getResult(param1 int32, param2 int32) int32 {
	return param1 / param2
}
