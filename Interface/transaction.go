package Interface

type Transaction interface {
	BeginTransaction() (transactionId int)
	Commit(transactionId int) error
	Rollback(transactionId int) error
}
