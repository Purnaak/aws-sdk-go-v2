// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Previews the agents installed on the EC2 instances that are part of the
// specified assessment target.
func (c *Client) PreviewAgents(ctx context.Context, params *PreviewAgentsInput, optFns ...func(*Options)) (*PreviewAgentsOutput, error) {
	if params == nil {
		params = &PreviewAgentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PreviewAgents", params, optFns, addOperationPreviewAgentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PreviewAgentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PreviewAgentsInput struct {

	// The ARN of the assessment target whose agents you want to preview.
	//
	// This member is required.
	PreviewAgentsArn *string

	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 10. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the PreviewAgents action. Subsequent
	// calls to the action fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string
}

type PreviewAgentsOutput struct {

	// The resulting list of agents.
	//
	// This member is required.
	AgentPreviews []*types.AgentPreview

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to be
	// listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPreviewAgentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPreviewAgents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPreviewAgents{}, middleware.After)
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
	if err = addOpPreviewAgentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPreviewAgents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPreviewAgents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "PreviewAgents",
	}
}