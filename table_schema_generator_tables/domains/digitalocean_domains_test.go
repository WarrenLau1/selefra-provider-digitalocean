package domains

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"

	"github.com/selefra/selefra-provider-digitalocean/faker"

	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createRecords(t *testing.T, m *mocks.MockDomainsService) {

	var data []godo.DomainRecord
	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().Records(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}

func createDomains(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockDomainsService(ctrl)

	var data []godo.Domain

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	createRecords(t, m)

	return digitalocean_client.Services{

		Domains: m,
	}

}

func TestDomains(t *testing.T) {
	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanDomainsGenerator{}), createDomains, digitalocean_client.TestOptions{})
}
