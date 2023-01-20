package main

import (
	"os"

	breezewareecs "breezeware-aws-cdk-patterns-samples/ecs"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	ecs "github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	// breezewareecs "github.com/deepak-palanisamy/breezeware-aws-cdk-patterns/ecs"
)

type BreezewareAwsCdkPatternsSamplesStackProps struct {
	awscdk.StackProps
}

func Service(scope constructs.Construct, id string, props *BreezewareAwsCdkPatternsSamplesStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	breezewareecs.NewLoadBalancedEc2Service(stack, jsii.String("BreezewareAwsCdkPatternsSampleLoadBalancedEc2Service"), &breezewareecs.LoadBalancedEc2ServiceProps{
		Cluster: breezewareecs.ClusterProps{
			ClusterName: "ClusterGoLang",
			Vpc: breezewareecs.VpcProps{
				IsDefault: true,
				Id:        "vpc-535bd136",
			},
			SecurityGroups: []ec2.ISecurityGroup{
				ec2.SecurityGroup_FromLookupById(
					stack, jsii.String("VpcSecurityGroup"), jsii.String("sg-0ca01451382c5594c"),
				),
			},
		},
		CapacityProviderStrategies: []string{
			// "GoLangMicroAsgCapacityProvider",
			"GoLangSmallAsgCapacityProvider",
		},
		LogGroupName:              "Service",
		EnableTracing:             true,
		DesiredTaskCount:          1,
		IsServiceDiscoveryEnabled: false,
		// ServiceDiscovery: breezewareecs.ServiceDiscoveryOptions{
		// 	NamespaceName: "brz.demo",
		// 	NamespaceId:   "ns-kbnyr3owncs5677p",
		// 	NamespaceArn:  "arn:aws:servicediscovery:us-east-1:305251478828:namespace/ns-kbnyr3owncs5677p",
		// },
		RoutePriority:               1,
		RoutePath:                   "/api/*",
		RoutePort:                   8443,
		Host:                        "nginx.dynamostack.com",
		IsLoadBalancerEnabled:       true,
		LoadBalancerListenerArn:     "arn:aws:elasticloadbalancing:us-east-1:305251478828:listener/app/ClusterAlb/fbc3ee80eebe84ab/cdbbb1a8a52abce3",
		LoadBalancerSecurityGroupId: "sg-076574aabd3b41d78",
		LoadBalancerHealthCheckPath: "/api/health-status",
		LoadBalancerTargetOptions: ecs.LoadBalancerTargetOptions{
			ContainerName: jsii.String("rpc-service"),
			ContainerPort: jsii.Number(8443),
			Protocol:      ecs.Protocol_TCP,
		},
		// LoadBalancerTargetOptions: breezewareecs.LoadBalancerTargetOptions{

		// 	ContainerName: "demo-app",
		// 	Port:          80,
		// 	Protocol:      breezewareecs.LOAD_BALANCER_TARGET_PROTOCOL_TCP,
		// },
		TaskDefinition: breezewareecs.TaskDefinition{
			FamilyName: "rpc-service",
			// Cpu:         "1024",
			// MemoryInMiB: "2048",
			NetworkMode: breezewareecs.TASK_DEFINTION_NETWORK_MODE_BRIDGE,
			EnvironmentFile: breezewareecs.EnvironmentFile{
				BucketName: "golang-cdk-construct-demo-bucket",
				BucketArn:  "arn:aws:s3:::golang-cdk-construct-demo-bucket",
			},
			RequiresVolume: false,
			// Volumes: []breezewareecs.Volume{
			// 	{
			// 		Name: "demo-volume",
			// 		Size: "10",
			// 	},
			// },
			ApplicationContainers: []breezewareecs.ContainerDefinition{
				{
					ContainerName: "rpc-service",
					Image:         "rpc-service",
					RegistryType:  breezewareecs.CONTAINER_DEFINITION_REGISTRY_AWS_ECR,
					ImageTag:      "latest",
					IsEssential:   true,
					// Commands: []string{
					// 	"--enable-cors", "--api-key=rpcproductionkey", "--data-dir=/data",
					// 	"--cors-domains=https://revolution.film,https://www.revolution.film,http://localhost:3000",
					// },
					// EntryPointCommands: []string{
					// 	"",
					// },
					Cpu:                      512,
					Memory:                   1500,
					EnvironmentFileObjectKey: "rpc-service/prod/app.env",
					// VolumeMountPoint: []ecs.MountPoint{
					// 	{
					// 		ContainerPath: jsii.String("/usr/share/nginx/html"),
					// 		ReadOnly:      jsii.Bool(false),
					// 		SourceVolume:  jsii.String("demo-volume"),
					// 	},
					// },
					PortMappings: []ecs.PortMapping{
						{
							// HostPort:      jsii.Number(8080),
							ContainerPort: jsii.Number(8443),
							Protocol:      ecs.Protocol_TCP,
						},
						// {
						// 	ContainerPort: jsii.Number(443),
						// 	Protocol:      ecs.Protocol_TCP,
						// },
					},
				},
			},
			// TaskPolicy: iam.NewPolicyDocument(&iam.PolicyDocumentProps{
			// 	AssignSids: jsii.Bool(true),
			// 	Statements: &[]iam.PolicyStatement{
			// 		iam.NewPolicyStatement(
			// 			&iam.PolicyStatementProps{
			// 				Actions: &[]*string{
			// 					jsii.String("s3:Get*"),
			// 				},
			// 				Effect: iam.Effect_ALLOW,
			// 				Resources: &[]*string{
			// 					jsii.String("arn:aws:s3:::canoja-ecs-cluster-environment-properties-us-east-1"),
			// 				},
			// 			},
			// 		),
			// 	},
			// }),
		},
	})

	return stack
}

func Db(scope constructs.Construct, id string, props *BreezewareAwsCdkPatternsSamplesStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	breezewareecs.NewLoadBalancedEc2Service(stack, jsii.String("BreezewareAwsCdkPatternsSampleLoadBalancedEc2Service"), &breezewareecs.LoadBalancedEc2ServiceProps{
		Cluster: breezewareecs.ClusterProps{
			ClusterName: "ClusterGoLang",
			Vpc: breezewareecs.VpcProps{
				IsDefault: true,
				Id:        "vpc-535bd136",
			},
			SecurityGroups: []ec2.ISecurityGroup{
				ec2.SecurityGroup_FromLookupById(
					stack, jsii.String("VpcSecurityGroup"), jsii.String("sg-071e1eff270b753c9"),
				),
			},
		},
		CapacityProviderStrategies: []string{
			"GoLangMicroAsgCapacityProvider",
			// "GoLangSmallAsgCapacityProvider",
		},
		LogGroupName:              "Db",
		EnableTracing:             false,
		DesiredTaskCount:          1,
		IsServiceDiscoveryEnabled: true,
		ServiceDiscovery: breezewareecs.ServiceDiscoveryOptions{
			NamespaceName: "brz.demo",
			NamespaceId:   "ns-kbnyr3owncs5677p",
			NamespaceArn:  "arn:aws:servicediscovery:us-east-1:305251478828:namespace/ns-kbnyr3owncs5677p",
		},
		// RoutePriority:         2,
		// RoutePath:             "/*",
		// RoutePort:             80,
		RouteName: "rpc-service-db",
		// Host:                  "nginx1.dynamostack.com",
		IsLoadBalancerEnabled: false,
		// LoadBalancerListenerArn:     "arn:aws:elasticloadbalancing:us-east-1:305251478828:listener/app/ClusterAlb/fbc3ee80eebe84ab/cdbbb1a8a52abce3",
		// LoadBalancerSecurityGroupId: "sg-076574aabd3b41d78",
		// LoadBalancerTargetOptions: ecs.LoadBalancerTargetOptions{
		// 	ContainerName: jsii.String("rpc-service-db"),
		// 	ContainerPort: jsii.Number(80),
		// 	Protocol:      ecs.Protocol_TCP,
		// },
		// LoadBalancerTargetOptions: breezewareecs.LoadBalancerTargetOptions{

		// 	ContainerName: "demo-app",
		// 	Port:          80,
		// 	Protocol:      breezewareecs.LOAD_BALANCER_TARGET_PROTOCOL_TCP,
		// },
		TaskDefinition: breezewareecs.TaskDefinition{
			FamilyName: "rpc-service-db",
			// Cpu:         "1024",
			// MemoryInMiB: "2048",
			NetworkMode: breezewareecs.TASK_DEFINTION_NETWORK_MODE_AWS_VPC,
			EnvironmentFile: breezewareecs.EnvironmentFile{
				BucketName: "golang-cdk-construct-demo-bucket",
				BucketArn:  "arn:aws:s3:::golang-cdk-construct-demo-bucket",
			},
			RequiresVolume: true,
			Volumes: []breezewareecs.Volume{
				{
					Name: "rpc-service-db",
					Size: "10",
				},
			},
			ApplicationContainers: []breezewareecs.ContainerDefinition{
				{
					ContainerName: "rpc-service-db",
					Image:         "rpc-service-db",
					RegistryType:  breezewareecs.CONTAINER_DEFINITION_REGISTRY_AWS_ECR,
					ImageTag:      "latest",
					IsEssential:   true,
					// Commands: []string{
					// 	"--enable-cors", "--api-key=rpcproductionkey", "--data-dir=/data",
					// 	"--cors-domains=https://revolution.film,https://www.revolution.film,http://localhost:3000",
					// },
					// EntryPointCommands: []string{
					// 	"",
					// },
					Cpu:                      512,
					Memory:                   900,
					EnvironmentFileObjectKey: "rpc-service-db/prod/db.env",
					VolumeMountPoint: []ecs.MountPoint{
						{
							ContainerPath: jsii.String("/var/lib/postgresql/data"),
							ReadOnly:      jsii.Bool(false),
							SourceVolume:  jsii.String("rpc-service-db"),
						},
					},
					PortMappings: []ecs.PortMapping{
						{
							// HostPort:      jsii.Number(8080),
							ContainerPort: jsii.Number(5432),
							Protocol:      ecs.Protocol_TCP,
						},
						// {
						// 	ContainerPort: jsii.Number(443),
						// 	Protocol:      ecs.Protocol_TCP,
						// },
					},
				},
			},
			// TaskPolicy: iam.NewPolicyDocument(&iam.PolicyDocumentProps{
			// 	AssignSids: jsii.Bool(true),
			// 	Statements: &[]iam.PolicyStatement{
			// 		iam.NewPolicyStatement(
			// 			&iam.PolicyStatementProps{
			// 				Actions: &[]*string{
			// 					jsii.String("s3:Get*"),
			// 				},
			// 				Effect: iam.Effect_ALLOW,
			// 				Resources: &[]*string{
			// 					jsii.String("arn:aws:s3:::canoja-ecs-cluster-environment-properties-us-east-1"),
			// 				},
			// 			},
			// 		),
			// 	},
			// }),
		},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	Service(app, "Service", &BreezewareAwsCdkPatternsSamplesStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	Db(app, "Db", &BreezewareAwsCdkPatternsSamplesStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
