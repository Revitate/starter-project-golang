package datasource

type EntityDataSource interface {

}

type KeyValueDataSource interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}