// Copyright 2015 CoreOS, Inc.
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

package backend

import (
	"github.com/coreos/flannel/Godeps/_workspace/src/golang.org/x/net/context"

	"github.com/coreos/flannel/subnet"
)

type SubnetDef struct {
	Lease *subnet.Lease
	MTU   int
}

// Besides the entry points in the Backend interface, the backend's New()
// function receives static network interface information (like internal and
// external IP addresses, MTU, etc) which it should cache for later use if
// needed.
//
// To implement a singleton backend which manages multiple networks, the
// New() function should create the singleton backend object once, and return
// that object on on further calls to New().  The backend is guaranteed that
// the arguments passed via New() will not change across invocations.  Also,
// since multiple RegisterNetwork() and Run() calls may be in-flight at any
// given time for a singleton backend, it must protect these calls with a mutex.
type Backend interface {
	// Called when the backend should create or begin managing a new network
	RegisterNetwork(ctx context.Context, network string, config *subnet.Config) (*SubnetDef, error)
	// Called after the backend's first network has been registered to
	// allow the plugin to watch dynamic events
	Run(ctx context.Context)
	// Called to clean up any network resources or operations
	UnregisterNetwork(ctx context.Context, network string)
}
