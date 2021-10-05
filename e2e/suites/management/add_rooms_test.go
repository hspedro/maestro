// MIT License
//
// Copyright (c) 2021 TFG Co
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package management

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	instanceStorageRedis "github.com/topfreegames/maestro/internal/adapters/instance_storage/redis"
	roomStorageRedis "github.com/topfreegames/maestro/internal/adapters/room_storage/redis"
	"github.com/topfreegames/maestro/internal/core/entities/game_room"

	"github.com/stretchr/testify/require"

	"github.com/topfreegames/maestro/e2e/framework/maestro"

	"github.com/topfreegames/maestro/e2e/framework"
	maestroApiV1 "github.com/topfreegames/maestro/pkg/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func TestAddRooms(t *testing.T) {
	framework.WithClients(t, func(apiClient *framework.APIClient, kubeclient kubernetes.Interface, redisClient *redis.Client, maestro *maestro.MaestroInstance) {
		roomsStorage := roomStorageRedis.NewRedisStateStorage(redisClient)
		instanceStorage := instanceStorageRedis.NewRedisInstanceStorage(redisClient, 10)

		t.Run("when created rooms does not reply its state back then it finishes the operation successfully", func(t *testing.T) {
			t.Parallel()

			schedulerName, err := createSchedulerAndWaitForIt(
				t,
				maestro,
				apiClient,
				kubeclient,
				[]string{"sh", "-c", "tail -f /dev/null"},
			)

			addRoomsRequest := &maestroApiV1.AddRoomsRequest{SchedulerName: schedulerName, Amount: 1}
			addRoomsResponse := &maestroApiV1.AddRoomsResponse{}
			err = apiClient.Do("POST", fmt.Sprintf("/schedulers/%s/add-rooms", schedulerName), addRoomsRequest, addRoomsResponse)

			require.Eventually(t, func() bool {
				listOperationsRequest := &maestroApiV1.ListOperationsRequest{}
				listOperationsResponse := &maestroApiV1.ListOperationsResponse{}
				err = apiClient.Do("GET", fmt.Sprintf("/schedulers/%s/operations", schedulerName), listOperationsRequest, listOperationsResponse)
				require.NoError(t, err)

				if len(listOperationsResponse.FinishedOperations) < 2 {
					return false
				}

				require.Equal(t, "add_rooms", listOperationsResponse.FinishedOperations[1].DefinitionName)
				return true
			}, 240*time.Second, time.Second)

			pods, err := kubeclient.CoreV1().Pods(schedulerName).List(context.Background(), metav1.ListOptions{})
			require.NoError(t, err)
			require.NotEmpty(t, pods.Items)

			require.Eventually(t, func() bool {
				instance, err := instanceStorage.GetInstance(context.Background(), schedulerName, pods.Items[0].ObjectMeta.Name)

				return err == nil && instance.Status.Type == game_room.InstanceReady
			}, time.Minute, time.Second)
		})

		t.Run("when created rooms replies its state back then it finishes the operation successfully", func(t *testing.T) {
			t.Parallel()

			schedulerName, err := createSchedulerAndWaitForIt(t,
				maestro,
				apiClient,
				kubeclient,
				[]string{"/bin/sh", "-c", "apk add curl && curl --request POST " +
					"$ROOMS_API_ADDRESS:9097/scheduler/$MAESTRO_SCHEDULER_NAME/rooms/$MAESTRO_ROOM_ID/ping " +
					"--data-raw '{\"status\": \"ready\",\"timestamp\": \"12312312313\"}'"})

			addRoomsRequest := &maestroApiV1.AddRoomsRequest{SchedulerName: schedulerName, Amount: 1}
			addRoomsResponse := &maestroApiV1.AddRoomsResponse{}
			err = apiClient.Do("POST", fmt.Sprintf("/schedulers/%s/add-rooms", schedulerName), addRoomsRequest, addRoomsResponse)

			require.Eventually(t, func() bool {
				listOperationsRequest := &maestroApiV1.ListOperationsRequest{}
				listOperationsResponse := &maestroApiV1.ListOperationsResponse{}
				err = apiClient.Do("GET", fmt.Sprintf("/schedulers/%s/operations", schedulerName), listOperationsRequest, listOperationsResponse)
				require.NoError(t, err)

				if len(listOperationsResponse.FinishedOperations) < 2 {
					return false
				}

				require.Equal(t, "add_rooms", listOperationsResponse.FinishedOperations[1].DefinitionName)
				return true
			}, 240*time.Second, time.Second)

			pods, err := kubeclient.CoreV1().Pods(schedulerName).List(context.Background(), metav1.ListOptions{})
			require.NoError(t, err)
			require.NotEmpty(t, pods.Items)

			require.Eventually(t, func() bool {
				room, err := roomsStorage.GetRoom(context.Background(), schedulerName, pods.Items[0].ObjectMeta.Name)

				return err == nil && room.Status == game_room.GameStatusReady
			}, time.Minute, time.Second)
		})

		t.Run("when trying to add rooms to a nonexistent scheduler then the operation fails", func(t *testing.T) {
			t.Parallel()

			schedulerName := "NonExistent"

			addRoomsRequest := &maestroApiV1.AddRoomsRequest{SchedulerName: schedulerName, Amount: 1}
			addRoomsResponse := &maestroApiV1.AddRoomsResponse{}
			err := apiClient.Do("POST", fmt.Sprintf("/schedulers/%s/add-rooms", schedulerName), addRoomsRequest, addRoomsResponse)
			require.Error(t, err, "failed with status 404, response body: {\"code\":5, \"message\":\"no "+
				"scheduler found to add rooms on it: scheduler NonExistent not found\", \"details\":[]}")
		})
	})

}
