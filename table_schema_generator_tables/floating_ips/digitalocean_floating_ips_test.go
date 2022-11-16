package floating_ips

import (
	"testing"

	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"
	"github.com/selefra/selefra-provider-digitalocean/faker"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createFloatingIps(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockFloatingIpsService(ctrl)

	var data []godo.FloatingIP

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	return digitalocean_client.Services{
		FloatingIps: m,
	}

}

func TestFloatingIps(t *testing.T) {

	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanFloatingIpsGenerator{}), createFloatingIps, digitalocean_client.TestOptions{})
}
