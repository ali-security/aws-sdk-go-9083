// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codebuildiface provides an interface to enable mocking the AWS CodeBuild service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codebuildiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

// CodeBuildAPI provides an interface to enable mocking the
// codebuild.CodeBuild service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS CodeBuild.
//	func myFunc(svc codebuildiface.CodeBuildAPI) bool {
//	    // Make svc.BatchDeleteBuilds request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := codebuild.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCodeBuildClient struct {
//	    codebuildiface.CodeBuildAPI
//	}
//	func (m *mockCodeBuildClient) BatchDeleteBuilds(input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCodeBuildClient{}
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
type CodeBuildAPI interface {
	BatchDeleteBuilds(*codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error)
	BatchDeleteBuildsWithContext(aws.Context, *codebuild.BatchDeleteBuildsInput, ...request.Option) (*codebuild.BatchDeleteBuildsOutput, error)
	BatchDeleteBuildsRequest(*codebuild.BatchDeleteBuildsInput) (*request.Request, *codebuild.BatchDeleteBuildsOutput)

	BatchGetBuildBatches(*codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error)
	BatchGetBuildBatchesWithContext(aws.Context, *codebuild.BatchGetBuildBatchesInput, ...request.Option) (*codebuild.BatchGetBuildBatchesOutput, error)
	BatchGetBuildBatchesRequest(*codebuild.BatchGetBuildBatchesInput) (*request.Request, *codebuild.BatchGetBuildBatchesOutput)

	BatchGetBuilds(*codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error)
	BatchGetBuildsWithContext(aws.Context, *codebuild.BatchGetBuildsInput, ...request.Option) (*codebuild.BatchGetBuildsOutput, error)
	BatchGetBuildsRequest(*codebuild.BatchGetBuildsInput) (*request.Request, *codebuild.BatchGetBuildsOutput)

	BatchGetFleets(*codebuild.BatchGetFleetsInput) (*codebuild.BatchGetFleetsOutput, error)
	BatchGetFleetsWithContext(aws.Context, *codebuild.BatchGetFleetsInput, ...request.Option) (*codebuild.BatchGetFleetsOutput, error)
	BatchGetFleetsRequest(*codebuild.BatchGetFleetsInput) (*request.Request, *codebuild.BatchGetFleetsOutput)

	BatchGetProjects(*codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error)
	BatchGetProjectsWithContext(aws.Context, *codebuild.BatchGetProjectsInput, ...request.Option) (*codebuild.BatchGetProjectsOutput, error)
	BatchGetProjectsRequest(*codebuild.BatchGetProjectsInput) (*request.Request, *codebuild.BatchGetProjectsOutput)

	BatchGetReportGroups(*codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error)
	BatchGetReportGroupsWithContext(aws.Context, *codebuild.BatchGetReportGroupsInput, ...request.Option) (*codebuild.BatchGetReportGroupsOutput, error)
	BatchGetReportGroupsRequest(*codebuild.BatchGetReportGroupsInput) (*request.Request, *codebuild.BatchGetReportGroupsOutput)

	BatchGetReports(*codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error)
	BatchGetReportsWithContext(aws.Context, *codebuild.BatchGetReportsInput, ...request.Option) (*codebuild.BatchGetReportsOutput, error)
	BatchGetReportsRequest(*codebuild.BatchGetReportsInput) (*request.Request, *codebuild.BatchGetReportsOutput)

	CreateFleet(*codebuild.CreateFleetInput) (*codebuild.CreateFleetOutput, error)
	CreateFleetWithContext(aws.Context, *codebuild.CreateFleetInput, ...request.Option) (*codebuild.CreateFleetOutput, error)
	CreateFleetRequest(*codebuild.CreateFleetInput) (*request.Request, *codebuild.CreateFleetOutput)

	CreateProject(*codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error)
	CreateProjectWithContext(aws.Context, *codebuild.CreateProjectInput, ...request.Option) (*codebuild.CreateProjectOutput, error)
	CreateProjectRequest(*codebuild.CreateProjectInput) (*request.Request, *codebuild.CreateProjectOutput)

	CreateReportGroup(*codebuild.CreateReportGroupInput) (*codebuild.CreateReportGroupOutput, error)
	CreateReportGroupWithContext(aws.Context, *codebuild.CreateReportGroupInput, ...request.Option) (*codebuild.CreateReportGroupOutput, error)
	CreateReportGroupRequest(*codebuild.CreateReportGroupInput) (*request.Request, *codebuild.CreateReportGroupOutput)

	CreateWebhook(*codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error)
	CreateWebhookWithContext(aws.Context, *codebuild.CreateWebhookInput, ...request.Option) (*codebuild.CreateWebhookOutput, error)
	CreateWebhookRequest(*codebuild.CreateWebhookInput) (*request.Request, *codebuild.CreateWebhookOutput)

	DeleteBuildBatch(*codebuild.DeleteBuildBatchInput) (*codebuild.DeleteBuildBatchOutput, error)
	DeleteBuildBatchWithContext(aws.Context, *codebuild.DeleteBuildBatchInput, ...request.Option) (*codebuild.DeleteBuildBatchOutput, error)
	DeleteBuildBatchRequest(*codebuild.DeleteBuildBatchInput) (*request.Request, *codebuild.DeleteBuildBatchOutput)

	DeleteFleet(*codebuild.DeleteFleetInput) (*codebuild.DeleteFleetOutput, error)
	DeleteFleetWithContext(aws.Context, *codebuild.DeleteFleetInput, ...request.Option) (*codebuild.DeleteFleetOutput, error)
	DeleteFleetRequest(*codebuild.DeleteFleetInput) (*request.Request, *codebuild.DeleteFleetOutput)

	DeleteProject(*codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error)
	DeleteProjectWithContext(aws.Context, *codebuild.DeleteProjectInput, ...request.Option) (*codebuild.DeleteProjectOutput, error)
	DeleteProjectRequest(*codebuild.DeleteProjectInput) (*request.Request, *codebuild.DeleteProjectOutput)

	DeleteReport(*codebuild.DeleteReportInput) (*codebuild.DeleteReportOutput, error)
	DeleteReportWithContext(aws.Context, *codebuild.DeleteReportInput, ...request.Option) (*codebuild.DeleteReportOutput, error)
	DeleteReportRequest(*codebuild.DeleteReportInput) (*request.Request, *codebuild.DeleteReportOutput)

	DeleteReportGroup(*codebuild.DeleteReportGroupInput) (*codebuild.DeleteReportGroupOutput, error)
	DeleteReportGroupWithContext(aws.Context, *codebuild.DeleteReportGroupInput, ...request.Option) (*codebuild.DeleteReportGroupOutput, error)
	DeleteReportGroupRequest(*codebuild.DeleteReportGroupInput) (*request.Request, *codebuild.DeleteReportGroupOutput)

	DeleteResourcePolicy(*codebuild.DeleteResourcePolicyInput) (*codebuild.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyWithContext(aws.Context, *codebuild.DeleteResourcePolicyInput, ...request.Option) (*codebuild.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyRequest(*codebuild.DeleteResourcePolicyInput) (*request.Request, *codebuild.DeleteResourcePolicyOutput)

	DeleteSourceCredentials(*codebuild.DeleteSourceCredentialsInput) (*codebuild.DeleteSourceCredentialsOutput, error)
	DeleteSourceCredentialsWithContext(aws.Context, *codebuild.DeleteSourceCredentialsInput, ...request.Option) (*codebuild.DeleteSourceCredentialsOutput, error)
	DeleteSourceCredentialsRequest(*codebuild.DeleteSourceCredentialsInput) (*request.Request, *codebuild.DeleteSourceCredentialsOutput)

	DeleteWebhook(*codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error)
	DeleteWebhookWithContext(aws.Context, *codebuild.DeleteWebhookInput, ...request.Option) (*codebuild.DeleteWebhookOutput, error)
	DeleteWebhookRequest(*codebuild.DeleteWebhookInput) (*request.Request, *codebuild.DeleteWebhookOutput)

	DescribeCodeCoverages(*codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error)
	DescribeCodeCoveragesWithContext(aws.Context, *codebuild.DescribeCodeCoveragesInput, ...request.Option) (*codebuild.DescribeCodeCoveragesOutput, error)
	DescribeCodeCoveragesRequest(*codebuild.DescribeCodeCoveragesInput) (*request.Request, *codebuild.DescribeCodeCoveragesOutput)

	DescribeCodeCoveragesPages(*codebuild.DescribeCodeCoveragesInput, func(*codebuild.DescribeCodeCoveragesOutput, bool) bool) error
	DescribeCodeCoveragesPagesWithContext(aws.Context, *codebuild.DescribeCodeCoveragesInput, func(*codebuild.DescribeCodeCoveragesOutput, bool) bool, ...request.Option) error

	DescribeTestCases(*codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error)
	DescribeTestCasesWithContext(aws.Context, *codebuild.DescribeTestCasesInput, ...request.Option) (*codebuild.DescribeTestCasesOutput, error)
	DescribeTestCasesRequest(*codebuild.DescribeTestCasesInput) (*request.Request, *codebuild.DescribeTestCasesOutput)

	DescribeTestCasesPages(*codebuild.DescribeTestCasesInput, func(*codebuild.DescribeTestCasesOutput, bool) bool) error
	DescribeTestCasesPagesWithContext(aws.Context, *codebuild.DescribeTestCasesInput, func(*codebuild.DescribeTestCasesOutput, bool) bool, ...request.Option) error

	GetReportGroupTrend(*codebuild.GetReportGroupTrendInput) (*codebuild.GetReportGroupTrendOutput, error)
	GetReportGroupTrendWithContext(aws.Context, *codebuild.GetReportGroupTrendInput, ...request.Option) (*codebuild.GetReportGroupTrendOutput, error)
	GetReportGroupTrendRequest(*codebuild.GetReportGroupTrendInput) (*request.Request, *codebuild.GetReportGroupTrendOutput)

	GetResourcePolicy(*codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error)
	GetResourcePolicyWithContext(aws.Context, *codebuild.GetResourcePolicyInput, ...request.Option) (*codebuild.GetResourcePolicyOutput, error)
	GetResourcePolicyRequest(*codebuild.GetResourcePolicyInput) (*request.Request, *codebuild.GetResourcePolicyOutput)

	ImportSourceCredentials(*codebuild.ImportSourceCredentialsInput) (*codebuild.ImportSourceCredentialsOutput, error)
	ImportSourceCredentialsWithContext(aws.Context, *codebuild.ImportSourceCredentialsInput, ...request.Option) (*codebuild.ImportSourceCredentialsOutput, error)
	ImportSourceCredentialsRequest(*codebuild.ImportSourceCredentialsInput) (*request.Request, *codebuild.ImportSourceCredentialsOutput)

	InvalidateProjectCache(*codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error)
	InvalidateProjectCacheWithContext(aws.Context, *codebuild.InvalidateProjectCacheInput, ...request.Option) (*codebuild.InvalidateProjectCacheOutput, error)
	InvalidateProjectCacheRequest(*codebuild.InvalidateProjectCacheInput) (*request.Request, *codebuild.InvalidateProjectCacheOutput)

	ListBuildBatches(*codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error)
	ListBuildBatchesWithContext(aws.Context, *codebuild.ListBuildBatchesInput, ...request.Option) (*codebuild.ListBuildBatchesOutput, error)
	ListBuildBatchesRequest(*codebuild.ListBuildBatchesInput) (*request.Request, *codebuild.ListBuildBatchesOutput)

	ListBuildBatchesPages(*codebuild.ListBuildBatchesInput, func(*codebuild.ListBuildBatchesOutput, bool) bool) error
	ListBuildBatchesPagesWithContext(aws.Context, *codebuild.ListBuildBatchesInput, func(*codebuild.ListBuildBatchesOutput, bool) bool, ...request.Option) error

	ListBuildBatchesForProject(*codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error)
	ListBuildBatchesForProjectWithContext(aws.Context, *codebuild.ListBuildBatchesForProjectInput, ...request.Option) (*codebuild.ListBuildBatchesForProjectOutput, error)
	ListBuildBatchesForProjectRequest(*codebuild.ListBuildBatchesForProjectInput) (*request.Request, *codebuild.ListBuildBatchesForProjectOutput)

	ListBuildBatchesForProjectPages(*codebuild.ListBuildBatchesForProjectInput, func(*codebuild.ListBuildBatchesForProjectOutput, bool) bool) error
	ListBuildBatchesForProjectPagesWithContext(aws.Context, *codebuild.ListBuildBatchesForProjectInput, func(*codebuild.ListBuildBatchesForProjectOutput, bool) bool, ...request.Option) error

	ListBuilds(*codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error)
	ListBuildsWithContext(aws.Context, *codebuild.ListBuildsInput, ...request.Option) (*codebuild.ListBuildsOutput, error)
	ListBuildsRequest(*codebuild.ListBuildsInput) (*request.Request, *codebuild.ListBuildsOutput)

	ListBuildsPages(*codebuild.ListBuildsInput, func(*codebuild.ListBuildsOutput, bool) bool) error
	ListBuildsPagesWithContext(aws.Context, *codebuild.ListBuildsInput, func(*codebuild.ListBuildsOutput, bool) bool, ...request.Option) error

	ListBuildsForProject(*codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error)
	ListBuildsForProjectWithContext(aws.Context, *codebuild.ListBuildsForProjectInput, ...request.Option) (*codebuild.ListBuildsForProjectOutput, error)
	ListBuildsForProjectRequest(*codebuild.ListBuildsForProjectInput) (*request.Request, *codebuild.ListBuildsForProjectOutput)

	ListBuildsForProjectPages(*codebuild.ListBuildsForProjectInput, func(*codebuild.ListBuildsForProjectOutput, bool) bool) error
	ListBuildsForProjectPagesWithContext(aws.Context, *codebuild.ListBuildsForProjectInput, func(*codebuild.ListBuildsForProjectOutput, bool) bool, ...request.Option) error

	ListCuratedEnvironmentImages(*codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error)
	ListCuratedEnvironmentImagesWithContext(aws.Context, *codebuild.ListCuratedEnvironmentImagesInput, ...request.Option) (*codebuild.ListCuratedEnvironmentImagesOutput, error)
	ListCuratedEnvironmentImagesRequest(*codebuild.ListCuratedEnvironmentImagesInput) (*request.Request, *codebuild.ListCuratedEnvironmentImagesOutput)

	ListFleets(*codebuild.ListFleetsInput) (*codebuild.ListFleetsOutput, error)
	ListFleetsWithContext(aws.Context, *codebuild.ListFleetsInput, ...request.Option) (*codebuild.ListFleetsOutput, error)
	ListFleetsRequest(*codebuild.ListFleetsInput) (*request.Request, *codebuild.ListFleetsOutput)

	ListFleetsPages(*codebuild.ListFleetsInput, func(*codebuild.ListFleetsOutput, bool) bool) error
	ListFleetsPagesWithContext(aws.Context, *codebuild.ListFleetsInput, func(*codebuild.ListFleetsOutput, bool) bool, ...request.Option) error

	ListProjects(*codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error)
	ListProjectsWithContext(aws.Context, *codebuild.ListProjectsInput, ...request.Option) (*codebuild.ListProjectsOutput, error)
	ListProjectsRequest(*codebuild.ListProjectsInput) (*request.Request, *codebuild.ListProjectsOutput)

	ListProjectsPages(*codebuild.ListProjectsInput, func(*codebuild.ListProjectsOutput, bool) bool) error
	ListProjectsPagesWithContext(aws.Context, *codebuild.ListProjectsInput, func(*codebuild.ListProjectsOutput, bool) bool, ...request.Option) error

	ListReportGroups(*codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error)
	ListReportGroupsWithContext(aws.Context, *codebuild.ListReportGroupsInput, ...request.Option) (*codebuild.ListReportGroupsOutput, error)
	ListReportGroupsRequest(*codebuild.ListReportGroupsInput) (*request.Request, *codebuild.ListReportGroupsOutput)

	ListReportGroupsPages(*codebuild.ListReportGroupsInput, func(*codebuild.ListReportGroupsOutput, bool) bool) error
	ListReportGroupsPagesWithContext(aws.Context, *codebuild.ListReportGroupsInput, func(*codebuild.ListReportGroupsOutput, bool) bool, ...request.Option) error

	ListReports(*codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error)
	ListReportsWithContext(aws.Context, *codebuild.ListReportsInput, ...request.Option) (*codebuild.ListReportsOutput, error)
	ListReportsRequest(*codebuild.ListReportsInput) (*request.Request, *codebuild.ListReportsOutput)

	ListReportsPages(*codebuild.ListReportsInput, func(*codebuild.ListReportsOutput, bool) bool) error
	ListReportsPagesWithContext(aws.Context, *codebuild.ListReportsInput, func(*codebuild.ListReportsOutput, bool) bool, ...request.Option) error

	ListReportsForReportGroup(*codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error)
	ListReportsForReportGroupWithContext(aws.Context, *codebuild.ListReportsForReportGroupInput, ...request.Option) (*codebuild.ListReportsForReportGroupOutput, error)
	ListReportsForReportGroupRequest(*codebuild.ListReportsForReportGroupInput) (*request.Request, *codebuild.ListReportsForReportGroupOutput)

	ListReportsForReportGroupPages(*codebuild.ListReportsForReportGroupInput, func(*codebuild.ListReportsForReportGroupOutput, bool) bool) error
	ListReportsForReportGroupPagesWithContext(aws.Context, *codebuild.ListReportsForReportGroupInput, func(*codebuild.ListReportsForReportGroupOutput, bool) bool, ...request.Option) error

	ListSharedProjects(*codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error)
	ListSharedProjectsWithContext(aws.Context, *codebuild.ListSharedProjectsInput, ...request.Option) (*codebuild.ListSharedProjectsOutput, error)
	ListSharedProjectsRequest(*codebuild.ListSharedProjectsInput) (*request.Request, *codebuild.ListSharedProjectsOutput)

	ListSharedProjectsPages(*codebuild.ListSharedProjectsInput, func(*codebuild.ListSharedProjectsOutput, bool) bool) error
	ListSharedProjectsPagesWithContext(aws.Context, *codebuild.ListSharedProjectsInput, func(*codebuild.ListSharedProjectsOutput, bool) bool, ...request.Option) error

	ListSharedReportGroups(*codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error)
	ListSharedReportGroupsWithContext(aws.Context, *codebuild.ListSharedReportGroupsInput, ...request.Option) (*codebuild.ListSharedReportGroupsOutput, error)
	ListSharedReportGroupsRequest(*codebuild.ListSharedReportGroupsInput) (*request.Request, *codebuild.ListSharedReportGroupsOutput)

	ListSharedReportGroupsPages(*codebuild.ListSharedReportGroupsInput, func(*codebuild.ListSharedReportGroupsOutput, bool) bool) error
	ListSharedReportGroupsPagesWithContext(aws.Context, *codebuild.ListSharedReportGroupsInput, func(*codebuild.ListSharedReportGroupsOutput, bool) bool, ...request.Option) error

	ListSourceCredentials(*codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error)
	ListSourceCredentialsWithContext(aws.Context, *codebuild.ListSourceCredentialsInput, ...request.Option) (*codebuild.ListSourceCredentialsOutput, error)
	ListSourceCredentialsRequest(*codebuild.ListSourceCredentialsInput) (*request.Request, *codebuild.ListSourceCredentialsOutput)

	PutResourcePolicy(*codebuild.PutResourcePolicyInput) (*codebuild.PutResourcePolicyOutput, error)
	PutResourcePolicyWithContext(aws.Context, *codebuild.PutResourcePolicyInput, ...request.Option) (*codebuild.PutResourcePolicyOutput, error)
	PutResourcePolicyRequest(*codebuild.PutResourcePolicyInput) (*request.Request, *codebuild.PutResourcePolicyOutput)

	RetryBuild(*codebuild.RetryBuildInput) (*codebuild.RetryBuildOutput, error)
	RetryBuildWithContext(aws.Context, *codebuild.RetryBuildInput, ...request.Option) (*codebuild.RetryBuildOutput, error)
	RetryBuildRequest(*codebuild.RetryBuildInput) (*request.Request, *codebuild.RetryBuildOutput)

	RetryBuildBatch(*codebuild.RetryBuildBatchInput) (*codebuild.RetryBuildBatchOutput, error)
	RetryBuildBatchWithContext(aws.Context, *codebuild.RetryBuildBatchInput, ...request.Option) (*codebuild.RetryBuildBatchOutput, error)
	RetryBuildBatchRequest(*codebuild.RetryBuildBatchInput) (*request.Request, *codebuild.RetryBuildBatchOutput)

	StartBuild(*codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error)
	StartBuildWithContext(aws.Context, *codebuild.StartBuildInput, ...request.Option) (*codebuild.StartBuildOutput, error)
	StartBuildRequest(*codebuild.StartBuildInput) (*request.Request, *codebuild.StartBuildOutput)

	StartBuildBatch(*codebuild.StartBuildBatchInput) (*codebuild.StartBuildBatchOutput, error)
	StartBuildBatchWithContext(aws.Context, *codebuild.StartBuildBatchInput, ...request.Option) (*codebuild.StartBuildBatchOutput, error)
	StartBuildBatchRequest(*codebuild.StartBuildBatchInput) (*request.Request, *codebuild.StartBuildBatchOutput)

	StopBuild(*codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error)
	StopBuildWithContext(aws.Context, *codebuild.StopBuildInput, ...request.Option) (*codebuild.StopBuildOutput, error)
	StopBuildRequest(*codebuild.StopBuildInput) (*request.Request, *codebuild.StopBuildOutput)

	StopBuildBatch(*codebuild.StopBuildBatchInput) (*codebuild.StopBuildBatchOutput, error)
	StopBuildBatchWithContext(aws.Context, *codebuild.StopBuildBatchInput, ...request.Option) (*codebuild.StopBuildBatchOutput, error)
	StopBuildBatchRequest(*codebuild.StopBuildBatchInput) (*request.Request, *codebuild.StopBuildBatchOutput)

	UpdateFleet(*codebuild.UpdateFleetInput) (*codebuild.UpdateFleetOutput, error)
	UpdateFleetWithContext(aws.Context, *codebuild.UpdateFleetInput, ...request.Option) (*codebuild.UpdateFleetOutput, error)
	UpdateFleetRequest(*codebuild.UpdateFleetInput) (*request.Request, *codebuild.UpdateFleetOutput)

	UpdateProject(*codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error)
	UpdateProjectWithContext(aws.Context, *codebuild.UpdateProjectInput, ...request.Option) (*codebuild.UpdateProjectOutput, error)
	UpdateProjectRequest(*codebuild.UpdateProjectInput) (*request.Request, *codebuild.UpdateProjectOutput)

	UpdateProjectVisibility(*codebuild.UpdateProjectVisibilityInput) (*codebuild.UpdateProjectVisibilityOutput, error)
	UpdateProjectVisibilityWithContext(aws.Context, *codebuild.UpdateProjectVisibilityInput, ...request.Option) (*codebuild.UpdateProjectVisibilityOutput, error)
	UpdateProjectVisibilityRequest(*codebuild.UpdateProjectVisibilityInput) (*request.Request, *codebuild.UpdateProjectVisibilityOutput)

	UpdateReportGroup(*codebuild.UpdateReportGroupInput) (*codebuild.UpdateReportGroupOutput, error)
	UpdateReportGroupWithContext(aws.Context, *codebuild.UpdateReportGroupInput, ...request.Option) (*codebuild.UpdateReportGroupOutput, error)
	UpdateReportGroupRequest(*codebuild.UpdateReportGroupInput) (*request.Request, *codebuild.UpdateReportGroupOutput)

	UpdateWebhook(*codebuild.UpdateWebhookInput) (*codebuild.UpdateWebhookOutput, error)
	UpdateWebhookWithContext(aws.Context, *codebuild.UpdateWebhookInput, ...request.Option) (*codebuild.UpdateWebhookOutput, error)
	UpdateWebhookRequest(*codebuild.UpdateWebhookInput) (*request.Request, *codebuild.UpdateWebhookOutput)
}

var _ CodeBuildAPI = (*codebuild.CodeBuild)(nil)
