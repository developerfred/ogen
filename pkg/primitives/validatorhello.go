package primitives

import (
	"encoding/binary"
)

const (
	// MaxValidatorHelloMessageSize is the maximum amount of bytes a CombinedSignature can contain.
	MaxValidatorHelloMessageSize = 160
)

// ValidatorHelloMessage is a message sent by validators to indicate that they are coming online.
type ValidatorHelloMessage struct {
	PublicKey [48]byte
	Timestamp uint64
	Nonce     uint64
	Signature [96]byte
}

// SignatureMessage gets the signed portion of the message.
func (v *ValidatorHelloMessage) SignatureMessage() []byte {
	timeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBytes, v.Timestamp)

	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, v.Nonce)

	var msg []byte
	msg = append(msg, v.PublicKey[:]...)
	msg = append(msg, timeBytes...)
	msg = append(msg, nonceBytes...)

	return msg
}

// Marshal serializes the hello message to the given writer.
func (v *ValidatorHelloMessage) Marshal() ([]byte, error) {
	return v.MarshalSSZ()
}

// Unmarshal deserializes the validator hello message from the reader.
func (v *ValidatorHelloMessage) Unmarshal(b []byte) error {
	return v.UnmarshalSSZ(b)
}
