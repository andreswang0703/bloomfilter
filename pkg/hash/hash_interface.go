package hash

type Hash_Interface interface {
	Hash(input string) uint64
}
