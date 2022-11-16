package vpcs

import (
	"testing"

	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"

	"github.com/selefra/selefra-provider-digitalocean/faker"

	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createMembers(t *testing.T, m *mocks.MockVpcsService) {

	var data []*godo.VPCMember

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

}

func createVpcs(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockVpcsService(ctrl)

	var data []*godo.VPC

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	createMembers(t, m)

	return digitalocean_client.Services{

		Vpcs: m,
	}

}

func TestVpcs(t *testing.T) {
	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanVpcsGenerator{}), createVpcs, digitalocean_client.TestOptions{})

}
