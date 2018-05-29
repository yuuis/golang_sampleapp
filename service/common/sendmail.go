package common

import (
	"../../config"
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/smtp"
	"strings"
)

/*
 * メール送信
 */
func SendMail(
	mailRecepterAddr string,
	mailCCRecepterAddr string,
	mailSenderAddr string,
	mailSubject string,
	mailText string) {
	c, err := smtp.Dial(config.STMPSERVERADDR)
	if err != nil {
		WriteErrorLog(config.ERROR, err, nil)
		return
	}
	// メール送信者(送信元)の設定
	c.Mail(mailSenderAddr)
	// メール受信者(送信先)の設定
	c.Rcpt(mailRecepterAddr)
	if mailCCRecepterAddr != "" {
		// CCでのメール受信者(CC送信先)の設定
		c.Rcpt(mailCCRecepterAddr)
	}

	// メールのボディを作成
	wc, err := c.Data()
	if err != nil {
		WriteErrorLog(config.ERROR, err, nil)
		return
	}
	defer wc.Close()

	// バッファでボディの内容を構成していく
	buf := bytes.NewBufferString("")

	buf.WriteString("Content-Type: text/plain; charset=\"ISO-2022-JP\"")
	buf.WriteString("\r\n")

	// Toを指定
	buf.WriteString("To:" + mailRecepterAddr)
	buf.WriteString("\r\n") // To終わり
	// CCを指定
	if mailCCRecepterAddr != "" {
		buf.WriteString("Cc:" + mailCCRecepterAddr)
		buf.WriteString("\r\n") // CC終わり
	}

	// タイトル(件名)指定
	buf.WriteString("Subject:" + toISO2022JP(mailSubject))
	buf.WriteString("\r\n") // タイトル終わり

	// 送信元の指定
	buf.WriteString("From:" + mailSenderAddr)
	buf.WriteString("\r\n") // From終わり

	// メールヘッダ終わり
	buf.WriteString("\r\n")

	// 本文指定
	buf.WriteString(toISO2022JP(mailText))
	if _, err = buf.WriteTo(wc); err != nil {
		WriteErrorLog(config.ERROR, err, nil)
		return
	}

	// メールセッション終了
	c.Quit()
	return
}

/*
 * UTF8をISO2022JPへエンコード
 */
func toISO2022JP(s string) string {
	if msg, err := ioutil.ReadAll(
		transform.NewReader(strings.NewReader(s), japanese.ISO2022JP.NewEncoder())); err == nil {
		return string(msg)
	}
	return ""
}
