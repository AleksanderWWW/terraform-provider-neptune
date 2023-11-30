// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"terraform-provider-neptune/neptune"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &NeptuneProvider{}

// ScaffoldingProvider defines the provider implementation.
type NeptuneProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ScaffoldingProviderModel describes the provider data model.
type NeptuneProviderModel struct {
	ApiToken types.String `tfsdk:"api_token"`
	Timeout  types.Int64  `tfsdk:"timeout"`
}

func (p *NeptuneProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "neptune"
	resp.Version = p.version
}

func (p *NeptuneProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				MarkdownDescription: "Your Neptune.ai api token.",
				Optional:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout for neptune client.",
				Optional:            true,
			},
		},
	}
}

func (p *NeptuneProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data NeptuneProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.ApiToken.IsNull() {
		data.ApiToken = types.StringValue("")
	}

	if data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(30)
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client, err := neptune.NewNeptuneClient(data.ApiToken.ValueString(), data.Timeout.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Neptune API Client",
			"An unexpected error occurred when creating the Neptune API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Neptune Client Error: "+err.Error(),
		)
		return
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *NeptuneProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewNeptuneProjectResource,
		NewNeptuneProjectMemberResource,
	}
}

func (p *NeptuneProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NeptuneProvider{
			version: version,
		}
	}
}
