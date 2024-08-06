// Copyright 2020 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	"golang.org/x/crypto/sha3"
)

// func u64(val uint64) *uint64 { return &val }

// TestStateProcessorErrors tests the output from the 'core' errors
// as defined in core/error.go. These errors are generated when the
// blockchain imports bad blocks, meaning blocks which have valid headers but
// contain invalid transactions
func TestStateProcessorErrors(t *testing.T) {
	var (
		config = &params.ChainConfig{
			ChainID:                       big.NewInt(1),
			HomesteadBlock:                big.NewInt(0),
			EIP150Block:                   big.NewInt(0),
			EIP155Block:                   big.NewInt(0),
			EIP158Block:                   big.NewInt(0),
			ByzantiumBlock:                big.NewInt(0),
			ConstantinopleBlock:           big.NewInt(0),
			PetersburgBlock:               big.NewInt(0),
			IstanbulBlock:                 big.NewInt(0),
			MuirGlacierBlock:              big.NewInt(0),
			BerlinBlock:                   big.NewInt(0),
			LondonBlock:                   big.NewInt(0),
			Ethash:                        new(params.EthashConfig),
			TerminalTotalDifficulty:       big.NewInt(0),
			TerminalTotalDifficultyPassed: true,
			ShanghaiTime:                  new(uint64),
			CancunTime:                    new(uint64),
		}
		signer  = types.LatestSigner(config)
		key1, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	)

	var mkDynamicTx = func(nonce uint64, to common.Address, gasLimit uint64, gasTipCap, gasFeeCap *big.Int) *types.Transaction {
		tx, _ := types.SignTx(types.NewTx(&types.DynamicFeeTx{
			Nonce:     nonce,
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     big.NewInt(0),
		}), signer, key1)
		return tx
	}

	// ErrTxTypeNotSupported, For this, we need an older chain
	{
		var (
			db    = rawdb.NewMemoryDatabase()
			gspec = &Genesis{
				Config: &params.ChainConfig{
					ChainID:             big.NewInt(1),
					HomesteadBlock:      big.NewInt(0),
					EIP150Block:         big.NewInt(0),
					EIP155Block:         big.NewInt(0),
					EIP158Block:         big.NewInt(0),
					ByzantiumBlock:      big.NewInt(0),
					ConstantinopleBlock: big.NewInt(0),
					PetersburgBlock:     big.NewInt(0),
					IstanbulBlock:       big.NewInt(0),
					MuirGlacierBlock:    big.NewInt(0),
				},
				Alloc: types.GenesisAlloc{
					common.HexToAddress("0x71562b71999873DB5b286dF957af199Ec94617F7"): types.Account{
						Balance: big.NewInt(1000000000000000000), // 1 ether
						Nonce:   0,
					},
				},
			}
			blockchain, _ = NewBlockChain(db, nil, gspec, nil, ethash.NewFaker(), vm.Config{}, nil, nil)
		)
		defer blockchain.Stop()
		for i, tt := range []struct {
			txs  []*types.Transaction
			want string
		}{
			{ // ErrTxTypeNotSupported
				txs: []*types.Transaction{
					mkDynamicTx(0, common.Address{}, params.TxGas-1000, big.NewInt(0), big.NewInt(0)),
				},
				want: "could not apply tx 0 [0x88626ac0d53cb65308f2416103c62bb1f18b805573d4f96a3640bbbfff13c14f]: transaction type not supported",
			},
		} {
			block := GenerateBadBlock(gspec.ToBlock(), ethash.NewFaker(), tt.txs, gspec.Config)
			_, err := blockchain.InsertChain(types.Blocks{block})
			if err == nil {
				t.Fatal("block imported without errors")
			}
			if have, want := err.Error(), tt.want; have != want {
				t.Errorf("test %d:\nhave \"%v\"\nwant \"%v\"\n", i, have, want)
			}
		}
	}
}

// GenerateBadBlock constructs a "block" which contains the transactions. The transactions are not expected to be
// valid, and no proper post-state can be made. But from the perspective of the blockchain, the block is sufficiently
// valid to be considered for import:
// - valid pow (fake), ancestry, difficulty, gaslimit etc
func GenerateBadBlock(parent *types.Block, engine consensus.Engine, txs types.Transactions, config *params.ChainConfig) *types.Block {
	difficulty := big.NewInt(0)
	if !config.TerminalTotalDifficultyPassed {
		fakeChainReader := newChainMaker(nil, config, engine)
		difficulty = engine.CalcDifficulty(fakeChainReader, parent.Time()+10, &types.Header{
			Number:     parent.Number(),
			Time:       parent.Time(),
			Difficulty: parent.Difficulty(),
			UncleHash:  parent.UncleHash(),
		})
	}

	header := &types.Header{
		ParentHash: parent.Hash(),
		Coinbase:   parent.Coinbase(),
		Difficulty: difficulty,
		GasLimit:   parent.GasLimit(),
		Number:     new(big.Int).Add(parent.Number(), common.Big1.ToBig()),
		Time:       parent.Time() + 10,
		UncleHash:  types.EmptyUncleHash,
	}
	if config.IsLondon(header.Number) {
		header.BaseFee = eip1559.CalcBaseFee(config, parent.Header())
	}
	if config.IsShanghai(header.Number, header.Time) {
		header.WithdrawalsHash = &types.EmptyWithdrawalsHash
	}
	var receipts []*types.Receipt
	// The post-state result doesn't need to be correct (this is a bad block), but we do need something there
	// Preferably something unique. So let's use a combo of blocknum + txhash
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(header.Number.Bytes())
	var cumulativeGas uint64
	var nBlobs int
	for _, tx := range txs {
		txh := tx.Hash()
		hasher.Write(txh[:])
		receipt := types.NewReceipt(nil, false, cumulativeGas+tx.Gas())
		receipt.TxHash = tx.Hash()
		receipt.GasUsed = tx.Gas()
		receipts = append(receipts, receipt)
		cumulativeGas += tx.Gas()
		nBlobs += len(tx.BlobHashes())
	}
	header.Root = common.BytesToHash(hasher.Sum(nil))
	if config.IsCancun(header.Number, header.Time) {
		var pExcess, pUsed = uint64(0), uint64(0)
		if parent.ExcessBlobGas() != nil {
			pExcess = *parent.ExcessBlobGas()
			pUsed = *parent.BlobGasUsed()
		}
		excess := eip4844.CalcExcessBlobGas(pExcess, pUsed)
		used := uint64(nBlobs * params.BlobTxBlobGasPerBlob)
		header.ExcessBlobGas = &excess
		header.BlobGasUsed = &used

		beaconRoot := common.HexToHash("0xbeac00")
		header.ParentBeaconRoot = &beaconRoot
	}
	// Assemble and return the final block for sealing
	if config.IsShanghai(header.Number, header.Time) {
		return types.NewBlockWithWithdrawals(header, txs, nil, receipts, []*types.Withdrawal{}, trie.NewStackTrie(nil))
	}
	return types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil))
}
