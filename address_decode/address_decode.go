package address_decode

import (
	"crypto/sha512"
	"encoding/base32"
	"errors"
	"github.com/stellar/go/crc16"
)

/**
Start with 32 bytes
Add a byte of 0x30 as prefix 'G' (now you have 33 bytes)
Calculate the checksum (two bytes)
Add the checksum as suffix (now you have 35 bytes)
Convert them to base32
That's your public key
Apply the same but using 'S' (byte 0x90) as prefix for secret keys
*/

// DigestSize is the number of bytes in the preferred hash Digest used here.
const DigestSize = sha512.Size256
const PublicKeySize = 32
const ChecksumLength = 2

// Digest represents a 32-byte value holding the 256-bit Hash digest.
type Digest [DigestSize]byte
type PublicKey [PublicKeySize]byte

type (
	Address Digest
)

type ChecksumAddress struct {
	shortAddress Address
	checksum     []byte
}

// AddressDecoderV2
type AddressDecoderV2 struct {
}

var (
	Default = AddressDecoderV2{}
)

var (
	ErrorInvalidHashLength = errors.New("Invalid hash length!")
	ErrorInvalidAddress    = errors.New("Invalid address!")
)

//AddressEncode encode address bytes
func (dec *AddressDecoderV2) AddressEncode(address []byte) (string, error) {
	var pk PublicKey

	if len(pk) != len(address) {
		return "", ErrorInvalidHashLength
	}

	for i := range pk {
		pk[i] = address[i]
	}
	publicKeyChecksummed := Address(pk).GetChecksumAddress().String()
	return publicKeyChecksummed, nil
}

var prefix = []byte{0x30}

// GetChecksumAddress returns the short address with its checksum as a string
// Checksum in  traim are the last 2  bytes of the checksum
func (addr Address) GetChecksumAddress() *ChecksumAddress {
	rawData := append(prefix, addr[:]...)
	checkSum := checkSum(rawData)
	return &ChecksumAddress{addr, checkSum[len(checkSum)-ChecksumLength:]}
}

// String returns a string representation of ChecksumAddress
func (addr *ChecksumAddress) String() string {
	var addrWithChecksum []byte
	addrWithChecksum = append(prefix, addr.shortAddress[:]...)
	addrWithChecksum = append(addrWithChecksum, addr.checksum...)
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(addrWithChecksum)
}

//
func checkSum(data []byte) []byte {
	return crc16.Checksum(data)
}
