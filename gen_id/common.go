package main

const (
	NodeIdBits uint8 = 10 // the biggest is 1024
	IdBits uint8 = 12 // the maximum every millisecond generate is 4096
	NodeIDMax   int64 = -1 ^ (-1 << NodeIdBits) // the maximum of node id
	IdMask  int64 = -1 ^ (-1 << IdBits) // the maximum of id, as mask off code
	TimeShift   uint8 = NodeIdBits + IdBits // timestamp left offset
	NodeIdShift uint8 = IdBits              // node id left offset
	Epoch int64 = 1584583903000 // this is the time I write this lib(millisecond)
)



