package event

import (
	"github.com/bytedance/sonic"
	"time"
)

type Base struct {
	Id         int       `json:"id"`
	Source     string    `json:"source"`
	Data       string    `json:"data"`
	CreateTime time.Time `json:"create_time"`
}

func NewBasePublisher() PublisherI {
	return new(Base)
}

func (b *Base) Publish() error {
	// send to mq
	bt, err := sonic.Marshal(b)
	if err != nil {
		return err
	}
	println(bt)

	return nil
}
