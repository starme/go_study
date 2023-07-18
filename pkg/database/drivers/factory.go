package drivers

func Factory(c Config) Connector {
	switch c.GetDriverName() {
	case "mysql":
		return makeMysqlConnector(c)
	//case "sqlite":
	//	return makeSqliteConnector(c)
	//case "postgres":
	//	return makePostgresConnector(c)
	//case "sqlserver":
	//	return makeSqlServerConnector(c)
	//case "tidb":
	//	return makeTidbConnector(c)
	default:
		panic("unknown driver")
	}
}
