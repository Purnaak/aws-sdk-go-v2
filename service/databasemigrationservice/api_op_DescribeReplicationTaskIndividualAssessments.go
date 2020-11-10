// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a paginated list of individual assessments based on filter settings.
// These filter settings can specify a combination of premigration assessment runs,
// migration tasks, and assessment status values.
func (c *Client) DescribeReplicationTaskIndividualAssessments(ctx context.Context, params *DescribeReplicationTaskIndividualAssessmentsInput, optFns ...func(*Options)) (*DescribeReplicationTaskIndividualAssessmentsOutput, error) {
	if params == nil {
		params = &DescribeReplicationTaskIndividualAssessmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationTaskIndividualAssessments", params, optFns, addOperationDescribeReplicationTaskIndividualAssessmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationTaskIndividualAssessmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReplicationTaskIndividualAssessmentsInput struct {

	// Filters applied to the individual assessments described in the form of key-value
	// pairs. Valid filter names: replication-task-assessment-run-arn,
	// replication-task-arn, status
	Filters []*types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32
}

//
type DescribeReplicationTaskIndividualAssessmentsOutput struct {

	// A pagination token returned for you to pass to a subsequent request. If you pass
	// this token as the Marker value in a subsequent request, the response includes
	// only records beyond the marker, up to the value specified in the request by
	// MaxRecords.
	Marker *string

	// One or more individual assessments as specified by Filters.
	ReplicationTaskIndividualAssessments []*types.ReplicationTaskIndividualAssessment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReplicationTaskIndividualAssessmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationTaskIndividualAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationTaskIndividualAssessments{}, middleware.After)
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
	if err = addOpDescribeReplicationTaskIndividualAssessmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationTaskIndividualAssessments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReplicationTaskIndividualAssessments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeReplicationTaskIndividualAssessments",
	}
}