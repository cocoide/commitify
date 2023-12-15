package usecase

import (
	"fmt"
	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/gateway"
	"net/url"
	"strconv"
	"time"
)

const (
	GithubClientID = "b27d87c28752d2363922"
	GithubScope    = "repo"
	GrantType      = "urn:ietf:params:oauth:grant-type:device_code"
)

type LoginCmdUsecase struct {
	http *gateway.HttpClient
}

func NewLoginCmdUsecase(http *gateway.HttpClient) *LoginCmdUsecase {
	http.WithBaseURL("https://github.com/login")
	return &LoginCmdUsecase{http: http}
}

type BeginGithubSSOResponse struct {
	DeviceCode string
	UserCode   string
	Interval   int
	ExpiresIn  int
}

func (u *LoginCmdUsecase) BeginGithubSSO() (*BeginGithubSSOResponse, error) {
	b, err := u.http.WithPath("device/code").
		WithParam("client_id", GithubClientID).
		WithParam("scope", GithubScope).
		Execute(gateway.POST)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}
	deviceCode := values.Get("device_code")
	userCode := values.Get("user_code")
	expiresIn, err := strconv.Atoi(values.Get("expires_in"))
	if err != nil {
		return nil, err
	}
	interval, err := strconv.Atoi(values.Get("interval"))
	if err != nil {
		return nil, err
	}
	if deviceCode == "" || userCode == "" {
		return nil, fmt.Errorf("failed to parse code")
	}
	return &BeginGithubSSOResponse{
		DeviceCode: deviceCode,
		UserCode:   userCode,
		ExpiresIn:  expiresIn,
		Interval:   interval,
	}, nil
}

type ScheduleVerifyAuthRequest struct {
	DeviceCode string
	Interval   int
	ExpiresIn  int
}

func (u *LoginCmdUsecase) ScheduleVerifyAuth(req *ScheduleVerifyAuthRequest) error {
	u.http = gateway.NewHttpClient().
		WithBaseURL("https://github.com/login").
		WithPath("oauth/access_token").
		WithParam("client_id", GithubClientID).
		WithParam("device_code", req.DeviceCode).
		WithParam("grant_type", GrantType)

	timeout := time.After(time.Duration(req.ExpiresIn) * time.Second)
	ticker := time.NewTicker(time.Duration(req.Interval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return fmt.Errorf("認証プロセスがタイムアウトしました")
		case <-ticker.C:
			b, err := u.http.Execute(gateway.POST)
			if err != nil {
				return err
			}
			values, err := url.ParseQuery(string(b))
			if err != nil {
				return err
			}
			accessToken := values.Get("access_token")
			if accessToken != "" {
				config, err := entity.ReadConfig()
				if err != nil {
					return err
				}
				config.WithGithubToken(accessToken)
				if err := config.WriteConfig(); err != nil {
					return err
				}
				return nil
			}
			if newIntervalStr := values.Get("interval"); newIntervalStr != "" {
				newInterval, err := strconv.Atoi(newIntervalStr)
				if err != nil {
					return err
				}
				ticker.Stop()
				ticker = time.NewTicker(time.Duration(newInterval) * time.Second)
			}
		}
	}
}
