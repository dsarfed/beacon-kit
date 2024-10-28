// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package e2e_test

import (
	beaconapi "github.com/attestantio/go-eth2-client/api"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/berachain/beacon-kit/testing/e2e/config"
	"github.com/berachain/beacon-kit/testing/e2e/suite/types"
)

// TODO: Extract a helper function to initialize the consensus client so that
// it can be used for different API endpoints, instead of defining separately.
// initNodeTest initializes any tests for the node api.
func (s *BeaconKitE2ESuite) initNodeTest() *types.ConsensusClient {
	// Wait for execution block 5.
	err := s.WaitForFinalizedBlockNumber(5)
	s.Require().NoError(err)

	// Get the consensus client.
	client := s.ConsensusClients()[config.DefaultClient]
	s.Require().NotNil(client)

	return client
}

// TestNodeVersion tests the node api for beacon version of the node.
func (s *BeaconKitE2ESuite) TestNodeVersion() {
	client := s.initNodeTest()

	version, err := client.NodeVersion(s.Ctx(),
		&beaconapi.NodeVersionOpts{})
	s.Require().NoError(err)
	s.Require().NotEmpty(version)
	versionStr := version.Data
	s.Require().NotEmpty(versionStr)
}

// TestNodeSyncing tests the node api for syncing status of the node.
func (s *BeaconKitE2ESuite) TestNodeSyncing() {
	client := s.initNodeTest()

	syncing, err := client.NodeSyncing(s.Ctx(),
		&beaconapi.NodeSyncingOpts{})
	s.Require().NoError(err)
	s.Require().NotNil(syncing)
	syncData := syncing.Data
	s.Require().NotEmpty(syncData.HeadSlot)
	s.Require().Greater(syncData.HeadSlot, phase0.Slot(0))
	s.Require().NotNil(syncData.SyncDistance)

	s.Require().NotNil(syncData.IsSyncing)
	s.Require().True(syncData.IsOptimistic)

	// TODO: Add more assertions.
}
