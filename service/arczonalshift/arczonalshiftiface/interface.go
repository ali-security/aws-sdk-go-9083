// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package arczonalshiftiface provides an interface to enable mocking the AWS ARC - Zonal Shift service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package arczonalshiftiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/arczonalshift"
)

// ARCZonalShiftAPI provides an interface to enable mocking the
// arczonalshift.ARCZonalShift service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS ARC - Zonal Shift.
//	func myFunc(svc arczonalshiftiface.ARCZonalShiftAPI) bool {
//	    // Make svc.CancelZonalShift request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := arczonalshift.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockARCZonalShiftClient struct {
//	    arczonalshiftiface.ARCZonalShiftAPI
//	}
//	func (m *mockARCZonalShiftClient) CancelZonalShift(input *arczonalshift.CancelZonalShiftInput) (*arczonalshift.CancelZonalShiftOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockARCZonalShiftClient{}
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
type ARCZonalShiftAPI interface {
	CancelZonalShift(*arczonalshift.CancelZonalShiftInput) (*arczonalshift.CancelZonalShiftOutput, error)
	CancelZonalShiftWithContext(aws.Context, *arczonalshift.CancelZonalShiftInput, ...request.Option) (*arczonalshift.CancelZonalShiftOutput, error)
	CancelZonalShiftRequest(*arczonalshift.CancelZonalShiftInput) (*request.Request, *arczonalshift.CancelZonalShiftOutput)

	CreatePracticeRunConfiguration(*arczonalshift.CreatePracticeRunConfigurationInput) (*arczonalshift.CreatePracticeRunConfigurationOutput, error)
	CreatePracticeRunConfigurationWithContext(aws.Context, *arczonalshift.CreatePracticeRunConfigurationInput, ...request.Option) (*arczonalshift.CreatePracticeRunConfigurationOutput, error)
	CreatePracticeRunConfigurationRequest(*arczonalshift.CreatePracticeRunConfigurationInput) (*request.Request, *arczonalshift.CreatePracticeRunConfigurationOutput)

	DeletePracticeRunConfiguration(*arczonalshift.DeletePracticeRunConfigurationInput) (*arczonalshift.DeletePracticeRunConfigurationOutput, error)
	DeletePracticeRunConfigurationWithContext(aws.Context, *arczonalshift.DeletePracticeRunConfigurationInput, ...request.Option) (*arczonalshift.DeletePracticeRunConfigurationOutput, error)
	DeletePracticeRunConfigurationRequest(*arczonalshift.DeletePracticeRunConfigurationInput) (*request.Request, *arczonalshift.DeletePracticeRunConfigurationOutput)

	GetAutoshiftObserverNotificationStatus(*arczonalshift.GetAutoshiftObserverNotificationStatusInput) (*arczonalshift.GetAutoshiftObserverNotificationStatusOutput, error)
	GetAutoshiftObserverNotificationStatusWithContext(aws.Context, *arczonalshift.GetAutoshiftObserverNotificationStatusInput, ...request.Option) (*arczonalshift.GetAutoshiftObserverNotificationStatusOutput, error)
	GetAutoshiftObserverNotificationStatusRequest(*arczonalshift.GetAutoshiftObserverNotificationStatusInput) (*request.Request, *arczonalshift.GetAutoshiftObserverNotificationStatusOutput)

	GetManagedResource(*arczonalshift.GetManagedResourceInput) (*arczonalshift.GetManagedResourceOutput, error)
	GetManagedResourceWithContext(aws.Context, *arczonalshift.GetManagedResourceInput, ...request.Option) (*arczonalshift.GetManagedResourceOutput, error)
	GetManagedResourceRequest(*arczonalshift.GetManagedResourceInput) (*request.Request, *arczonalshift.GetManagedResourceOutput)

	ListAutoshifts(*arczonalshift.ListAutoshiftsInput) (*arczonalshift.ListAutoshiftsOutput, error)
	ListAutoshiftsWithContext(aws.Context, *arczonalshift.ListAutoshiftsInput, ...request.Option) (*arczonalshift.ListAutoshiftsOutput, error)
	ListAutoshiftsRequest(*arczonalshift.ListAutoshiftsInput) (*request.Request, *arczonalshift.ListAutoshiftsOutput)

	ListAutoshiftsPages(*arczonalshift.ListAutoshiftsInput, func(*arczonalshift.ListAutoshiftsOutput, bool) bool) error
	ListAutoshiftsPagesWithContext(aws.Context, *arczonalshift.ListAutoshiftsInput, func(*arczonalshift.ListAutoshiftsOutput, bool) bool, ...request.Option) error

	ListManagedResources(*arczonalshift.ListManagedResourcesInput) (*arczonalshift.ListManagedResourcesOutput, error)
	ListManagedResourcesWithContext(aws.Context, *arczonalshift.ListManagedResourcesInput, ...request.Option) (*arczonalshift.ListManagedResourcesOutput, error)
	ListManagedResourcesRequest(*arczonalshift.ListManagedResourcesInput) (*request.Request, *arczonalshift.ListManagedResourcesOutput)

	ListManagedResourcesPages(*arczonalshift.ListManagedResourcesInput, func(*arczonalshift.ListManagedResourcesOutput, bool) bool) error
	ListManagedResourcesPagesWithContext(aws.Context, *arczonalshift.ListManagedResourcesInput, func(*arczonalshift.ListManagedResourcesOutput, bool) bool, ...request.Option) error

	ListZonalShifts(*arczonalshift.ListZonalShiftsInput) (*arczonalshift.ListZonalShiftsOutput, error)
	ListZonalShiftsWithContext(aws.Context, *arczonalshift.ListZonalShiftsInput, ...request.Option) (*arczonalshift.ListZonalShiftsOutput, error)
	ListZonalShiftsRequest(*arczonalshift.ListZonalShiftsInput) (*request.Request, *arczonalshift.ListZonalShiftsOutput)

	ListZonalShiftsPages(*arczonalshift.ListZonalShiftsInput, func(*arczonalshift.ListZonalShiftsOutput, bool) bool) error
	ListZonalShiftsPagesWithContext(aws.Context, *arczonalshift.ListZonalShiftsInput, func(*arczonalshift.ListZonalShiftsOutput, bool) bool, ...request.Option) error

	StartZonalShift(*arczonalshift.StartZonalShiftInput) (*arczonalshift.StartZonalShiftOutput, error)
	StartZonalShiftWithContext(aws.Context, *arczonalshift.StartZonalShiftInput, ...request.Option) (*arczonalshift.StartZonalShiftOutput, error)
	StartZonalShiftRequest(*arczonalshift.StartZonalShiftInput) (*request.Request, *arczonalshift.StartZonalShiftOutput)

	UpdateAutoshiftObserverNotificationStatus(*arczonalshift.UpdateAutoshiftObserverNotificationStatusInput) (*arczonalshift.UpdateAutoshiftObserverNotificationStatusOutput, error)
	UpdateAutoshiftObserverNotificationStatusWithContext(aws.Context, *arczonalshift.UpdateAutoshiftObserverNotificationStatusInput, ...request.Option) (*arczonalshift.UpdateAutoshiftObserverNotificationStatusOutput, error)
	UpdateAutoshiftObserverNotificationStatusRequest(*arczonalshift.UpdateAutoshiftObserverNotificationStatusInput) (*request.Request, *arczonalshift.UpdateAutoshiftObserverNotificationStatusOutput)

	UpdatePracticeRunConfiguration(*arczonalshift.UpdatePracticeRunConfigurationInput) (*arczonalshift.UpdatePracticeRunConfigurationOutput, error)
	UpdatePracticeRunConfigurationWithContext(aws.Context, *arczonalshift.UpdatePracticeRunConfigurationInput, ...request.Option) (*arczonalshift.UpdatePracticeRunConfigurationOutput, error)
	UpdatePracticeRunConfigurationRequest(*arczonalshift.UpdatePracticeRunConfigurationInput) (*request.Request, *arczonalshift.UpdatePracticeRunConfigurationOutput)

	UpdateZonalAutoshiftConfiguration(*arczonalshift.UpdateZonalAutoshiftConfigurationInput) (*arczonalshift.UpdateZonalAutoshiftConfigurationOutput, error)
	UpdateZonalAutoshiftConfigurationWithContext(aws.Context, *arczonalshift.UpdateZonalAutoshiftConfigurationInput, ...request.Option) (*arczonalshift.UpdateZonalAutoshiftConfigurationOutput, error)
	UpdateZonalAutoshiftConfigurationRequest(*arczonalshift.UpdateZonalAutoshiftConfigurationInput) (*request.Request, *arczonalshift.UpdateZonalAutoshiftConfigurationOutput)

	UpdateZonalShift(*arczonalshift.UpdateZonalShiftInput) (*arczonalshift.UpdateZonalShiftOutput, error)
	UpdateZonalShiftWithContext(aws.Context, *arczonalshift.UpdateZonalShiftInput, ...request.Option) (*arczonalshift.UpdateZonalShiftOutput, error)
	UpdateZonalShiftRequest(*arczonalshift.UpdateZonalShiftInput) (*request.Request, *arczonalshift.UpdateZonalShiftOutput)
}

var _ ARCZonalShiftAPI = (*arczonalshift.ARCZonalShift)(nil)
