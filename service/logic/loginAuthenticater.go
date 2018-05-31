package logic

import (
  // "errors"
  "log"
  "crypto/sha256"
  "encoding/hex"
  "../model"
  "../dto"
)

func Authenticate(user dto.UserForm) (bool, string) {
  ok := model.IsExistUserName(user.Username)
  log.Println("ok : ", ok)
  if ok {
    pwd := model.GetUserPassword(user.Username)

    s := sha256.New()
    s.Write([]byte(user.Password))
    hash := s.Sum(nil)
    hexhash := hex.EncodeToString(hash)

    log.Println("pwd: ", hexhash)
    if hexhash == pwd {
      // ログイン成功
      return true, ""
    } else {
      // パスワードが違った場合
      // return false, errors.New("wrong password")
      return false, "wrong password"
    }
  } else {
    // ユーザネームが違った場合
    return false, "this username is not exist"
  }
}