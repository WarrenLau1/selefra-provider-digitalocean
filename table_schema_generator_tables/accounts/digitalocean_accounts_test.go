package accounts

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"

	"github.com/selefra/selefra-provider-digitalocean/faker"

	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createAccount(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockAccountService(ctrl)

	var data godo.Account

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}
	m.EXPECT().Get(gomock.Any()).AnyTimes().Return(&data, nil, nil)

	return digitalocean_client.Services{

		Account: m,
	}

}

func TestAccount(t *testing.T) {

	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanAccountsGenerator{}), createAccount, digitalocean_client.TestOptions{})
}
