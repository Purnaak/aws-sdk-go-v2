// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists processing jobs that satisfy various filters.
func (c *Client) ListProcessingJobs(ctx context.Context, params *ListProcessingJobsInput, optFns ...func(*Options)) (*ListProcessingJobsOutput, error) {
	if params == nil {
		params = &ListProcessingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProcessingJobs", params, optFns, addOperationListProcessingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProcessingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProcessingJobsInput struct {

	// A filter that returns only processing jobs created after the specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only processing jobs created after the specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only processing jobs modified after the specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only processing jobs modified before the specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of processing jobs to return in the response.
	MaxResults *int32

	// A string in the processing job name. This filter returns only processing jobs
	// whose name contains the specified string.
	NameContains *string

	// If the result of the previous ListProcessingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of processing jobs, use
	// the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.SortBy

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that retrieves only processing jobs with a specific status.
	StatusEquals types.ProcessingJobStatus
}

type ListProcessingJobsOutput struct {

	// An array of ProcessingJobSummary objects, each listing a processing job.
	//
	// This member is required.
	ProcessingJobSummaries []*types.ProcessingJobSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of processing jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProcessingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProcessingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProcessingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProcessingJobs(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListProcessingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListProcessingJobs",
	}
}