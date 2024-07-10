package unsafe

type OneInterface interface {
	action()
}

type One struct {
}

func (o One) action() {
	//TODO implement me
	panic("implement me")
}

var _ OneInterface = (*One)(nil)
