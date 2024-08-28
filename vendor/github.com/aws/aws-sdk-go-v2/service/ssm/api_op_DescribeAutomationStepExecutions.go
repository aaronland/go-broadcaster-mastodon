// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Information about all active and terminated step executions in an Automation
// workflow.
func (c *Client) DescribeAutomationStepExecutions(ctx context.Context, params *DescribeAutomationStepExecutionsInput, optFns ...func(*Options)) (*DescribeAutomationStepExecutionsOutput, error) {
	if params == nil {
		params = &DescribeAutomationStepExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutomationStepExecutions", params, optFns, c.addOperationDescribeAutomationStepExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutomationStepExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutomationStepExecutionsInput struct {

	// The Automation execution ID for which you want step execution descriptions.
	//
	// This member is required.
	AutomationExecutionId *string

	// One or more filters to limit the number of step executions returned by the
	// request.
	Filters []types.StepExecutionFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// Indicates whether to list step executions in reverse order by start time. The
	// default value is 'false'.
	ReverseOrder *bool

	noSmithyDocumentSerde
}

type DescribeAutomationStepExecutionsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// A list of details about the current state of all steps that make up an
	// execution.
	StepExecutions []types.StepExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAutomationStepExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAutomationStepExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAutomationStepExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAutomationStepExecutions"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeAutomationStepExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutomationStepExecutions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeAutomationStepExecutionsAPIClient is a client that implements the
// DescribeAutomationStepExecutions operation.
type DescribeAutomationStepExecutionsAPIClient interface {
	DescribeAutomationStepExecutions(context.Context, *DescribeAutomationStepExecutionsInput, ...func(*Options)) (*DescribeAutomationStepExecutionsOutput, error)
}

var _ DescribeAutomationStepExecutionsAPIClient = (*Client)(nil)

// DescribeAutomationStepExecutionsPaginatorOptions is the paginator options for
// DescribeAutomationStepExecutions
type DescribeAutomationStepExecutionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAutomationStepExecutionsPaginator is a paginator for
// DescribeAutomationStepExecutions
type DescribeAutomationStepExecutionsPaginator struct {
	options   DescribeAutomationStepExecutionsPaginatorOptions
	client    DescribeAutomationStepExecutionsAPIClient
	params    *DescribeAutomationStepExecutionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAutomationStepExecutionsPaginator returns a new
// DescribeAutomationStepExecutionsPaginator
func NewDescribeAutomationStepExecutionsPaginator(client DescribeAutomationStepExecutionsAPIClient, params *DescribeAutomationStepExecutionsInput, optFns ...func(*DescribeAutomationStepExecutionsPaginatorOptions)) *DescribeAutomationStepExecutionsPaginator {
	if params == nil {
		params = &DescribeAutomationStepExecutionsInput{}
	}

	options := DescribeAutomationStepExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAutomationStepExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAutomationStepExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAutomationStepExecutions page.
func (p *DescribeAutomationStepExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAutomationStepExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeAutomationStepExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeAutomationStepExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAutomationStepExecutions",
	}
}
