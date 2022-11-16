package registries

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"

	"github.com/selefra/selefra-provider-digitalocean/faker"

	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createRepositories(t *testing.T, m *mocks.MockRegistryService) {

	var data []*godo.Repository

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().ListRepositories(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

}

func createRegistry(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockRegistryService(ctrl)

	var data godo.Registry

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().Get(gomock.Any()).AnyTimes().Return(&data, nil, nil)

	createRepositories(t, m)

	return digitalocean_client.Services{

		Registry: m,
	}
}

func TestRegistry(t *testing.T) {
	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanRegistriesGenerator{}), createRegistry, digitalocean_client.TestOptions{})
}
