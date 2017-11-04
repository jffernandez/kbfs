// Copyright 2017 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package libkbfs

import (
	"testing"

	"github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/kbfs/kbfsmd"
	"github.com/stretchr/testify/require"
)

// We should unwrap a kbfsmd.ServerErrorUnauthorized{} as a
// mdServerErrorUnauthorized{}.
//
// Ideally, we'd be able to test this by instantiating a
// BServerRemote, but there's no convenient way to maintain that. So
// manually verify that MDServerRemote uses mdServerErrorUnwrapper.
func TestMDServerUnwrapErrorUnauthorized(t *testing.T) {
	var eu mdServerErrorUnwrapper
	status := keybase1.Status{
		Code: kbfsmd.StatusCodeServerErrorUnauthorized,
	}
	ae, de := eu.UnwrapError(&status)
	require.Equal(t, mdServerErrorUnauthorized{
		kbfsmd.ServerErrorUnauthorized{},
	}, ae)
	require.NoError(t, de)
}
