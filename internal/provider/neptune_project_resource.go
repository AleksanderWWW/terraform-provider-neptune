package provider

import (
	"context"
	"fmt"
	"strings"
	"terraform-provider-neptune/internal/neptune"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NeptuneProjectResource{}
var _ resource.ResourceWithImportState = &NeptuneProjectResource{}

func NewNeptuneProjectResource() resource.Resource {
	return &NeptuneProjectResource{}
}

// ExampleResource defines the resource implementation.
type NeptuneProjectResource struct {
	apiToken *string
}

// ExampleResourceModel describes the resource data model.
type NeptuneProjectResourceModel struct {
	Name      types.String `tfsdk:"name"`
	Workspace types.String `tfsdk:"workspace"`
	Key       types.String `tfsdk:"key"`
	Vis       types.String `tfsdk:"vis"`
}

func (r *NeptuneProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (r *NeptuneProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Neptune project",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the project you want to create",
				Required:            true,
			},
			"workspace": schema.StringAttribute{
				MarkdownDescription: "Your Neptune workspace",
				Required:            true,
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "Project 3-letter key (default: uppercased first 3 letters of the project name)",
				Optional:            true,
			},
			"vis": schema.StringAttribute{
				MarkdownDescription: "Project visibility ('private' or 'public', default: 'private')",
				Optional:            true,
			},
		},
	}
}

func (r *NeptuneProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	apiToken, ok := req.ProviderData.(*string)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *string, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.apiToken = apiToken
}

func (r *NeptuneProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NeptuneProjectResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	err := neptune.CreateProject(
		data.Name.ValueString(),
		data.Workspace.ValueString(),
		data.Vis.ValueString(),
		r.apiToken,
	)

	if err != nil {
		if strings.Contains(err.Error(), "createProject (status 409)") {
			resp.Diagnostics.AddWarning("Project already exists", err.Error())
		} else {
			resp.Diagnostics.AddError(
				"Error creating project",
				"Could not create project, unexpected error: "+err.Error(),
			)
			return
		}
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, fmt.Sprintf("Project %s/%s created", data.Workspace.ValueString(), data.Name.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NeptuneProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NeptuneProjectResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NeptuneProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NeptuneProjectResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NeptuneProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NeptuneProjectResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := neptune.DeleteProject(data.Name.ValueString(), data.Workspace.ValueString(), r.apiToken)

	if err != nil {
		if strings.Contains(err.Error(), "deleteProjectNotFound") {
			resp.Diagnostics.AddWarning("Project does not exist", err.Error())
		} else {
			resp.Diagnostics.AddError(
				"Error deleting project",
				"Could not delete project, unexpected error: "+err.Error(),
			)
			return
		}
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (r *NeptuneProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
