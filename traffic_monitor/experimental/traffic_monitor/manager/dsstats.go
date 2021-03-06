package manager

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 * 
 *   http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */


import (
	ds "github.com/apache/incubator-trafficcontrol/traffic_monitor/experimental/traffic_monitor/deliveryservice"
	dsdata "github.com/apache/incubator-trafficcontrol/traffic_monitor/experimental/traffic_monitor/deliveryservicedata"
	"sync"
)

// DSStatsThreadsafe wraps a deliveryservice.Stats object to be safe for multiple reader goroutines and a single writer.
type DSStatsThreadsafe struct {
	dsStats *ds.Stats
	m       *sync.RWMutex
}

// DSStatsReader permits reading of a dsdata.Stats object, but not writing. This is designed so a Stats object can safely be passed to multiple goroutines, without worry one may unsafely write.
type DSStatsReader interface {
	Get() dsdata.StatsReadonly
}

// NewDSStatsThreadsafe returns a deliveryservice.Stats object wrapped to be safe for multiple readers and a single writer.
func NewDSStatsThreadsafe() DSStatsThreadsafe {
	s := ds.NewStats()
	return DSStatsThreadsafe{m: &sync.RWMutex{}, dsStats: &s}
}

// Get returns a Stats object safe for reading by multiple goroutines
func (o *DSStatsThreadsafe) Get() dsdata.StatsReadonly {
	o.m.RLock()
	defer o.m.RUnlock()
	return *o.dsStats
}

// Set sets the internal Stats object. This MUST NOT be called by multiple goroutines.
func (o *DSStatsThreadsafe) Set(newDsStats ds.Stats) {
	o.m.Lock()
	*o.dsStats = newDsStats
	o.m.Unlock()
}
