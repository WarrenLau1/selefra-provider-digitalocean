package monitoring

import (
	"testing"

	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"
	"github.com/selefra/selefra-provider-digitalocean/faker"

	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createAlertPolicies(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {
	m := mocks.NewMockMonitoringService(ctrl)

	var data []godo.AlertPolicy
	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)
	}

	m.EXPECT().ListAlertPolicies(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	return digitalocean_client.Services{

		Monitoring: m,
	}
}

func TestAlertPolicies(t *testing.T) {

	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanMonitoringAlertPoliciesGenerator{}), createAlertPolicies, digitalocean_client.TestOptions{})

}
