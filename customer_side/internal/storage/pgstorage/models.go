package PGStorage

type GGOrderInfo struct {
	ID           uint64 `db:"id"`
	CustomerName string `db:"customer_name"`
	Email        string `db:"email"`
	Details      string `db:"details"`
}

const (
	tableName              = "ggorderInfo"
	IDСolumnName           = "id"
	CustomerNameСolumnName = "customer_name"
	EmailСolumnName        = "email"
	DetailsСolumnName      = "details"
)
