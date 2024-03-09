package roc

type addrImpl struct {
	addrMap map[uint64]string
}

func NewAddrImpl() *addrImpl {
	return &addrImpl{
		addrMap: make(map[uint64]string),
	}
}

func (this *addrImpl) UserAddr(uid uint64) string {
	return "127.0.0.1:7890"
}

func (this *addrImpl) AddUserAddr(uid uint64, addr string) {
	this.addrMap[uid] = addr
}

func (this *addrImpl) RemoveUserAddr(uid uint64) {
	delete(this.addrMap, uid)
}

func (this *addrImpl) RedisAddr() string {
	return "127.0.0.1:6379"
}
