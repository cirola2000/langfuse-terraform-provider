package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure LangfuseProvider satisfies various provider interfaces.
var _ provider.Provider = &LangfuseProvider{}

// LangfuseProvider defines the provider implementation.
type LangfuseProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// LangfuseProviderModel describes the provider data model.
type LangfuseProviderModel struct {
	ApiHost   types.String `tfsdk:"api_host"`
	SecretKey types.String `tfsdk:"secret_key"`
	PublicKey types.String `tfsdk:"public_key"`
}

func (p *LangfuseProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "langfuse"
	resp.Version = p.version
}

func (p *LangfuseProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_host": schema.StringAttribute{
				MarkdownDescription: "Langfuse API host URL",
				Optional:            true,
			},
			"secret_key": schema.StringAttribute{
				MarkdownDescription: "Langfuse secret key for authentication",
				Optional:            true,
				Sensitive:           true,
			},
			"public_key": schema.StringAttribute{
				MarkdownDescription: "Langfuse public key for authentication",
				Optional:            true,
			},
		},
	}
}

func (p *LangfuseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data LangfuseProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// Example code to configure a HTTP client...

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	apiHost := os.Getenv("LANGFUSE_API_HOST")
	secretKey := os.Getenv("LANGFUSE_SECRET_KEY")
	publicKey := os.Getenv("LANGFUSE_PUBLIC_KEY")

	if !data.ApiHost.IsNull() {
		apiHost = data.ApiHost.ValueString()
	}

	if !data.SecretKey.IsNull() {
		secretKey = data.SecretKey.ValueString()
	}

	if !data.PublicKey.IsNull() {
		publicKey = data.PublicKey.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if apiHost == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_host"),
			"Missing Langfuse API Host",
			"The provider cannot create the Langfuse API client as there is a missing or empty value for the Langfuse API host. "+
				"Set the api_host value in the configuration or use the LANGFUSE_API_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if secretKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("secret_key"),
			"Missing Langfuse Secret Key",
			"The provider cannot create the Langfuse API client as there is a missing or empty value for the Langfuse secret key. "+
				"Set the secret_key value in the configuration or use the LANGFUSE_SECRET_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if publicKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("public_key"),
			"Missing Langfuse Public Key",
			"The provider cannot create the Langfuse API client as there is a missing or empty value for the Langfuse public key. "+
				"Set the public_key value in the configuration or use the LANGFUSE_PUBLIC_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "langfuse_api_host", apiHost)
	ctx = tflog.SetField(ctx, "langfuse_secret_key", secretKey)
	ctx = tflog.SetField(ctx, "langfuse_public_key", publicKey)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "langfuse_secret_key")

	tflog.Debug(ctx, "Creating Langfuse client")

	// Create a new Langfuse client using the configuration values
	client := NewClient(apiHost, secretKey, publicKey)

	// Make the Langfuse client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Langfuse client", map[string]any{"success": true})
}

func (p *LangfuseProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProjectResource,
		NewProjectApiKeyResource,
	}
}

func (p *LangfuseProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Add data sources here if needed
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &LangfuseProvider{
			version: version,
		}
	}
} 