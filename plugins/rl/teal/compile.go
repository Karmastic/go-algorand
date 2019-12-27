package teal

import (
	"github.com/algorand/go-algorand/data/transactions"
	"github.com/algorand/go-algorand/data/transactions/logic"
	"github.com/algorand/go-algorand/protocol"
)

func DisassembleFile(program []byte) (string, *logic.DebugData, error) {
	// try parsing it as a msgpack LogicSig
	var lsig transactions.LogicSig
	err := protocol.Decode(program, &lsig)
	if err == nil {
		// success, extract program to disassemble
		program = lsig.Logic
	}
	text, dd, err := logic.DisassembleWithSymbols(program)
	if err != nil {
		return "", dd, err
	} else {
		dd.ErrorLoc = logic.CodePoint{}
	}
	return text, dd, nil
}
