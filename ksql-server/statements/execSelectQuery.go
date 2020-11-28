package statements

import (
	config "../../configManagement"
	visualize "../../visuailize"
	model "../model"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexeyco/simpletable"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)


func Execute(query string) {
	ctx, cancel := context.WithCancel(context.Background())
	r := model.Request{
		Ksql: strings.TrimSpace(query),
	}
	byteQuery, _ := json.Marshal(r)

	req, err := http.NewRequest("POST", config.DefaultKsqlServerURL+"/query", bytes.NewBuffer(bytes.ToLower(byteQuery)))
	req = req.WithContext(ctx)

	if err != nil {
		panic(err)
	}

	if config.DefaultAuthentication != "" {
		req.Header.Set("Authentication", config.DefaultAuthentication)
	}
	req.Header.Set("Accept", "application/vnd.ksql.v1+json")

	client := &http.Client{}
	res, err := client.Do(req)

	reader := bufio.NewReader(res.Body)

	var isHeader bool
	var table = simpletable.New()
	for {

		line, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}
		row := strings.Trim(string(line), " ")
		if !isHeader {
			headerResponse, err := convertResponseHeader(row)

			if err != nil {
				panic(err)
			}

			splitSchemaValues := strings.Split(headerResponse.Header.Schema, ",")
			isHeader = true
			visualize.VisualizeSelectStatementHeader(splitSchemaValues, table)
		} else if row != "" {
			bodyResponse := convertResponseBody(row)

			if len(bodyResponse.Row.Columns) > 1 {
				visualize.VisualizeSelectStatementBody(bodyResponse, table)
			}

		}

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			fmt.Println()
			cancel()
		}()

	}
}

func convertResponseHeader(row string) (model.SelectStatementHeaderResponse, error) {
	replacer := strings.NewReplacer("[", "", "`", "")
	replacedHeader := replacer.Replace(row)
	suffix := strings.TrimSuffix(replacedHeader, ",\n")

	var headerResponse model.SelectStatementHeaderResponse
	err := json.Unmarshal([]byte(suffix), &headerResponse)
	return headerResponse, err
}

func convertResponseBody(row string) model.SelectStatementBodyResponse {
	suffix := strings.TrimSuffix(row, ",\n")
	var bodyResponse model.SelectStatementBodyResponse
	err := json.Unmarshal([]byte(suffix), &bodyResponse)

	if err != nil {

	}
	return bodyResponse
}
