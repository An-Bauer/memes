package qrcode

type Matrix [21][21]bool

type EccLevel int

const (
	L = iota
	M
	Q
	H
)
