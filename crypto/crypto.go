package crypto

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"hash"
	"io/ioutil"
)

type Crypto struct {
	private_key *rsa.PrivateKey
	public_key  *rsa.PublicKey
}

func NewCrypto() (Crypto, error) {
	var c Crypto
	var err error
	var block *pem.Block
	var private_key *rsa.PrivateKey
	var public_key *rsa.PublicKey
	var pem_data []byte
	pem_file_path := "../keys/id_rsa"
	if pem_data, err = ioutil.ReadFile(pem_file_path); err != nil {
		return c, err
	}

	if block, _ = pem.Decode(pem_data); block == nil || block.Type != "RSA PRIVATE KEY" {
		return c, err
	}

	if private_key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return c, err
	}

	public_key = &private_key.PublicKey
	c = Crypto{private_key: private_key, public_key: public_key}
	return c, err
}

func (c Crypto) Encrypt(plain_text, label []byte) (encrypted []byte, e error) {
	var err error
	var md5_hash hash.Hash

	md5_hash = md5.New()
	if encrypted, err = rsa.EncryptOAEP(md5_hash, rand.Reader, c.public_key, plain_text, label); err != nil {
		e = err
	}
	return
}

func (c Crypto) Decrypt(cipher_text, label []byte) (decrypted []byte, e error) {
	var md5_hash hash.Hash

	md5_hash = md5.New()
	decrypted, e = rsa.DecryptOAEP(md5_hash, rand.Reader, c.private_key, cipher_text, label)
	return
}
