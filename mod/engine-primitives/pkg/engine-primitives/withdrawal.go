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

package engineprimitives

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/constants"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/ssz"
)

// Withdrawal represents a validator withdrawal from the consensus layer.
//
//go:generate go run github.com/ferranbt/fastssz/sszgen -path withdrawal.go -objs Withdrawal -include ../../../primitives/pkg/math,../../../primitives/pkg/common,$GETH_PKG_INCLUDE/common,$GETH_PKG_INCLUDE/common/hexutil -output withdrawal.ssz.go
type Withdrawal struct {
	// Index is the unique identifier for the withdrawal.
	Index math.U64 `json:"index"`
	// Validator is the index of the validator initiating the withdrawal.
	Validator math.ValidatorIndex `json:"validatorIndex"`
	// Address is the execution address where the withdrawal will be sent.
	// It has a fixed size of 20 bytes.
	Address common.ExecutionAddress `json:"address"        ssz-size:"20"`
	// Amount is the amount of Gwei to be withdrawn.
	Amount math.Gwei `json:"amount"`
}

// Equals returns true if the Withdrawal is equal to the other.
func (w *Withdrawal) Equals(other *Withdrawal) bool {
	return w.Index == other.Index &&
		w.Validator == other.Validator &&
		w.Address == other.Address &&
		w.Amount == other.Amount
}

// Withdrawals represents a slice of withdrawals.
type Withdrawals []*Withdrawal

// HashTreeRoot returns the hash tree root of the Withdrawals list.
func (w Withdrawals) HashTreeRoot() (common.Root, error) {
	// TODO: read max withdrawals from the chain spec.
	return ssz.MerkleizeListComposite[any, math.U64](
		w, constants.MaxWithdrawalsPerPayload,
	)
}