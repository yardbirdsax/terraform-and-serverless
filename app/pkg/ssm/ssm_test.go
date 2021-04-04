package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type mockSSMClient struct {
	ssmiface.SSMAPI
}

func (m *mockSSMClient) GetParameter(*ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	str := "MockValue"
	val := ssm.GetParameterOutput {
		Parameter: &ssm.Parameter{
			Value: &str,
		},
	}
	return &val, nil
}

func TestSSMGetParameter(t *testing.T) {
	
	expected := "MockValue"
	
	mockSSMClient := &mockSSMClient{}

	val, err := GetSSMParameterValue("param", mockSSMClient)
	if err != nil {
		t.Fatalf("Func returned error: %v",err)
	}

	if val != expected {
		t.Fatalf("Expected: %v; Actual: %v",expected,val)
	}
}