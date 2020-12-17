package KyleDB

import(
	"github.com/spf13/afero"
)

type DB struct {
	Path string
}

var kyledb = afero.NewMemMapFs()

func Open(path string) (DB,error) {
	isExist,err := afero.DirExists(kyledb,path)
	if err != nil {
		return DB{},err
	}
	if isExist == false {
		_ = kyledb.Mkdir(path,0700)
	}
	return DB{Path : path},nil
}

func (Db DB) Get(key string) ([]byte,error) {
	value,err := afero.ReadFile(kyledb,Db.Path + "/" + key)
	return value, err
}

func (Db DB) Put(key string,value []byte) error {
	err := afero.WriteFile(kyledb,Db.Path + "/" + key,value,700)
	return err
}

func (Db DB) Keys() ([]string,error) {
	DBKeys, err := kyledb.Open(Db.Path)
	if err != nil {
		return []string{},err
	}
	keys, err2 := DBKeys.Readdirnames(0)
	return keys,err2
}
