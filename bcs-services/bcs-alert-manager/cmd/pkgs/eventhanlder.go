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

package pkgs

import (
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/config"
	"sync"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/handler/eventhandler"
)

const (
	EventHandleConcurrencyNum = 100
)

var (
	eventHandlerOnce sync.Once
	eventHandler     *eventhandler.SyncEventHandler
)

func GetEventSyncHandler(options config.AlertManagerOptions) *eventhandler.SyncEventHandler {
	eventHandlerOnce.Do(func() {
		eventHandler = eventhandler.NewSyncEventHandler(eventhandler.Options{
			ConcurrencyNum: EventHandleConcurrencyNum,
			Client:         GetAlertClient(options),
		})
		if eventHandler == nil {
			panic("init NewSyncEventHandler failed")
		}
	})

	return eventHandler
}
