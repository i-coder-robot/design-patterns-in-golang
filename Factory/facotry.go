package Factory

type (
	MongoDB struct {
		Database map[string]string
	}
	Sqlite struct {
		Database map[string]string
	}

	Database interface {
		GetData(string) string
		PutData(string, string)
	}
)

func (mdb *MongoDB) GetData(query string) string {
	result, ok := mdb.Database[query]
	if !ok {
		return ""
	}
	return result
}

func (sql *Sqlite) GetData(query string) string {
	result, ok := sql.Database[query]
	if !ok {
		return ""
	}
	return result
}

func (mdb *MongoDB) PutData(query, data string) {
	mdb.Database[query] = data
}

func (sql *Sqlite) PutData(query, data string) {
	sql.Database[query] = data
}

func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return &MongoDB{Database: make(map[string]string)}
	case "development":
		return &Sqlite{Database: make(map[string]string)}
	default:
		return nil
	}
}
