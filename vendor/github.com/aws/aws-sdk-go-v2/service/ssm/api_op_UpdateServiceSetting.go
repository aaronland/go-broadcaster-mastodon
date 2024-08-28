// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// ServiceSetting is an account-level setting for an Amazon Web Services service.
// This setting defines how a user interacts with or uses a service or a feature of
// a service. For example, if an Amazon Web Services service charges money to the
// account based on feature or service usage, then the Amazon Web Services service
// team might create a default setting of "false". This means the user can't use
// this feature unless they change the setting to "true" and intentionally opt in
// for a paid feature. Services map a SettingId object to a setting value. Amazon
// Web Services services teams define the default value for a SettingId . You can't
// create a new SettingId , but you can overwrite the default value if you have the
// ssm:UpdateServiceSetting permission for the setting. Use the GetServiceSetting
// API operation to view the current value. Or, use the ResetServiceSetting to
// change the value back to the original value defined by the Amazon Web Services
// service team. Update the service setting for the account.
func (c *Client) UpdateServiceSetting(ctx context.Context, params *UpdateServiceSettingInput, optFns ...func(*Options)) (*UpdateServiceSettingOutput, error) {
	if params == nil {
		params = &UpdateServiceSettingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServiceSetting", params, optFns, c.addOperationUpdateServiceSettingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body of the UpdateServiceSetting API operation.
type UpdateServiceSettingInput struct {

	// The Amazon Resource Name (ARN) of the service setting to update. For example,
	// arn:aws:ssm:us-east-1:111122223333:servicesetting/ssm/parameter-store/high-throughput-enabled
	// . The setting ID can be one of the following.
	//   - /ssm/managed-instance/default-ec2-instance-management-role
	//   - /ssm/automation/customer-script-log-destination
	//   - /ssm/automation/customer-script-log-group-name
	//   - /ssm/documents/console/public-sharing-permission
	//   - /ssm/managed-instance/activation-tier
	//   - /ssm/opsinsights/opscenter
	//   - /ssm/parameter-store/default-parameter-tier
	//   - /ssm/parameter-store/high-throughput-enabled
	// Permissions to update the
	// /ssm/managed-instance/default-ec2-instance-management-role setting should only
	// be provided to administrators. Implement least privilege access when allowing
	// individuals to configure or modify the Default Host Management Configuration.
	//
	// This member is required.
	SettingId *string

	// The new value to specify for the service setting. The following list specifies
	// the available values for each setting.
	//   - For /ssm/managed-instance/default-ec2-instance-management-role , enter the
	//   name of an IAM role.
	//   - For /ssm/automation/customer-script-log-destination , enter CloudWatch .
	//   - For /ssm/automation/customer-script-log-group-name , enter the name of an
	//   Amazon CloudWatch Logs log group.
	//   - For /ssm/documents/console/public-sharing-permission , enter Enable or
	//   Disable .
	//   - For /ssm/managed-instance/activation-tier , enter standard or advanced .
	//   - For /ssm/opsinsights/opscenter , enter Enabled or Disabled .
	//   - For /ssm/parameter-store/default-parameter-tier , enter Standard , Advanced
	//   , or Intelligent-Tiering
	//   - For /ssm/parameter-store/high-throughput-enabled , enter true or false .
	//
	// This member is required.
	SettingValue *string

	noSmithyDocumentSerde
}

// The result body of the UpdateServiceSetting API operation.
type UpdateServiceSettingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceSettingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServiceSetting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServiceSetting{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateServiceSetting"); err != nil {
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
	if err = addOpUpdateServiceSettingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceSetting(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServiceSetting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateServiceSetting",
	}
}
