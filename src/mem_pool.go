package src

//	singleton pattern
var mempool MemPool

type MemPool interface {
	GetTrxs() []*trx
	AddTrx(t *trx)
	Clear()
}

type memPool struct {
	CurrentTrxs []*trx
}

//	return mempool instance
func GetMemPool() MemPool {
	if mempool == nil {
		mempool = &memPool{CurrentTrxs: make([]*trx, 0)}
	}
	return mempool
}

//	Delete all trxs in mempool
func (m *memPool) Clear() {
	m.CurrentTrxs = make([]*trx, 0)
}

//	return current trxs in mempool
func (m *memPool) GetTrxs() []*trx {
	return m.CurrentTrxs
}

//	Add trx to the mempool
func (m *memPool) AddTrx(t *trx) {
	m.CurrentTrxs = append(m.CurrentTrxs, t)
}
