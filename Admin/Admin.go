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

func (rc *Administration) Init(hash string) bool {
	rc.apiUrl = "https://developers.auth.gg"
	rc.apiKey = hash
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/USERS?type=count&authorization=%v", rc.apiUrl, rc.apiKey), nil)
	client := &http.Client{}
	_, e := client.Do(r)
	if e != nil {
		log.Println(e.Error())
		return false
	}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	if strings.Contains(string(responseData), "failed") == true && strings.Contains(string(responseData), "No application found") == true {
		log.Println(string(responseData))
		return false
	}
	return true
}

// fetchOne - Common method for various similar stuff
func (rc Administration) fetchOne(userName string, which string) map[string]interface{} {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/%v?type=fetch&authorization=%v&user=%v", rc.apiUrl, which, rc.apiKey, userName), nil)
	r.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent":   []string{"AuthGGo - smartass08"},
	}
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil
	}
	if err != nil {
		return nil
	}
	return v.(map[string]interface{})
}

// fetchAll - common method for fetching group of stuff
func (rc *Administration) fetchAll(which string) map[string]interface{} {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/%v?type=fetchall&authorization=%v", rc.apiUrl, which, rc.apiKey), nil)
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil
	}
	if err != nil {
		return nil
	}
	return v.(map[string]interface{})
}

// fetchCount -  Common method for grabbing all count
func (rc Administration) fetchCount(which string) int {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/%v?type=count&authorization=%v", rc.apiUrl, which, rc.apiKey), nil)
	r.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent":   []string{"AuthGGo - smartass08"},
	}
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return 0
	}
	if err != nil {
		return 0
	}
	val, err := strconv.Atoi(v.(map[string]interface{})["value"].(string))
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return val
}

// GenerateLicense :- License generator - Max 9998 days
func (rc Administration) GenerateLicense(amount int, days int, prefix string) map[string]interface{} {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/LICENSES?type=generate&&authorization=%v&amount=%v&days=%v&level=1&format=3&prefix=%v&length=0", rc.apiUrl, rc.apiKey, amount, days, prefix), nil)
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil
	}
	if err != nil {
		return nil
	}
	return v.(map[string]interface{})
}
func (rc Administration) changeLicense(license string, which string) map[string]interface{} {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/LICENSES?type=%v&authorization=%v&license=%v", rc.apiUrl, which, rc.apiKey, license), nil)
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil
	}
	if err != nil {
		return nil
	}
	return v.(map[string]interface{})
}
func (rc Administration) updateHwid(license string, which string, which2 string) map[string]interface{} {
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/HWID?type=%v&authorization=%v&%s=%v", rc.apiUrl, which, rc.apiKey, which2, license), nil)
	client := &http.Client{}
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var v interface{}
	err = json.Unmarshal(responseData, &v)
	if err != nil {
		return nil
	}
	if err != nil {
		return nil
	}
	return v.(map[string]interface{})
}

func (rc *Administration) FetchAllUsedLicenses(user string) []string {
	var allUselessData []string
	uselessData := rc.FetchAllLicenseInfo()
	for _, v := range uselessData {
		moreUselessData := v.(map[string]interface{})
		if moreUselessData["used"] == "1" && strings.Contains(moreUselessData["used_by"].(string), user) == true {
			allUselessData = append(allUselessData, moreUselessData["token"].(string))
		}
	}
	return allUselessData
}

func (rc *Administration) FetchAllUserInfo() map[string]interface{} {
	return rc.fetchAll("USERS")
}

func (rc *Administration) FetchAllLicenseInfo() map[string]interface{} {
	return rc.fetchAll("LICENSES")
}

func (rc Administration) FetchUserCount() int {
	return rc.fetchCount("USERS")
}

func (rc Administration) FetchLicenseCount() int {
	return rc.fetchCount("LICENSES")
}

func (rc Administration) FetchUserInfo(username string) map[string]interface{} {
	return rc.fetchOne(username, "USERS")
}

func (rc Administration) FetchLicenseInfo(license string) map[string]interface{} {
	return rc.fetchOne(license, "LICENSES")
}

func (rc Administration) DeleteKey(licenseKey string) map[string]interface{} {
	return rc.changeLicense(licenseKey, "delete")
}

func (rc Administration) UseKey(licenseKey string) map[string]interface{} {
	return rc.changeLicense(licenseKey, "use")
}

func (rc Administration) UnUseKey(licenseKey string) map[string]interface{} {
	return rc.changeLicense(licenseKey, "unuse")
}

func (rc Administration) FetchHwid(username string) map[string]interface{} {
	return rc.updateHwid(username, "fetch", "user")
}

func (rc Administration) ResetHwid(username string) map[string]interface{} {
	return rc.updateHwid(username, "reset", "user")
}
