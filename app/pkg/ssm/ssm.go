package ssm

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

func GetSSMParameterValue (parameterName string, client ssmiface.SSMAPI) (string, error) {
	withDecryption := true
	val, err := client.GetParameter(&ssm.GetParameterInput{
		Name: &parameterName,
		WithDecryption: &withDecryption,
	})
	if err != nil {
		return "",err
	}
	return *val.Parameter.Value,nil

}

