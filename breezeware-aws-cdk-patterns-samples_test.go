package main

// import (
// 	"testing"

// 	"github.com/aws/aws-cdk-go/awscdk/v2"
// 	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
// 	"github.com/aws/jsii-runtime-go"
// )

// example tests. To run these tests, uncomment this file along with the
// example resource in breezeware-aws-cdk-patterns-samples_test.go
// func TestBreezewareAwsCdkPatternsSamplesStack(t *testing.T) {
// 	// GIVEN
// 	app := awscdk.NewApp(nil)

// 	// WHEN
// 	stack := NewBreezewareAwsCdkPatternsSamplesStack(app, "MyStack", nil)

// 	// THEN
// 	template := assertions.Template_FromStack(stack)

// 	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
// 		"VisibilityTimeout": 300,
// 	})
// }
