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

// Ideally, we'd be able to run the tests below by instantiating an
// MDServerRemote, but there's no convenient way to do that. So
// manually verify that MDServerRemote uses mdServerErrorUnwrapper.

// We should unwrap a kbfsmd.ServerErrorUnauthorized{} as a
// mdServerErrorUnauthorized{}.
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

// We should unwrap a kbfsmd.ServerErrorWriteAccess{} as a
// mdServerErrorWriteAccess{}.
func TestMDServerUnwrapErrorWriteAccess(t *testing.T) {
	var eu mdServerErrorUnwrapper
	status := keybase1.Status{
		Code: kbfsmd.StatusCodeServerErrorWriteAccess,
	}
	ae, de := eu.UnwrapError(&status)
	require.Equal(t, mdServerErrorWriteAccess{
		kbfsmd.ServerErrorWriteAccess{},
	}, ae)
	require.NoError(t, de)
}
