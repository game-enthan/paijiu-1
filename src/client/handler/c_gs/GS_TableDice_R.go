package c_gs

import (
	"client/msg"
)

func GS_TableDice_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableDice_R)
	log.Info("Res:", req.ErrorCode, req.DiceNum)
}
