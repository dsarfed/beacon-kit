// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package deposit

import (
	"context"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
)

type StorageBackend[
	AvailabilityStoreT any,
	BeaconStateT any,
	BlobSidecarsT any,
	DepositStoreT Store,
] interface {
	// DepositStore returns the deposit store for the given context.
	DepositStore(context.Context) DepositStoreT
}

// Contract is the ABI for the deposit contract.
type Contract interface {
	GetDeposits(
		ctx context.Context,
		blockNumber uint64,
	) ([]*types.Deposit, error)
}

// Store defines the interface for managing deposit operations.
type Store interface {
	// PruneToIndex prunes the deposit store up to the specified index.
	PruneToIndex(index uint64) error
	// EnqueueDeposits adds a list of deposits to the deposit store.
	EnqueueDeposits(deposits []*types.Deposit) error
}