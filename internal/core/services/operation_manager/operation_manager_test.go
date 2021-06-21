package operation_manager

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	opstorage "github.com/topfreegames/maestro/internal/adapters/operation_storage/mock"
	"github.com/topfreegames/maestro/internal/core/entities/operation"
	porterrors "github.com/topfreegames/maestro/internal/core/ports/errors"
	"github.com/topfreegames/maestro/internal/core/services/operations_registry"
)

type testOperationDefinition struct {
	marshalResult   []byte
	unmarshalResult error
}

func (d *testOperationDefinition) Marshal() []byte            { return d.marshalResult }
func (d *testOperationDefinition) Unmarshal(raw []byte) error { return d.unmarshalResult }
func (d *testOperationDefinition) Name() string               { return "testOperationDefinition" }

type opMatcher struct {
	status operation.Status
	def operation.Definition
}

func (m *opMatcher) Matches(x interface{}) bool {
	op, _ := x.(*operation.Operation)
	_, err := uuid.Parse(op.ID)
	return err == nil && op.Status == m.status && m.def.Name() == op.DefinitionName
}

func (m *opMatcher) String() string {
	return fmt.Sprintf("a operation with definition \"%s\"", m.def.Name())
}

func TestCreateOperation(t *testing.T) {
	cases := map[string]struct {
		definition operation.Definition
		storageErr error
	}{
		"create without errors": {
			definition: &testOperationDefinition{marshalResult: []byte("test")},
		},
		"create with storage errors": {
			definition: &testOperationDefinition{},
			storageErr: porterrors.ErrUnexpected,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
			opManager := New(operationStorage)

			ctx := context.Background()
			testDefinition, _ := test.definition.(*testOperationDefinition)
			operationStorage.EXPECT().CreateOperation(ctx, &opMatcher{operation.StatusPending, test.definition}, testDefinition.marshalResult).Return(test.storageErr)

			op, err := opManager.CreateOperation(ctx, test.definition)
			if test.storageErr != nil {
				require.ErrorIs(t, err, test.storageErr)
				require.Nil(t, op)
				return
			}

			require.NotNil(t, op)
			require.Equal(t, operation.StatusPending, op.Status)
		})
	}
}

func TestGetOperation(t *testing.T) {
	t.Run("find operation", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operation.Definition { return &testOperationDefinition{} }
		registry := operations_registry.NewRegistry()
		registry.Register(defFunc().Name(), defFunc)

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := NewWithRegistry(operationStorage, registry)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			nil,
		)

		op, definition, err := opManager.GetOperation(ctx, schedulerName, operationID)
		require.NoError(t, err)
		require.NotNil(t, op)
		require.Equal(t, operationID, op.ID)
		require.Equal(t, schedulerName, op.SchedulerName)
		require.IsType(t, defFunc(), definition)
	})

	t.Run("defition not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operation.Definition { return &testOperationDefinition{} }
		registry := operations_registry.NewRegistry()

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := NewWithRegistry(operationStorage, registry)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			nil,
		)

		_, _, err := opManager.GetOperation(ctx, schedulerName, operationID)
		require.Error(t, err)
	})

	t.Run("operation not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operation.Definition { return &testOperationDefinition{} }
		registry := operations_registry.NewRegistry()

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := NewWithRegistry(operationStorage, registry)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			porterrors.ErrNotFound,
		)

		_, _, err := opManager.GetOperation(ctx, schedulerName, operationID)
		require.Error(t, err)
	})

	t.Run("unmarshal error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operation.Definition { return &testOperationDefinition{unmarshalResult: errors.New("invalid")} }
		registry := operations_registry.NewRegistry()

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := NewWithRegistry(operationStorage, registry)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			porterrors.ErrNotFound,
		)

		_, _, err := opManager.GetOperation(ctx, schedulerName, operationID)
		require.Error(t, err)
	})
}

func TestNextSchedulerOperation(t *testing.T) {
	t.Run("fetch operation", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operation.Definition { return &testOperationDefinition{} }
		registry := operations_registry.NewRegistry()
		registry.Register(defFunc().Name(), defFunc)

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := NewWithRegistry(operationStorage, registry)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"

		operationStorage.EXPECT().NextSchedulerOperationID(ctx, schedulerName).Return(operationID, nil)
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			&operation.Operation{ID: operationID, SchedulerName: schedulerName, DefinitionName: defFunc().Name()},
			[]byte{},
			nil,
		)

		op, definition, err := opManager.NextSchedulerOperation(ctx, schedulerName)
		require.NoError(t, err)
		require.NotNil(t, op)
		require.Equal(t, operationID, op.ID)
		require.Equal(t, schedulerName, op.SchedulerName)
		require.IsType(t, defFunc(), definition)
	})

	t.Run("no next operation", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operation.Definition { return &testOperationDefinition{} }
		registry := operations_registry.NewRegistry()
		registry.Register(defFunc().Name(), defFunc)

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := NewWithRegistry(operationStorage, registry)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationStorage.EXPECT().NextSchedulerOperationID(ctx, schedulerName).Return("", porterrors.ErrUnexpected)

		_, _, err := opManager.NextSchedulerOperation(ctx, schedulerName)
		require.Error(t, err)
	})

	t.Run("operation not found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		defFunc := func() operation.Definition { return &testOperationDefinition{} }
		registry := operations_registry.NewRegistry()
		registry.Register(defFunc().Name(), defFunc)

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := NewWithRegistry(operationStorage, registry)

		ctx := context.Background()
		schedulerName := "test-scheduler"
		operationID := "some-op-id"

		operationStorage.EXPECT().NextSchedulerOperationID(ctx, schedulerName).Return(operationID, nil)
		operationStorage.EXPECT().GetOperation(ctx, schedulerName, operationID).Return(
			nil,
			[]byte{},
			porterrors.ErrNotFound,
		)

		_, _, err := opManager.NextSchedulerOperation(ctx, schedulerName)
		require.Error(t, err)
	})
}

func TestStartOperation(t *testing.T) {
	t.Run("sets active", func (t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		operationStorage := opstorage.NewMockOperationStorage(mockCtrl)
		opManager := NewWithRegistry(operationStorage, operations_registry.NewRegistry())

		ctx := context.Background()
		op := &operation.Operation{ID: uuid.NewString(), DefinitionName: (&testOperationDefinition{}).Name()}

		operationStorage.EXPECT().SetOperationActive(ctx, &opMatcher{operation.StatusInProgress, &testOperationDefinition{}}).Return(nil)
		err := opManager.StartOperation(ctx, op)
		require.NoError(t, err)
	})
}
