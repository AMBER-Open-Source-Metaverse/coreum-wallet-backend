package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/CoreumFoundation/coreum/app"
	coreumconfig "github.com/CoreumFoundation/coreum/pkg/config"
	"github.com/CoreumFoundation/coreum/pkg/tx"

	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func main() {

	network, err := coreumconfig.NetworkByChainID(coreumconfig.ChainID("coreum-devnet-1"))

	checkErr(err)

	network.SetSDKConfig()

	// Init the mux router
	router := mux.NewRouter()

	// Get Balance of Account.
	router.HandleFunc("/get-balance/{address}", GetBalance).Methods("GET")

	// Create a new wallet
	router.HandleFunc("/create-new-wallet", CreateNewWallet).Methods("GET")

	// Recover wallet from Mnemonic
	router.HandleFunc("/recovery-wallet", RecoveryWallet).Methods("POST")

	// serve the app
	fmt.Println("Server at 5432")
	log.Fatal(http.ListenAndServe(":5432", router))
}

// type concurrentSafeKeyring struct {
// 	mu *sync.RWMutex
// 	kr keyring.Keyring
// }

// func newConcurrentSafeKeyring(kr keyring.Keyring) concurrentSafeKeyring {
// 	return concurrentSafeKeyring{
// 		mu: &sync.RWMutex{},
// 		kr: kr,
// 	}
// }

func GetBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	address := params["address"]

	type JsonResponse struct {
		Type    string `json:"type"`
		Data    string `json:"data"`
		Message string `json:"message"`
	}

	var response = JsonResponse{}

	if address == "" {
		response = JsonResponse{Type: "error", Message: "You are missing wallet address parameter."}
	} else {

		rpcAddress := "https://s-0.devnet-1.coreum.dev:443/"

		ctx := context.Background()

		rpcClient, err := cosmosclient.NewClientFromNode(rpcAddress)
		clientCtx := tx.NewClientContext(app.ModuleBasics).WithChainID("coreum-devnet-1").
			WithClient(rpcClient).
			WithKeyring(keyring.NewInMemory()).
			WithBroadcastMode(flags.BroadcastBlock)

		checkErr(err)

		// Getting Balance from Coreum Blockchain

		bankClient := banktypes.NewQueryClient(clientCtx)

		balance, err := bankClient.Balance(ctx, &banktypes.QueryBalanceRequest{
			Address: address,
			Denom:   "ducore",
		})

		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The Balance of the wallet is gotten successfully!", Data: balance.Balance.Amount.String()}
	}

	json.NewEncoder(w).Encode(response)
}

func CreateNewWallet(w http.ResponseWriter, r *http.Request) {
	kr := keyring.NewInMemory()
	info, mnemonic, err := kr.NewMnemonic("", keyring.English, sdk.GetConfig().GetFullBIP44Path(), "", hd.Secp256k1)

	checkErr(err)

	type JsonResponse struct {
		Mnemonic string `json:"mnemonic"`
		Address  string `json:"address"`
	}

	sdkAddr := info.GetAddress()

	var response = JsonResponse{Mnemonic: mnemonic, Address: sdkAddr.String()}

	json.NewEncoder(w).Encode(response)

}

func RecoveryWallet(w http.ResponseWriter, r *http.Request) {

	mnemonic := r.FormValue("mnemonic")

	kr := keyring.NewInMemory()

	info, err := kr.NewAccount("wallet", mnemonic, "", sdk.GetConfig().GetFullBIP44Path(), hd.Secp256k1)

	checkErr(err)

	json.NewEncoder(w).Encode(info.GetAddress())
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
