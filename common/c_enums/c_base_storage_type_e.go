//go:generate stringer -type=EStorageType -output=c_base_storage_type_e_string.go
package c_base

type EStorageType string

const (
	EStorageTypeInfluxDB1 EStorageType = "idb1"
	EStorageTypeInfluxDB2 EStorageType = "idb2"
	EStorageTypePebbledb  EStorageType = "pdb"
	EStorageTypeTsdb      EStorageType = "tsdb"
)
