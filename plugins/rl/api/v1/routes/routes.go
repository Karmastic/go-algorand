package routes

import (
	"github.com/algorand/go-algorand/daemon/algod/api/server/lib"
	"github.com/algorand/go-algorand/plugins/rl/api/v1/handlers"
)

// Routes contains all routes for /rl/v1
var Routes = lib.Routes{
	lib.Route{
		Name:        "info",
		Method:      "GET",
		Path:        "/info",
		HandlerFunc: handlers.Info,
	},
	lib.Route{
		Name:        "teal-disassemble",
		Method:      "POST",
		Path:        "/teal/disassemble",
		HandlerFunc: handlers.DisassembleTeal,
	},
}

func GetRoutes() (lib.Routes, string, error) {
	return Routes, "rl/v1", nil
}
