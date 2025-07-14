package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ProjectApiKeyResource{}
var _ resource.ResourceWithImportState = &ProjectApiKeyResource{}

func NewProjectApiKeyResource() resource.Resource {
	return &ProjectApiKeyResource{}
}

// ProjectApiKeyResource defines the resource implementation.
type ProjectApiKeyResource struct {
	client *Client
}

// ProjectApiKeyResourceModel describes the resource data model.
type ProjectApiKeyResourceModel struct {
	ID               types.String `tfsdk:"id"`
	ProjectID        types.String `tfsdk:"project_id"`
	Note             types.String `tfsdk:"note"`
	PublicKey        types.String `tfsdk:"public_key"`
	SecretKey        types.String `tfsdk:"secret_key"`
	DisplaySecretKey types.String `tfsdk:"display_secret_key"`
	CreatedAt        types.String `tfsdk:"created_at"`
}

func (r *ProjectApiKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_api_key"
}

func (r *ProjectApiKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Project API Key resource for managing Langfuse project API keys.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "API key identifier",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "Project identifier that the API key belongs to",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"note": schema.StringAttribute{
				MarkdownDescription: "Optional note for the API key",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"public_key": schema.StringAttribute{
				MarkdownDescription: "Public key for the API key",
				Computed:            true,
			},
			"secret_key": schema.StringAttribute{
				MarkdownDescription: "Secret key for the API key (only available on creation)",
				Computed:            true,
				Sensitive:           true,
			},
			"display_secret_key": schema.StringAttribute{
				MarkdownDescription: "Display version of the secret key",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the API key was created",
				Computed:            true,
			},
		},
	}
}

func (r *ProjectApiKeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ProjectApiKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ProjectApiKeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API request
	createReq := CreateApiKeyRequest{}

	// Handle note
	if !data.Note.IsNull() && !data.Note.IsUnknown() {
		note := data.Note.ValueString()
		createReq.Note = &note
	}

	// Create API key
	apiKey, err := r.client.CreateApiKey(data.ProjectID.ValueString(), createReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create API key, got error: %s", err))
		return
	}

	// Update model with response data
	data.ID = types.StringValue(apiKey.ID)
	data.PublicKey = types.StringValue(apiKey.PublicKey)
	data.CreatedAt = types.StringValue(apiKey.CreatedAt)

	// Set the secret key (only available on creation)
	if apiKey.SecretKey != "" {
		data.SecretKey = types.StringValue(apiKey.SecretKey)
	}

	// Set display secret key
	if apiKey.DisplaySecretKey != "" {
		data.DisplaySecretKey = types.StringValue(apiKey.DisplaySecretKey)
	}

	// Handle optional fields
	if apiKey.Note != nil {
		data.Note = types.StringValue(*apiKey.Note)
	}

	// Write logs using the tflog package
	tflog.Trace(ctx, "created an API key resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectApiKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ProjectApiKeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get API key from API
	apiKey, err := r.client.GetApiKey(data.ProjectID.ValueString(), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read API key, got error: %s", err))
		return
	}

	// Update model with fresh data
	data.ID = types.StringValue(apiKey.ID)
	data.PublicKey = types.StringValue(apiKey.PublicKey)
	data.CreatedAt = types.StringValue(apiKey.CreatedAt)

	// Set display secret key
	if apiKey.DisplaySecretKey != "" {
		data.DisplaySecretKey = types.StringValue(apiKey.DisplaySecretKey)
	}

	// Handle optional fields
	if apiKey.Note != nil {
		data.Note = types.StringValue(*apiKey.Note)
	}

	// Note: SecretKey is not returned from read operations, so we keep the state value

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectApiKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// API keys are immutable in Langfuse API - any change requires replacement
	// This is handled by the RequiresReplace plan modifiers in the schema
	resp.Diagnostics.AddError(
		"Update Not Supported",
		"API keys cannot be updated. Any changes require replacement of the resource.",
	)
}

func (r *ProjectApiKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ProjectApiKeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API key
	err := r.client.DeleteApiKey(data.ProjectID.ValueString(), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete API key, got error: %s", err))
		return
	}
}

func (r *ProjectApiKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: project_id:api_key_id
	// This allows importing using both the project ID and API key ID
	// Example: terraform import langfuse_project_api_key.example project-123:api-key-456
	
	// For now, we'll use a simple ID format and expect users to set project_id separately
	// This could be enhanced to parse "project_id:api_key_id" format if needed
	resp.Diagnostics.AddError(
		"Import Not Implemented",
		"Import is not yet implemented for API key resources. Please recreate the resource instead.",
	)
} 