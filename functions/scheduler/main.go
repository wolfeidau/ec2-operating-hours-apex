package main

import (
	"encoding/json"
	"os"

	apex "github.com/apex/go-apex"
	"github.com/apex/log"
	jsonlog "github.com/apex/log/handlers/json"
	"github.com/aws/aws-sdk-go/aws/session"
	operatinghours "github.com/wolfeidau/ec2-operating-hours"
)

type message struct {
	Hello string `json:"hello"`
}

func main() {
	log.SetHandler(jsonlog.New(os.Stderr))

	ohrs := operatinghours.NewOperatingHours(session.New())

	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {

		var m message
		err := ohrs.Check("Australia/Melbourne")
		if err != nil {
			log.WithError(err).Error("failed to run check")
		}

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		return m, nil
	})
}
