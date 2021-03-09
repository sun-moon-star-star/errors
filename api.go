package errors

import "error"

type RetryConfig struct{
	Req []byte
	Resp []byte

	RetryTimes int16 // -1 means until success
	RetryInterval int32  // millisecond
	
	IP string
	Port uint16
	Protocol string
}

type API struct {
	Record ([]byte) error // Record error
	Retry(*RetryConfig) error // 
}
