package repository

type File interface {
	Get()
	GetAll()
	Update()
	Create()
}

type Repository struct {
	File
}
