package main

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/google/uuid"
)

type ErrOnlyOpenRecsCanAgglutinate struct {
	RecID     uuid.UUID
	RecStatus string
}

func NewErrOnlyOpenRecsCanAgglutinate(recID uuid.UUID, recStatus string) *ErrOnlyOpenRecsCanAgglutinate {
	return &ErrOnlyOpenRecsCanAgglutinate{
		RecID:     recID,
		RecStatus: recStatus,
	}
}

func (e *ErrOnlyOpenRecsCanAgglutinate) Error() string {
	return fmt.Sprintf("cannot agglutinate: only OPEN receivables can be agglutinated, receivable %s is in status %s", e.RecID.String(), e.RecStatus)
}

var (
	ErrOnlyTuitionCanAgglutinate = errors.New("cannot agglutinate: only tuition installments can be agglutinated")
)

func testZap() error {
	err := NewErrOnlyOpenRecsCanAgglutinate(uuid.New(), "canceled")

	_ = NewErrOnlyOpenRecsCanAgglutinate(uuid.Nil, "")

	return err
}

func testandoEncode() {
	guardianIDS := []string{uuid.NewString(), uuid.NewString(), uuid.NewString()}

	params := make(url.Values)

	var guardiamnIds string
	for _, id := range guardianIDS {
		guardiamnIds += "," + id
	}

	params.Add("guardian_id", guardiamnIds)

	fmt.Println(params.Encode())
}

func main() {
	testandoEncode()
}
