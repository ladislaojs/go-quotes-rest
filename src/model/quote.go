package model

import "math/big"

type Quote struct {
	ID     big.Int `json:"id"`
	Author string  `json:"author"`
	Quote  string  `json:"quote"`
}
