package handlers

import "log"

type (
	Products   struct{ l *log.Logger }
	KeyProduct struct{}
)

func NewProducts(l *log.Logger) *Products { return &Products{l} }
