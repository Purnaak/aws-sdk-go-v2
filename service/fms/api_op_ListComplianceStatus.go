// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of PolicyComplianceStatus objects. Use PolicyComplianceStatus
// to get a summary of which member accounts are protected by the specified policy.
func (c *Client) ListComplianceStatus(ctx context.Context, params *ListComplianceStatusInput, optFns ...func(*Options)) (*ListComplianceStatusOutput, error) {
	if params == nil {
		params = &ListComplianceStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComplianceStatus", params, optFns, addOperationListComplianceStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComplianceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComplianceStatusInput struct {

	// The ID of the AWS Firewall Manager policy that you want the details for.
	//
	// This member is required.
	PolicyId *string

	// Specifies the number of PolicyComplianceStatus objects that you want AWS
	// Firewall Manager to return for this request. If you have more
	// PolicyComplianceStatus objects than the number that you specify for MaxResults,
	// the response includes a NextToken value that you can use to get another batch of
	// PolicyComplianceStatus objects.
	MaxResults *int32

	// If you specify a value for MaxResults and you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, AWS Firewall Manager
	// returns a NextToken value in the response that allows you to list another group
	// of PolicyComplianceStatus objects. For the second and subsequent
	// ListComplianceStatus requests, specify the value of NextToken from the previous
	// response to get information about another batch of PolicyComplianceStatus
	// objects.
	NextToken *string
}

type ListComplianceStatusOutput struct {

	// If you have more PolicyComplianceStatus objects than the number that you
	// specified for MaxResults in the request, the response includes a NextToken
	// value. To list more PolicyComplianceStatus objects, submit another
	// ListComplianceStatus request, and specify the NextToken value from the response
	// in the NextToken value in the next request.
	NextToken *string

	// An array of PolicyComplianceStatus objects.
	PolicyComplianceStatusList []*types.PolicyComplianceStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListComplianceStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListComplianceStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListComplianceStatus{}, middleware.After)
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
	if err = addOpListComplianceStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComplianceStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListComplianceStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListComplianceStatus",
	}
}