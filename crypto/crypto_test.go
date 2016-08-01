package crypto

import (
	"bytes"
	"testing"
)

var c Crypto
var enc, plain_text, dec, label []byte

func TestNewCrypto(t *testing.T) {
	var err error
	c, err = NewCrypto()
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func TestEncode(t *testing.T) {
	var err error

	plain_text = []byte("Hello!")
	enc, err = c.Encrypt(plain_text, label)
	if err != nil {
		t.Error("Expected no error encrypting, got ", err)
	}

	if bytes.Equal(enc, plain_text) {
		t.Error("Failed to encrypt ", enc)
	}
}

func TestDecrypt(t *testing.T) {
	var err error
	dec, err = c.Decrypt(enc, label)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	if !bytes.Equal(dec, plain_text) {
		t.Error("Failed to decrypt", dec)
	}
}
