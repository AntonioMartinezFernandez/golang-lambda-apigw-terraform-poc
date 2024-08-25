package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

const UlidRegex = "[0-7][0-9A-HJKMNP-TV-Z]{25}"

type Ulid string

func (u Ulid) String() string {
	return string(u)
}

func NewUlid() Ulid {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return Ulid(ulid.MustNew(ulid.Timestamp(t), entropy).String())
}

func guardULID(rawUlid string) error {
	_, err := ulid.Parse(rawUlid)
	return err
}

func UlidFromString(rawUlid string) (Ulid, error) {
	var ulid Ulid
	if err := guardULID(rawUlid); err != nil {
		return ulid, err
	}

	return Ulid(rawUlid), nil
}

func IsValidUlid(rawUlid string) bool {
	return guardULID(rawUlid) == nil
}

type UlidProvider interface {
	New() Ulid
}

type RandomUlidProvider struct {
}

func NewRandomUlidProvider() *RandomUlidProvider {
	return &RandomUlidProvider{}
}

func (up RandomUlidProvider) New() Ulid {
	ulid := NewUlid()

	return ulid
}

type FixedUlidProvider struct {
	ulid Ulid
}

func NewFixedUlidProvider() *FixedUlidProvider {
	return &FixedUlidProvider{}
}

func (up *FixedUlidProvider) New() Ulid {
	if up.ulid == "" {
		up.ulid = NewUlid()
	}

	return up.ulid
}
