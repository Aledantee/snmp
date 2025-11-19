package snmp

type Value interface {
	Tag() TypeTag

	sealed()
}

type Opaque []byte

func (o Opaque) sealed() {}
