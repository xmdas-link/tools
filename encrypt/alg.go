package encrypt

type encrypt interface {
	encode(string) string
	hash(string) []byte
}
