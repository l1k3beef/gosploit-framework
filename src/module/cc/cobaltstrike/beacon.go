package cobaltstrike

type IBeacon interface {
}

type Beacon struct {
}

func (cc *Beacon) Generate() []byte {
	return nil
}
