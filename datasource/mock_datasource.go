package datasource

import "starter-project/connector"

func NewMockDataSource() KeyValueDataSource {
	return connector.NewInMemoryKeyValue()
}
