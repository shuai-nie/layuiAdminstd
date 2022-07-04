package model

type Model struct {
	ID uint32
	CreatedBy string
	ModifiedBy string
	CreatedOn uint32
	ModifiedOn uint32
	DeletedOn uint32
	IsDel uint8
}
