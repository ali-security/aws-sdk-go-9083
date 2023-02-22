// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package opensearchserviceiface provides an interface to enable mocking the Amazon OpenSearch Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package opensearchserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/opensearchservice"
)

// OpenSearchServiceAPI provides an interface to enable mocking the
// opensearchservice.OpenSearchService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon OpenSearch Service.
//	func myFunc(svc opensearchserviceiface.OpenSearchServiceAPI) bool {
//	    // Make svc.AcceptInboundConnection request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := opensearchservice.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockOpenSearchServiceClient struct {
//	    opensearchserviceiface.OpenSearchServiceAPI
//	}
//	func (m *mockOpenSearchServiceClient) AcceptInboundConnection(input *opensearchservice.AcceptInboundConnectionInput) (*opensearchservice.AcceptInboundConnectionOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockOpenSearchServiceClient{}
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
type OpenSearchServiceAPI interface {
	AcceptInboundConnection(*opensearchservice.AcceptInboundConnectionInput) (*opensearchservice.AcceptInboundConnectionOutput, error)
	AcceptInboundConnectionWithContext(aws.Context, *opensearchservice.AcceptInboundConnectionInput, ...request.Option) (*opensearchservice.AcceptInboundConnectionOutput, error)
	AcceptInboundConnectionRequest(*opensearchservice.AcceptInboundConnectionInput) (*request.Request, *opensearchservice.AcceptInboundConnectionOutput)

	AddTags(*opensearchservice.AddTagsInput) (*opensearchservice.AddTagsOutput, error)
	AddTagsWithContext(aws.Context, *opensearchservice.AddTagsInput, ...request.Option) (*opensearchservice.AddTagsOutput, error)
	AddTagsRequest(*opensearchservice.AddTagsInput) (*request.Request, *opensearchservice.AddTagsOutput)

	AssociatePackage(*opensearchservice.AssociatePackageInput) (*opensearchservice.AssociatePackageOutput, error)
	AssociatePackageWithContext(aws.Context, *opensearchservice.AssociatePackageInput, ...request.Option) (*opensearchservice.AssociatePackageOutput, error)
	AssociatePackageRequest(*opensearchservice.AssociatePackageInput) (*request.Request, *opensearchservice.AssociatePackageOutput)

	AuthorizeVpcEndpointAccess(*opensearchservice.AuthorizeVpcEndpointAccessInput) (*opensearchservice.AuthorizeVpcEndpointAccessOutput, error)
	AuthorizeVpcEndpointAccessWithContext(aws.Context, *opensearchservice.AuthorizeVpcEndpointAccessInput, ...request.Option) (*opensearchservice.AuthorizeVpcEndpointAccessOutput, error)
	AuthorizeVpcEndpointAccessRequest(*opensearchservice.AuthorizeVpcEndpointAccessInput) (*request.Request, *opensearchservice.AuthorizeVpcEndpointAccessOutput)

	CancelServiceSoftwareUpdate(*opensearchservice.CancelServiceSoftwareUpdateInput) (*opensearchservice.CancelServiceSoftwareUpdateOutput, error)
	CancelServiceSoftwareUpdateWithContext(aws.Context, *opensearchservice.CancelServiceSoftwareUpdateInput, ...request.Option) (*opensearchservice.CancelServiceSoftwareUpdateOutput, error)
	CancelServiceSoftwareUpdateRequest(*opensearchservice.CancelServiceSoftwareUpdateInput) (*request.Request, *opensearchservice.CancelServiceSoftwareUpdateOutput)

	CreateDomain(*opensearchservice.CreateDomainInput) (*opensearchservice.CreateDomainOutput, error)
	CreateDomainWithContext(aws.Context, *opensearchservice.CreateDomainInput, ...request.Option) (*opensearchservice.CreateDomainOutput, error)
	CreateDomainRequest(*opensearchservice.CreateDomainInput) (*request.Request, *opensearchservice.CreateDomainOutput)

	CreateOutboundConnection(*opensearchservice.CreateOutboundConnectionInput) (*opensearchservice.CreateOutboundConnectionOutput, error)
	CreateOutboundConnectionWithContext(aws.Context, *opensearchservice.CreateOutboundConnectionInput, ...request.Option) (*opensearchservice.CreateOutboundConnectionOutput, error)
	CreateOutboundConnectionRequest(*opensearchservice.CreateOutboundConnectionInput) (*request.Request, *opensearchservice.CreateOutboundConnectionOutput)

	CreatePackage(*opensearchservice.CreatePackageInput) (*opensearchservice.CreatePackageOutput, error)
	CreatePackageWithContext(aws.Context, *opensearchservice.CreatePackageInput, ...request.Option) (*opensearchservice.CreatePackageOutput, error)
	CreatePackageRequest(*opensearchservice.CreatePackageInput) (*request.Request, *opensearchservice.CreatePackageOutput)

	CreateVpcEndpoint(*opensearchservice.CreateVpcEndpointInput) (*opensearchservice.CreateVpcEndpointOutput, error)
	CreateVpcEndpointWithContext(aws.Context, *opensearchservice.CreateVpcEndpointInput, ...request.Option) (*opensearchservice.CreateVpcEndpointOutput, error)
	CreateVpcEndpointRequest(*opensearchservice.CreateVpcEndpointInput) (*request.Request, *opensearchservice.CreateVpcEndpointOutput)

	DeleteDomain(*opensearchservice.DeleteDomainInput) (*opensearchservice.DeleteDomainOutput, error)
	DeleteDomainWithContext(aws.Context, *opensearchservice.DeleteDomainInput, ...request.Option) (*opensearchservice.DeleteDomainOutput, error)
	DeleteDomainRequest(*opensearchservice.DeleteDomainInput) (*request.Request, *opensearchservice.DeleteDomainOutput)

	DeleteInboundConnection(*opensearchservice.DeleteInboundConnectionInput) (*opensearchservice.DeleteInboundConnectionOutput, error)
	DeleteInboundConnectionWithContext(aws.Context, *opensearchservice.DeleteInboundConnectionInput, ...request.Option) (*opensearchservice.DeleteInboundConnectionOutput, error)
	DeleteInboundConnectionRequest(*opensearchservice.DeleteInboundConnectionInput) (*request.Request, *opensearchservice.DeleteInboundConnectionOutput)

	DeleteOutboundConnection(*opensearchservice.DeleteOutboundConnectionInput) (*opensearchservice.DeleteOutboundConnectionOutput, error)
	DeleteOutboundConnectionWithContext(aws.Context, *opensearchservice.DeleteOutboundConnectionInput, ...request.Option) (*opensearchservice.DeleteOutboundConnectionOutput, error)
	DeleteOutboundConnectionRequest(*opensearchservice.DeleteOutboundConnectionInput) (*request.Request, *opensearchservice.DeleteOutboundConnectionOutput)

	DeletePackage(*opensearchservice.DeletePackageInput) (*opensearchservice.DeletePackageOutput, error)
	DeletePackageWithContext(aws.Context, *opensearchservice.DeletePackageInput, ...request.Option) (*opensearchservice.DeletePackageOutput, error)
	DeletePackageRequest(*opensearchservice.DeletePackageInput) (*request.Request, *opensearchservice.DeletePackageOutput)

	DeleteVpcEndpoint(*opensearchservice.DeleteVpcEndpointInput) (*opensearchservice.DeleteVpcEndpointOutput, error)
	DeleteVpcEndpointWithContext(aws.Context, *opensearchservice.DeleteVpcEndpointInput, ...request.Option) (*opensearchservice.DeleteVpcEndpointOutput, error)
	DeleteVpcEndpointRequest(*opensearchservice.DeleteVpcEndpointInput) (*request.Request, *opensearchservice.DeleteVpcEndpointOutput)

	DescribeDomain(*opensearchservice.DescribeDomainInput) (*opensearchservice.DescribeDomainOutput, error)
	DescribeDomainWithContext(aws.Context, *opensearchservice.DescribeDomainInput, ...request.Option) (*opensearchservice.DescribeDomainOutput, error)
	DescribeDomainRequest(*opensearchservice.DescribeDomainInput) (*request.Request, *opensearchservice.DescribeDomainOutput)

	DescribeDomainAutoTunes(*opensearchservice.DescribeDomainAutoTunesInput) (*opensearchservice.DescribeDomainAutoTunesOutput, error)
	DescribeDomainAutoTunesWithContext(aws.Context, *opensearchservice.DescribeDomainAutoTunesInput, ...request.Option) (*opensearchservice.DescribeDomainAutoTunesOutput, error)
	DescribeDomainAutoTunesRequest(*opensearchservice.DescribeDomainAutoTunesInput) (*request.Request, *opensearchservice.DescribeDomainAutoTunesOutput)

	DescribeDomainAutoTunesPages(*opensearchservice.DescribeDomainAutoTunesInput, func(*opensearchservice.DescribeDomainAutoTunesOutput, bool) bool) error
	DescribeDomainAutoTunesPagesWithContext(aws.Context, *opensearchservice.DescribeDomainAutoTunesInput, func(*opensearchservice.DescribeDomainAutoTunesOutput, bool) bool, ...request.Option) error

	DescribeDomainChangeProgress(*opensearchservice.DescribeDomainChangeProgressInput) (*opensearchservice.DescribeDomainChangeProgressOutput, error)
	DescribeDomainChangeProgressWithContext(aws.Context, *opensearchservice.DescribeDomainChangeProgressInput, ...request.Option) (*opensearchservice.DescribeDomainChangeProgressOutput, error)
	DescribeDomainChangeProgressRequest(*opensearchservice.DescribeDomainChangeProgressInput) (*request.Request, *opensearchservice.DescribeDomainChangeProgressOutput)

	DescribeDomainConfig(*opensearchservice.DescribeDomainConfigInput) (*opensearchservice.DescribeDomainConfigOutput, error)
	DescribeDomainConfigWithContext(aws.Context, *opensearchservice.DescribeDomainConfigInput, ...request.Option) (*opensearchservice.DescribeDomainConfigOutput, error)
	DescribeDomainConfigRequest(*opensearchservice.DescribeDomainConfigInput) (*request.Request, *opensearchservice.DescribeDomainConfigOutput)

	DescribeDomains(*opensearchservice.DescribeDomainsInput) (*opensearchservice.DescribeDomainsOutput, error)
	DescribeDomainsWithContext(aws.Context, *opensearchservice.DescribeDomainsInput, ...request.Option) (*opensearchservice.DescribeDomainsOutput, error)
	DescribeDomainsRequest(*opensearchservice.DescribeDomainsInput) (*request.Request, *opensearchservice.DescribeDomainsOutput)

	DescribeDryRunProgress(*opensearchservice.DescribeDryRunProgressInput) (*opensearchservice.DescribeDryRunProgressOutput, error)
	DescribeDryRunProgressWithContext(aws.Context, *opensearchservice.DescribeDryRunProgressInput, ...request.Option) (*opensearchservice.DescribeDryRunProgressOutput, error)
	DescribeDryRunProgressRequest(*opensearchservice.DescribeDryRunProgressInput) (*request.Request, *opensearchservice.DescribeDryRunProgressOutput)

	DescribeInboundConnections(*opensearchservice.DescribeInboundConnectionsInput) (*opensearchservice.DescribeInboundConnectionsOutput, error)
	DescribeInboundConnectionsWithContext(aws.Context, *opensearchservice.DescribeInboundConnectionsInput, ...request.Option) (*opensearchservice.DescribeInboundConnectionsOutput, error)
	DescribeInboundConnectionsRequest(*opensearchservice.DescribeInboundConnectionsInput) (*request.Request, *opensearchservice.DescribeInboundConnectionsOutput)

	DescribeInboundConnectionsPages(*opensearchservice.DescribeInboundConnectionsInput, func(*opensearchservice.DescribeInboundConnectionsOutput, bool) bool) error
	DescribeInboundConnectionsPagesWithContext(aws.Context, *opensearchservice.DescribeInboundConnectionsInput, func(*opensearchservice.DescribeInboundConnectionsOutput, bool) bool, ...request.Option) error

	DescribeInstanceTypeLimits(*opensearchservice.DescribeInstanceTypeLimitsInput) (*opensearchservice.DescribeInstanceTypeLimitsOutput, error)
	DescribeInstanceTypeLimitsWithContext(aws.Context, *opensearchservice.DescribeInstanceTypeLimitsInput, ...request.Option) (*opensearchservice.DescribeInstanceTypeLimitsOutput, error)
	DescribeInstanceTypeLimitsRequest(*opensearchservice.DescribeInstanceTypeLimitsInput) (*request.Request, *opensearchservice.DescribeInstanceTypeLimitsOutput)

	DescribeOutboundConnections(*opensearchservice.DescribeOutboundConnectionsInput) (*opensearchservice.DescribeOutboundConnectionsOutput, error)
	DescribeOutboundConnectionsWithContext(aws.Context, *opensearchservice.DescribeOutboundConnectionsInput, ...request.Option) (*opensearchservice.DescribeOutboundConnectionsOutput, error)
	DescribeOutboundConnectionsRequest(*opensearchservice.DescribeOutboundConnectionsInput) (*request.Request, *opensearchservice.DescribeOutboundConnectionsOutput)

	DescribeOutboundConnectionsPages(*opensearchservice.DescribeOutboundConnectionsInput, func(*opensearchservice.DescribeOutboundConnectionsOutput, bool) bool) error
	DescribeOutboundConnectionsPagesWithContext(aws.Context, *opensearchservice.DescribeOutboundConnectionsInput, func(*opensearchservice.DescribeOutboundConnectionsOutput, bool) bool, ...request.Option) error

	DescribePackages(*opensearchservice.DescribePackagesInput) (*opensearchservice.DescribePackagesOutput, error)
	DescribePackagesWithContext(aws.Context, *opensearchservice.DescribePackagesInput, ...request.Option) (*opensearchservice.DescribePackagesOutput, error)
	DescribePackagesRequest(*opensearchservice.DescribePackagesInput) (*request.Request, *opensearchservice.DescribePackagesOutput)

	DescribePackagesPages(*opensearchservice.DescribePackagesInput, func(*opensearchservice.DescribePackagesOutput, bool) bool) error
	DescribePackagesPagesWithContext(aws.Context, *opensearchservice.DescribePackagesInput, func(*opensearchservice.DescribePackagesOutput, bool) bool, ...request.Option) error

	DescribeReservedInstanceOfferings(*opensearchservice.DescribeReservedInstanceOfferingsInput) (*opensearchservice.DescribeReservedInstanceOfferingsOutput, error)
	DescribeReservedInstanceOfferingsWithContext(aws.Context, *opensearchservice.DescribeReservedInstanceOfferingsInput, ...request.Option) (*opensearchservice.DescribeReservedInstanceOfferingsOutput, error)
	DescribeReservedInstanceOfferingsRequest(*opensearchservice.DescribeReservedInstanceOfferingsInput) (*request.Request, *opensearchservice.DescribeReservedInstanceOfferingsOutput)

	DescribeReservedInstanceOfferingsPages(*opensearchservice.DescribeReservedInstanceOfferingsInput, func(*opensearchservice.DescribeReservedInstanceOfferingsOutput, bool) bool) error
	DescribeReservedInstanceOfferingsPagesWithContext(aws.Context, *opensearchservice.DescribeReservedInstanceOfferingsInput, func(*opensearchservice.DescribeReservedInstanceOfferingsOutput, bool) bool, ...request.Option) error

	DescribeReservedInstances(*opensearchservice.DescribeReservedInstancesInput) (*opensearchservice.DescribeReservedInstancesOutput, error)
	DescribeReservedInstancesWithContext(aws.Context, *opensearchservice.DescribeReservedInstancesInput, ...request.Option) (*opensearchservice.DescribeReservedInstancesOutput, error)
	DescribeReservedInstancesRequest(*opensearchservice.DescribeReservedInstancesInput) (*request.Request, *opensearchservice.DescribeReservedInstancesOutput)

	DescribeReservedInstancesPages(*opensearchservice.DescribeReservedInstancesInput, func(*opensearchservice.DescribeReservedInstancesOutput, bool) bool) error
	DescribeReservedInstancesPagesWithContext(aws.Context, *opensearchservice.DescribeReservedInstancesInput, func(*opensearchservice.DescribeReservedInstancesOutput, bool) bool, ...request.Option) error

	DescribeVpcEndpoints(*opensearchservice.DescribeVpcEndpointsInput) (*opensearchservice.DescribeVpcEndpointsOutput, error)
	DescribeVpcEndpointsWithContext(aws.Context, *opensearchservice.DescribeVpcEndpointsInput, ...request.Option) (*opensearchservice.DescribeVpcEndpointsOutput, error)
	DescribeVpcEndpointsRequest(*opensearchservice.DescribeVpcEndpointsInput) (*request.Request, *opensearchservice.DescribeVpcEndpointsOutput)

	DissociatePackage(*opensearchservice.DissociatePackageInput) (*opensearchservice.DissociatePackageOutput, error)
	DissociatePackageWithContext(aws.Context, *opensearchservice.DissociatePackageInput, ...request.Option) (*opensearchservice.DissociatePackageOutput, error)
	DissociatePackageRequest(*opensearchservice.DissociatePackageInput) (*request.Request, *opensearchservice.DissociatePackageOutput)

	GetCompatibleVersions(*opensearchservice.GetCompatibleVersionsInput) (*opensearchservice.GetCompatibleVersionsOutput, error)
	GetCompatibleVersionsWithContext(aws.Context, *opensearchservice.GetCompatibleVersionsInput, ...request.Option) (*opensearchservice.GetCompatibleVersionsOutput, error)
	GetCompatibleVersionsRequest(*opensearchservice.GetCompatibleVersionsInput) (*request.Request, *opensearchservice.GetCompatibleVersionsOutput)

	GetPackageVersionHistory(*opensearchservice.GetPackageVersionHistoryInput) (*opensearchservice.GetPackageVersionHistoryOutput, error)
	GetPackageVersionHistoryWithContext(aws.Context, *opensearchservice.GetPackageVersionHistoryInput, ...request.Option) (*opensearchservice.GetPackageVersionHistoryOutput, error)
	GetPackageVersionHistoryRequest(*opensearchservice.GetPackageVersionHistoryInput) (*request.Request, *opensearchservice.GetPackageVersionHistoryOutput)

	GetPackageVersionHistoryPages(*opensearchservice.GetPackageVersionHistoryInput, func(*opensearchservice.GetPackageVersionHistoryOutput, bool) bool) error
	GetPackageVersionHistoryPagesWithContext(aws.Context, *opensearchservice.GetPackageVersionHistoryInput, func(*opensearchservice.GetPackageVersionHistoryOutput, bool) bool, ...request.Option) error

	GetUpgradeHistory(*opensearchservice.GetUpgradeHistoryInput) (*opensearchservice.GetUpgradeHistoryOutput, error)
	GetUpgradeHistoryWithContext(aws.Context, *opensearchservice.GetUpgradeHistoryInput, ...request.Option) (*opensearchservice.GetUpgradeHistoryOutput, error)
	GetUpgradeHistoryRequest(*opensearchservice.GetUpgradeHistoryInput) (*request.Request, *opensearchservice.GetUpgradeHistoryOutput)

	GetUpgradeHistoryPages(*opensearchservice.GetUpgradeHistoryInput, func(*opensearchservice.GetUpgradeHistoryOutput, bool) bool) error
	GetUpgradeHistoryPagesWithContext(aws.Context, *opensearchservice.GetUpgradeHistoryInput, func(*opensearchservice.GetUpgradeHistoryOutput, bool) bool, ...request.Option) error

	GetUpgradeStatus(*opensearchservice.GetUpgradeStatusInput) (*opensearchservice.GetUpgradeStatusOutput, error)
	GetUpgradeStatusWithContext(aws.Context, *opensearchservice.GetUpgradeStatusInput, ...request.Option) (*opensearchservice.GetUpgradeStatusOutput, error)
	GetUpgradeStatusRequest(*opensearchservice.GetUpgradeStatusInput) (*request.Request, *opensearchservice.GetUpgradeStatusOutput)

	ListDomainNames(*opensearchservice.ListDomainNamesInput) (*opensearchservice.ListDomainNamesOutput, error)
	ListDomainNamesWithContext(aws.Context, *opensearchservice.ListDomainNamesInput, ...request.Option) (*opensearchservice.ListDomainNamesOutput, error)
	ListDomainNamesRequest(*opensearchservice.ListDomainNamesInput) (*request.Request, *opensearchservice.ListDomainNamesOutput)

	ListDomainsForPackage(*opensearchservice.ListDomainsForPackageInput) (*opensearchservice.ListDomainsForPackageOutput, error)
	ListDomainsForPackageWithContext(aws.Context, *opensearchservice.ListDomainsForPackageInput, ...request.Option) (*opensearchservice.ListDomainsForPackageOutput, error)
	ListDomainsForPackageRequest(*opensearchservice.ListDomainsForPackageInput) (*request.Request, *opensearchservice.ListDomainsForPackageOutput)

	ListDomainsForPackagePages(*opensearchservice.ListDomainsForPackageInput, func(*opensearchservice.ListDomainsForPackageOutput, bool) bool) error
	ListDomainsForPackagePagesWithContext(aws.Context, *opensearchservice.ListDomainsForPackageInput, func(*opensearchservice.ListDomainsForPackageOutput, bool) bool, ...request.Option) error

	ListInstanceTypeDetails(*opensearchservice.ListInstanceTypeDetailsInput) (*opensearchservice.ListInstanceTypeDetailsOutput, error)
	ListInstanceTypeDetailsWithContext(aws.Context, *opensearchservice.ListInstanceTypeDetailsInput, ...request.Option) (*opensearchservice.ListInstanceTypeDetailsOutput, error)
	ListInstanceTypeDetailsRequest(*opensearchservice.ListInstanceTypeDetailsInput) (*request.Request, *opensearchservice.ListInstanceTypeDetailsOutput)

	ListInstanceTypeDetailsPages(*opensearchservice.ListInstanceTypeDetailsInput, func(*opensearchservice.ListInstanceTypeDetailsOutput, bool) bool) error
	ListInstanceTypeDetailsPagesWithContext(aws.Context, *opensearchservice.ListInstanceTypeDetailsInput, func(*opensearchservice.ListInstanceTypeDetailsOutput, bool) bool, ...request.Option) error

	ListPackagesForDomain(*opensearchservice.ListPackagesForDomainInput) (*opensearchservice.ListPackagesForDomainOutput, error)
	ListPackagesForDomainWithContext(aws.Context, *opensearchservice.ListPackagesForDomainInput, ...request.Option) (*opensearchservice.ListPackagesForDomainOutput, error)
	ListPackagesForDomainRequest(*opensearchservice.ListPackagesForDomainInput) (*request.Request, *opensearchservice.ListPackagesForDomainOutput)

	ListPackagesForDomainPages(*opensearchservice.ListPackagesForDomainInput, func(*opensearchservice.ListPackagesForDomainOutput, bool) bool) error
	ListPackagesForDomainPagesWithContext(aws.Context, *opensearchservice.ListPackagesForDomainInput, func(*opensearchservice.ListPackagesForDomainOutput, bool) bool, ...request.Option) error

	ListScheduledActions(*opensearchservice.ListScheduledActionsInput) (*opensearchservice.ListScheduledActionsOutput, error)
	ListScheduledActionsWithContext(aws.Context, *opensearchservice.ListScheduledActionsInput, ...request.Option) (*opensearchservice.ListScheduledActionsOutput, error)
	ListScheduledActionsRequest(*opensearchservice.ListScheduledActionsInput) (*request.Request, *opensearchservice.ListScheduledActionsOutput)

	ListScheduledActionsPages(*opensearchservice.ListScheduledActionsInput, func(*opensearchservice.ListScheduledActionsOutput, bool) bool) error
	ListScheduledActionsPagesWithContext(aws.Context, *opensearchservice.ListScheduledActionsInput, func(*opensearchservice.ListScheduledActionsOutput, bool) bool, ...request.Option) error

	ListTags(*opensearchservice.ListTagsInput) (*opensearchservice.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *opensearchservice.ListTagsInput, ...request.Option) (*opensearchservice.ListTagsOutput, error)
	ListTagsRequest(*opensearchservice.ListTagsInput) (*request.Request, *opensearchservice.ListTagsOutput)

	ListVersions(*opensearchservice.ListVersionsInput) (*opensearchservice.ListVersionsOutput, error)
	ListVersionsWithContext(aws.Context, *opensearchservice.ListVersionsInput, ...request.Option) (*opensearchservice.ListVersionsOutput, error)
	ListVersionsRequest(*opensearchservice.ListVersionsInput) (*request.Request, *opensearchservice.ListVersionsOutput)

	ListVersionsPages(*opensearchservice.ListVersionsInput, func(*opensearchservice.ListVersionsOutput, bool) bool) error
	ListVersionsPagesWithContext(aws.Context, *opensearchservice.ListVersionsInput, func(*opensearchservice.ListVersionsOutput, bool) bool, ...request.Option) error

	ListVpcEndpointAccess(*opensearchservice.ListVpcEndpointAccessInput) (*opensearchservice.ListVpcEndpointAccessOutput, error)
	ListVpcEndpointAccessWithContext(aws.Context, *opensearchservice.ListVpcEndpointAccessInput, ...request.Option) (*opensearchservice.ListVpcEndpointAccessOutput, error)
	ListVpcEndpointAccessRequest(*opensearchservice.ListVpcEndpointAccessInput) (*request.Request, *opensearchservice.ListVpcEndpointAccessOutput)

	ListVpcEndpoints(*opensearchservice.ListVpcEndpointsInput) (*opensearchservice.ListVpcEndpointsOutput, error)
	ListVpcEndpointsWithContext(aws.Context, *opensearchservice.ListVpcEndpointsInput, ...request.Option) (*opensearchservice.ListVpcEndpointsOutput, error)
	ListVpcEndpointsRequest(*opensearchservice.ListVpcEndpointsInput) (*request.Request, *opensearchservice.ListVpcEndpointsOutput)

	ListVpcEndpointsForDomain(*opensearchservice.ListVpcEndpointsForDomainInput) (*opensearchservice.ListVpcEndpointsForDomainOutput, error)
	ListVpcEndpointsForDomainWithContext(aws.Context, *opensearchservice.ListVpcEndpointsForDomainInput, ...request.Option) (*opensearchservice.ListVpcEndpointsForDomainOutput, error)
	ListVpcEndpointsForDomainRequest(*opensearchservice.ListVpcEndpointsForDomainInput) (*request.Request, *opensearchservice.ListVpcEndpointsForDomainOutput)

	PurchaseReservedInstanceOffering(*opensearchservice.PurchaseReservedInstanceOfferingInput) (*opensearchservice.PurchaseReservedInstanceOfferingOutput, error)
	PurchaseReservedInstanceOfferingWithContext(aws.Context, *opensearchservice.PurchaseReservedInstanceOfferingInput, ...request.Option) (*opensearchservice.PurchaseReservedInstanceOfferingOutput, error)
	PurchaseReservedInstanceOfferingRequest(*opensearchservice.PurchaseReservedInstanceOfferingInput) (*request.Request, *opensearchservice.PurchaseReservedInstanceOfferingOutput)

	RejectInboundConnection(*opensearchservice.RejectInboundConnectionInput) (*opensearchservice.RejectInboundConnectionOutput, error)
	RejectInboundConnectionWithContext(aws.Context, *opensearchservice.RejectInboundConnectionInput, ...request.Option) (*opensearchservice.RejectInboundConnectionOutput, error)
	RejectInboundConnectionRequest(*opensearchservice.RejectInboundConnectionInput) (*request.Request, *opensearchservice.RejectInboundConnectionOutput)

	RemoveTags(*opensearchservice.RemoveTagsInput) (*opensearchservice.RemoveTagsOutput, error)
	RemoveTagsWithContext(aws.Context, *opensearchservice.RemoveTagsInput, ...request.Option) (*opensearchservice.RemoveTagsOutput, error)
	RemoveTagsRequest(*opensearchservice.RemoveTagsInput) (*request.Request, *opensearchservice.RemoveTagsOutput)

	RevokeVpcEndpointAccess(*opensearchservice.RevokeVpcEndpointAccessInput) (*opensearchservice.RevokeVpcEndpointAccessOutput, error)
	RevokeVpcEndpointAccessWithContext(aws.Context, *opensearchservice.RevokeVpcEndpointAccessInput, ...request.Option) (*opensearchservice.RevokeVpcEndpointAccessOutput, error)
	RevokeVpcEndpointAccessRequest(*opensearchservice.RevokeVpcEndpointAccessInput) (*request.Request, *opensearchservice.RevokeVpcEndpointAccessOutput)

	StartServiceSoftwareUpdate(*opensearchservice.StartServiceSoftwareUpdateInput) (*opensearchservice.StartServiceSoftwareUpdateOutput, error)
	StartServiceSoftwareUpdateWithContext(aws.Context, *opensearchservice.StartServiceSoftwareUpdateInput, ...request.Option) (*opensearchservice.StartServiceSoftwareUpdateOutput, error)
	StartServiceSoftwareUpdateRequest(*opensearchservice.StartServiceSoftwareUpdateInput) (*request.Request, *opensearchservice.StartServiceSoftwareUpdateOutput)

	UpdateDomainConfig(*opensearchservice.UpdateDomainConfigInput) (*opensearchservice.UpdateDomainConfigOutput, error)
	UpdateDomainConfigWithContext(aws.Context, *opensearchservice.UpdateDomainConfigInput, ...request.Option) (*opensearchservice.UpdateDomainConfigOutput, error)
	UpdateDomainConfigRequest(*opensearchservice.UpdateDomainConfigInput) (*request.Request, *opensearchservice.UpdateDomainConfigOutput)

	UpdatePackage(*opensearchservice.UpdatePackageInput) (*opensearchservice.UpdatePackageOutput, error)
	UpdatePackageWithContext(aws.Context, *opensearchservice.UpdatePackageInput, ...request.Option) (*opensearchservice.UpdatePackageOutput, error)
	UpdatePackageRequest(*opensearchservice.UpdatePackageInput) (*request.Request, *opensearchservice.UpdatePackageOutput)

	UpdateScheduledAction(*opensearchservice.UpdateScheduledActionInput) (*opensearchservice.UpdateScheduledActionOutput, error)
	UpdateScheduledActionWithContext(aws.Context, *opensearchservice.UpdateScheduledActionInput, ...request.Option) (*opensearchservice.UpdateScheduledActionOutput, error)
	UpdateScheduledActionRequest(*opensearchservice.UpdateScheduledActionInput) (*request.Request, *opensearchservice.UpdateScheduledActionOutput)

	UpdateVpcEndpoint(*opensearchservice.UpdateVpcEndpointInput) (*opensearchservice.UpdateVpcEndpointOutput, error)
	UpdateVpcEndpointWithContext(aws.Context, *opensearchservice.UpdateVpcEndpointInput, ...request.Option) (*opensearchservice.UpdateVpcEndpointOutput, error)
	UpdateVpcEndpointRequest(*opensearchservice.UpdateVpcEndpointInput) (*request.Request, *opensearchservice.UpdateVpcEndpointOutput)

	UpgradeDomain(*opensearchservice.UpgradeDomainInput) (*opensearchservice.UpgradeDomainOutput, error)
	UpgradeDomainWithContext(aws.Context, *opensearchservice.UpgradeDomainInput, ...request.Option) (*opensearchservice.UpgradeDomainOutput, error)
	UpgradeDomainRequest(*opensearchservice.UpgradeDomainInput) (*request.Request, *opensearchservice.UpgradeDomainOutput)
}

var _ OpenSearchServiceAPI = (*opensearchservice.OpenSearchService)(nil)
