// +build ignore

package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/terraform-providers/terraform-provider-aws/aws/internal/keyvaluetags"
)

const filename = `list_tags_gen.go`

var serviceNames = []string{
	"acmpca",
	"amplify",
	"appmesh",
	"appstream",
	"appsync",
	"athena",
	"backup",
	"cloudhsmv2",
	"cloudwatch",
	"cloudwatchevents",
	"codecommit",
	"codedeploy",
	"codepipeline",
	"cognitoidentity",
	"cognitoidentityprovider",
	"configservice",
	"databasemigrationservice",
	"datasync",
	"dax",
	"devicefarm",
	"directoryservice",
	"docdb",
	"dynamodb",
	"ecr",
	"ecs",
	"efs",
	"eks",
	"elasticache",
	"elasticbeanstalk",
	"elasticsearchservice",
	"firehose",
	"fsx",
	"glue",
	"guardduty",
	"inspector",
	"iot",
	"iotanalytics",
	"iotevents",
	"kafka",
	"kinesisanalytics",
	"kinesisanalyticsv2",
	"kms",
	"lambda",
	"licensemanager",
	"mediaconnect",
	"medialive",
	"mediapackage",
	"mediastore",
	"mq",
	"neptune",
	"opsworks",
	"organizations",
	"qldb",
	"rds",
	"route53resolver",
	"sagemaker",
	"securityhub",
	"sfn",
	"sns",
	"ssm",
	"storagegateway",
	"swf",
	"transfer",
	"waf",
	"workspaces",
}

type TemplateData struct {
	ServiceNames []string
}

func main() {
	// Always sort to reduce any potential generation churn
	sort.Strings(serviceNames)

	templateData := TemplateData{
		ServiceNames: serviceNames,
	}
	templateFuncMap := template.FuncMap{
		"ClientType":                     keyvaluetags.ServiceClientType,
		"ListTagsFunction":               ServiceListTagsFunction,
		"ListTagsInputIdentifierField":   ServiceListTagsInputIdentifierField,
		"ListTagsInputResourceTypeField": ServiceListTagsInputResourceTypeField,
		"ListTagsOutputTagsField":        ServiceListTagsOutputTagsField,
		"Title":                          strings.Title,
	}

	tmpl, err := template.New("listtags").Funcs(templateFuncMap).Parse(templateBody)

	if err != nil {
		log.Fatalf("error parsing template: %s", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		log.Fatalf("error executing template: %s", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		log.Fatalf("error formatting generated file: %s", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		log.Fatalf("error creating file (%s): %s", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		log.Fatalf("error writing to file (%s): %s", filename, err)
	}
}

var templateBody = `
// Code generated by generators/listtags/main.go; DO NOT EDIT.

package keyvaluetags

import (
	"github.com/aws/aws-sdk-go/aws"
{{- range .ServiceNames }}
	"github.com/aws/aws-sdk-go/service/{{ . }}"
{{- end }}
)
{{ range .ServiceNames }}

// {{ . | Title }}ListTags lists {{ . }} service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func {{ . | Title }}ListTags(conn {{ . | ClientType }}, identifier string{{ if . | ListTagsInputResourceTypeField }}, resourceType string{{ end }}) (KeyValueTags, error) {
	input := &{{ . }}.{{ . | ListTagsFunction }}Input{
		{{ . | ListTagsInputIdentifierField }}:   aws.String(identifier),
		{{- if . | ListTagsInputResourceTypeField }}
		{{ . | ListTagsInputResourceTypeField }}: aws.String(resourceType),
		{{- end }}
	}

	output, err := conn.{{ . | ListTagsFunction }}(input)

	if err != nil {
		return New(nil), err
	}

	return {{ . | Title }}KeyValueTags(output.{{ . | ListTagsOutputTagsField }}), nil
}
{{- end }}
`

// ServiceListTagsFunction determines the service tagging function.
func ServiceListTagsFunction(serviceName string) string {
	switch serviceName {
	case "acmpca":
		return "ListTags"
	case "backup":
		return "ListTags"
	case "cloudhsmv2":
		return "ListTags"
	case "dax":
		return "ListTags"
	case "dynamodb":
		return "ListTagsOfResource"
	case "efs":
		return "DescribeTags"
	case "elasticsearchservice":
		return "ListTags"
	case "firehose":
		return "ListTagsForDeliveryStream"
	case "glue":
		return "GetTags"
	case "kms":
		return "ListResourceTags"
	case "lambda":
		return "ListTags"
	case "mq":
		return "ListTags"
	case "opsworks":
		return "ListTags"
	case "redshift":
		return "DescribeTags"
	case "sagemaker":
		return "ListTags"
	case "workspaces":
		return "DescribeTags"
	default:
		return "ListTagsForResource"
	}
}

// ServiceListTagsInputIdentifierField determines the service tag identifier field.
func ServiceListTagsInputIdentifierField(serviceName string) string {
	switch serviceName {
	case "acmpca":
		return "CertificateAuthorityArn"
	case "athena":
		return "ResourceARN"
	case "cloudhsmv2":
		return "ResourceId"
	case "cloudwatch":
		return "ResourceARN"
	case "cloudwatchevents":
		return "ResourceARN"
	case "dax":
		return "ResourceName"
	case "devicefarm":
		return "ResourceARN"
	case "directoryservice":
		return "ResourceId"
	case "docdb":
		return "ResourceName"
	case "efs":
		return "FileSystemId"
	case "elasticache":
		return "ResourceName"
	case "elasticsearchservice":
		return "ARN"
	case "firehose":
		return "DeliveryStreamName"
	case "fsx":
		return "ResourceARN"
	case "kinesisanalytics":
		return "ResourceARN"
	case "kinesisanalyticsv2":
		return "ResourceARN"
	case "kms":
		return "KeyId"
	case "lambda":
		return "Resource"
	case "mediastore":
		return "Resource"
	case "neptune":
		return "ResourceName"
	case "organizations":
		return "ResourceId"
	case "rds":
		return "ResourceName"
	case "redshift":
		return "ResourceName"
	case "ssm":
		return "ResourceId"
	case "storagegateway":
		return "ResourceARN"
	case "transfer":
		return "Arn"
	case "workspaces":
		return "ResourceId"
	case "waf":
		return "ResourceARN"
	default:
		return "ResourceArn"
	}
}

// ServiceListTagsInputResourceTypeField determines the service tagging resource type field.
func ServiceListTagsInputResourceTypeField(serviceName string) string {
	switch serviceName {
	case "ssm":
		return "ResourceType"
	default:
		return ""
	}
}

// ServiceListTagsOutputTagsField determines the service tag field.
func ServiceListTagsOutputTagsField(serviceName string) string {
	switch serviceName {
	case "waf":
		return "TagInfoForResource.TagList"
	case "cloudhsmv2":
		return "TagList"
	case "databasemigrationservice":
		return "TagList"
	case "docdb":
		return "TagList"
	case "elasticache":
		return "TagList"
	case "elasticbeanstalk":
		return "ResourceTags"
	case "elasticsearchservice":
		return "TagList"
	case "neptune":
		return "TagList"
	case "rds":
		return "TagList"
	case "ssm":
		return "TagList"
	case "workspaces":
		return "TagList"
	default:
		return "Tags"
	}
}
