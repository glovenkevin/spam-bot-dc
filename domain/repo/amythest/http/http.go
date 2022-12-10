package amythest_http

import (
	"context"
	"discord-spam-bot/domain/repo"
	"discord-spam-bot/lib/constant"
	http_helper "discord-spam-bot/lib/pkg/http"
	"discord-spam-bot/lib/pkg/loggerext"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type AmythestRepoHttp struct {
	log        loggerext.LoggerInterface
	httpClient http_helper.HttpClientInterface
}

func NewAmythestRepoHttp(l loggerext.LoggerInterface, http http_helper.HttpClientInterface) repo.AmythestRepoInterface {
	return &AmythestRepoHttp{
		log:        l,
		httpClient: http,
	}
}

func (r *AmythestRepoHttp) GenerateWanted(ctx context.Context, imageLink string) ([]byte, error) {
	tracestr := "repo.amythest.http.GenerateWanted"
	select {
	case <-ctx.Done():
		r.log.Error(tracestr, ctx.Err())
		return nil, ctx.Err()
	default:
	}

	data := url.Values{}
	data.Set("url", imageLink)

	req, err := http.NewRequest(http.MethodPost, os.Getenv(constant.AmythestBaseURL)+generateWantedURI, strings.NewReader(data.Encode()))
	if err != nil {
		r.log.Error(tracestr, err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv(constant.AmythestToken)))
	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.log.Errorf("%s.http.Do: %v", tracestr, err)
		return nil, err
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.log.Errorf("%s.Response: %v", tracestr, err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		r.log.Errorf("%s.Response: %s", tracestr, string(respByte))
		return nil, err
	}

	return respByte, nil
}

func (r *AmythestRepoHttp) GenerateCircle(ctx context.Context, imageLink string) ([]byte, error) {
	tracestr := "repo.amythest.http.GenerateCircle"
	select {
	case <-ctx.Done():
		r.log.Error(tracestr, ctx.Err())
		return nil, ctx.Err()
	default:
	}

	data := url.Values{}
	data.Set("url", imageLink)

	req, err := http.NewRequest(http.MethodPost, os.Getenv(constant.AmythestBaseURL)+generateCircleURI, strings.NewReader(data.Encode()))
	if err != nil {
		r.log.Error(tracestr, err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv(constant.AmythestToken)))
	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.log.Errorf("%s.http.Do: %v", tracestr, err)
		return nil, err
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.log.Errorf("%s.Response: %v", tracestr, err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		r.log.Errorf("%s.Response: %s", tracestr, string(respByte))
		return nil, err
	}

	return respByte, nil
}
