package str

import "golang.org/x/crypto/bcrypt"

// 密码加密
func PasswordHash(passwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(bytes), err
}

// 密码解密
func PasswordVerify(passwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
	return err == nil
}
