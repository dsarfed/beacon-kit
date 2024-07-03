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

package engineprimitives

import (
	"unsafe"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/constants"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/ssz"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/ssz/merkleizer"
)

// Transactions is a typealias for [][]byte, which is how transactions are
// received in the execution payload.
//
// TODO: make it take a generic SpecT type.
type Transactions [][]byte

// HashTreeRoot returns the hash tree root of the Transactions list.
//
// NOTE: Uses a new merkleizer for each call.
func (txs Transactions) HashTreeRoot() (common.Root, error) {
	return txs.HashTreeRootWith(
		merkleizer.New[[32]byte, common.Root](),
	)
}

type BartioTransactions = ssz.List[ssz.Vector[ssz.Byte]]

// BartioTransactionsFromBytes creates a Transactions object from a byte slice.
func BartioTransactionsFromBytes(data [][]byte) *BartioTransactions {
	return ssz.ListFromElements(
		// TODO: Move this value to chain spec.
		constants.MaxTxsPerPayload,
		//#nosec:G103 // todo fix later.
		*(*[]ssz.Vector[ssz.Byte])(unsafe.Pointer(&data))...)
}

type ProperTransactions = ssz.List[*ssz.List[ssz.Byte]]

// ProperTransactionsFromBytes creates a Transactions object from a byte slice.
func ProperTransactionsFromBytes(data [][]byte) *ProperTransactions {
	txs := make([]*ssz.List[ssz.Byte], len(data))
	for i, tx := range data {
		//nolint:mnd // unhood later.
		txs[i] = ssz.ByteListFromBytes(tx, 1073741824)
	}

	y := ssz.ListFromElements(constants.MaxTxsPerPayload, txs...)

	return ssz.ListFromElements(
		constants.MaxTxsPerPayload,
		y.Elements()...,
	)
}

// TxsMerkleizer is a ssz merkleizer used for transactions.
//
// TODO: make the ChainSpec a generic on this type.
type TxsMerkleizer merkleizer.Merkleizer[[32]byte, common.Root]

// HashTreeRootWith returns the hash tree root of the Transactions list
// using the given merkleizer.
func (txs Transactions) HashTreeRootWith(
	merkleizer *merkleizer.Merkleizer[[32]byte, common.Root],
) (common.Root, error) {
	var (
		err   error
		roots = make([]common.Root, len(txs))
	)

	for i, tx := range txs {
		roots[i], err = merkleizer.MerkleizeByteSlice(tx)
		if err != nil {
			return common.Root{}, err
		}
	}

	return merkleizer.MerkleizeListComposite(roots, constants.MaxTxsPerPayload)
}
