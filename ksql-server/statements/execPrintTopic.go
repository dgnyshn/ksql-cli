package statements

import (
	config "../../configManagement"
	model "../model"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func ExecutePrintTopic(query string) {
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
	for {

		line, err := reader.ReadBytes('\n')

		if len(string(line)) > 1 {
			fmt.Println(string(line))
			fmt.Println(len(string(line)))
		}

		if err != nil {
			return
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
