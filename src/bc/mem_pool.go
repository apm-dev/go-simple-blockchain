package bc

type MemPool interface {
	getTrxs() []*trx
	addTrx(t *trx)
	clear()
}

type memPool struct {
	CurrentTrxs []*trx
}

//	return mempool instance
func newMemPool() MemPool {
	return &memPool{CurrentTrxs: make([]*trx, 0)}
}

//	Delete all trxs in mempool
func (m *memPool) clear() {
	m.CurrentTrxs = make([]*trx, 0)
}

//	return current trxs in mempool
func (m *memPool) getTrxs() []*trx {
	return m.CurrentTrxs
}

//	Add trx to the mempool
func (m *memPool) addTrx(t *trx) {
	m.CurrentTrxs = append(m.CurrentTrxs, t)
}
