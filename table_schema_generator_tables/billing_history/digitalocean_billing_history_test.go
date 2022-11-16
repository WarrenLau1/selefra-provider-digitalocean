package billing_history

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"

	"github.com/selefra/selefra-provider-digitalocean/faker"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createBillingHistory(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockBillingHistoryService(ctrl)

	data := &godo.BillingHistory{}

	if err := faker.FakeObject(data); err != nil {
		t.Fatal(err)
	}
	data.Links = nil

	m.EXPECT().List(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	return digitalocean_client.Services{

		BillingHistory: m,
	}

}

func TestBillingHistory(t *testing.T) {

	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanBillingHistoryGenerator{}), createBillingHistory, digitalocean_client.TestOptions{})

}
