// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the ID format for the specified resource on a per-Region basis. You can
// specify that resources should receive longer IDs (17-character IDs) when they
// are created. This request can only be used to modify longer ID settings for
// resource types that are within the opt-in period. Resources currently in their
// opt-in period include: bundle | conversion-task | customer-gateway |
// dhcp-options | elastic-ip-allocation | elastic-ip-association | export-task |
// flow-log | image | import-task | internet-gateway | network-acl |
// network-acl-association | network-interface | network-interface-attachment |
// prefix-list | route-table | route-table-association | security-group | subnet |
// subnet-cidr-block-association | vpc | vpc-cidr-block-association | vpc-endpoint
// | vpc-peering-connection | vpn-connection | vpn-gateway. This setting applies to
// the IAM user who makes the request; it does not apply to the entire AWS account.
// By default, an IAM user defaults to the same settings as the root user. If
// you're using this action as the root user, then these settings apply to the
// entire account, unless an IAM user explicitly overrides these settings for
// themselves. For more information, see Resource IDs
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/resource-ids.html) in the
// Amazon Elastic Compute Cloud User Guide. Resources created with longer IDs are
// visible to all IAM roles and users, regardless of these settings and provided
// that they have permission to use the relevant Describe command for the resource
// type.
func (c *Client) ModifyIdFormat(ctx context.Context, params *ModifyIdFormatInput, optFns ...func(*Options)) (*ModifyIdFormatOutput, error) {
	if params == nil {
		params = &ModifyIdFormatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyIdFormat", params, optFns, addOperationModifyIdFormatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyIdFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyIdFormatInput struct {

	// The type of resource: bundle | conversion-task | customer-gateway | dhcp-options
	// | elastic-ip-allocation | elastic-ip-association | export-task | flow-log |
	// image | import-task | internet-gateway | network-acl | network-acl-association |
	// network-interface | network-interface-attachment | prefix-list | route-table |
	// route-table-association | security-group | subnet |
	// subnet-cidr-block-association | vpc | vpc-cidr-block-association | vpc-endpoint
	// | vpc-peering-connection | vpn-connection | vpn-gateway. Alternatively, use the
	// all-current option to include all resource types that are currently within their
	// opt-in period for longer IDs.
	//
	// This member is required.
	Resource *string

	// Indicate whether the resource should use longer IDs (17-character IDs).
	//
	// This member is required.
	UseLongIds *bool
}

type ModifyIdFormatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyIdFormatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyIdFormat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyIdFormat{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpModifyIdFormatValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyIdFormat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyIdFormat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyIdFormat",
	}
}
