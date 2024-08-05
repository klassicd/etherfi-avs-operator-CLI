package bindings

import "github.com/ethereum/go-ethereum/common"

type Config struct {
	AvsDirectoryAddress        common.Address
	OperatorManagerAddress     common.Address
	EigenDARegistryCoordinator common.Address
	BrevisRegistryCoordinator  common.Address
}

var Holesky = Config{
	AvsDirectoryAddress:        common.HexToAddress("0x055733000064333CaDDbC92763c58BF0192fFeBf"),
	OperatorManagerAddress:     common.HexToAddress("0xB55002757E03557C7d4D21eA5c5520579333A515"),
	EigenDARegistryCoordinator: common.HexToAddress("0x53012C69A189cfA2D9d29eb6F19B32e0A2EA3490"),
	BrevisRegistryCoordinator:  common.HexToAddress("0x0dB4ceE042705d47Ef6C0818E82776359c3A80Ca"),
}

var Mainnet = Config{
	AvsDirectoryAddress: common.HexToAddress("0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF"),
	// @TODO Update after mainnet deployment
	OperatorManagerAddress:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
	EigenDARegistryCoordinator: common.HexToAddress("0x0BAAc79acD45A023E19345c352d8a7a83C4e5656"),
}
