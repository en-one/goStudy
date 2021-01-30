package server

//远程服务对象
type MathService struct {
}

//参数
type Args struct {
	A, B int
}

func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
