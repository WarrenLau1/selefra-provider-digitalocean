package spaces

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"

	"github.com/selefra/selefra-provider-digitalocean/faker"

	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createSpaces(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockSpacesService(ctrl)

	var data *s3.ListBucketsOutput
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)

	}
	m.EXPECT().ListBuckets(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(data, nil)

	var cors *s3.GetBucketCorsOutput

	if err := faker.FakeObject(&cors); err != nil {

		t.Fatal(err)
	}
	m.EXPECT().GetBucketCors(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(cors, nil)

	var acl *s3.GetBucketAclOutput
	if err := faker.FakeObject(&acl); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetBucketAcl(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(acl, nil)

	return digitalocean_client.Services{

		Spaces: m,
	}
}

func TestSpaces(t *testing.T) {
	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanSpacesGenerator{}), createSpaces, digitalocean_client.TestOptions{})

}
