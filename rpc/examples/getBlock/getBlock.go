// Copyright 2021 github.com/gagliardetto
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.DevNet_RPC
	client := rpc.New(endpoint)

	// 获得失败GetRecentBlockhash  没有此方法
	// example, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentProcessed)
	// if err != nil {
	// 	panic(err)
	// }
	// spew.Dump(example)
	{
		out, err := client.GetBlock(context.TODO(), uint64(319836958))
		if err != nil {
			panic(err)
		}
		spew.Dump(out) // NOTE: This generates a lot of output.

	}

	// {
	// 	includeRewards := false
	// 	out, err := client.GetBlockWithOpts(
	// 		context.TODO(),
	// 		uint64(example.Context.Slot),
	// 		// You can specify more options here:
	// 		&rpc.GetBlockOpts{
	// 			Encoding:   solana.EncodingBase64,
	// 			Commitment: rpc.CommitmentFinalized,
	// 			// Get only signatures:
	// 			TransactionDetails: rpc.TransactionDetailsSignatures,
	// 			// Exclude rewards:
	// 			Rewards: &includeRewards,
	// 		},
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	spew.Dump(out)
	// }
}

// (int) 12

// 输出
// spew.Dump(len(out.Transactions))
// 		spew.Dump(out.Transactions[0])
// (rpc.TransactionWithMeta) {
//  Slot: (uint64) 0,
//  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
//  Transaction: (*rpc.DataBytesOrJSON)(0xc000168820)({
//   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
//   asDecodedBinary: (solana.Data) AUoh2k3J/d7ZAckTp6W4w/8wkvpREncJA6AzXIrUCa3Z02EY3YhTiG0hmkIaJXSsBmbxCqysYdV7wRDJ5IWVTQ0BAAEDadhaiV3m4mMXKduJBqVIEV/3Pe+rRUwKbHyNNCDGyX4Q3mfdhGp/Yagr6sXMgwqLob5VxcWGzw0/7jr+58gfbAdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAA+7g8n8Om33uqYi7mtuVyhjI2ivvPYpR6WzqxD3KZrYMBAgIBAHQMAAAA/VIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGSrnG2t6u6JScNTfE8RMHsAdSEl7fAG4dOD81PyZXP+QFqxMJmAAAAAA==,
//   asJSON: (json.RawMessage) <nil>
//  }),
//  Meta: (*rpc.TransactionMeta)(0xc00019ab00)({
//   Err: (interface {}) <nil>,
//   Fee: (uint64) 5000,
//   PreBalances: ([]uint64) (len=3 cap=4) {
//    (uint64) 11161445714088,
//    (uint64) 773079315933,
//    (uint64) 1
//   },
//   PostBalances: ([]uint64) (len=3 cap=4) {
//    (uint64) 11161445709088,
//    (uint64) 773079315933,
//    (uint64) 1
//   },
//   InnerInstructions: ([]rpc.InnerInstruction) {
//   },
//   PreTokenBalances: ([]rpc.TokenBalance) {
//   },
//   PostTokenBalances: ([]rpc.TokenBalance) {
//   },
//   LogMessages: ([]string) (len=2 cap=2) {
//    (string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
//    (string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
//   },
//   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
//    (string) (len=2) "Ok": (interface {}) <nil>
//   },
//   Rewards: ([]rpc.BlockReward) {
//   },
//   LoadedAddresses: (rpc.LoadedAddresses) {
//    ReadOnly: (solana.PublicKeySlice) {
//    },
//    Writable: (solana.PublicKeySlice) {
//    }
//   },
//   ReturnData: (rpc.ReturnData) {
//    ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
//    Data: (solana.Data)
//   },
//   ComputeUnitsConsumed: (*uint64)(0xc0003eaa60)(2100)
//  }),
//  Version: (rpc.TransactionVersion) 0
// }

// 输出out

// (*rpc.GetBlockResult)(0xc0000c1a40)({
// 	Blockhash: (solana.Hash) (len=32 cap=32) 5vKYdK5Ye62CJeZGV4EEbQTsij8FuJR8kHVoEqpXqgP6,
// 	PreviousBlockhash: (solana.Hash) (len=32 cap=32) A4ZpZKMZzg23hQQesUxAwZaQGsHZ21KeuExMrHPhMYth,
// 	ParentSlot: (uint64) 319836957,
// 	Transactions: ([]rpc.TransactionWithMeta) (len=12 cap=16) {
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8820)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) AUoh2k3J/d7ZAckTp6W4w/8wkvpREncJA6AzXIrUCa3Z02EY3YhTiG0hmkIaJXSsBmbxCqysYdV7wRDJ5IWVTQ0BAAEDadhaiV3m4mMXKduJBqVIEV/3Pe+rRUwKbHyNNCDGyX4Q3mfdhGp/Yagr6sXMgwqLob5VxcWGzw0/7jr+58gfbAdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAA+7g8n8Om33uqYi7mtuVyhjI2ivvPYpR6WzqxD3KZrYMBAgIBAHQMAAAA/VIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGSrnG2t6u6JScNTfE8RMHsAdSEl7fAG4dOD81PyZXP+QFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012ab00)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 11161445714088,
// 		(uint64) 773079315933,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 11161445709088,
// 		(uint64) 773079315933,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003daa70)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f88c0)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) Af8qo+hYp0EJASfpAGBgKGszshjv0BMkRkYFgqLo9oHrszttYcqfd+6QjS5zFhUJzDzPkAYJeBN89+jDlE6Q9AABAAEDCXTkQLNJkdxKN7M4VtovPqyif2Ylkmd1//X41GWjvOcNwJEeinmQHle5sclss9S2jEuTwcDy2hwDp7Lxd76F3wdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012ac60)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 211323810415084,
// 		(uint64) 2285633188123568,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 211323810410084,
// 		(uint64) 2285633188123568,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003dab48)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8910)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) ATgSHSLQ5anBTJQ2Qsr3MfYwrqYhAllJI+hT/qKD55j/rfUoAVauluRgSAveUvixYTKSgvbbOdlYnu8CIvxR+QEBAAEDCXTl4uVueDgPjMDFzva6DzaSSvZOO4/TWEL4CerBhiQPae3eTL/R+1wk3j98gr0FHPS5+SGBCkbgHx/BEC/6sgdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012adc0)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 1411608746497276,
// 		(uint64) 2285646897290820,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 1411608746492276,
// 		(uint64) 2285646897290820,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003dac08)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8960)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) ATUuzNz36u7Tlxewc7xRIwN7fJkcPhA5dGO06LWvDlMf2ggT+W02AlUZD8Xot0oHybXUdfgjmxCDoyUrkcZYcgEBAAEDCXTeVPd6+prT1HSQn2zglbliOjWaXQjMgHuKkt1eH5YKiAaVN9osvzAWT9Si70lMmML05TxzF4KLoavohN5bhAdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012af20)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 1010764103743199,
// 		(uint64) 2285729220287594,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 1010764103738199,
// 		(uint64) 2285729220287594,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003dacd8)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f89b0)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) AdHqU/q5A81/ajg/9ERJuNAFy5GIp/Bg8HdT83c3eUigxgpGpJJTL3tCE10lHzlbe6WQWp4qGJZbbfy+fv4hkgUBAAED3qnGiHxsd4x3P3TGh/mxlOxIveo0AAr8FvpC21b+5rS1i/Vu3jlqDPCWfz+vXdzeX8+o1Rd78z1wkVgpDDsKjgdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012b080)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 30624888603,
// 		(uint64) 1366093348,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 30624883603,
// 		(uint64) 1366093348,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003dad98)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8a00)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) Ab5xjCcPAsKWnb3Ht/ZydBTS4taUy1WPtpAaDdLMCXa+Ef/L3g+JqYsvxyG95ZF8YjuGxHhxFt9SyBLsBnbqAgQBAAEDsUwo4/srKBSDxYknHsvJEe7pdB81Mh3+dSyGa6l332Bbg4/CwZ3FEuIbep8hEGH5HyN1Q2EqTA4mNJCG8n4I6gdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012b1e0)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 783593464507,
// 		(uint64) 308914197381346,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 783593459507,
// 		(uint64) 308914197381346,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003dae58)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8a50)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) AWZivHoA2dRYTAypdZOAaeNqXxVGb8eFY/Xs12DtaOE0fyLv0CORNTYp/0qamjFYPYkiu1iZo4dtHr/V0T9soQIBAAEDeIqSmkGF70CXUdSlfTJAR7AJkBbmZVPrPbBMjBMDEF24zdnZsBMfZN4FEzCj+m93AIBuRoMEhsve3hYVj0ZolQdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012b340)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 102094187809,
// 		(uint64) 224548278801117,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 102094182809,
// 		(uint64) 224548278801117,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003daf18)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8aa0)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) Aaw4qvRN3UaCtrfr+uxf5c+ZHyBEKDrogFr2bUfEVK8o4wmzS56Jcz9YbQ8J1OadncgHGnhOW07ffxr3v2TVgAEBAAED6vMdzHFHy5+UhZpxwL6b7Uq0uifl4gVjOC6v0KOWORhGHpqyAvIkIVo+tW8+HcrUwMCAeh5LT0WAthYMGZQeuwdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012b4a0)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 1364587818692,
// 		(uint64) 564074200958,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 1364587813692,
// 		(uint64) 564074200958,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003dafd8)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8af0)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) AR9xuMjJpbFwn+BfvqBXTik5EPp9eC0oomCT2uINS0totYtdxJfthsxhyZWMmwVW+WDEfXWlumxJtcRjlnwIwQYBAAEDtZKCjs2FQelpBuUcqubagIIM806+OB7KU/pLN1fVlKPQZLXQphQLO9iiP6R9Q8oBbJUmoW/TIfYA1IQrQjLXQwdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012b600)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 10997620600,
// 		(uint64) 173502878,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 10997615600,
// 		(uint64) 173502878,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003db098)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8b40)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) AQzCZ7/5ClMW/TUCqt0Dpi2wPRc/RKlUrwLVEvo93DH459QG4iVrpQoezBOUCuIu3mFBxUZl8w/QsrwrI2cqawoBAAEDqCSTnytuM7EVv2DOGQyF7pPfAoVpHVPqLG979vQPp9vENTAKSdXsY3/NgSRMofYRkgf4vD/udXJpddx+GHCvkwdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012b760)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 1348652453975,
// 		(uint64) 246216169220,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 1348652448975,
// 		(uint64) 246216169220,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003db158)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8b90)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) AVNvQ5urSiBxJaRzl9YIV4Er6PgfvLlrMQVCC+3e8NgCU6GuCZS7ufpcTEmCWSr4IQDIdv5bq0Rk0QSSN0lL8woBAAEDCXTY4+NwbM72hpR0lLBqKGZCnDhrN4yKGNUOJGWn6cVDwuUEDggs5o/YD1j44aSkcLenTTBQ6mjIfH04DlaYZQdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAAhqMIzo1W0+CNNm0v++zWUw2auYMRQ+zu2wF3IYctKPIBAgIBAHQMAAAA/lIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGZ2Np4B3jaB+QayKBXv8/rARnWdm2F4N4aup7xkxJ83AFqxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012b8c0)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 5951859630740,
// 		(uint64) 2213918269545815,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 5951859625740,
// 		(uint64) 2213918269545815,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003db218)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 },
// 	 (rpc.TransactionWithMeta) {
// 	  Slot: (uint64) 0,
// 	  BlockTime: (*solana.UnixTimeSeconds)(<nil>),
// 	  Transaction: (*rpc.DataBytesOrJSON)(0xc0000f8be0)({
// 	   rawDataEncoding: (solana.EncodingType) (len=6) "base64",
// 	   asDecodedBinary: (solana.Data) AV/FLWOG1bYXoJXJNf9eqgqiqExSioTOIzyQUQzSG9QuGGOHSXHTUJHJV1/TdJ6kVkm2Pf0/SpyoHYF5/dl/6Q4BAAED0zDq8eLi+YgYYFXL1iL4OQ3JmJYfpkL47RtGexA4gAjd9CoEgApU3i5YP5TxewiXJbdy0TM1JicSQVMndtL/xgdhSB01dHS7fE12JOvTvbPYNV5z0RBD/A2jU4AAAAAA+7g8n8Om33uqYi7mtuVyhjI2ivvPYpR6WzqxD3KZrYMBAgIBAHQMAAAA/VIQEwAAAAAfAR8BHgEdARwBGwEaARkBGAEXARYBFQEUARMBEgERARABDwEOAQ0BDAELAQoBCQEIAQcBBgEFAQQBAwECAQGSrnG2t6u6JScNTfE8RMHsAdSEl7fAG4dOD81PyZXP+QFrxMJmAAAAAA==,
// 	   asJSON: (json.RawMessage) <nil>
// 	  }),
// 	  Meta: (*rpc.TransactionMeta)(0xc00012ba20)({
// 	   Err: (interface {}) <nil>,
// 	   Fee: (uint64) 5000,
// 	   PreBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 2460680209086,
// 		(uint64) 1759663000380,
// 		(uint64) 1
// 	   },
// 	   PostBalances: ([]uint64) (len=3 cap=4) {
// 		(uint64) 2460680204086,
// 		(uint64) 1759663000380,
// 		(uint64) 1
// 	   },
// 	   InnerInstructions: ([]rpc.InnerInstruction) {
// 	   },
// 	   PreTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   PostTokenBalances: ([]rpc.TokenBalance) {
// 	   },
// 	   LogMessages: ([]string) (len=2 cap=2) {
// 		(string) (len=62) "Program Vote111111111111111111111111111111111111111 invoke [1]",
// 		(string) (len=59) "Program Vote111111111111111111111111111111111111111 success"
// 	   },
// 	   Status: (rpc.DeprecatedTransactionMetaStatus) (len=1) {
// 		(string) (len=2) "Ok": (interface {}) <nil>
// 	   },
// 	   Rewards: ([]rpc.BlockReward) {
// 	   },
// 	   LoadedAddresses: (rpc.LoadedAddresses) {
// 		ReadOnly: (solana.PublicKeySlice) {
// 		},
// 		Writable: (solana.PublicKeySlice) {
// 		}
// 	   },
// 	   ReturnData: (rpc.ReturnData) {
// 		ProgramId: (solana.PublicKey) (len=32 cap=32) 11111111111111111111111111111111,
// 		Data: (solana.Data)
// 	   },
// 	   ComputeUnitsConsumed: (*uint64)(0xc0003db2d8)(2100)
// 	  }),
// 	  Version: (rpc.TransactionVersion) 0
// 	 }
// 	},
// 	Signatures: ([]solana.Signature) <nil>,
// 	Rewards: ([]rpc.BlockReward) (len=6 cap=8) {
// 	 (rpc.BlockReward) {
// 	  Pubkey: (solana.PublicKey) (len=32 cap=32) 7G5idnXooSt7JQKCuxMUYJF4mh2FhqR7QG1u7jumebDc,
// 	  Lamports: (int64) -46,
// 	  PostBalance: (uint64) 71356,
// 	  RewardType: (rpc.RewardType) (len=4) "Rent",
// 	  Commission: (*uint8)(<nil>)
// 	 },
// 	 (rpc.BlockReward) {
// 	  Pubkey: (solana.PublicKey) (len=32 cap=32) dv3qDFk1DTF36Z62bNvrCXe9sKATA6xvVy6A798xxAS,
// 	  Lamports: (int64) 30000,
// 	  PostBalance: (uint64) 211323810440084,
// 	  RewardType: (rpc.RewardType) (len=3) "Fee",
// 	  Commission: (*uint8)(<nil>)
// 	 },
// 	 (rpc.BlockReward) {
// 	  Pubkey: (solana.PublicKey) (len=32 cap=32) dv3qDFk1DTF36Z62bNvrCXe9sKATA6xvVy6A798xxAS,
// 	  Lamports: (int64) 6,
// 	  PostBalance: (uint64) 211323810440090,
// 	  RewardType: (rpc.RewardType) (len=4) "Rent",
// 	  Commission: (*uint8)(<nil>)
// 	 },
// 	 (rpc.BlockReward) {
// 	  Pubkey: (solana.PublicKey) (len=32 cap=32) dv1ZAGvdsz5hHLwWXsVnM94hWf1pjbKVau1QVkaMJ92,
// 	  Lamports: (int64) 6,
// 	  PostBalance: (uint64) 5951859625746,
// 	  RewardType: (rpc.RewardType) (len=4) "Rent",
// 	  Commission: (*uint8)(<nil>)
// 	 },
// 	 (rpc.BlockReward) {
// 	  Pubkey: (solana.PublicKey) (len=32 cap=32) dv2eQHeP4RFrJZ6UeiZWoc3XTtmtZCUKxxCApCDcRNV,
// 	  Lamports: (int64) 6,
// 	  PostBalance: (uint64) 1010764103738205,
// 	  RewardType: (rpc.RewardType) (len=4) "Rent",
// 	  Commission: (*uint8)(<nil>)
// 	 },
// 	 (rpc.BlockReward) {
// 	  Pubkey: (solana.PublicKey) (len=32 cap=32) dv4ACNkpYPcE3aKmYDqZm9G5EB3J4MRoeE7WNDRBVJB,
// 	  Lamports: (int64) 5,
// 	  PostBalance: (uint64) 1411608746492281,
// 	  RewardType: (rpc.RewardType) (len=4) "Rent",
// 	  Commission: (*uint8)(<nil>)
// 	 }
// 	},
// 	BlockTime: (*solana.UnixTimeSeconds)(0xc0003da8b0)(2024-08-19 12:04:58 +0800 CST),
// 	BlockHeight: (*uint64)(0xc0003da888)(308061968)
//    })
