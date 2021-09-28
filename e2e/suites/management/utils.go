package management

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/topfreegames/maestro/e2e/framework"
	"github.com/topfreegames/maestro/e2e/framework/maestro"
	maestroApiV1 "github.com/topfreegames/maestro/pkg/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createSchedulerAndWaitForIt(
	t *testing.T,
	maestro *maestro.MaestroInstance,
	apiClient *framework.APIClient,
	kubeclient kubernetes.Interface,
	gruCommand []string) (string, error) {
	schedulerName := framework.GenerateSchedulerName()
	createRequest := &maestroApiV1.CreateSchedulerRequest{
		Name:                   schedulerName,
		Game:                   "test",
		Version:                "v1.1",
		TerminationGracePeriod: 15,
		Containers: []*maestroApiV1.Container{
			{
				Name:            "example",
				Image:           "alpine",
				Command:         gruCommand,
				ImagePullPolicy: "Always",
				Environment: []*maestroApiV1.ContainerEnvironment{
					{
						Name:  "ROOMS_API_ADDRESS",
						Value: maestro.RoomsApiServer.ContainerInternalAddress,
					},
				},
				Requests: &maestroApiV1.ContainerResources{
					Memory: "20Mi",
					Cpu:    "10m",
				},
				Limits: &maestroApiV1.ContainerResources{
					Memory: "20Mi",
					Cpu:    "10m",
				},
				Ports: []*maestroApiV1.ContainerPort{
					{
						Name:     "default",
						Protocol: "tcp",
						Port:     80,
					},
				},
			},
		},
	}

	createResponse := &maestroApiV1.CreateSchedulerResponse{}
	err := apiClient.Do("POST", "/schedulers", createRequest, createResponse)
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		listOperationsRequest := &maestroApiV1.ListOperationsRequest{}
		listOperationsResponse := &maestroApiV1.ListOperationsResponse{}
		err = apiClient.Do("GET", fmt.Sprintf("/schedulers/%s/operations", schedulerName), listOperationsRequest, listOperationsResponse)
		require.NoError(t, err)

		if len(listOperationsResponse.FinishedOperations) == 0 {
			return false
		}

		require.Equal(t, "create_scheduler", listOperationsResponse.FinishedOperations[0].DefinitionName)
		require.Equal(t, "finished", listOperationsResponse.FinishedOperations[0].Status)
		return true
	}, 30*time.Second, time.Second)

	// Check on kubernetes that the scheduler namespace was created.
	_, err = kubeclient.CoreV1().Namespaces().Get(context.Background(), schedulerName, metav1.GetOptions{})
	require.NoError(t, err)

	// wait for service account to be created
	// TODO: check if we need to wait the service account to be created on internal/adapters/runtime/kubernetes/scheduler.go
	// we were having errors when not waiting for this in this test, reported in this issue https://github.com/kubernetes/kubernetes/issues/66689
	require.Eventually(t, func() bool {
		svcAccs, err := kubeclient.CoreV1().ServiceAccounts(schedulerName).List(context.Background(), metav1.ListOptions{})
		require.NoError(t, err)

		return len(svcAccs.Items) > 0
	}, 5*time.Second, time.Second)
	return schedulerName, err
}
