package types

const (
	// ModuleName defines the module name
	ModuleName = "authentication"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_authentication"
)

var (
	ParamsKey = []byte("p_authentication")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
