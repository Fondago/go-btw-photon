// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chainbook

import (
	"errors"

	"github.com/BTWhite/go-btw-photon/account"
	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	// ErrTxInvalidPrevTx is returned if the transaction specified an invalid
	// previous transaction hash.
	ErrTxInvalidPrevTx = errors.New("Invalid previous tx")
)

type DefaultProcessor struct {
	db *leveldb.Db
	am *account.AccountManager
}

func NewProcessor(db *leveldb.Db) *DefaultProcessor {
	dp := &DefaultProcessor{}
	dp.db = db
	dp.am = account.NewAccountManager(db)
	return dp
}

func (p *DefaultProcessor) Process(tx *types.Tx, ch *chain.Chain) error {
	a := p.am.Get(tx.RecipientId)
	a.Balance = types.NewCoin(a.Balance.Uint64() + tx.Amount.Uint64())

	return p.am.Save(a)
}

func (p *DefaultProcessor) Validate(tx *types.Tx, ch *chain.Chain) error {
	if !tx.PreviousTx.Equals(ch.LastTx()) {
		return ErrTxInvalidPrevTx
	}

	return nil
}