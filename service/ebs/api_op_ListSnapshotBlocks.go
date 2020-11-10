// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the blocks in an Amazon Elastic Block Store snapshot.
func (c *Client) ListSnapshotBlocks(ctx context.Context, params *ListSnapshotBlocksInput, optFns ...func(*Options)) (*ListSnapshotBlocksOutput, error) {
	if params == nil {
		params = &ListSnapshotBlocksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSnapshotBlocks", params, optFns, addOperationListSnapshotBlocksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSnapshotBlocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSnapshotBlocksInput struct {

	// The ID of the snapshot from which to get block indexes and block tokens.
	//
	// This member is required.
	SnapshotId *string

	// The number of results to return.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The block index from which the list should start. The list in the response will
	// start from this block index or the next valid block index in the snapshot.
	StartingBlockIndex *int32
}

type ListSnapshotBlocksOutput struct {

	// The size of the block.
	BlockSize *int32

	// An array of objects containing information about the blocks.
	Blocks []*types.Block

	// The time when the BlockToken expires.
	ExpiryTime *time.Time

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The size of the volume in GB.
	VolumeSize *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSnapshotBlocksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSnapshotBlocks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSnapshotBlocks{}, middleware.After)
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
	if err = addOpListSnapshotBlocksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSnapshotBlocks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSnapshotBlocks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "ListSnapshotBlocks",
	}
}