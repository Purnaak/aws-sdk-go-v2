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

// Lists all the experiments in your account. The list can be filtered to show only
// experiments that were created in a specific time range. The list can be sorted
// by experiment name or creation time.
func (c *Client) ListExperiments(ctx context.Context, params *ListExperimentsInput, optFns ...func(*Options)) (*ListExperimentsOutput, error) {
	if params == nil {
		params = &ListExperimentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExperiments", params, optFns, addOperationListExperimentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExperimentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExperimentsInput struct {

	// A filter that returns only experiments created after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only experiments created before the specified time.
	CreatedBefore *time.Time

	// The maximum number of experiments to return in the response. The default value
	// is 10.
	MaxResults *int32

	// If the previous call to ListExperiments didn't return the full set of
	// experiments, the call returns a token for getting the next set of experiments.
	NextToken *string

	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortExperimentsBy

	// The sort order. The default value is Descending.
	SortOrder types.SortOrder
}

type ListExperimentsOutput struct {

	// A list of the summaries of your experiments.
	ExperimentSummaries []*types.ExperimentSummary

	// A token for getting the next set of experiments, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListExperimentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListExperiments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExperiments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExperiments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListExperiments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListExperiments",
	}
}