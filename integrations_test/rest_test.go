package integrations

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DATA-DOG/godog"
	_ "github.com/jackc/pgx/stdlib"
)

func (test *ServerTest) iAddBannerRestrequestTo(arg1 string) error {
	// Добавление события
	body := []byte(`{ "id_banner": 1, "id_slot": 1}`)
	url := arg1 + ":" + restPort + "/api/add_banner/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	test.responseBody = body
	return nil
}

func (test *ServerTest) theJsonResponseAddBannerMustContainStatus() error {
	if string(test.responseBody) != `{"status": true}` {
		return errors.New("wrong status")
	}
	return nil
}

func (test *ServerTest) iDelBannerRestrequestTo(arg1 string) error {
	// Удаление баннера из ротации
	url := arg1 + ":" + restPort + "/api/del_banner/1/"
	req, err := http.NewRequest("DELETE", url, nil)
	vars := map[string]string{"id": "1"}
	req = mux.SetURLVars(req, vars)
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	test.responseBody = body
	return nil
}

func (test *ServerTest) theJsonResponseDelBannerMustContainStatus() error {
	if string(test.responseBody) != `{"status": true}` {
		return errors.New("wrong status")
	}
	return nil
}

func (test *ServerTest) iCountTransitionBannerRestrequestTo(arg1 string) error {
	// Клик по баннеру
	body := []byte(`{"id_banner": 1,"id_slot": 1,"id_soc_dem": 1}`)
	url := arg1 + ":" + restPort + "/api/count_transition/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	test.responseBody = body
	return nil
}

func (test *ServerTest) theJsonResponseCountTransitionMustContainStatus() error {
	if string(test.responseBody) != `{"status": true}` {
		return errors.New("wrong status")
	}
	return nil
}

func (test *ServerTest) iGetBannerBannerRestrequestTo(arg1 string) error {
	// Удаление баннера из ротации
	url := arg1 + ":" + restPort + "/api/get_banner/1/1/"
	req, err := http.NewRequest("GET", url, nil)
	vars := map[string]string{"idSlot": "1", "idSocDemGroup": "1"}
	req = mux.SetURLVars(req, vars)
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	test.responseBody = body
	return nil
}

func (test *ServerTest) theJsonResponseMustContainIdBanner() error {
	if string(test.responseBody) != `{"status": true,"id_banner": 1}` {
		return errors.New("wrong status")
	}
	return nil
}

func RestContext(s *godog.Suite) {
	test := new(ServerTest)

	s.BeforeFeature(test.connectDB)
	s.BeforeFeature(test.truncateDb)
	s.AfterFeature(test.truncateDb)

	s.Step(`^I add banner rest-request to "([^"]*)"$`, test.iAddBannerRestrequestTo)
	s.Step(`^The json response add banner must contain status$`, test.theJsonResponseAddBannerMustContainStatus)
	s.Step(`^I del banner rest-request to "([^"]*)"$`, test.iDelBannerRestrequestTo)
	s.Step(`^The json response del banner must contain status$`, test.theJsonResponseDelBannerMustContainStatus)
	s.Step(`^I count transition banner rest-request to "([^"]*)"$`, test.iCountTransitionBannerRestrequestTo)
	s.Step(`^The json response  count transition must contain status$`, test.theJsonResponseCountTransitionMustContainStatus)
	s.Step(`^I Get banner banner rest-request to "([^"]*)"$`, test.iGetBannerBannerRestrequestTo)
	s.Step(`^The json response must contain id banner$`, test.theJsonResponseMustContainIdBanner)
}
