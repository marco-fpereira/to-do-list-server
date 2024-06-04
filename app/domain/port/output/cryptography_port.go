package output

type CryptographyPort interface {
	EncryptKey(rawKey string) (string, error)
	VerifyEncryptedKey(rawKey string, encryptedKey string) bool
}
