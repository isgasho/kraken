// Copyright (c) 2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package networkevent

import "sync"

// TestProducer records all produced events.
type TestProducer struct {
	sync.Mutex
	events []*Event
}

// NewTestProducer returns a new TestProducer.
func NewTestProducer() *TestProducer {
	return &TestProducer{}
}

// Produce records e.
func (p *TestProducer) Produce(e *Event) {
	p.Lock()
	defer p.Unlock()

	p.events = append(p.events, e)
}

// Close noops.
func (p *TestProducer) Close() error { return nil }

// Events returns all currently recorded events.
func (p *TestProducer) Events() []*Event {
	p.Lock()
	defer p.Unlock()

	res := make([]*Event, len(p.events))
	copy(res, p.events)
	return res
}
