package main

import (
	"fmt"
	"html"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/carltonsemple/domain-bot-stackoverflow-poller/discovery"
	"github.com/carltonsemple/domain-bot-stackoverflow-poller/stackoverflow"
)

const ()

var (
	daysInHistoryToQuery     = 0
	secondsBetweenRequests   = 0
	stackexchangeAccessToken = os.Getenv("STACKEXCHANGE_ACCESS_TOKEN")
	stackexchangeKey         = os.Getenv("STACKEXCHANGE_API_KEY") // not considered a secret and may safely be embedded in code
	discoveryUsername        = os.Getenv("DISCOVERY_USERNAME")
	discoveryPassword        = os.Getenv("DISCOVERY_PASSWORD")
	discoveryEnvironmentID   = os.Getenv("DISCOVERY_ENVIRONMENT_ID")
	discoveryCollectionID    = os.Getenv("DISCOVERY_COLLECTION_ID")
	tags                     = []string{}
)

func setEnvVars() {
	requestDelay, err := strconv.Atoi(os.Getenv("SECONDS_BETWEEN_REQUESTS"))
	if err != nil {
		log.Println("error with env SECONDS_BETWEEN_REQUESTS", err)
		return
	}
	secondsBetweenRequests = requestDelay
	daysToQuery, err := strconv.Atoi(os.Getenv("QUERY_DAYS_PAST"))
	if err != nil {
		log.Println("error with env QUERY_DAYS_PAST", err)
		return
	}
	daysInHistoryToQuery = daysToQuery
	tagsEnv := os.Getenv("TAGS")
	tags = strings.Split(tagsEnv, ",")
}

func main() {
	setEnvVars()
	questions := []stackoverflow.QuestionInfo{}
	for _, tag := range tags {
		log.Println(`tag: `, tag)
		xQuestions, err := GetQuestionsFromLastXDays(daysInHistoryToQuery, tag)
		if err != nil {
			log.Println(err)
			return
		}
		questions = append(questions, xQuestions...)
	}
	log.Println(len(questions), ` questions found`)
	for qNum, question := range questions {
		if question.IsAnswered {
			if answer, err := stackoverflow.GetAcceptedAnswer(question.QuestionID, stackexchangeAccessToken, stackexchangeKey); err != nil {
				log.Println("GetAcceptedAnswer: ", err)
			} else {
				if (stackoverflow.Answer{} != answer) {
					log.Println("----------------------------------")
					log.Println(question.Title)
					body := answer.BodyMarkdown
					body = string(blackfriday.Run([]byte(body)))
					body = html.UnescapeString(body)
					log.Println("\n" + body + "\n" + question.Link + "/" + strconv.Itoa(answer.AnswerID) + "#" + strconv.Itoa(answer.AnswerID))
					log.Println("\n\n\n ")
					pushQuestionAnswerToDiscovery(question, answer)
				}
			}
			log.Println("sleeping after question ", qNum, " out of ", len(questions), " questions")
			time.Sleep(time.Duration(secondsBetweenRequests) * time.Second)
		}
	}
	log.Println("finished")
}

func pushQuestionAnswerToDiscovery(question stackoverflow.QuestionInfo, answer stackoverflow.Answer) error {
	url := question.Link + "/" + strconv.Itoa(answer.AnswerID) + "#" + strconv.Itoa(answer.AnswerID)
	documentID := discovery.StackoverflowURLToDiscoveryID(url)
	discoveryDoc := discovery.DiscoveryDocAdapter{
		EnvironmentID:     discoveryEnvironmentID,
		CollectionID:      discoveryCollectionID,
		DocumentID:        documentID,
		Content:           fmt.Sprint(question.Title, ":\n\n", html.UnescapeString(answer.BodyMarkdown)),
		URL:               url,
		RepositoryAccount: "stackexchange",
		RepositoryName:    "stackoverflow",
	}

	return discovery.UpdateDocument(discoveryDoc, discoveryUsername, discoveryPassword)
}

// GetQuestionsFromLastXDays ...
func GetQuestionsFromLastXDays(numDays int, questionTag string) ([]stackoverflow.QuestionInfo, error) {
	startTime := time.Now().AddDate(0, 0, -numDays)
	return stackoverflow.GetTaggedQuestions(questionTag, startTime, time.Now(), secondsBetweenRequests, stackexchangeAccessToken, stackexchangeKey)
}
