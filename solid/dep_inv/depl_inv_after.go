package dep_inv

import "fmt"

type MySQL struct{}

func (db MySQL) Query() interface{} {
	return []string{"mozilla", "edge", "chrome"}
}

type PostgreSQL struct{}

func (db PostgreSQL) Query() interface{} {
	return map[string]string{"a": "A", "b": "B", "c": "C"}
}

type DBConn interface {
	Query() interface{}
}

type UserRepository struct {
	db DBConn
}

func (r UserRepository) GetUsers() []string {
	var users []string
	res := r.db.Query()
	switch res.(type) {
	case map[string]string:
		for _, u := range res.(map[string]string) {
			users = append(users, u)
		}
		return users
	case []string:
		return res.([]string)
	}
	return []string{}
}

func DepInv() {

	mysqlDB := MySQL{}
	pgsqlDB := PostgreSQL{}

	repo1 := UserRepository{db: mysqlDB}
	repo2 := UserRepository{db: pgsqlDB}

	fmt.Println("mysql db : ", repo1.GetUsers())
	fmt.Println("pgsql db : ", repo2.GetUsers())
}
