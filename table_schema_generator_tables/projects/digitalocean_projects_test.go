package projects

import (
	"testing"

	"github.com/digitalocean/godo"

	"github.com/golang/mock/gomock"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client/mocks"

	"github.com/selefra/selefra-provider-digitalocean/faker"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
)

func createResources(t *testing.T, m *mocks.MockProjectsService) {

	var data []godo.ProjectResource

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)

	}

	m.EXPECT().ListResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}

func createProjects(t *testing.T, ctrl *gomock.Controller) digitalocean_client.Services {

	m := mocks.NewMockProjectsService(ctrl)

	var data []godo.Project

	if err := faker.FakeObject(&data); err != nil {

		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).AnyTimes().Return(data, &godo.Response{}, nil)

	createResources(t, m)

	return digitalocean_client.Services{

		Projects: m,
	}
}

func TestProjects(t *testing.T) {

	digitalocean_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableDigitaloceanProjectsGenerator{}), createProjects, digitalocean_client.TestOptions{})
}
