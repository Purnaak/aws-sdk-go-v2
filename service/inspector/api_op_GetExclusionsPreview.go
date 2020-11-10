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

// Retrieves the exclusions preview (a list of ExclusionPreview objects) specified
// by the preview token. You can obtain the preview token by running the
// CreateExclusionsPreview API.
func (c *Client) GetExclusionsPreview(ctx context.Context, params *GetExclusionsPreviewInput, optFns ...func(*Options)) (*GetExclusionsPreviewOutput, error) {
	if params == nil {
		params = &GetExclusionsPreviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExclusionsPreview", params, optFns, addOperationGetExclusionsPreviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExclusionsPreviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExclusionsPreviewInput struct {

	// The ARN that specifies the assessment template for which the exclusions preview
	// was requested.
	//
	// This member is required.
	AssessmentTemplateArn *string

	// The unique identifier associated of the exclusions preview.
	//
	// This member is required.
	PreviewToken *string

	// The locale into which you want to translate the exclusion's title, description,
	// and recommendation.
	Locale types.Locale

	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 100. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the GetExclusionsPreviewRequest action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// nextToken from the previous response to continue listing data.
	NextToken *string
}

type GetExclusionsPreviewOutput struct {

	// Specifies the status of the request to generate an exclusions preview.
	//
	// This member is required.
	PreviewStatus types.PreviewStatus

	// Information about the exclusions included in the preview.
	ExclusionPreviews []*types.ExclusionPreview

	// When a response is generated, if there is more data to be listed, this
	// parameters is present in the response and contains the value to use for the
	// nextToken parameter in a subsequent pagination request. If there is no more data
	// to be listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetExclusionsPreviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetExclusionsPreview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetExclusionsPreview{}, middleware.After)
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
	if err = addOpGetExclusionsPreviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExclusionsPreview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetExclusionsPreview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "GetExclusionsPreview",
	}
}