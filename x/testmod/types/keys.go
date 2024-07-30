package types

const (
	// ModuleName defines the module name
	ModuleName = "testmod"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_testmod"
)

var (
	ParamsKey = []byte("p_testmod")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
