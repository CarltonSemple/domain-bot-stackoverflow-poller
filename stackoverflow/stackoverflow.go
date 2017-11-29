package stackoverflow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const apiVersion = "2.2"

// GetAcceptedAnswer ...
func GetAcceptedAnswer(questionID int, accessToken, apiKey string) (Answer, error) {
	resp, err := http.Get("https://api.stackexchange.com/" + apiVersion + "/questions/" + strconv.Itoa(questionID) + "/answers?order=desc&sort=activity&site=stackoverflow&filter=!9YdnSLiq6" +
		"&access_token=" + accessToken + "&key=" + apiKey)
	if err != nil {
		return Answer{}, errors.Wrapf(err, "")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Println(resp.Status)
		log.Println(string(body))
	}

	answersCont := answersContainer{}
	json.Unmarshal(body, &answersCont)

	for _, answer := range answersCont.Items {
		if answer.IsAccepted {
			return answer, nil
		}
	}
	return Answer{}, nil
}

// GetTaggedQuestions ...
func GetTaggedQuestions(tag string, fromDate, toDate time.Time, secondsBetweenRequests int, accessToken, apiKey string) ([]QuestionInfo, error) {
	questions := []QuestionInfo{}
	page := 1
	log.Printf("page %v", page)
	for {
		resp, err := http.Get("https://api.stackexchange.com/" + apiVersion + "/questions?order=desc&sort=activity&tagged=" +
			tag + "&site=stackoverflow&pagesize=100&page=" + strconv.Itoa(page) + "&fromdate=" + fmt.Sprintf("%v", fromDate.Unix()) + "&todate=" + fmt.Sprintf("%v", toDate.Unix()) +
			"&access_token=" + accessToken + "&key=" + apiKey)
		if err != nil {
			return questions, errors.Wrapf(err, "")
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != 200 {
			log.Println(resp.Status)
			log.Println(string(body))
		}
		//log.Println(string(body))
		questionsCon := questionsContainer{}
		json.Unmarshal(body, &questionsCon)

		questions = append(questions, questionsCon.Items...)
		page++

		if !questionsCon.HasMore {
			break
		}
		log.Println("sleeping")
		time.Sleep(time.Duration(secondsBetweenRequests) * time.Second)
		log.Printf(", %v ", page)
	}
	log.Printf("\n")
	return questions, nil
}
