package databases

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"
	"github.com/selefra/selefra-provider-digitalocean/faker"

	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createFirewallRules(t *testing.T, m *mocks.MockDatabasesService) {

	var data []godo.DatabaseFirewallRule
	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().GetFirewallRules(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

}

func createReplicas(t *testing.T, m *mocks.MockDatabasesService) {
	var data []godo.DatabaseReplica

	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListReplicas(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

}
func createBackups(t *testing.T, m *mocks.MockDatabasesService) {
	var data []godo.DatabaseBackup

	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)

	}

	m.EXPECT().ListBackups(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}

func createDatabases(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockDatabasesService(ctrl)

	var data []godo.Database

	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	createFirewallRules(t, m)
	createReplicas(t, m)

	createBackups(t, m)

	return digitalocean_client.Services{
		Databases: m,
	}
}

func TestDatabases(t *testing.T) {

	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanDatabasesGenerator{}), createDatabases, digitalocean_client.TestOptions{})

}
