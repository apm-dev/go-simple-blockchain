package bc

//	singleton pattern
var mempool MemPool

type MemPool interface {
	GetTrxs() []*Trx
	AddTrx(t *Trx)
	Clear()
}

type memPool struct {
	CurrentTrxs []*Trx
}

//	return mempool instance
func GetMemPool() MemPool {
	if mempool == nil {
		mempool = &memPool{CurrentTrxs: make([]*Trx, 0)}
	}
	return mempool
}

//	Delete all trxs in mempool
func (m *memPool) Clear() {
	m.CurrentTrxs = make([]*Trx, 0)
}

//	return current trxs in mempool
func (m *memPool) GetTrxs() []*Trx {
	return m.CurrentTrxs
}

//	Add Trx to the mempool
func (m *memPool) AddTrx(t *Trx) {
	m.CurrentTrxs = append(m.CurrentTrxs, t)
}
