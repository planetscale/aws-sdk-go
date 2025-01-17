// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cleanroomsmliface provides an interface to enable mocking the cleanrooms-ml service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cleanroomsmliface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cleanroomsml"
)

// CleanRoomsMLAPI provides an interface to enable mocking the
// cleanroomsml.CleanRoomsML service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// cleanrooms-ml.
//	func myFunc(svc cleanroomsmliface.CleanRoomsMLAPI) bool {
//	    // Make svc.CreateAudienceModel request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := cleanroomsml.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCleanRoomsMLClient struct {
//	    cleanroomsmliface.CleanRoomsMLAPI
//	}
//	func (m *mockCleanRoomsMLClient) CreateAudienceModel(input *cleanroomsml.CreateAudienceModelInput) (*cleanroomsml.CreateAudienceModelOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCleanRoomsMLClient{}
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
type CleanRoomsMLAPI interface {
	CreateAudienceModel(*cleanroomsml.CreateAudienceModelInput) (*cleanroomsml.CreateAudienceModelOutput, error)
	CreateAudienceModelWithContext(aws.Context, *cleanroomsml.CreateAudienceModelInput, ...request.Option) (*cleanroomsml.CreateAudienceModelOutput, error)
	CreateAudienceModelRequest(*cleanroomsml.CreateAudienceModelInput) (*request.Request, *cleanroomsml.CreateAudienceModelOutput)

	CreateConfiguredAudienceModel(*cleanroomsml.CreateConfiguredAudienceModelInput) (*cleanroomsml.CreateConfiguredAudienceModelOutput, error)
	CreateConfiguredAudienceModelWithContext(aws.Context, *cleanroomsml.CreateConfiguredAudienceModelInput, ...request.Option) (*cleanroomsml.CreateConfiguredAudienceModelOutput, error)
	CreateConfiguredAudienceModelRequest(*cleanroomsml.CreateConfiguredAudienceModelInput) (*request.Request, *cleanroomsml.CreateConfiguredAudienceModelOutput)

	CreateTrainingDataset(*cleanroomsml.CreateTrainingDatasetInput) (*cleanroomsml.CreateTrainingDatasetOutput, error)
	CreateTrainingDatasetWithContext(aws.Context, *cleanroomsml.CreateTrainingDatasetInput, ...request.Option) (*cleanroomsml.CreateTrainingDatasetOutput, error)
	CreateTrainingDatasetRequest(*cleanroomsml.CreateTrainingDatasetInput) (*request.Request, *cleanroomsml.CreateTrainingDatasetOutput)

	DeleteAudienceGenerationJob(*cleanroomsml.DeleteAudienceGenerationJobInput) (*cleanroomsml.DeleteAudienceGenerationJobOutput, error)
	DeleteAudienceGenerationJobWithContext(aws.Context, *cleanroomsml.DeleteAudienceGenerationJobInput, ...request.Option) (*cleanroomsml.DeleteAudienceGenerationJobOutput, error)
	DeleteAudienceGenerationJobRequest(*cleanroomsml.DeleteAudienceGenerationJobInput) (*request.Request, *cleanroomsml.DeleteAudienceGenerationJobOutput)

	DeleteAudienceModel(*cleanroomsml.DeleteAudienceModelInput) (*cleanroomsml.DeleteAudienceModelOutput, error)
	DeleteAudienceModelWithContext(aws.Context, *cleanroomsml.DeleteAudienceModelInput, ...request.Option) (*cleanroomsml.DeleteAudienceModelOutput, error)
	DeleteAudienceModelRequest(*cleanroomsml.DeleteAudienceModelInput) (*request.Request, *cleanroomsml.DeleteAudienceModelOutput)

	DeleteConfiguredAudienceModel(*cleanroomsml.DeleteConfiguredAudienceModelInput) (*cleanroomsml.DeleteConfiguredAudienceModelOutput, error)
	DeleteConfiguredAudienceModelWithContext(aws.Context, *cleanroomsml.DeleteConfiguredAudienceModelInput, ...request.Option) (*cleanroomsml.DeleteConfiguredAudienceModelOutput, error)
	DeleteConfiguredAudienceModelRequest(*cleanroomsml.DeleteConfiguredAudienceModelInput) (*request.Request, *cleanroomsml.DeleteConfiguredAudienceModelOutput)

	DeleteConfiguredAudienceModelPolicy(*cleanroomsml.DeleteConfiguredAudienceModelPolicyInput) (*cleanroomsml.DeleteConfiguredAudienceModelPolicyOutput, error)
	DeleteConfiguredAudienceModelPolicyWithContext(aws.Context, *cleanroomsml.DeleteConfiguredAudienceModelPolicyInput, ...request.Option) (*cleanroomsml.DeleteConfiguredAudienceModelPolicyOutput, error)
	DeleteConfiguredAudienceModelPolicyRequest(*cleanroomsml.DeleteConfiguredAudienceModelPolicyInput) (*request.Request, *cleanroomsml.DeleteConfiguredAudienceModelPolicyOutput)

	DeleteTrainingDataset(*cleanroomsml.DeleteTrainingDatasetInput) (*cleanroomsml.DeleteTrainingDatasetOutput, error)
	DeleteTrainingDatasetWithContext(aws.Context, *cleanroomsml.DeleteTrainingDatasetInput, ...request.Option) (*cleanroomsml.DeleteTrainingDatasetOutput, error)
	DeleteTrainingDatasetRequest(*cleanroomsml.DeleteTrainingDatasetInput) (*request.Request, *cleanroomsml.DeleteTrainingDatasetOutput)

	GetAudienceGenerationJob(*cleanroomsml.GetAudienceGenerationJobInput) (*cleanroomsml.GetAudienceGenerationJobOutput, error)
	GetAudienceGenerationJobWithContext(aws.Context, *cleanroomsml.GetAudienceGenerationJobInput, ...request.Option) (*cleanroomsml.GetAudienceGenerationJobOutput, error)
	GetAudienceGenerationJobRequest(*cleanroomsml.GetAudienceGenerationJobInput) (*request.Request, *cleanroomsml.GetAudienceGenerationJobOutput)

	GetAudienceModel(*cleanroomsml.GetAudienceModelInput) (*cleanroomsml.GetAudienceModelOutput, error)
	GetAudienceModelWithContext(aws.Context, *cleanroomsml.GetAudienceModelInput, ...request.Option) (*cleanroomsml.GetAudienceModelOutput, error)
	GetAudienceModelRequest(*cleanroomsml.GetAudienceModelInput) (*request.Request, *cleanroomsml.GetAudienceModelOutput)

	GetConfiguredAudienceModel(*cleanroomsml.GetConfiguredAudienceModelInput) (*cleanroomsml.GetConfiguredAudienceModelOutput, error)
	GetConfiguredAudienceModelWithContext(aws.Context, *cleanroomsml.GetConfiguredAudienceModelInput, ...request.Option) (*cleanroomsml.GetConfiguredAudienceModelOutput, error)
	GetConfiguredAudienceModelRequest(*cleanroomsml.GetConfiguredAudienceModelInput) (*request.Request, *cleanroomsml.GetConfiguredAudienceModelOutput)

	GetConfiguredAudienceModelPolicy(*cleanroomsml.GetConfiguredAudienceModelPolicyInput) (*cleanroomsml.GetConfiguredAudienceModelPolicyOutput, error)
	GetConfiguredAudienceModelPolicyWithContext(aws.Context, *cleanroomsml.GetConfiguredAudienceModelPolicyInput, ...request.Option) (*cleanroomsml.GetConfiguredAudienceModelPolicyOutput, error)
	GetConfiguredAudienceModelPolicyRequest(*cleanroomsml.GetConfiguredAudienceModelPolicyInput) (*request.Request, *cleanroomsml.GetConfiguredAudienceModelPolicyOutput)

	GetTrainingDataset(*cleanroomsml.GetTrainingDatasetInput) (*cleanroomsml.GetTrainingDatasetOutput, error)
	GetTrainingDatasetWithContext(aws.Context, *cleanroomsml.GetTrainingDatasetInput, ...request.Option) (*cleanroomsml.GetTrainingDatasetOutput, error)
	GetTrainingDatasetRequest(*cleanroomsml.GetTrainingDatasetInput) (*request.Request, *cleanroomsml.GetTrainingDatasetOutput)

	ListAudienceExportJobs(*cleanroomsml.ListAudienceExportJobsInput) (*cleanroomsml.ListAudienceExportJobsOutput, error)
	ListAudienceExportJobsWithContext(aws.Context, *cleanroomsml.ListAudienceExportJobsInput, ...request.Option) (*cleanroomsml.ListAudienceExportJobsOutput, error)
	ListAudienceExportJobsRequest(*cleanroomsml.ListAudienceExportJobsInput) (*request.Request, *cleanroomsml.ListAudienceExportJobsOutput)

	ListAudienceExportJobsPages(*cleanroomsml.ListAudienceExportJobsInput, func(*cleanroomsml.ListAudienceExportJobsOutput, bool) bool) error
	ListAudienceExportJobsPagesWithContext(aws.Context, *cleanroomsml.ListAudienceExportJobsInput, func(*cleanroomsml.ListAudienceExportJobsOutput, bool) bool, ...request.Option) error

	ListAudienceGenerationJobs(*cleanroomsml.ListAudienceGenerationJobsInput) (*cleanroomsml.ListAudienceGenerationJobsOutput, error)
	ListAudienceGenerationJobsWithContext(aws.Context, *cleanroomsml.ListAudienceGenerationJobsInput, ...request.Option) (*cleanroomsml.ListAudienceGenerationJobsOutput, error)
	ListAudienceGenerationJobsRequest(*cleanroomsml.ListAudienceGenerationJobsInput) (*request.Request, *cleanroomsml.ListAudienceGenerationJobsOutput)

	ListAudienceGenerationJobsPages(*cleanroomsml.ListAudienceGenerationJobsInput, func(*cleanroomsml.ListAudienceGenerationJobsOutput, bool) bool) error
	ListAudienceGenerationJobsPagesWithContext(aws.Context, *cleanroomsml.ListAudienceGenerationJobsInput, func(*cleanroomsml.ListAudienceGenerationJobsOutput, bool) bool, ...request.Option) error

	ListAudienceModels(*cleanroomsml.ListAudienceModelsInput) (*cleanroomsml.ListAudienceModelsOutput, error)
	ListAudienceModelsWithContext(aws.Context, *cleanroomsml.ListAudienceModelsInput, ...request.Option) (*cleanroomsml.ListAudienceModelsOutput, error)
	ListAudienceModelsRequest(*cleanroomsml.ListAudienceModelsInput) (*request.Request, *cleanroomsml.ListAudienceModelsOutput)

	ListAudienceModelsPages(*cleanroomsml.ListAudienceModelsInput, func(*cleanroomsml.ListAudienceModelsOutput, bool) bool) error
	ListAudienceModelsPagesWithContext(aws.Context, *cleanroomsml.ListAudienceModelsInput, func(*cleanroomsml.ListAudienceModelsOutput, bool) bool, ...request.Option) error

	ListConfiguredAudienceModels(*cleanroomsml.ListConfiguredAudienceModelsInput) (*cleanroomsml.ListConfiguredAudienceModelsOutput, error)
	ListConfiguredAudienceModelsWithContext(aws.Context, *cleanroomsml.ListConfiguredAudienceModelsInput, ...request.Option) (*cleanroomsml.ListConfiguredAudienceModelsOutput, error)
	ListConfiguredAudienceModelsRequest(*cleanroomsml.ListConfiguredAudienceModelsInput) (*request.Request, *cleanroomsml.ListConfiguredAudienceModelsOutput)

	ListConfiguredAudienceModelsPages(*cleanroomsml.ListConfiguredAudienceModelsInput, func(*cleanroomsml.ListConfiguredAudienceModelsOutput, bool) bool) error
	ListConfiguredAudienceModelsPagesWithContext(aws.Context, *cleanroomsml.ListConfiguredAudienceModelsInput, func(*cleanroomsml.ListConfiguredAudienceModelsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*cleanroomsml.ListTagsForResourceInput) (*cleanroomsml.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *cleanroomsml.ListTagsForResourceInput, ...request.Option) (*cleanroomsml.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*cleanroomsml.ListTagsForResourceInput) (*request.Request, *cleanroomsml.ListTagsForResourceOutput)

	ListTrainingDatasets(*cleanroomsml.ListTrainingDatasetsInput) (*cleanroomsml.ListTrainingDatasetsOutput, error)
	ListTrainingDatasetsWithContext(aws.Context, *cleanroomsml.ListTrainingDatasetsInput, ...request.Option) (*cleanroomsml.ListTrainingDatasetsOutput, error)
	ListTrainingDatasetsRequest(*cleanroomsml.ListTrainingDatasetsInput) (*request.Request, *cleanroomsml.ListTrainingDatasetsOutput)

	ListTrainingDatasetsPages(*cleanroomsml.ListTrainingDatasetsInput, func(*cleanroomsml.ListTrainingDatasetsOutput, bool) bool) error
	ListTrainingDatasetsPagesWithContext(aws.Context, *cleanroomsml.ListTrainingDatasetsInput, func(*cleanroomsml.ListTrainingDatasetsOutput, bool) bool, ...request.Option) error

	PutConfiguredAudienceModelPolicy(*cleanroomsml.PutConfiguredAudienceModelPolicyInput) (*cleanroomsml.PutConfiguredAudienceModelPolicyOutput, error)
	PutConfiguredAudienceModelPolicyWithContext(aws.Context, *cleanroomsml.PutConfiguredAudienceModelPolicyInput, ...request.Option) (*cleanroomsml.PutConfiguredAudienceModelPolicyOutput, error)
	PutConfiguredAudienceModelPolicyRequest(*cleanroomsml.PutConfiguredAudienceModelPolicyInput) (*request.Request, *cleanroomsml.PutConfiguredAudienceModelPolicyOutput)

	StartAudienceExportJob(*cleanroomsml.StartAudienceExportJobInput) (*cleanroomsml.StartAudienceExportJobOutput, error)
	StartAudienceExportJobWithContext(aws.Context, *cleanroomsml.StartAudienceExportJobInput, ...request.Option) (*cleanroomsml.StartAudienceExportJobOutput, error)
	StartAudienceExportJobRequest(*cleanroomsml.StartAudienceExportJobInput) (*request.Request, *cleanroomsml.StartAudienceExportJobOutput)

	StartAudienceGenerationJob(*cleanroomsml.StartAudienceGenerationJobInput) (*cleanroomsml.StartAudienceGenerationJobOutput, error)
	StartAudienceGenerationJobWithContext(aws.Context, *cleanroomsml.StartAudienceGenerationJobInput, ...request.Option) (*cleanroomsml.StartAudienceGenerationJobOutput, error)
	StartAudienceGenerationJobRequest(*cleanroomsml.StartAudienceGenerationJobInput) (*request.Request, *cleanroomsml.StartAudienceGenerationJobOutput)

	TagResource(*cleanroomsml.TagResourceInput) (*cleanroomsml.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *cleanroomsml.TagResourceInput, ...request.Option) (*cleanroomsml.TagResourceOutput, error)
	TagResourceRequest(*cleanroomsml.TagResourceInput) (*request.Request, *cleanroomsml.TagResourceOutput)

	UntagResource(*cleanroomsml.UntagResourceInput) (*cleanroomsml.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *cleanroomsml.UntagResourceInput, ...request.Option) (*cleanroomsml.UntagResourceOutput, error)
	UntagResourceRequest(*cleanroomsml.UntagResourceInput) (*request.Request, *cleanroomsml.UntagResourceOutput)

	UpdateConfiguredAudienceModel(*cleanroomsml.UpdateConfiguredAudienceModelInput) (*cleanroomsml.UpdateConfiguredAudienceModelOutput, error)
	UpdateConfiguredAudienceModelWithContext(aws.Context, *cleanroomsml.UpdateConfiguredAudienceModelInput, ...request.Option) (*cleanroomsml.UpdateConfiguredAudienceModelOutput, error)
	UpdateConfiguredAudienceModelRequest(*cleanroomsml.UpdateConfiguredAudienceModelInput) (*request.Request, *cleanroomsml.UpdateConfiguredAudienceModelOutput)
}

var _ CleanRoomsMLAPI = (*cleanroomsml.CleanRoomsML)(nil)
