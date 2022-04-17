package presenter

import (
	"fmt"
	"net/http"

	"auth-parts/entity"
	"auth-parts/usecase/port"
)

type User struct {
	w http.ResponseWriter
}

// NewUserOutputPort はUserOutputPortを取得する．
func NewUserOutputPort(w http.ResponseWriter) port.UserOutputPort {
	return &User{
		w: w,
	}
}

// usecase.UserOutputPortを実装している
// Render はNameを出力する．
func (u *User) Render(user *entity.User) {
	u.w.WriteHeader(http.StatusOK)
	// httpでentity.User.Nameを出力
	fmt.Fprint(u.w, user.Name)
}

// RenderError はErrorを出力する．
func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}
