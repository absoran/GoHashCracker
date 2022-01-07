package shared

var Config = configuration{
	POSTGRESURL: "postgres://yigwmsvo:MPFlYrQLOm9dBw0x07q38oQMTadXtHbN@kashin.db.elephantsql.com/yigwmsvo",
}

type configuration struct {
	POSTGRESURL string
}
