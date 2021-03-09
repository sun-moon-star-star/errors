package caller

type RequestConfig struct {
	ServiceName string // naming service

	IP   string
	Port uint16

	Timeout uint16 // millisecond

	Protocol string
}

type RetryConfig struct {
	RetryTimes    int16 // -1 means until success
	RetryInterval int32 // millisecond
}

type Callback interface {
	Success(*Caller) error
	Failure(*Caller) error
}

type Protocol interface {
	Call(*Caller) error
}

type Caller struct {
	Req  []byte
	Resp []byte

	RequestConfig RequestConfig
	RetryConfig   RetryConfig

	TryTimes int16
	Callback Callback

	Protocol    Protocol
	PrivateData interface{}
}

func (c *Caller) Call() error {
	if c.Protocol == nil {
		return DefaultCaller().Call(c)
	}
	return c.Protocol.Call(c)
}

func DefaultCaller() Protocol {
	return &Http{}
}
