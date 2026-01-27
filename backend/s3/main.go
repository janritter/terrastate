package s3

import (
	"github.com/janritter/terrastate/backend/types"
	helperAPI "github.com/janritter/terrastate/helper"
	creatorAPI "github.com/janritter/terrastate/helper/creator"
	parserAPI "github.com/janritter/terrastate/helper/parser"
)

var backendAttributes = types.BackendAttributes{
	StateFileAttributes: []*types.StateFileAttribute{
		// Credentials and Shared Configuration
		{
			AttributeKey: "region",
			VarKey:       "region",
			ExpectedType: "string",
			Required:     true,
		},
		{
			AttributeKey: "access_key",
			VarKey:       "state_access_key",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "secret_key",
			VarKey:       "state_secret_key",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "iam_endpoint",
			VarKey:       "state_iam_endpoint",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "max_retries",
			VarKey:       "state_max_retries",
			ExpectedType: "int",
			Required:     false,
		},
		{
			AttributeKey: "profile",
			VarKey:       "state_profile",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "shared_credentials_file",
			VarKey:       "state_shared_credentials_file",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "skip_credentials_validation",
			VarKey:       "state_skip_credentials_validation",
			ExpectedType: "bool",
			Required:     false,
		},
		{
			AttributeKey: "skip_region_validation",
			VarKey:       "state_skip_region_validation",
			ExpectedType: "bool",
			Required:     false,
		},
		{
			AttributeKey: "skip_metadata_api_check",
			VarKey:       "state_skip_metadata_api_check",
			ExpectedType: "bool",
			Required:     false,
		},
		{
			AttributeKey: "sts_endpoint",
			VarKey:       "state_sts_endpoint",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "token",
			VarKey:       "state_token",
			ExpectedType: "string",
			Required:     false,
		},
		// Assume Role Configuration
		{
			AttributeKey: "assume_role_duration_seconds",
			VarKey:       "state_assume_role_duration_seconds",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "assume_role_policy",
			VarKey:       "state_assume_role_policy",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "assume_role_policy_arns",
			VarKey:       "state_assume_role_policy_arns",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "assume_role_tags",
			VarKey:       "state_assume_role_tags",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "assume_role_transitive_tag_keys",
			VarKey:       "state_assume_role_transitive_tag_keys",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "external_id",
			VarKey:       "state_external_id",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "role_arn",
			VarKey:       "state_role_arn",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "session_name",
			VarKey:       "state_session_name",
			ExpectedType: "string",
			Required:     false,
		},
		// S3 State Storage
		{
			AttributeKey: "bucket",
			VarKey:       "state_bucket",
			ExpectedType: "string",
			Required:     true,
		},
		{
			AttributeKey: "key",
			VarKey:       "state_key",
			ExpectedType: "string",
			Required:     true,
		},
		{
			AttributeKey: "acl",
			VarKey:       "state_acl",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "encrypt",
			VarKey:       "state_encrypt",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "endpoint",
			VarKey:       "state_endpoint",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "force_path_style",
			VarKey:       "state_force_path_style",
			ExpectedType: "bool",
			Required:     false,
		},
		{
			AttributeKey: "kms_key_id",
			VarKey:       "state_kms_key_id",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "sse_customer_key",
			VarKey:       "state_sse_customer_key",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "use_lockfile",
			VarKey:       "state_use_lockfile",
			ExpectedType: "bool",
			Required:     false,
		},
		{
			AttributeKey: "workspace_key_prefix",
			VarKey:       "state_workspace_key_prefix",
			ExpectedType: "string",
			Required:     false,
		},
		// DynamoDB State Locking
		{
			AttributeKey: "dynamodb_endpoint",
			VarKey:       "state_dynamodb_endpoint",
			ExpectedType: "string",
			Required:     false,
		},
		{
			AttributeKey: "dynamodb_table",
			VarKey:       "state_dynamodb_table",
			ExpectedType: "string",
			Required:     false,
		},
	},
}

type S3BackendAPI interface {
	Generate() error
}

type S3Backend struct {
	parser  parserAPI.ParserAPI
	creator creatorAPI.CreatorAPI
	helper  helperAPI.HelperAPI
}

func NewS3Backend(parser parserAPI.ParserAPI, creator creatorAPI.CreatorAPI, helper helperAPI.HelperAPI, terrastateAttributes map[string]*types.TerrastateAttribute) *S3Backend {
	backendAttributes.TerrastateAttributes = terrastateAttributes

	return &S3Backend{
		parser:  parser,
		creator: creator,
		helper:  helper,
	}
}

func (backend *S3Backend) Generate() {
	backend.parser.Gather(backendAttributes.StateFileAttributes)

	backend.helper.PrintStateFileAttributes(backendAttributes.StateFileAttributes)

	backend.helper.RemoveDotTerraformFolder(backendAttributes.TerrastateAttributes["state_auto_remove_old"].Value.(bool))

	backend.creator.Create(backendAttributes.StateFileAttributes, "s3")
}
