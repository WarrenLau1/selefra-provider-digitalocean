package digitalocean_client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/digitalocean/godo"
)

type Services struct {
	Account        AccountService
	Cdn            CdnService
	BillingHistory BillingHistoryService
	Monitoring     MonitoringService
	Balance        BalanceService
	Certificates   CertificatesService
	Databases      DatabasesService
	Domains        DomainsService
	Droplets       DropletsService
	Firewalls      FirewallsService
	FloatingIps    FloatingIpsService
	Images         ImagesService
	Keys           KeysService
	LoadBalancers  LoadBalancersService
	Projects       ProjectsService
	Registry       RegistryService
	Sizes          SizesService
	Snapshots      SnapshotsService
	Storage        StorageService
	Vpcs           VpcsService
	Spaces         SpacesService
}

type AccountService interface {
	Get(context.Context) (*godo.Account, *godo.Response, error)
}

type CdnService interface {
	List(context.Context, *godo.ListOptions) ([]godo.CDN, *godo.Response, error)
}

type BillingHistoryService interface {
	List(context.Context, *godo.ListOptions) (*godo.BillingHistory, *godo.Response, error)
}

type MonitoringService interface {
	ListAlertPolicies(context.Context, *godo.ListOptions) ([]godo.AlertPolicy, *godo.Response, error)
}

type BalanceService interface {
	Get(context.Context) (*godo.Balance, *godo.Response, error)
}

type CertificatesService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Certificate, *godo.Response, error)
}

type DatabasesService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Database, *godo.Response, error)
	ListBackups(context.Context, string, *godo.ListOptions) ([]godo.DatabaseBackup, *godo.Response, error)
	ListReplicas(context.Context, string, *godo.ListOptions) ([]godo.DatabaseReplica, *godo.Response, error)
	GetFirewallRules(context.Context, string) ([]godo.DatabaseFirewallRule, *godo.Response, error)
}

type DomainsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Domain, *godo.Response, error)
	Records(context.Context, string, *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error)
}

type DropletsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Droplet, *godo.Response, error)
	Neighbors(context.Context, int) ([]godo.Droplet, *godo.Response, error)
}

type FirewallsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Firewall, *godo.Response, error)
}

type FloatingIpsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.FloatingIP, *godo.Response, error)
}

type ImagesService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Image, *godo.Response, error)
}

type KeysService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Key, *godo.Response, error)
}

type LoadBalancersService interface {
	List(context.Context, *godo.ListOptions) ([]godo.LoadBalancer, *godo.Response, error)
}

type ProjectsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Project, *godo.Response, error)
	ListResources(context.Context, string, *godo.ListOptions) ([]godo.ProjectResource, *godo.Response, error)
}

type RegionsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Region, *godo.Response, error)
}

type RegistryService interface {
	Get(context.Context) (*godo.Registry, *godo.Response, error)
	ListRepositories(context.Context, string, *godo.ListOptions) ([]*godo.Repository, *godo.Response, error)
}

type SizesService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Size, *godo.Response, error)
}

type SnapshotsService interface {
	List(context.Context, *godo.ListOptions) ([]godo.Snapshot, *godo.Response, error)
}

type StorageService interface {
	ListVolumes(context.Context, *godo.ListVolumeParams) ([]godo.Volume, *godo.Response, error)
}

type VpcsService interface {
	List(context.Context, *godo.ListOptions) ([]*godo.VPC, *godo.Response, error)
	ListMembers(context.Context, string, *godo.VPCListMembersRequest, *godo.ListOptions) ([]*godo.VPCMember, *godo.Response, error)
}

type SpacesService interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	GetBucketCors(ctx context.Context, params *s3.GetBucketCorsInput, optFns ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error)
	GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
}
