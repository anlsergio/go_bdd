package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/anlsergio/go_bdd/parking_lot/config"
	"github.com/cucumber/godog"
)

var (
	apiURL = fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
	res    *http.Response
)

func aRequestIsSentToTheEndpoint(httpMethod, endpoint string) error {
	reader := strings.NewReader("")
	req, err := http.NewRequest(
		httpMethod,
		apiURL+endpoint,
		reader)
	if err != nil {
		return fmt.Errorf("could not create request %s", err.Error())
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request %s", err.Error())
	}
	return nil
}

func theHTTPResponseCodeShouldBe(wantCode int) error {
	if wantCode != res.StatusCode {
		return fmt.Errorf("the status code is different than the expected one. Want '%d', got '%d'", wantCode, res.StatusCode)
	}
	return nil
}

func theResponseContentShouldBe(wantContent string) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not read from the response body: %s", err)
	}

	content := strings.TrimSpace(string(body))
	if wantContent != content {
		return fmt.Errorf("the response content is different than the expected one. Want '%s', got '%s'", wantContent, content)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a "([^"]*)" request is sent to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)
	ctx.Step(`^the HTTP response code should be (\d+)$`, theHTTPResponseCodeShouldBe)
	ctx.Step(`^the response content should be "([^"]*)"$`, theResponseContentShouldBe)
}
