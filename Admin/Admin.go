package Admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var HEADERS = map[string]string{
	"Content-Type": "application/x-www-form-urlencoded",
	"User-Agent":   "AuthGGo - smartass08"}

/*type LicencesAll struct {
	Token  string `json:"token"`
	Rank   string `json:"rank"`
	Used   string `json:"used"`
	UsedBy string `json:"used_by"`
	Days   string `json:"days"`
}*/

type Administration struct {
	apiKey     string
	apiUrl     string
	apiHeaders string
}

func (rc *Administration) Init(hash string) error {
	rc.apiUrl = "https://developers.auth.gg"
	rc.apiKey = hash
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/USERS?type=count&authorization=%v", rc.apiUrl, rc.apiKey), nil)
	client := &http.Client{}
	_, err := client.Do(r)
	if err != nil {
		return err
	}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	if strings.Contains(string(responseData), "failed") == true && strings.Contains(string(responseData), "No application found") == true {
		log.Println(string(responseData))
		return err
	}
	return nil
}

// fetchOne - Common method for various similar stuff
func (rc Administration) fetchOne(userName string, who string, which string ) (map[string]interface{}, error) {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/%v?type=fetch&authorization=%v&%v=%v", rc.apiUrl, who, rc.apiKey,which, userName), nil)
	r.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent":   []string{"AuthGGo - smartass08"},
	}
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil, err
	}
	return v.(map[string]interface{}), nil
}

// fetchAll - common method for fetching group of stuff
func (rc *Administration) fetchAll(which string) (map[string]interface{}, error) {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/%v?type=fetchall&authorization=%v", rc.apiUrl, which, rc.apiKey), nil)
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil, err
	}
	return v.(map[string]interface{}) , nil
}

// fetchCount -  Common method for grabbing all count
func (rc Administration) fetchCount(which string) (int, error) {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/%v?type=count&authorization=%v", rc.apiUrl, which, rc.apiKey), nil)
	r.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent":   []string{"AuthGGo - smartass08"},
	}
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		return 0, err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return 0, err
	}
	val, err := strconv.Atoi(v.(map[string]interface{})["value"].(string))
	if err != nil {
		return 0, err
	}
	return val, nil
}

// GenerateLicense :- License generator - Max 9998 days
func (rc Administration) GenerateLicense(amount int, days int, prefix string) (map[string]interface{}, error) {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/LICENSES?type=generate&&authorization=%v&amount=%v&days=%v&level=0&format=3&prefix=%v&length=0", rc.apiUrl, rc.apiKey, amount, days, prefix), nil)
	r.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent":   []string{"AuthGGo - smartass08"},
	}
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil, err
	}
	return v.(map[string]interface{}), nil
}
func (rc Administration) changeLicense(license string, which string) (map[string]interface{}, error) {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/LICENSES?type=%v&authorization=%v&license=%v", rc.apiUrl, which, rc.apiKey, license), nil)
	r.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent":   []string{"AuthGGo - smartass08"},
	}
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil, err
	}
	return v.(map[string]interface{}), nil
}
func (rc Administration) updateHwid(license string, which string, which2 string) (map[string]interface{}, error) {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/HWID?type=%v&authorization=%v&%s=%v", rc.apiUrl, which, rc.apiKey, which2, license), nil)
	r.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent":   []string{"AuthGGo - smartass08"},
	}
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil, err
	}
	return v.(map[string]interface{}), nil
}

func (rc *Administration) FetchAllUsedLicenses(user string) ([]string, error) {
	var allUselessData []string
	allLicenses, err:= rc.FetchAllLicenseInfo()
	if err != nil {
		return allUselessData, err
	}
	for _, v := range  allLicenses {
		moreUselessData := v.(map[string]interface{})
		if moreUselessData["used"] == "1" && moreUselessData["used_by"].(string) == user {
			allUselessData = append(allUselessData, moreUselessData["token"].(string))
		}
	}
	return allUselessData, nil
}

func (rc *Administration) FetchAllUserInfo() (map[string]interface{}, error) {
	return rc.fetchAll("USERS")
}

func (rc *Administration) FetchAllLicenseInfo() (map[string]interface{}, error) {
	return rc.fetchAll("LICENSES")
}

func (rc Administration) FetchUserCount() (int, error) {
	return rc.fetchCount("USERS")
}

func (rc Administration) FetchLicenseCount() (int, error) {
	return rc.fetchCount("LICENSES")
}

func (rc Administration) FetchUserInfo(username string) (map[string]interface{}, error) {
	return rc.fetchOne(username, "USERS", "user")
}

func (rc Administration) FetchLicenseInfo(license string) (map[string]interface{}, error) {
	return rc.fetchOne(license, "LICENSES", "license")
}

func (rc Administration) DeleteKey(licenseKey string) (map[string]interface{}, error) {
	return rc.changeLicense(licenseKey, "delete")
}

func (rc Administration) UseKey(licenseKey string) (map[string]interface{}, error) {
	return rc.changeLicense(licenseKey, "use")
}

func (rc Administration) UnUseKey(licenseKey string) (map[string]interface{}, error) {
	return rc.changeLicense(licenseKey, "unuse")
}

func (rc Administration) FetchHwid(username string) (map[string]interface{}, error) {
	return rc.updateHwid(username, "fetch", "user")
}

func (rc Administration) ResetHwid(username string) (map[string]interface{}, error) {
	return rc.updateHwid(username, "reset", "user")
}
