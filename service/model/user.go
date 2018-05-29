package model

import (
  "log"
  "../db"
)

func IsExistUserName(name string) bool {
  db := db.DbConn()
  sql := "select name from m_test_user where name = ? and delete_flag = ?"
  stmt, err := db.Prepare(sql)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  rows, err := stmt.Query(name, 0)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  result := rows.Next()
  return result
}

func GetUserPassword(name string) string {
  db := db.DbConn()
  sql := "select password from m_test_user where name = ? and delete_flag = ?"
  stmt, err := db.Prepare(sql)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  var password string
  row := stmt.QueryRow(name, 0)

  err = row.Scan(&password)
  if err != nil {
    log.Fatal(err)    
  }
  return password
}
