package requirements

type RequiredDatabase interface {
	RequiredUsers
	RequiredTrasactions
}
