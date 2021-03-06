package broker

import (
	"context"
	"encoding/json"
)

type PublishOptions struct {
	ExchangeName     string
	ExchangeType     string
	ExchangeAD       bool // exchange auto delete
	ExchangeDuration bool // exchange duration
	Context          context.Context
	QOS              int
	Retained         bool
	// kafka
	Part    int
	Replica int
}

func (p *PublishOptions) String() string {
	return p.Marshal()
}
func (p *PublishOptions) Marshal() string {
	bs, err := json.Marshal(p)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}

type SubscribeOptions struct {
	AutoAck          bool
	AutoDel          bool
	Duration         bool
	Queue            string
	ExchangeName     string
	ExchangeType     string
	ExchangeAD       bool // exchange auto delete
	ExchangeDuration bool // exchange duration
	QOS              int
	Context          context.Context
	// kafka
	Part    int
	Replica int
}

func (s *SubscribeOptions) String() string {
	return s.Marshal()
}
func (s *SubscribeOptions) Marshal() string {
	bs, err := json.Marshal(s)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}

type SubscribeOption func(*SubscribeOptions)

type PublishOption func(*PublishOptions)

// Set SubscribeOption
// SetSubPart
func SetSubPart(p int) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Part = p
	}
}

// SetSubReplica
func SetSubReplica(r int) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Replica = r
	}
}

// SetSubDuration
func SetSubDuration(duration bool) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Duration = duration
	}
}

// SetSubQOS
func SetSubQOS(qos int) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.QOS = qos
	}
}

// SetSubAutoAck
func SetSubAutoAck(autoAck bool) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.AutoAck = autoAck
	}
}

// SetSubAutoDel
func SetSubAutoDel(autoDel bool) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.AutoDel = autoDel
	}
}

// SetSubQueue
func SetSubQueue(queue string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Queue = queue
	}
}

// SetSubContext
func SetSubContext(cxt context.Context) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Context = cxt
	}
}

// SetSubExchangeType
func SetSubExchangeType(et string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.ExchangeType = et
	}
}

// SetSubExchangeName
func SetSubExchangeName(en string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.ExchangeName = en
	}
}

// SetSubExchangeAD
func SetSubExchangeAD(ad bool) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.ExchangeAD = ad
	}
}

// SetSubExchangeDuration
func SetSubExchangeDuration(duration bool) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.ExchangeDuration = duration
	}
}

// Set PublishOption
// SetPubRetained
func SetPubRetained(r bool) PublishOption {
	return func(o *PublishOptions) {
		o.Retained = r
	}
}

// Set PublishOption
// SetPubQOS
func SetPubQOS(qos int) PublishOption {
	return func(o *PublishOptions) {
		o.QOS = qos
	}
}

// SetPubContext
func SetPubContext(cxt context.Context) PublishOption {
	return func(o *PublishOptions) {
		o.Context = cxt
	}
}

// SetPubExchangeName
func SetPubExchangeName(en string) PublishOption {
	return func(o *PublishOptions) {
		o.ExchangeName = en
	}
}

// SetPubExchangeType
func SetPubExchangeType(et string) PublishOption {
	return func(o *PublishOptions) {
		o.ExchangeType = et
	}
}

// SetPubExchangeAD
func SetPubExchangeAD(ad bool) PublishOption {
	return func(o *PublishOptions) {
		o.ExchangeAD = ad
	}
}

// SetPubExchangeDuration
func SetPubExchangeDuration(duration bool) PublishOption {
	return func(o *PublishOptions) {
		o.ExchangeDuration = duration
	}
}

// SetPubPart
func SetPubPart(p int) PublishOption {
	return func(o *PublishOptions) {
		o.Part = p
	}
}

// SetPubReplica
func SetPubReplica(r int) PublishOption {
	return func(o *PublishOptions) {
		o.Replica = r
	}
}
