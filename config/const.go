package config

/*
 * セッション関連
 */
const (
	SESSION_NAME           = "session-xseadfef"
	SESSION_MAX_AGE        = 7200
	SESSION_PATH           = "/"
	SESSION_MAX_CONNECTION = 10
)

/*
 * DB接続
 */
const (
	DBTYPE               = "mysql"
	USER                 = "root"
	PASSWORD             = "password"
	HOSTNAME             = "localhost"
	HOSTPORT             = "3306"
	DBNAME               = "golang_sample"
	DBMAXIDOLECONNECTION = 10
	DBMAXCONNECTION      = 15
)

/*
 * Memcachedのホスト名
 */
const (
	MEMCACHEDHOSTNAME     = "test-memcached.ccc1kl.cfg.apne1.cache.amazonaws.com:11211"
	MEMCACHEMAXCONNECTION = 10
)

/*
 * Redisのホスト名
 */
const (
	REDISHOSTNAME      = "localhost:6379"
	REDISMAXCONNECTION = 10
)

/*
 * メール接続
 */
const (
	STMPSERVERADDR = "172.31.28.95:25"
)

/*
 * ボタンタイプ
 */
const (
	BUTTON_SEARCH   = "search"
	BUTTON_SHOW     = "show"
	BUTTON_REGISTER = "register"
	BUTTON_BACK     = "back"
	BUTTON_DELETE   = "delete"
)

/*
 * 結果成否
 */
const (
	RESULT_SUCCESS = 1
	RESULT_ERROR   = 2
)
