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

package ports

import (
	"context"

	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/filters"
)

type SchedulerStorage interface {
	GetScheduler(ctx context.Context, name string) (*entities.Scheduler, error)
	GetSchedulerWithFilter(ctx context.Context, schedulerFilter *filters.SchedulerFilter) (*entities.Scheduler, error)
	GetSchedulerVersions(ctx context.Context, name string) ([]*entities.SchedulerVersion, error)
	GetSchedulers(ctx context.Context, names []string) ([]*entities.Scheduler, error)
	GetAllSchedulers(ctx context.Context) ([]*entities.Scheduler, error)
	CreateScheduler(ctx context.Context, scheduler *entities.Scheduler) error
	UpdateScheduler(ctx context.Context, scheduler *entities.Scheduler) error
	DeleteScheduler(ctx context.Context, scheduler *entities.Scheduler) error
	CreateSchedulerVersion(ctx context.Context, transactionID TransactionID, scheduler *entities.Scheduler) error
	RunWithTransaction(ctx context.Context, transactionFunc func(transactionId TransactionID) error) error
}

type TransactionID string
