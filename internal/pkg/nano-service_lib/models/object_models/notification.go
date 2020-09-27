
package object_models

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type NotificationType string

const (
	NotificationTypeError    NotificationType = "error"
	NotificationTypeTechInfo NotificationType = "tech_info"
)

type Notification struct {
	Id               string
	NotificationType NotificationType
	Header           string
	Message          string
}

func NewNotification(notificationType NotificationType, header string, message string, id string) *Notification {
	return &Notification{NotificationType: notificationType, Header: header, Message: message, Id: id}
}

func (n Notification) LockKey() string {
	var arrBytes []byte
	jsonBytes, _ := json.Marshal(n)
	arrBytes = append(arrBytes, jsonBytes...)

	return fmt.Sprintf("%x", md5.Sum(arrBytes))
}
