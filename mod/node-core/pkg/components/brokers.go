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

package components

import (
	"github.com/berachain/beacon-kit/mod/async/pkg/broker"
	asynctypes "github.com/berachain/beacon-kit/mod/async/pkg/types"
)

// ProvideBlobBroker provides a blob feed for the depinject framework.
func ProvideBlobBroker[
	SidecarT any,
]() *broker.Broker[*asynctypes.Event[SidecarT]] {
	return broker.New[*asynctypes.Event[SidecarT]](
		"blob-broker",
	)
}

// ProvideBlockBroker provides a block feed for the depinject framework.
func ProvideBlockBroker[
	BeaconBlockT any,
]() *broker.Broker[*asynctypes.Event[BeaconBlockT]] {
	return broker.New[*asynctypes.Event[BeaconBlockT]](
		"blk-broker",
	)
}

// ProvideGenesisBroker provides a genesis feed for the depinject framework.
func ProvideGenesisBroker[
	GenesisT any,
]() *broker.Broker[*asynctypes.Event[GenesisT]] {
	return broker.New[*asynctypes.Event[GenesisT]](
		"genesis-broker",
	)
}

// ProvideSlotBroker provides a slot feed for the depinject framework.
func ProvideSlotBroker[
	SlotDataT any,
]() *broker.Broker[*asynctypes.Event[SlotDataT]] {
	return broker.New[*asynctypes.Event[SlotDataT]](
		"slot-broker",
	)
}

// ProvideStatusBroker provides a status feed.
func ProvideStatusBroker[
	StatusEventT any,
]() *broker.Broker[*asynctypes.Event[StatusEventT]] {
	return broker.New[*asynctypes.Event[StatusEventT]](
		"status-broker",
	)
}

// ProvideValidatorUpdateBroker provides a validator updates feed.
func ProvideValidatorUpdateBroker[
	ValidatorUpdatesT any,
]() *broker.Broker[*asynctypes.Event[ValidatorUpdatesT]] {
	return broker.New[*asynctypes.Event[ValidatorUpdatesT]](
		"validator-updates-broker",
	)
}

// DefaultBrokerProviders returns a slice of the default broker providers.
func DefaultBrokerProviders[
	SidecarT any,
	BeaconBlockT any,
	GenesisT any,
	SlotDataT any,
	StatusEventT any,
	ValidatorUpdatesT any,
]() []interface{} {
	return []interface{}{
		ProvideBlobBroker[SidecarT],
		ProvideBlockBroker[BeaconBlockT],
		ProvideGenesisBroker[GenesisT],
		ProvideSlotBroker[SlotDataT],
		ProvideStatusBroker[StatusEventT],
		ProvideValidatorUpdateBroker[ValidatorUpdatesT],
	}
}