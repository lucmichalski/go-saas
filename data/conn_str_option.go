package data

type ConnStrings map[string]string

const Default = "Default"

type ConnStrOption struct {
	//connection string map
	Conn ConnStrings
}

func NewConnStrOption() *ConnStrOption {
	return &ConnStrOption{
		Conn: make(ConnStrings),
	}
}

func (c ConnStrings) Default() string {
	return c[Default]
}

func (c ConnStrings) GetOrDefault(k string) string {
	ret := c[k]
	if ret == "" {
		return c.Default()
	}
	return ret
}
func (c ConnStrings) SetDefault(value string) {
	c[Default] = value
}
