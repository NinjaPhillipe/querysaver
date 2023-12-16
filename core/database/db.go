package database

func Add(a, b int) int {
	return a + b
}

type sqlDatabase interface {
	Connect()
	Close()
}

// make a new DbCOnnection
// func NewDbConnection(db sqlDatabase) *DbConnection {
// 	res := &DbConnection{db}
// 	res.db.Connect()
// 	return res
// }
