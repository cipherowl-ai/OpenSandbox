// Copyright 2026 Alibaba Group Holding Ltd.
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

package constants

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAPIProxyInternalSuffixes_DefaultOnly(t *testing.T) {
	t.Setenv(EnvEgressAPIProxyInternalSuffixes, "")
	require.Equal(t, []string{DefaultAPIProxyInternalSuffix}, APIProxyInternalSuffixes())
}

func TestAPIProxyInternalSuffixes_TrimsLowersAndAdds(t *testing.T) {
	t.Setenv(EnvEgressAPIProxyInternalSuffixes, " .AWS.Cipherowl.net , .Internal.Example.com ")
	require.Equal(t,
		[]string{DefaultAPIProxyInternalSuffix, ".aws.cipherowl.net", ".internal.example.com"},
		APIProxyInternalSuffixes(),
	)
}

func TestAPIProxyInternalSuffixes_DropsInvalidEntries(t *testing.T) {
	// Missing leading dot, blank, and a duplicate of the default are all dropped
	// so a stray value cannot widen the internal match to arbitrary hosts.
	t.Setenv(EnvEgressAPIProxyInternalSuffixes, "aws.cipherowl.net,, .svc.cluster.local , .ok.net")
	require.Equal(t, []string{DefaultAPIProxyInternalSuffix, ".ok.net"}, APIProxyInternalSuffixes())
}
