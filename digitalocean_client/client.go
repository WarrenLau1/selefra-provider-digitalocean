package digitalocean_client

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	retry "github.com/avast/retry-go"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/digitalocean/godo"
	"github.com/pkg/errors"
)

var defaultSpacesRegions = []string{constants.Nyc, constants.Sfo, constants.Ams, constants.Sgp, constants.Fra}

var MaxItemsPerPage = 200

type Client struct {
	DoClient         *godo.Client
	Regions          []string
	SpacesRegion     string
	CredentialStatus DoCredentialStruct
	Services         *Services
}

type DoCredentialStruct struct {
	Api    bool
	Spaces bool
}

type SpacesEndpointResolver struct{}

func (SpacesEndpointResolver) ResolveEndpoint(_, region string) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:    fmt.Sprintf(constants.Httpssdigitaloceanspacescom, region),
		Source: aws.EndpointSourceCustom,
	}, nil
}

type SpacesCredentialsProvider struct {
	SpacesAccessKey   string
	SpacesAccessKeyId string
}

func (s SpacesCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     s.SpacesAccessKeyId,
		SecretAccessKey: s.SpacesAccessKey,
		Source:          constants.Digitalocean,
	}, nil
}

func (c *Client) WithSpacesRegion(region string) *Client {
	return &Client{
		DoClient:     c.DoClient,
		SpacesRegion: region,
		Services:     initServices(c.DoClient, c.Services.Spaces),
	}
}

func (c *Client) ID() string {
	return c.SpacesRegion
}

type ServicesRegionMap map[string]*Services

type ServicesManager struct {
	services ServicesRegionMap
}

func (s *ServicesManager) ServicesByRegion(region string) *Services {
	return s.services[region]
}

func initServices(doClient *godo.Client, spacesService SpacesService) *Services {
	return &Services{
		Account:        doClient.Account,
		Cdn:            doClient.CDNs,
		BillingHistory: doClient.BillingHistory,
		Monitoring:     doClient.Monitoring,
		Balance:        doClient.Balance,
		Certificates:   doClient.Certificates,
		Databases:      doClient.Databases,
		Domains:        doClient.Domains,
		Droplets:       doClient.Droplets,
		Firewalls:      doClient.Firewalls,
		FloatingIps:    doClient.FloatingIPs,
		Images:         doClient.Images,
		Keys:           doClient.Keys,
		LoadBalancers:  doClient.LoadBalancers,
		Projects:       doClient.Projects,
		Registry:       doClient.Registry,
		Sizes:          doClient.Sizes,
		Snapshots:      doClient.Snapshots,
		Storage:        doClient.Storage,
		Vpcs:           doClient.VPCs,
		Spaces:         spacesService,
	}
}

func getTokenFromEnv() string {
	doToken := os.Getenv(constants.DIGITALOCEANTOKEN)
	doAccessToken := os.Getenv(constants.DIGITALOCEANACCESSTOKEN)
	if doToken != constants.Constants_0 {
		return doToken
	}
	if doAccessToken != constants.Constants_1 {
		return doAccessToken
	}
	return constants.Constants_2
}

func getSpacesTokenFromEnv() (string, string) {
	spacesAccessKey := os.Getenv(constants.SPACESACCESSKEYID)
	spacesSecretKey := os.Getenv(constants.SPACESSECRETACCESSKEY)
	if spacesAccessKey == constants.Constants_3 {
		return constants.Constants_4, constants.Constants_5
	}
	if spacesSecretKey == constants.Constants_6 {
		return constants.Constants_7, constants.Constants_8
	}
	return spacesAccessKey, spacesSecretKey
}

func NewClients(config Config) ([]*Client, error) {
	client, err := newClient(config)
	if err != nil {
		return nil, err
	}
	return []*Client{client}, nil
}

func newClient(providerConfig Config) (*Client, error) {
	if providerConfig.Token == constants.Constants_9 {
		providerConfig.Token = getTokenFromEnv()
	}
	if providerConfig.Token == constants.Constants_10 {
		return nil, errors.New(constants.Tokenfailedtoget)
	}
	credStatus := DoCredentialStruct{
		Api:    true,
		Spaces: true,
	}
	if providerConfig.SpacesAccessKey == constants.Constants_11 || providerConfig.SpacesAccessKeyId == constants.Constants_12 {
		providerConfig.SpacesAccessKey, providerConfig.SpacesAccessKeyId = getSpacesTokenFromEnv()
	}
	if providerConfig.SpacesAccessKey == constants.Constants_13 || providerConfig.SpacesAccessKeyId == constants.Constants_14 {
		credStatus.Spaces = false
	}
	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(SpacesCredentialsProvider{providerConfig.SpacesAccessKey, providerConfig.SpacesAccessKeyId}),
		config.WithEndpointResolver(SpacesEndpointResolver{}),
	)
	if err != nil {
		return nil, err
	}
	spacesRegions := defaultSpacesRegions
	if len(providerConfig.SpacesRegions) > 0 {
		spacesRegions = providerConfig.SpacesRegions
	}
	client := godo.NewFromToken(providerConfig.Token)
	c := Client{
		DoClient:         godo.NewFromToken(providerConfig.Token),
		Regions:          spacesRegions,
		SpacesRegion:     constants.Nyc,
		CredentialStatus: credStatus,
		Services:         initServices(client, s3.NewFromConfig(awsCfg)),
	}
	return &c, nil
}

func IsLimitReached(err error) bool {
	unwrapped := errors.Unwrap(err)
	er, ok := unwrapped.(*godo.ErrorResponse)
	if !ok {
		return false
	}
	return er.Message == constants.Toomanyrequests
}

func ThrottleWrapper(ctx context.Context, client *Client, doFunc retry.RetryableFunc) error {
	err := retry.Do(
		doFunc,
		retry.OnRetry(func(n uint, err error) {
			rate := client.DoClient.GetRate()
			fmt.Println(constants.CurrentAPIratelimits, constants.Limit, rate.Limit, constants.Remaining, rate.Remaining, constants.Reset, rate.Reset.Time)
		}),
		retry.RetryIf(IsLimitReached),
		retry.Attempts(5),
		retry.Context(ctx),
		retry.Delay(time.Second+5),
	)
	return err
}
