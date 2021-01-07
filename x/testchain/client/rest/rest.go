package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers testchain-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/testchain/item", createItemHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/testchain/item", listItemHandler(cliCtx, "testchain")).Methods("GET")
		r.HandleFunc("/testchain/item/{key}", getItemHandler(cliCtx, "testchain")).Methods("GET")
		r.HandleFunc("/testchain/item", setItemHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/testchain/item", deleteItemHandler(cliCtx)).Methods("DELETE")

		
}
