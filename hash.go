package conv

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// AES crypt data with key, iv
// key is either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
// iv is recommended 16 bytes.
func AES(data []byte, key []byte, iv []byte) error {
	n := len(key)
	if len(data) == 0 {
		return errors.New("no data")
	}

	if n != 16 && n != 24 && n != 32 {
		return errors.New("invalid key")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(data, data)
	return nil
}

func XOR(data []byte, key []byte) error {
	n := len(key)
	if n == 0 {
		return errors.New("key is empty")
	}

	if len(data) == 0 {
		return errors.New("data is empty")
	}

	for i, b := range data {
		data[i] = b ^ key[i%n]
	}
	return nil
}

func MD5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}

func SHA1(str string) string {
	sha1er := sha1.New()
	b := []byte(str)
	for len(b) > 0 {
		n, err := sha1er.Write(b)
		if err != nil {
			panic(err)
		}
		b = b[n:]
	}
	return hex.EncodeToString(sha1er.Sum(nil))
}

func SHA256(str string) string {
	sha256er := sha256.New()
	b := []byte(str)
	for len(b) > 0 {
		n, err := sha256er.Write(b)
		if err != nil {
			panic(err)
		}
		b = b[n:]
	}
	return hex.EncodeToString(sha256er.Sum(nil))
}

func Hash32(b []byte) [32]byte {
	v := b
	sum := sha256.Sum256(v)
	for i := 0; i < 3; i++ {
		v = append(sum[:], v...)
		sum = sha256.Sum256(v)
	}
	return sum
}
