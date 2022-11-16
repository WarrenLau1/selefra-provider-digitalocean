package droplets

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"
	"github.com/selefra/selefra-provider-digitalocean/faker"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createNeighbors(t *testing.T, m *mocks.MockDropletsService) {

	var data []godo.Droplet

	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().Neighbors(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

}

func createDroplets(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {
	m := mocks.NewMockDropletsService(ctrl)

	var data []godo.Droplet

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	createNeighbors(t, m)

	return digitalocean_client.Services{

		Droplets: m,
	}

}

func TestDroplets(t *testing.T) {
	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanDropletsGenerator{}), createDroplets, digitalocean_client.TestOptions{})
}
