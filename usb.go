package main

type device struct {
	model int // 0 - planck // 1 - ergodox // 2 - moonlander
	bus   int
	port  int
}

const (
	vendorID1    int = 0xFEED
	vendorID2    int = 0x3297
	planckID     int = 0x6060
	ergodoxID    int = 0x1307
	moonlanderID int = 0x1969

	dfuSuffixVendorID  int = 0x83
	dfuSuffixProductID int = 0x11
	dfuVendorID        int = 0x0483
	dfuProductID       int = 0xdf11

	halfKayVendorID  int = 0x16C0
	halfKayProductID int = 0x0478

	ergodoxMaxMemorySize = 0x100000
	ergodoxCodeSize      = 32256
	ergodoxBlockSize     = 128

	dfuSuffixLength    = 16
	planckBlockSize    = 2048
	planckStartAddress = 0x08000000
	setAddress         = 0
	eraseAddress       = 1
	eraseFlash         = 2
)
