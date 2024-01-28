package PackageInitialization

var connection string

func init() {
	connection = "MYSQL"
}

func GetDatabase() string {
	return connection
}
