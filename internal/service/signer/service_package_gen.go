// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package signer

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	signer_sdkv2 "github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceSigningJob,
			TypeName: "aws_signer_signing_job",
		},
		{
			Factory:  DataSourceSigningProfile,
			TypeName: "aws_signer_signing_profile",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceSigningJob,
			TypeName: "aws_signer_signing_job",
		},
		{
			Factory:  ResourceSigningProfile,
			TypeName: "aws_signer_signing_profile",
			Name:     "Signing Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceSigningProfilePermission,
			TypeName: "aws_signer_signing_profile_permission",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Signer
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*signer_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return signer_sdkv2.NewFromConfig(cfg, func(o *signer_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
