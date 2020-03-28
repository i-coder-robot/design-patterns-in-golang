package Factory


type (
	mongoDB struct {
		Database map[string]string
	}
	sqlite struct {
		Database map[string]string
	}

	//以下是抽象工厂模式

	file struct {
		name    string
		content string
	}
	ntfs struct {
		Files map[string]file
	}
	ext4 struct {
		Files map[string]file
	}

	Database interface {
		GetData(string) string
		PutData(string, string)
	}

	FileSystem interface {
		CreateFile(string)
		FindFile(string) file
	}

	Factory func(string) interface{}
)

func (mdb mongoDB) GetData(query string) string {
	result, ok := mdb.Database[query]
	if !ok {
		return ""
	}
	return result
}

func (sql sqlite) GetData(query string) string {
	result, ok := sql.Database[query]
	if !ok {
		return ""
	}
	return result
}

func (mdb mongoDB) PutData(query, data string) {
	mdb.Database[query] = data
}

func (sql sqlite) PutData(query, data string) {
	sql.Database[query] = data
}

func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return mongoDB{Database: make(map[string]string)}
	case "development":
		return sqlite{Database: make(map[string]string)}
	default:
		return nil
	}
}

func (ntfs ntfs) CreateFile(path string) {
	f := file{
		name:    path,
		content: "ntfs file",
	}
	ntfs.Files[path] = f
}
func (ext ext4) CreateFile(path string) {
	f := file{
		name:    path,
		content: "EXT4 file",
	}
	ext.Files[path] = f
}

func (ntfs ntfs) FindFile(path string) file {
	f, ok := ntfs.Files[path]
	if !ok {
		return file{}
	}
	return f
}

func (ext ext4) FindFile(path string) file {
	f, ok := ext.Files[path]
	if !ok {
		return file{}
	}
	return f
}

func FilesystemFactory(env string) FileSystem {
	switch env {
	case "production":
		return ntfs{Files: make(map[string]file)}
	case "development":
		return ext4{Files: map[string]file{}}
	default:
		return nil
	}
}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FilesystemFactory
	default:
		return nil

	}
}
