// Code generated by fastssz. DO NOT EDIT.
package primitives

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the CommunityVoteDataInfo object
func (c *CommunityVoteDataInfo) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the CommunityVoteDataInfo object to a target array
func (c *CommunityVoteDataInfo) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(36)

	// Field (0) 'Hash'
	dst = append(dst, c.Hash[:]...)

	// Offset (1) 'Data'
	dst = ssz.WriteOffset(dst, offset)
	if c.Data == nil {
		c.Data = new(CommunityVoteData)
	}
	offset += c.Data.SizeSSZ()

	// Field (1) 'Data'
	if dst, err = c.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the CommunityVoteDataInfo object
func (c *CommunityVoteDataInfo) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 36 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Hash'
	copy(c.Hash[:], buf[0:32])

	// Offset (1) 'Data'
	if o1 = ssz.ReadOffset(buf[32:36]); o1 > size {
		return ssz.ErrOffset
	}

	// Field (1) 'Data'
	{
		buf = tail[o1:]
		if c.Data == nil {
			c.Data = new(CommunityVoteData)
		}
		if err = c.Data.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the CommunityVoteDataInfo object
func (c *CommunityVoteDataInfo) SizeSSZ() (size int) {
	size = 36

	// Field (1) 'Data'
	if c.Data == nil {
		c.Data = new(CommunityVoteData)
	}
	size += c.Data.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the CommunityVoteDataInfo object
func (c *CommunityVoteDataInfo) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the CommunityVoteDataInfo object with a hasher
func (c *CommunityVoteDataInfo) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Hash'
	hh.PutBytes(c.Hash[:])

	// Field (1) 'Data'
	if err = c.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the ReplacementVotes object
func (r *ReplacementVotes) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(r)
}

// MarshalSSZTo ssz marshals the ReplacementVotes object to a target array
func (r *ReplacementVotes) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Account'
	dst = append(dst, r.Account[:]...)

	// Field (1) 'Hash'
	dst = append(dst, r.Hash[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ReplacementVotes object
func (r *ReplacementVotes) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 52 {
		return ssz.ErrSize
	}

	// Field (0) 'Account'
	copy(r.Account[:], buf[0:20])

	// Field (1) 'Hash'
	copy(r.Hash[:], buf[20:52])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ReplacementVotes object
func (r *ReplacementVotes) SizeSSZ() (size int) {
	size = 52
	return
}

// HashTreeRoot ssz hashes the ReplacementVotes object
func (r *ReplacementVotes) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(r)
}

// HashTreeRootWith ssz hashes the ReplacementVotes object with a hasher
func (r *ReplacementVotes) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Account'
	hh.PutBytes(r.Account[:])

	// Field (1) 'Hash'
	hh.PutBytes(r.Hash[:])

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the GovernanceSerializable object
func (g *GovernanceSerializable) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(g)
}

// MarshalSSZTo ssz marshals the GovernanceSerializable object to a target array
func (g *GovernanceSerializable) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'ReplaceVotes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(g.ReplaceVotes) * 52

	// Offset (1) 'CommunityVotes'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(g.CommunityVotes); ii++ {
		offset += 4
		offset += g.CommunityVotes[ii].SizeSSZ()
	}

	// Field (0) 'ReplaceVotes'
	if len(g.ReplaceVotes) > 2097152 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(g.ReplaceVotes); ii++ {
		if dst, err = g.ReplaceVotes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (1) 'CommunityVotes'
	if len(g.CommunityVotes) > 2097152 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(g.CommunityVotes)
		for ii := 0; ii < len(g.CommunityVotes); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += g.CommunityVotes[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(g.CommunityVotes); ii++ {
		if dst, err = g.CommunityVotes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the GovernanceSerializable object
func (g *GovernanceSerializable) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'ReplaceVotes'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Offset (1) 'CommunityVotes'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'ReplaceVotes'
	{
		buf = tail[o0:o1]
		num, err := ssz.DivideInt2(len(buf), 52, 2097152)
		if err != nil {
			return err
		}
		g.ReplaceVotes = make([]*ReplacementVotes, num)
		for ii := 0; ii < num; ii++ {
			if g.ReplaceVotes[ii] == nil {
				g.ReplaceVotes[ii] = new(ReplacementVotes)
			}
			if err = g.ReplaceVotes[ii].UnmarshalSSZ(buf[ii*52 : (ii+1)*52]); err != nil {
				return err
			}
		}
	}

	// Field (1) 'CommunityVotes'
	{
		buf = tail[o1:]
		num, err := ssz.DecodeDynamicLength(buf, 2097152)
		if err != nil {
			return err
		}
		g.CommunityVotes = make([]*CommunityVoteDataInfo, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if g.CommunityVotes[indx] == nil {
				g.CommunityVotes[indx] = new(CommunityVoteDataInfo)
			}
			if err = g.CommunityVotes[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the GovernanceSerializable object
func (g *GovernanceSerializable) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'ReplaceVotes'
	size += len(g.ReplaceVotes) * 52

	// Field (1) 'CommunityVotes'
	for ii := 0; ii < len(g.CommunityVotes); ii++ {
		size += 4
		size += g.CommunityVotes[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the GovernanceSerializable object
func (g *GovernanceSerializable) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(g)
}

// HashTreeRootWith ssz hashes the GovernanceSerializable object with a hasher
func (g *GovernanceSerializable) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ReplaceVotes'
	{
		subIndx := hh.Index()
		num := uint64(len(g.ReplaceVotes))
		if num > 2097152 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = g.ReplaceVotes[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 2097152)
	}

	// Field (1) 'CommunityVotes'
	{
		subIndx := hh.Index()
		num := uint64(len(g.CommunityVotes))
		if num > 2097152 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for i := uint64(0); i < num; i++ {
			if err = g.CommunityVotes[i].HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 2097152)
	}

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the CommunityVoteData object
func (c *CommunityVoteData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the CommunityVoteData object to a target array
func (c *CommunityVoteData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'ReplacementCandidates'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.ReplacementCandidates) * 20

	// Field (0) 'ReplacementCandidates'
	if len(c.ReplacementCandidates) > 5 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(c.ReplacementCandidates); ii++ {
		dst = append(dst, c.ReplacementCandidates[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the CommunityVoteData object
func (c *CommunityVoteData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'ReplacementCandidates'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (0) 'ReplacementCandidates'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 20, 5)
		if err != nil {
			return err
		}
		c.ReplacementCandidates = make([][20]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(c.ReplacementCandidates[ii][:], buf[ii*20:(ii+1)*20])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the CommunityVoteData object
func (c *CommunityVoteData) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'ReplacementCandidates'
	size += len(c.ReplacementCandidates) * 20

	return
}

// HashTreeRoot ssz hashes the CommunityVoteData object
func (c *CommunityVoteData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the CommunityVoteData object with a hasher
func (c *CommunityVoteData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ReplacementCandidates'
	{
		if len(c.ReplacementCandidates) > 5 {
			err = ssz.ErrListTooBig
			return
		}
		subIndx := hh.Index()
		for _, i := range c.ReplacementCandidates {
			hh.Append(i[:])
		}
		numItems := uint64(len(c.ReplacementCandidates))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(5, numItems, 32))
	}

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the GovernanceVote object
func (g *GovernanceVote) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(g)
}

// MarshalSSZTo ssz marshals the GovernanceVote object to a target array
func (g *GovernanceVote) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Type'
	dst = ssz.MarshalUint64(dst, g.Type)

	// Field (1) 'Data'
	dst = append(dst, g.Data[:]...)

	// Field (2) 'FunctionalSig'
	dst = append(dst, g.FunctionalSig[:]...)

	// Field (3) 'VoteEpoch'
	dst = ssz.MarshalUint64(dst, g.VoteEpoch)

	return
}

// UnmarshalSSZ ssz unmarshals the GovernanceVote object
func (g *GovernanceVote) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 260 {
		return ssz.ErrSize
	}

	// Field (0) 'Type'
	g.Type = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Data'
	copy(g.Data[:], buf[8:108])

	// Field (2) 'FunctionalSig'
	copy(g.FunctionalSig[:], buf[108:252])

	// Field (3) 'VoteEpoch'
	g.VoteEpoch = ssz.UnmarshallUint64(buf[252:260])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the GovernanceVote object
func (g *GovernanceVote) SizeSSZ() (size int) {
	size = 260
	return
}

// HashTreeRoot ssz hashes the GovernanceVote object
func (g *GovernanceVote) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(g)
}

// HashTreeRootWith ssz hashes the GovernanceVote object with a hasher
func (g *GovernanceVote) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Type'
	hh.PutUint64(g.Type)

	// Field (1) 'Data'
	hh.PutBytes(g.Data[:])

	// Field (2) 'FunctionalSig'
	hh.PutBytes(g.FunctionalSig[:])

	// Field (3) 'VoteEpoch'
	hh.PutUint64(g.VoteEpoch)

	hh.Merkleize(indx)
	return
}