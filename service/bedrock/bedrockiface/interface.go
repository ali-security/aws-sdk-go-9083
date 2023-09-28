// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package bedrockiface provides an interface to enable mocking the Amazon Bedrock service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package bedrockiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/bedrock"
)

// BedrockAPI provides an interface to enable mocking the
// bedrock.Bedrock service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Bedrock.
//	func myFunc(svc bedrockiface.BedrockAPI) bool {
//	    // Make svc.CreateModelCustomizationJob request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := bedrock.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockBedrockClient struct {
//	    bedrockiface.BedrockAPI
//	}
//	func (m *mockBedrockClient) CreateModelCustomizationJob(input *bedrock.CreateModelCustomizationJobInput) (*bedrock.CreateModelCustomizationJobOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockBedrockClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type BedrockAPI interface {
	CreateModelCustomizationJob(*bedrock.CreateModelCustomizationJobInput) (*bedrock.CreateModelCustomizationJobOutput, error)
	CreateModelCustomizationJobWithContext(aws.Context, *bedrock.CreateModelCustomizationJobInput, ...request.Option) (*bedrock.CreateModelCustomizationJobOutput, error)
	CreateModelCustomizationJobRequest(*bedrock.CreateModelCustomizationJobInput) (*request.Request, *bedrock.CreateModelCustomizationJobOutput)

	DeleteCustomModel(*bedrock.DeleteCustomModelInput) (*bedrock.DeleteCustomModelOutput, error)
	DeleteCustomModelWithContext(aws.Context, *bedrock.DeleteCustomModelInput, ...request.Option) (*bedrock.DeleteCustomModelOutput, error)
	DeleteCustomModelRequest(*bedrock.DeleteCustomModelInput) (*request.Request, *bedrock.DeleteCustomModelOutput)

	DeleteModelInvocationLoggingConfiguration(*bedrock.DeleteModelInvocationLoggingConfigurationInput) (*bedrock.DeleteModelInvocationLoggingConfigurationOutput, error)
	DeleteModelInvocationLoggingConfigurationWithContext(aws.Context, *bedrock.DeleteModelInvocationLoggingConfigurationInput, ...request.Option) (*bedrock.DeleteModelInvocationLoggingConfigurationOutput, error)
	DeleteModelInvocationLoggingConfigurationRequest(*bedrock.DeleteModelInvocationLoggingConfigurationInput) (*request.Request, *bedrock.DeleteModelInvocationLoggingConfigurationOutput)

	GetCustomModel(*bedrock.GetCustomModelInput) (*bedrock.GetCustomModelOutput, error)
	GetCustomModelWithContext(aws.Context, *bedrock.GetCustomModelInput, ...request.Option) (*bedrock.GetCustomModelOutput, error)
	GetCustomModelRequest(*bedrock.GetCustomModelInput) (*request.Request, *bedrock.GetCustomModelOutput)

	GetFoundationModel(*bedrock.GetFoundationModelInput) (*bedrock.GetFoundationModelOutput, error)
	GetFoundationModelWithContext(aws.Context, *bedrock.GetFoundationModelInput, ...request.Option) (*bedrock.GetFoundationModelOutput, error)
	GetFoundationModelRequest(*bedrock.GetFoundationModelInput) (*request.Request, *bedrock.GetFoundationModelOutput)

	GetModelCustomizationJob(*bedrock.GetModelCustomizationJobInput) (*bedrock.GetModelCustomizationJobOutput, error)
	GetModelCustomizationJobWithContext(aws.Context, *bedrock.GetModelCustomizationJobInput, ...request.Option) (*bedrock.GetModelCustomizationJobOutput, error)
	GetModelCustomizationJobRequest(*bedrock.GetModelCustomizationJobInput) (*request.Request, *bedrock.GetModelCustomizationJobOutput)

	GetModelInvocationLoggingConfiguration(*bedrock.GetModelInvocationLoggingConfigurationInput) (*bedrock.GetModelInvocationLoggingConfigurationOutput, error)
	GetModelInvocationLoggingConfigurationWithContext(aws.Context, *bedrock.GetModelInvocationLoggingConfigurationInput, ...request.Option) (*bedrock.GetModelInvocationLoggingConfigurationOutput, error)
	GetModelInvocationLoggingConfigurationRequest(*bedrock.GetModelInvocationLoggingConfigurationInput) (*request.Request, *bedrock.GetModelInvocationLoggingConfigurationOutput)

	ListCustomModels(*bedrock.ListCustomModelsInput) (*bedrock.ListCustomModelsOutput, error)
	ListCustomModelsWithContext(aws.Context, *bedrock.ListCustomModelsInput, ...request.Option) (*bedrock.ListCustomModelsOutput, error)
	ListCustomModelsRequest(*bedrock.ListCustomModelsInput) (*request.Request, *bedrock.ListCustomModelsOutput)

	ListCustomModelsPages(*bedrock.ListCustomModelsInput, func(*bedrock.ListCustomModelsOutput, bool) bool) error
	ListCustomModelsPagesWithContext(aws.Context, *bedrock.ListCustomModelsInput, func(*bedrock.ListCustomModelsOutput, bool) bool, ...request.Option) error

	ListFoundationModels(*bedrock.ListFoundationModelsInput) (*bedrock.ListFoundationModelsOutput, error)
	ListFoundationModelsWithContext(aws.Context, *bedrock.ListFoundationModelsInput, ...request.Option) (*bedrock.ListFoundationModelsOutput, error)
	ListFoundationModelsRequest(*bedrock.ListFoundationModelsInput) (*request.Request, *bedrock.ListFoundationModelsOutput)

	ListModelCustomizationJobs(*bedrock.ListModelCustomizationJobsInput) (*bedrock.ListModelCustomizationJobsOutput, error)
	ListModelCustomizationJobsWithContext(aws.Context, *bedrock.ListModelCustomizationJobsInput, ...request.Option) (*bedrock.ListModelCustomizationJobsOutput, error)
	ListModelCustomizationJobsRequest(*bedrock.ListModelCustomizationJobsInput) (*request.Request, *bedrock.ListModelCustomizationJobsOutput)

	ListModelCustomizationJobsPages(*bedrock.ListModelCustomizationJobsInput, func(*bedrock.ListModelCustomizationJobsOutput, bool) bool) error
	ListModelCustomizationJobsPagesWithContext(aws.Context, *bedrock.ListModelCustomizationJobsInput, func(*bedrock.ListModelCustomizationJobsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*bedrock.ListTagsForResourceInput) (*bedrock.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *bedrock.ListTagsForResourceInput, ...request.Option) (*bedrock.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*bedrock.ListTagsForResourceInput) (*request.Request, *bedrock.ListTagsForResourceOutput)

	PutModelInvocationLoggingConfiguration(*bedrock.PutModelInvocationLoggingConfigurationInput) (*bedrock.PutModelInvocationLoggingConfigurationOutput, error)
	PutModelInvocationLoggingConfigurationWithContext(aws.Context, *bedrock.PutModelInvocationLoggingConfigurationInput, ...request.Option) (*bedrock.PutModelInvocationLoggingConfigurationOutput, error)
	PutModelInvocationLoggingConfigurationRequest(*bedrock.PutModelInvocationLoggingConfigurationInput) (*request.Request, *bedrock.PutModelInvocationLoggingConfigurationOutput)

	StopModelCustomizationJob(*bedrock.StopModelCustomizationJobInput) (*bedrock.StopModelCustomizationJobOutput, error)
	StopModelCustomizationJobWithContext(aws.Context, *bedrock.StopModelCustomizationJobInput, ...request.Option) (*bedrock.StopModelCustomizationJobOutput, error)
	StopModelCustomizationJobRequest(*bedrock.StopModelCustomizationJobInput) (*request.Request, *bedrock.StopModelCustomizationJobOutput)

	TagResource(*bedrock.TagResourceInput) (*bedrock.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *bedrock.TagResourceInput, ...request.Option) (*bedrock.TagResourceOutput, error)
	TagResourceRequest(*bedrock.TagResourceInput) (*request.Request, *bedrock.TagResourceOutput)

	UntagResource(*bedrock.UntagResourceInput) (*bedrock.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *bedrock.UntagResourceInput, ...request.Option) (*bedrock.UntagResourceOutput, error)
	UntagResourceRequest(*bedrock.UntagResourceInput) (*request.Request, *bedrock.UntagResourceOutput)
}

var _ BedrockAPI = (*bedrock.Bedrock)(nil)
