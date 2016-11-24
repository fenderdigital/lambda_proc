package main

import (
	"encoding/json"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/fenderdigital/lambda_proc"
	"github.com/joho/godotenv"
)

func main() {
	logrus.SetOutput(os.Stderr)

	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file to load")
	}

	lambda_proc.Run(func(context *lambda_proc.Context, eventJSON json.RawMessage) (interface{}, error) {
		var v map[string]interface{}
		if err := json.Unmarshal(eventJSON, &v); err != nil {
			return nil, err
		}
		logrus.Infof("env var: %s", os.Getenv("KEY"))
		logrus.Infof("context: %+v", *context)
		data, _ := json.MarshalIndent(&v, "", "    ")

		logrus.Infof("event: %+v", string(data))
		return v, nil
	})
}
