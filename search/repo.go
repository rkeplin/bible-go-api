package search

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/rkeplin/bible-go-api/common"
)

type Repository struct {
	translationFactory common.TranslationFactory
}

func (r Repository) Search(query, translation string, offset, limit int) (TextCollection, error) {
	collection := TextCollection{}

	result, err := r.makeRequest(query, translation, offset, limit)

	if err != nil {
		return collection, err
	}

	collection.Total = result.Hits.Total.Value

	for _, innerHit := range result.Hits.Hits {
		b := Book{
			ID:        innerHit.Source.BookID,
			Name:      innerHit.Source.BookName,
			Testament: innerHit.Source.Testament,
		}

		t := Text{
			ID:        innerHit.Source.ID,
			Book:      b,
			ChapterID: innerHit.Source.ChapterID,
			VerseID:   innerHit.Source.VerseID,
			Verse:     innerHit.Highlight.Verse[0],
		}

		collection.Items = append(collection.Items, t)
	}

	return collection, nil
}

func (r Repository) makeRequest(query, translation string, offset, limit int) (ESResult, error) {
	result := ESResult{}
	index := r.translationFactory.GetIndex(translation)

	url := os.Getenv("ES_URL") + "/" + index + "/_search"
	esQueryStr := r.getESQuery(query, offset, limit)

	fmt.Println("ES URL: ", url)
	fmt.Println("ES Query: ", query)

	var jsonStr = []byte(esQueryStr)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, errors.New("Non-200 Response From ES")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &result)

	return result, nil
}

func (r Repository) getESQuery(query string, offset, limit int) string {
	if offset < 0 {
		offset = 0
	}

	if limit == 0 {
		limit = 100
	}

	if limit > 1000 {
		limit = 1000
	}

	ESQuery := `{
		"from":` + strconv.Itoa(offset) + `,
		"size":` + strconv.Itoa(limit) + `,
		"query": {
			"match": {
				"verse":"` + query + `"
			}
		},
		"highlight":{
			"pre_tags": ["<span class=\"highlight\">"],
			"post_tags":["</span>"],
			"fields": {
				"verse": {
					"type": "plain"
				}
			}
		}
	}`

	return ESQuery
}
