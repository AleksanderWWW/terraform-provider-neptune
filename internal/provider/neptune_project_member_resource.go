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
var _ resource.Resource = &NeptuneProjectMemberResource{}
var _ resource.ResourceWithImportState = &NeptuneProjectMemberResource{}

func NewNeptuneProjectMemberResource() resource.Resource {
	return &NeptuneProjectMemberResource{}
}

// ExampleResource defines the resource implementation.
type NeptuneProjectMemberResource struct {
	client *neptune.NeptuneClient
}

// ExampleResourceModel describes the resource data model.
type NeptuneProjectMemberResourceModel struct {
	Project   types.String `tfsdk:"project"`
	Workspace types.String `tfsdk:"workspace"`
	Username  types.String `tfsdk:"username"`
	Role      types.String `tfsdk:"role"`
}

func (r *NeptuneProjectMemberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_member"
}

func (r *NeptuneProjectMemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Neptune project member",

		Attributes: map[string]schema.Attribute{
			"project": schema.StringAttribute{
				MarkdownDescription: "Name of the project",
				Required:            true,
			},
			"workspace": schema.StringAttribute{
				MarkdownDescription: "Your Neptune workspace",
				Required:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Name of the user in Neptune",
				Required:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: "User's role in the project",
				Required:            true,
			},
		},
	}
}

func (r *NeptuneProjectMemberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*neptune.NeptuneClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *neptune.NeptuneClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NeptuneProjectMemberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NeptuneProjectMemberResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	err := r.client.AddProjectMember(
		data.Project.ValueString(),
		data.Workspace.ValueString(),
		data.Username.ValueString(),
		data.Role.ValueString(),
	)

	if err != nil {
		if strings.Contains(err.Error(), "User already has permissions to the project") {
			resp.Diagnostics.AddWarning("Project member already exists", err.Error())
		} else {
			resp.Diagnostics.AddError(
				"Error adding project member",
				"Could not add project member, unexpected error: "+err.Error(),
			)
			return
		}
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, fmt.Sprintf("User %s added to project %s/%s", data.Username.ValueString(), data.Workspace.ValueString(), data.Project.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NeptuneProjectMemberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NeptuneProjectMemberResourceModel

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

func (r *NeptuneProjectMemberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NeptuneProjectMemberResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.UpdateProjectMember(
		data.Project.ValueString(),
		data.Workspace.ValueString(),
		data.Username.ValueString(),
		data.Role.ValueString(),
	); err != nil {
		if strings.Contains(err.Error(), "not found") {
			resp.Diagnostics.AddWarning("User or project not found", err.Error())
		} else {
			resp.Diagnostics.AddError(
				"Error updating project member",
				"Could not update project member, unexpected error: "+err.Error(),
			)
			return
		}
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

func (r *NeptuneProjectMemberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NeptuneProjectMemberResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteProjectMember(data.Project.ValueString(), data.Workspace.ValueString(), data.Username.ValueString())

	if err != nil {
		if strings.Contains(err.Error(), "USER_NOT_IN_PROJECT") {
			resp.Diagnostics.AddWarning("Project member not found", err.Error())
		} else if strings.Contains(err.Error(), "404") {
			return
		} else {
			resp.Diagnostics.AddError(
				"Error removing project member",
				"Could not remove project member, unexpected error: "+err.Error(),
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

func (r *NeptuneProjectMemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
