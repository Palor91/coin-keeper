package requirements

type RequiredDatabase interface {
	RequiredUsers
	RequiredTransactions
}
