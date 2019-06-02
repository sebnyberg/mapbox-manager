package writer

type WritableRow interface {
	AsRow() string
}
