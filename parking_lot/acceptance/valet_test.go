package acceptance

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/anlsergio/go_bdd/parking_lot/app/server"
	"github.com/anlsergio/go_bdd/parking_lot/config"
	"github.com/anlsergio/go_bdd/parking_lot/utils"
	"github.com/cucumber/godog"
)

var (
	apiURL = fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
	w      *httptest.ResponseRecorder
)

var apiServer *server.Server

func TestMain(m *testing.M) {
	apiServer = server.New()
	code := m.Run()
	os.Exit(code)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func aRequestIsSentToTheEndpoint(httpMethod, endpoint string) error {
	req, err := http.NewRequest(httpMethod, apiURL+endpoint, nil)
	if err != nil {
		return fmt.Errorf("could not create request %s", err.Error())
	}

	w = utils.ExecuteRequest(apiServer.Router, req)
	return nil
}

func theHTTPResponseCodeShouldBe(wantCode int) error {
	if wantCode != w.Code {
		return fmt.Errorf("the status code is different than the expected one. Want '%d', got '%d'", wantCode, w.Code)
	}
	return nil
}

func theResponseContentShouldBe(wantContent string) error {
	body, err := ioutil.ReadAll(w.Body)
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
