// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
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
			Factory:  DataSourceEngineVersion,
			TypeName: "aws_docdb_engine_version",
		},
		{
			Factory:  DataSourceOrderableDBInstance,
			TypeName: "aws_docdb_orderable_db_instance",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceCluster,
			TypeName: "aws_docdb_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceClusterInstance,
			TypeName: "aws_docdb_cluster_instance",
			Name:     "Cluster Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceClusterParameterGroup,
			TypeName: "aws_docdb_cluster_parameter_group",
			Name:     "Cluster Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceClusterSnapshot,
			TypeName: "aws_docdb_cluster_snapshot",
		},
		{
			Factory:  ResourceEventSubscription,
			TypeName: "aws_docdb_event_subscription",
			Name:     "Event Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceGlobalCluster,
			TypeName: "aws_docdb_global_cluster",
		},
		{
			Factory:  ResourceSubnetGroup,
			TypeName: "aws_docdb_subnet_group",
			Name:     "Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DocDB
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*docdb.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return docdb.NewFromConfig(cfg,
		docdb.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
