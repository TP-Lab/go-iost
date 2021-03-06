// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iwallet

import (
	"fmt"
	"github.com/iost-official/go-iost/sdk"

	"github.com/iost-official/go-iost/ilog"
	"github.com/iost-official/go-iost/rpc/pb"
	"github.com/spf13/cobra"
)

// callCmd represents the call command that call a contract with given actions.
var callCmd = &cobra.Command{
	Use:   "call [ACTION]...",
	Short: "Call the method in contracts",
	Long: `Call the method in contracts
	Would accept arguments as call actions or load transaction request directly from given file (which could be generated by "save" command).
	An ACTION is a group of 3 arguments: contract name, function name, method parameters.
	The method parameters should be a string with format '["arg0","arg1",...]'.`,
	Example: `  iwallet call "token.iost" "transfer" '["iost","user0001","user0002","123.45",""]' --account test0
  iwallet call --tx_file tx.json --account test0`,
	Args: func(cmd *cobra.Command, args []string) error {
		return checkAccount(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		trx := &rpcpb.TransactionRequest{}
		if txFile != "" {
			if len(args) != 0 {
				ilog.Warnf("load tx from file %v, will ignore cmd args %v", txFile, args)
			}
			err := sdk.LoadProtoStructFromJSONFile(txFile, trx)
			if err != nil {
				return err
			}
		} else {
			var actions []*rpcpb.Action
			actions, err := actionsFromFlags(args)
			if err != nil {
				return err
			}
			trx, err = iwalletSDK.CreateTxFromActions(actions)
			if err != nil {
				return err
			}
		}

		err := InitAccount()
		if err != nil {
			return fmt.Errorf("failed to load account: %v", err)
		}

		if err := checkSigners(signers); err != nil {
			return err
		}
		trx.Signers = signers

		if len(withSigns) != 0 || len(signKeys) != 0 {
			ilog.Infof("making multi sig...")
			err = handleMultiSig(trx, withSigns, signKeys)
			if err != nil {
				return fmt.Errorf("multi sig err %v", err)
			}
		}

		_, err = iwalletSDK.SendTx(trx)
		return err
	},
}

func init() {
	rootCmd.AddCommand(callCmd)
	callCmd.Flags().StringSliceVarP(&signKeys, "sign_keys", "", []string{}, "optional private key files used for signing, split by comma")
	callCmd.Flags().StringSliceVarP(&withSigns, "with_signs", "", []string{}, "optional signatures, split by comma")
	callCmd.Flags().StringVarP(&txFile, "tx_file", "", "", "load tx from this file")
}

var (
	// used for multi sig
	signKeys  []string
	withSigns []string
)
