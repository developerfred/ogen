package bls_test

import (
	"encoding/hex"
	"github.com/olympus-protocol/ogen/pkg/bls"
	testdata "github.com/olympus-protocol/ogen/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicKey_Cache(t *testing.T) {
	pub := bls.RandKey().PublicKey()
	pubBytes := pub.Marshal()

	pub1, err := bls.PublicKeyFromBytes(pubBytes)
	assert.NoError(t, err)

	pub2, err := bls.PublicKeyFromBytes(pubBytes)
	assert.NoError(t, err)

	assert.Equal(t, pub1, pub2)
}

func TestPublicKey_Hash(t *testing.T) {
	pubBytes, err := hex.DecodeString("8817994e67c131ed73d6ff851013be76322dccdeb78755bb341768b7e76ccd9a7982bcfd19d923a2c1f1556a45163695")
	assert.NoError(t, err)

	pub, err := bls.PublicKeyFromBytes(pubBytes)
	assert.NoError(t, err)

	hash, err := pub.Hash()
	assert.NoError(t, err)
	assert.Equal(t, "a47fad160040ba7b1b54b35dc74b1993664b522c", hex.EncodeToString(hash[:]))
}

func TestPublicKey_ToAccount(t *testing.T) {
	bls.Initialize(testdata.TestParams)

	pubBytes, err := hex.DecodeString("8817994e67c131ed73d6ff851013be76322dccdeb78755bb341768b7e76ccd9a7982bcfd19d923a2c1f1556a45163695")
	assert.NoError(t, err)

	pub, err := bls.PublicKeyFromBytes(pubBytes)
	assert.NoError(t, err)

	acc := pub.ToAccount()
	assert.Equal(t, "itpub153l669sqgza8kx65kdwuwjcejdnyk53v3nq49a", acc)
}

func TestPublicKeyFromBytes(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		err   error
	}{
		{
			name: "Nil",
			err:  bls.ErrorPubSize,
		},
		{
			name:  "Empty",
			input: []byte{},
			err:   bls.ErrorPubSize,
		},
		{
			name:  "Short",
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			err:   bls.ErrorPubSize,
		},
		{
			name:  "Long",
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			err:   bls.ErrorPubSize,
		},
		{
			name:  "Bad",
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			err:   bls.ErrorPubKeyUnmarshal,
		},
		{
			name:  "Good",
			input: []byte{0xa9, 0x9a, 0x76, 0xed, 0x77, 0x96, 0xf7, 0xbe, 0x22, 0xd5, 0xb7, 0xe8, 0x5d, 0xee, 0xb7, 0xc5, 0x67, 0x7e, 0x88, 0xe5, 0x11, 0xe0, 0xb3, 0x37, 0x61, 0x8f, 0x8c, 0x4e, 0xb6, 0x13, 0x49, 0xb4, 0xbf, 0x2d, 0x15, 0x3f, 0x64, 0x9f, 0x7b, 0x53, 0x35, 0x9f, 0xe8, 0xb9, 0x4a, 0x38, 0xe4, 0x4c},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := bls.PublicKeyFromBytes(test.input)
			if test.err != nil {
				assert.Equal(t, test.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.input, res.Marshal())
			}
		})
	}
}

func TestPublicKey_Copy(t *testing.T) {

	pubkeyA := bls.RandKey().PublicKey()
	pubkeyBytes := pubkeyA.Marshal()

	pubkeyB := pubkeyA.Copy()
	pubkeyB.Aggregate(bls.RandKey().PublicKey())

	assert.Equal(t, pubkeyA.Marshal(), pubkeyBytes, "Pubkey was mutated after copy")
}
