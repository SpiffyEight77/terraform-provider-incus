package image

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	provider_config "github.com/lxc/terraform-provider-incus/internal/provider-config"
)

// ImageAlias resource data model that matches the schema.
type ImageAlias struct {
}

// ImageAliasResource represent Incus cached image alias resource.
type ImageAliasResource struct {
	provider *provider_config.IncusProviderConfig
}

// NewImageAliasResource return new cached image alias resource.
func NewImageAliasResource() resource.Resource {
	return &ImageAliasResource{}
}

func (r ImageAliasResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_image", req.ProviderTypeName)
}

func (r ImageAliasResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (r *ImageAliasResource) Configure(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
}

func (r ImageAliasResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

func (r ImageAliasResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r ImageAliasResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r ImageAliasResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
