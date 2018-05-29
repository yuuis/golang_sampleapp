package config

/*
 * DB設定
 */
const (
	DEBUG = 1 //システムの動作状況に関する詳細な情報
	INFO  = 2 //実行時の何らかの注目すべき事象（開始や終了など）。メッセージ内容は簡潔に止める
	WARN  = 3 //廃要素となったAPIの使用、APIの不適切な使用、エラーに近い事象など。実行時に生じた異常とは言い切れないが正常とも異なる何らかの予期しない問題
	ERROR = 4 //予期しないその他の実行時エラー。コンソール等に即時出力することを想定
	FATAL = 5 //プログラムの異常終了を伴うようなもの。メールなどで通知することを想定

	DEBUGNAME = "DEBUG"
	INFONAME  = "INFO"
	WARNNAME  = "WARN"
	ERRORNAME = "ERROR"
	FATALNAME = "FATAL"
)
