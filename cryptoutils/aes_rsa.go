package cryptoutils

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func EncryptAsymmetricAES(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	encryptedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, key, nil)
	if err != nil {
		return nil, err
	}

	if len(encryptedKey) != 256 {
		return nil, fmt.Errorf("expected AES encrypted key to be 256 bytes long")
	}

	// Write encrypted key (256 bytes), original file size (8 bytes), and IV (16 bytes).
	if _, err := writer.Write(encryptedKey); err != nil {
		return nil, err
	}

	if err := binary.Write(writer, binary.LittleEndian, uint64(len(plaintext))); err != nil {
		return nil, err
	}

	if _, err = writer.Write(iv); err != nil {
		return nil, err
	}

	// Pad plaintext to a multiple of BlockSize with random padding.
	if len(plaintext)%aes.BlockSize != 0 {
		bytesToPad := aes.BlockSize - (len(plaintext) % aes.BlockSize)
		padding := make([]byte, bytesToPad)
		if _, err := rand.Read(padding); err != nil {
			return nil, err
		}

		plaintext = append(plaintext, padding...)
	}

	// Use AES implementation of the cipher.Block interface to encrypt the whole
	// file in CBC mode.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	if _, err = writer.Write(ciphertext); err != nil {
		return nil, err
	}

	if err := writer.Flush(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func DecryptAsymmetricAES(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	buf := bytes.NewReader(ciphertext)

	if len(ciphertext) < 256+8+aes.BlockSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}

	encryptedKey := make([]byte, 256)
	if _, err := buf.Read(encryptedKey); err != nil {
		return nil, err
	}

	var origSize uint64
	if err := binary.Read(buf, binary.LittleEndian, &origSize); err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := buf.Read(iv); err != nil {
		return nil, err
	}

	key, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedKey, nil)
	if err != nil {
		return nil, err
	}

	if len(key) != 32 {
		return nil, fmt.Errorf("expected AES key to be 32 bytes long")
	}

	// The remaining ciphertext has size=paddedSize.
	paddedSize := len(ciphertext) - 256 - 8 - aes.BlockSize
	if paddedSize%aes.BlockSize != 0 {
		return nil, fmt.Errorf("want padded plaintext size to be aligned to block size")
	}
	plaintext := make([]byte, paddedSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext[256+8+aes.BlockSize:])

	return plaintext[:origSize], nil
}

// PublicKeyAsymmetricAES returns the AES asymmentric (RSA) public key from a file.
func PublicKeyAsymmetricAES(filename string) (*rsa.PublicKey, error) {
	publicKey, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error while loading public key from PEM file: %v", err)
	}

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	var key rsa.PublicKey
	_, err = asn1.Unmarshal(block.Bytes, &key)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %v", err)
	}

	return &key, nil
}

// PrivateKeyAsymmetricAES returns the AES asymmentric (RSA) private key from a file.
func PrivateKeyAsymmetricAES(filename string) (*rsa.PrivateKey, error) {
	privateKey, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error while loading private key from PEM file: %v", err)
	}

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
