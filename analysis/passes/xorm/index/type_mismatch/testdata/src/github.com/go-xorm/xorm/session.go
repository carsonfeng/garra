package xorm

// Session keep a pointer to sql.DB and provides all execution of all
// kind of database operations.
type Session struct {
}

// Clone copy all the session's content and return a new session
func (session *Session) Clone() *Session {
	return nil
}

// Init reset the session as the init status.
func (session *Session) Init() {
	return
}

// Close release the connection from pool
func (session *Session) Close() {
	return
}

//// ContextCache enable context cache or not
//func (session *Session) ContextCache(context ContextCache) *Session {
//	return nil
//}

// IsClosed returns if session is closed
func (session *Session) IsClosed() bool {
	return false
}

func (session *Session) resetStatement() {
	return
}

// Prepare set a flag to session that should be prepare statement before execute query
func (session *Session) Prepare() *Session {
	return nil
}

// Before Apply before Processor, affected bean is passed to closure arg
func (session *Session) Before(closures func(interface{})) *Session {
	return nil
}

// After Apply after Processor, affected bean is passed to closure arg
func (session *Session) After(closures func(interface{})) *Session {
	return nil
}

// Table can input a string or pointer to struct for special a table to operate.
func (session *Session) Table(tableNameOrBean interface{}) *Session {
	return nil
}

// Alias set the table alias
func (session *Session) Alias(alias string) *Session {
	return nil
}

// NoCascade indicate that no cascade load child object
func (session *Session) NoCascade() *Session {
	return nil
}

// ForUpdate Set Read/Write locking for UPDATE
func (session *Session) ForUpdate() *Session {
	return nil
}

// NoAutoCondition disable generate SQL condition from beans
func (session *Session) NoAutoCondition(no ...bool) *Session {
	return nil
}

// Limit provide limit and offset query condition
func (session *Session) Limit(limit int, start ...int) *Session {
	return nil
}

// OrderBy provide order by query condition, the input parameter is the content
// after order by on a sql statement.
func (session *Session) OrderBy(order string) *Session {
	return nil
}

// Desc provide desc order by query condition, the input parameters are columns.
func (session *Session) Desc(colNames ...string) *Session {
	return nil
}

// Asc provide asc order by query condition, the input parameters are columns.
func (session *Session) Asc(colNames ...string) *Session {
	return nil
}

// StoreEngine is only avialble mysql dialect currently
func (session *Session) StoreEngine(storeEngine string) *Session {
	return nil
}

// Charset is only avialble mysql dialect currently
func (session *Session) Charset(charset string) *Session {
	return nil
}

// Cascade indicates if loading sub Struct
func (session *Session) Cascade(trueOrFalse ...bool) *Session {
	return nil
}

// NoCache ask this session do not retrieve data from cache system and
// get data from database directly.
func (session *Session) NoCache() *Session {
	return nil
}

// Join join_operator should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (session *Session) Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) *Session {
	return nil
}

// GroupBy Generate Group By statement
func (session *Session) GroupBy(keys string) *Session {
	return nil
}

// Having Generate Having statement
func (session *Session) Having(conditions string) *Session {
	return nil
}

//// DB db return the wrapper of sql.DB
//func (session *Session) DB() *core.DB {
//	return nil
//}

func cleanupProcessorsClosures(slices *[]func(interface{})) {
	return
}

func (session *Session) canCache() bool {
	return true
}

//func (session *Session) doPrepare(db *core.DB, sqlStr string) (stmt *core.Stmt, err error) {
//	return
//}

//func (session *Session) getField(dataStruct *reflect.Value, key string, table *core.Table, idx int) (*reflect.Value, error) {
//	return nil, nil
//}

// Cell cell is a result of one column field
type Cell *interface{}

//func (session *Session) rows2Beans(rows *core.Rows, fields []string,
//	table *core.Table, newElemFunc func([]string) reflect.Value,
//	sliceValueSetFunc func(*reflect.Value, core.PK) error) error {
//	return nil
//}

//func (session *Session) row2Slice(rows *core.Rows, fields []string, bean interface{}) ([]interface{}, error) {
//	return nil, nil
//}

//
//func (session *Session) slice2Bean(scanResults []interface{}, fields []string, bean interface{}, dataStruct *reflect.Value, table *core.Table) (core.PK, error) {
//
//	return pk, nil
//}

// saveLastSQL stores executed query information
func (session *Session) saveLastSQL(sql string, args ...interface{}) {

}

//// LastSQL returns last query information
//func (session *Session) LastSQL() (string, []interface{}) {
//	//return session.lastSQL, session.lastSQLArgs
//}

// Unscoped always disable struct tag "deleted"
func (session *Session) Unscoped() *Session {
	//session.statement.Unscoped()
	//return session
	return nil
}

//func (session *Session) incrVersionFieldValue(fieldValue *reflect.Value) {
//	return
//}

// Sql provides raw sql input parameter. When you have a complex SQL statement
// and cannot use Where, Id, In and etc. Methods to describe, you can use SQL.
//
// Deprecated: use SQL instead.
func (session *Session) Sql(query string, args ...interface{}) *Session {
	return nil
}

// SQL provides raw sql input parameter. When you have a complex SQL statement
// and cannot use Where, Id, In and etc. Methods to describe, you can use SQL.
func (session *Session) SQL(query interface{}, args ...interface{}) *Session {
	return nil
}

// Where provides custom query condition.
func (session *Session) Where(query interface{}, args ...interface{}) *Session {
	return nil
}

// And provides custom query condition.
func (session *Session) And(query interface{}, args ...interface{}) *Session {
	return nil
}

// Or provides custom query condition.
func (session *Session) Or(query interface{}, args ...interface{}) *Session {
	return nil
}

// Id provides converting id as a query condition
//
// Deprecated: use ID instead
func (session *Session) Id(id interface{}) *Session {
	return nil
}

// ID provides converting id as a query condition
func (session *Session) ID(id interface{}) *Session {
	return nil
}

// In provides a query string like "id in (1, 2, 3)"
func (session *Session) In(column string, args222 ...interface{}) *Session {
	return nil
}

// NotIn provides a query string like "id in (1, 2, 3)"
func (session *Session) NotIn(column string, args ...interface{}) *Session {
	return nil
}

//// Conds returns session query conditions except auto bean conditions
//func (session *Session) Conds() builder.Cond {
//	return nil
//}

// Find retrieve records from table, condiBeans's non-empty fields
// are conditions. beans could be []Struct, []*Struct, map[int64]Struct
// map[int64]*Struct
func (session *Session) Find(rowsSlicePtr interface{}, condiBean ...interface{}) error {
	return nil
}

// FindAndCount find the results and also return the counts
func (session *Session) FindAndCount(rowsSlicePtr interface{}, condiBean ...interface{}) (int64, error) {
	return 0, nil
}

func (session *Session) find(rowsSlicePtr interface{}, condiBean ...interface{}) error {
	return nil
}

//func (session *Session) noCacheFind(table *core.Table, containerValue reflect.Value, sqlStr string, args ...interface{}) error {
//	return nil
//}
//
//func convertPKToValue(table *core.Table, dst interface{}, pk core.PK) error {
//	return nil
//}

//func (session *Session) cacheFind(t reflect.Type, sqlStr string, rowsSlicePtr interface{}, args ...interface{}) (err error) {
//	return nil
//}
