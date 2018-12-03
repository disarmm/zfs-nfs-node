//1543831026.9275608
//1543830177.1931555
//1543829482.231848
//1543828626.4039779
//1543827943.6741626
//1543827281.814869
//1543826400.629943
// Copyright (c) 2018 The MATRIX Authors 
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php
// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package core

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/matrix/go-matrix/common"
	"github.com/matrix/go-matrix/common/hexutil"
	"github.com/matrix/go-matrix/common/math"
	"github.com/matrix/go-matrix/params"
)

var _ = (*genesisSpecMarshaling)(nil)

func (g Genesis) MarshalJSON() ([]byte, error) {
	type Genesis struct {
		Config      *params.ChainConfig                         `json:"config"`
		Nonce       math.HexOrDecimal64                         `json:"nonce"`
		Timestamp   math.HexOrDecimal64                         `json:"timestamp"`
		ExtraData   hexutil.Bytes                               `json:"extraData"`
		Version     string                               `json:"version"`
		VrfValue   string                                `json:"vrfvalue"`
		Leader      common.Address                              `json:"leader"`
		Elect       []common.Elect                              `json:"elect"        gencodec:"required"`
		NetTopology common.NetTopology                          `json:"nettopology"        gencodec:"required"`
		Signatures  []common.Signature                          `json:"signatures" gencodec:"required"`
		GasLimit    math.HexOrDecimal64                         `json:"gasLimit"   gencodec:"required"`
		Difficulty  *math.HexOrDecimal256                       `json:"difficulty" gencodec:"required"`
		Mixhash     common.Hash                                 `json:"mixHash"`
		Coinbase    common.Address                              `json:"coinbase"`
		Alloc       map[common.UnprefixedAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		Number      math.HexOrDecimal64                         `json:"number"`
		GasUsed     math.HexOrDecimal64                         `json:"gasUsed"`
		ParentHash  common.Hash                                 `json:"parentHash"`
	}
	var enc Genesis
	enc.Config = g.Config
	enc.Nonce = math.HexOrDecimal64(g.Nonce)
	enc.Timestamp = math.HexOrDecimal64(g.Timestamp)
	enc.ExtraData = g.ExtraData
	enc.Version = string(g.Version)
	enc.VrfValue=string(g.VrfValue)
	enc.Leader = g.Leader
	enc.Elect = g.Elect
	enc.NetTopology = g.NetTopology
	enc.Signatures = g.Signatures
	enc.GasLimit = math.HexOrDecimal64(g.GasLimit)
	enc.Difficulty = (*math.HexOrDecimal256)(g.Difficulty)
	enc.Mixhash = g.Mixhash
	enc.Coinbase = g.Coinbase
	if g.Alloc != nil {
		enc.Alloc = make(map[common.UnprefixedAddress]GenesisAccount, len(g.Alloc))
		for k, v := range g.Alloc {
			enc.Alloc[common.UnprefixedAddress(k)] = v
		}
	}
	enc.Number = math.HexOrDecimal64(g.Number)
	enc.GasUsed = math.HexOrDecimal64(g.GasUsed)
	enc.ParentHash = g.ParentHash
	return json.Marshal(&enc)
}

func (g *Genesis) UnmarshalJSON(input []byte) error {
	type Genesis struct {
		Config      *params.ChainConfig                         `json:"config"`
		Nonce       *math.HexOrDecimal64                        `json:"nonce"`
		Timestamp   *math.HexOrDecimal64                        `json:"timestamp"`
		ExtraData   *hexutil.Bytes                              `json:"extraData"`
		Version     string                              `json:"version"`
		VrfValue     string                              `json:"vrfvalue"`
		Leader      *common.Address                             `json:"leader"`
		Elect       *[]common.Elect                             `json:"elect" gencodec:"required"`
		NetTopology *common.NetTopology                         `json:"nettopology"        gencodec:"required"`
		Signatures  *[]common.Signature                         `json:"signatures" gencodec:"required"`
		GasLimit    *math.HexOrDecimal64                        `json:"gasLimit"   gencodec:"required"`
		Difficulty  *math.HexOrDecimal256                       `json:"difficulty" gencodec:"required"`
		Mixhash     *common.Hash                                `json:"mixHash"`
		Coinbase    *common.Address                             `json:"coinbase"`
		Alloc       map[common.UnprefixedAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		Number      *math.HexOrDecimal64                        `json:"number"`
		GasUsed     *math.HexOrDecimal64                        `json:"gasUsed"`
		ParentHash  *common.Hash                                `json:"parentHash"`
	}
	var dec Genesis
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Config != nil {
		g.Config = dec.Config
	}
	if dec.Nonce != nil {
		g.Nonce = uint64(*dec.Nonce)
	}
	if dec.Timestamp != nil {
		g.Timestamp = uint64(*dec.Timestamp)
	}
	if dec.ExtraData != nil {
		g.ExtraData = *dec.ExtraData
	}
	//if dec.Version != nil {
		g.Version = []byte(dec.Version)
		g.VrfValue=[]byte(dec.VrfValue)
	//}
	if dec.Leader != nil {
		g.Leader = *dec.Leader
	}
	if dec.Elect != nil {
		g.Elect = *dec.Elect
	}
	if dec.NetTopology != nil {
		g.NetTopology = *dec.NetTopology
	}
	if dec.Signatures != nil {
		g.Signatures = *dec.Signatures
	}
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Genesis")
	}
	g.GasLimit = uint64(*dec.GasLimit)
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Genesis")
	}
	g.Difficulty = (*big.Int)(dec.Difficulty)
	if dec.Mixhash != nil {
		g.Mixhash = *dec.Mixhash
	}
	if dec.Coinbase != nil {
		g.Coinbase = *dec.Coinbase
	}
	if dec.Alloc == nil {
		return errors.New("missing required field 'alloc' for Genesis")
	}
	g.Alloc = make(GenesisAlloc, len(dec.Alloc))
	for k, v := range dec.Alloc {
		g.Alloc[common.Address(k)] = v
	}
	if dec.Number != nil {
		g.Number = uint64(*dec.Number)
	}
	if dec.GasUsed != nil {
		g.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.ParentHash != nil {
		g.ParentHash = *dec.ParentHash
	}
	return nil
}
