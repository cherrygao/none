package hook

import (
	"bytes"
	"fmt"

	"github.com/qiniu/rpc"
)

func Push(appHost string, userId string, messageBody []byte) (err error) {
	u := fmt.Sprintf("%s/push?uid=%s", appHost, userId)
	err = rpc.DefaultClient.CallWith(nil, nil, u, "text/html; charset=utf-8", bytes.NewReader(messageBody), len(messageBody))
}

func Notify(appHost string, userId string, action string) (err error) {
	u := appHost + "/notify"
	err = rpc.DefaultClient.CallWithForm(nil, nil, u, map[string][]string{
		"uid":    {userId},
		"action": {action},
	})
}
