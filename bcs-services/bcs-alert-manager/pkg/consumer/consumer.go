/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package consumer

import (
	"context"
	glog "github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/msgqueue"
	"os"
	"os/signal"
	"syscall"
)

// Consumer xxx
type Consumer interface {
	Consume(ctx context.Context, queue msgqueue.MessageQueue) error
	Stop() error
}

// Consumers xxx
type Consumers struct {
	ctx       context.Context
	cancel    context.CancelFunc
	que       msgqueue.MessageQueue
	consumers []Consumer
}

func NewConsumers(consumers []Consumer, queue msgqueue.MessageQueue) *Consumers {
	c := &Consumers{
		que:       queue,
		consumers: consumers,
	}

	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c
}

func (c *Consumers) Run() {
	if c == nil {
		return
	}

	for idx := range c.consumers {
		go func(ctx context.Context, consumer Consumer) {
			defer func() {
				if r := recover(); r != nil {
					glog.Errorf("[monitor][panic] consumer panic: %v\n", r)
				}
			}()

			consumer.Consume(ctx, c.que)
		}(c.ctx, c.consumers[idx])
	}
}

func (c *Consumers) Wait() {
	signals := make(chan os.Signal, 10)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-signals:
		signal.Stop(signals)
		glog.Info("recv term signal")
		for idx := range c.consumers {
			c.consumers[idx].Stop()
		}
		c.cancel()
		c.que.Stop()
	}
}
