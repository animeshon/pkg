package cryptoutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteEncryption(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.Nil(t, err)

	// PRIVATE
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	privateFile, err := os.Create("testdata/fixtures/private-key.pem")
	require.Nil(t, err)
	defer privateFile.Close()

	require.Nil(t, pem.Encode(privateFile, privateKey))

	// PUBLIC
	body, err := asn1.Marshal(key.PublicKey)
	require.Nil(t, err)

	var publicKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: body,
	}

	publicFile, err := os.Create("testdata/fixtures/public-key.pem")
	require.Nil(t, err)
	defer publicFile.Close()

	require.Nil(t, pem.Encode(publicFile, publicKey))
}

func TestReadEncryption(t *testing.T) {
	// PUBLIC
	body, err := ioutil.ReadFile("testdata/fixtures/public-key.pem")
	require.Nil(t, err)

	block, _ := pem.Decode(body)
	assert.Equal(t, block.Type, "PUBLIC KEY")

	var publicKey rsa.PublicKey
	_, err = asn1.Unmarshal(block.Bytes, &publicKey)
	require.Nil(t, err)

	ciphertext, err := EncryptAsymmetricAES(&publicKey, []byte("secret message"))
	require.Nil(t, err)

	// PRIVATE
	body, err = ioutil.ReadFile("testdata/fixtures/private-key.pem")
	require.Nil(t, err)

	block, _ = pem.Decode(body)
	assert.Equal(t, block.Type, "PRIVATE KEY")

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	require.Nil(t, err)

	plaintext, err := DecryptAsymmetricAES(privateKey, ciphertext)
	require.Nil(t, err)

	assert.Equal(t, "secret message", string(plaintext))
}

// func TestReadFile(t *testing.T) {
// 	ciphertext, err := ioutil.ReadFile("test.enc")
// 	require.Nil(t, err)

// 	// PRIVATE
// 	body, err := ioutil.ReadFile("testdata/fixtures/private-key.pem")
// 	require.Nil(t, err)

// 	block, _ := pem.Decode(body)
// 	assert.Equal(t, block.Type, "PRIVATE KEY")

// 	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
// 	require.Nil(t, err)

// 	plaintext, err := DecryptAsymmetricAES(privateKey, ciphertext)
// 	require.Nil(t, err)

// 	err = ioutil.WriteFile("test.jpg", plaintext, 0664)
// 	require.Nil(t, err)
// }
