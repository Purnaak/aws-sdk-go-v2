// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new MLModel using the DataSource and the recipe as information
// sources. An MLModel is nearly immutable. Users can update only the MLModelName
// and the ScoreThreshold in an MLModel without creating a new MLModel.
// CreateMLModel is an asynchronous operation. In response to CreateMLModel, Amazon
// Machine Learning (Amazon ML) immediately returns and sets the MLModel status to
// PENDING. After the MLModel has been created and ready is for use, Amazon ML sets
// the status to COMPLETED. You can use the GetMLModel operation to check the
// progress of the MLModel during the creation operation. CreateMLModel requires a
// DataSource with computed statistics, which can be created by setting
// ComputeStatistics to true in CreateDataSourceFromRDS, CreateDataSourceFromS3, or
// CreateDataSourceFromRedshift operations.
func (c *Client) CreateMLModel(ctx context.Context, params *CreateMLModelInput, optFns ...func(*Options)) (*CreateMLModelOutput, error) {
	if params == nil {
		params = &CreateMLModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMLModel", params, optFns, addOperationCreateMLModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMLModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMLModelInput struct {

	// A user-supplied ID that uniquely identifies the MLModel.
	//
	// This member is required.
	MLModelId *string

	// The category of supervised learning that this MLModel will address. Choose from
	// the following types:
	//
	// * Choose REGRESSION if the MLModel will be used to predict
	// a numeric value.
	//
	// * Choose BINARY if the MLModel result has two possible
	// values.
	//
	// * Choose MULTICLASS if the MLModel result has a limited number of
	// values.
	//
	// For more information, see the Amazon Machine Learning Developer Guide
	// (https://docs.aws.amazon.com/machine-learning/latest/dg).
	//
	// This member is required.
	MLModelType types.MLModelType

	// The DataSource that points to the training data.
	//
	// This member is required.
	TrainingDataSourceId *string

	// A user-supplied name or description of the MLModel.
	MLModelName *string

	// A list of the training parameters in the MLModel. The list is implemented as a
	// map of key-value pairs. The following is the current set of training
	// parameters:
	//
	// * sgd.maxMLModelSizeInBytes - The maximum allowed size of the
	// model. Depending on the input data, the size of the model might affect its
	// performance. The value is an integer that ranges from 100000 to 2147483648. The
	// default value is 33554432.
	//
	// * sgd.maxPasses - The number of times that the
	// training process traverses the observations to build the MLModel. The value is
	// an integer that ranges from 1 to 100. The default value is 10.
	//
	// *
	// sgd.shuffleType - Whether Amazon ML shuffles the training data. Shuffling the
	// data improves a model's ability to find the optimal solution for a variety of
	// data types. The valid values are auto and none. The default value is none. We
	// strongly recommend that you shuffle your data.
	//
	// * sgd.l1RegularizationAmount -
	// The coefficient regularization L1 norm. It controls overfitting the data by
	// penalizing large coefficients. This tends to drive coefficients to zero,
	// resulting in a sparse feature set. If you use this parameter, start by
	// specifying a small value, such as 1.0E-08. The value is a double that ranges
	// from 0 to MAX_DOUBLE. The default is to not use L1 normalization. This parameter
	// can't be used when L2 is specified. Use this parameter sparingly.
	//
	// *
	// sgd.l2RegularizationAmount - The coefficient regularization L2 norm. It controls
	// overfitting the data by penalizing large coefficients. This tends to drive
	// coefficients to small, nonzero values. If you use this parameter, start by
	// specifying a small value, such as 1.0E-08. The value is a double that ranges
	// from 0 to MAX_DOUBLE. The default is to not use L2 normalization. This parameter
	// can't be used when L1 is specified. Use this parameter sparingly.
	Parameters map[string]*string

	// The data recipe for creating the MLModel. You must specify either the recipe or
	// its URI. If you don't specify a recipe or its URI, Amazon ML creates a default.
	Recipe *string

	// The Amazon Simple Storage Service (Amazon S3) location and file name that
	// contains the MLModel recipe. You must specify either the recipe or its URI. If
	// you don't specify a recipe or its URI, Amazon ML creates a default.
	RecipeUri *string
}

// Represents the output of a CreateMLModel operation, and is an acknowledgement
// that Amazon ML received the request. The CreateMLModel operation is
// asynchronous. You can poll for status updates by using the GetMLModel operation
// and checking the Status parameter.
type CreateMLModelOutput struct {

	// A user-supplied ID that uniquely identifies the MLModel. This value should be
	// identical to the value of the MLModelId in the request.
	MLModelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMLModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMLModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMLModel{}, middleware.After)
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
	if err = addOpCreateMLModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMLModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMLModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "CreateMLModel",
	}
}