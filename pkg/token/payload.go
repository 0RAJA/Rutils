package token

import (
	"time"

	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/google/uuid"
)

type Payload struct {
	//用于管理每个token
	ID       uuid.UUID
	UserName string `json:"user_name,omitempty"`
	//创建时间用于检验
	IssuedAt  time.Time `json:"issued-at"`
	ExpiredAt time.Time `json:"expired-at"`
}

func NewPayload(userName string, expireDate time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:        tokenID,
		UserName:  userName,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(expireDate),
	}, nil
}

func (p *Payload) Valid() error {
	if p.ExpiredAt.After(time.Now()) {
		return errcode.UnauthorizedTokenTimeoutErr
	}
	return nil
}
