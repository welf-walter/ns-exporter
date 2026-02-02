package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

type NSClient struct {
	nsUri   string
	nsToken string
	user    string
	jwt     string
	logging bool
}

type nsDeviceStatusResult struct {
	Status  int       `json:"status"`
	Records []NsEntry `json:"result"`
}
type nsTreatmentsResult struct {
	Status  int           `json:"status"`
	Records []NsTreatment `json:"result"`
}
type nsJwtResult struct {
	Token string `json:"token"`
}

func NewNSClient(uri string, token string, user string, logging bool) *NSClient {
	return &NSClient{
		nsUri:   strings.TrimRight(uri, "/"),
		nsToken: token,
		user:    user,
		logging: logging,
	}
}

func (c *NSClient) Authorize(_ context.Context) {
	client := resty.New()
	result := &nsJwtResult{}
	_, err := client.R().
		SetResult(result).
		SetHeader("Accept", "application/json").
		Get(c.nsUri + "/api/v2/authorization/request/" + c.nsToken)

	if err != nil {
		log.Fatal(err)
	}
	c.jwt = result.Token
}

func (c *NSClient) LoadDeviceStatuses(queue chan NsEntry, limit int64, skip int64, _ context.Context) {
	defer wg.Done()

	fmt.Println("LoadDeviceStatuses from NS, limit: ", limit, ", skip: ", skip)

	client := resty.New()

	entries := &nsDeviceStatusResult{}
	response, err := client.R().
		SetQueryParams(map[string]string{
			"skip":      strconv.FormatInt(skip, 10),
			"limit":     strconv.FormatInt(limit, 10),
			"sort$desc": "created_at",
		}).
		SetAuthScheme("Bearer").
		SetAuthToken(c.jwt).
		SetHeader("Accept", "application/json").
		SetResult(entries).
		Get(c.nsUri + "/api/v3/devicestatus")

	if c.logging {
		log.Printf("Response: %v", response)
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries.Records {
		if strings.HasPrefix(entry.Device, "openaps") {
			entry.User = c.user
			queue <- entry
		}
	}
}

func (c *NSClient) LoadTreatments(queue chan NsTreatment, limit int64, skip int64, _ context.Context) {
	defer wg.Done()

	fmt.Println("LoadTreatments from NS, limit: ", limit, ", skip: ", skip)

	client := resty.New()

	entries := &nsTreatmentsResult{}
	response, err := client.R().
		SetDebug(c.logging).
		SetQueryParams(map[string]string{
			"skip":      strconv.FormatInt(skip, 10),
			"limit":     strconv.FormatInt(limit, 10),
			"sort$desc": "created_at",
		}).
		SetResult(entries).
		SetHeader("Accept", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(c.jwt).
		Get(c.nsUri + "/api/v3/treatments")

	if c.logging {
		log.Printf("Response: %v", response)
	}

	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries.Records {
		entry.User = c.user
		queue <- entry
	}
}

func (c *NSClient) Close(_ context.Context) {}
