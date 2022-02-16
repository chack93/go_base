package session

import "github.com/chack93/go_base/internal/service/model"

type Session struct {
	model.Model
	JoinCode string `json:"joinCode"`
}
