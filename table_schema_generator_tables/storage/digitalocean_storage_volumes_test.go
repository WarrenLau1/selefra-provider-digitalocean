package storage

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"
	"github.com/selefra/selefra-provider-digitalocean/faker"

	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createVolumes(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {
	m := mocks.NewMockStorageService(ctrl)

	var data []godo.Volume

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().ListVolumes(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	return digitalocean_client.Services{

		Storage: m,
	}

}

func TestVolumes(t *testing.T) {

	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanStorageVolumesGenerator{}), createVolumes, digitalocean_client.TestOptions{})

}
