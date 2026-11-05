package top

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"most-active-github-users-counter/github"
	"most-active-github-users-counter/net"
)

func GithubTop(options Options) (github.GithubSearchResults, error) {
	var token = options.Token
	if token == "" {
		return github.GithubSearchResults{}, errors.New("Missing GITHUB token")
	}

	excludeStr := ""
	for _, location := range options.ExcludeLocations {
		excludeStr = fmt.Sprintf("%s -location:%s", excludeStr, location)
	}

	var client = github.NewGithubClient(net.TokenAuth(token))

	// Caso nao haja localizacoes especificadas (ex.: worldwide)
	if len(options.Locations) == 0 {
		query := fmt.Sprintf("type:user%s", excludeStr)
		return client.SearchUsers(github.UserSearchQuery{
			Q:        query,
			Sort:     "followers",
			Order:    "desc",
			MaxUsers: options.ConsiderNum,
		})
	}

	// Quando ha localizacoes, o GitHub Search API interpreta multiplos "location:" consecutivos
	// como conjuncao (AND). Por isso, consultamos as localizacoes separadamente e deduplicamos.
	seen := make(map[string]bool)
	var allUsers []github.User
	totalUserCount := 0
	minFollowers := -1

	for idx, loc := range options.Locations {
		// Delay defensivo de 1 segundo entre cidades para prevenir secondary rate limit do GitHub
		if idx > 0 {
			time.Sleep(1 * time.Second)
		}

		query := fmt.Sprintf("type:user location:%s%s", loc, excludeStr)
		res, err := client.SearchUsers(github.UserSearchQuery{
			Q:        query,
			Sort:     "followers",
			Order:    "desc",
			MaxUsers: options.ConsiderNum,
		})
		if err != nil {
			return github.GithubSearchResults{}, err
		}

		totalUserCount += res.TotalUserCount

		for _, u := range res.Users {
			if !seen[u.Login] {
				seen[u.Login] = true
				allUsers = append(allUsers, u)
			}
		}

		// Se ja atingimos usuarios suficientes para o corte inicial, otimizamos as chamadas
		if len(allUsers) >= options.ConsiderNum*2 {
			break
		}
	}

	// Ordena todos os usuarios unicos prioritariamente por Followers decrescente
	sort.Slice(allUsers, func(i, j int) bool {
		if allUsers[i].FollowerCount == allUsers[j].FollowerCount {
			return allUsers[i].ContributionCount > allUsers[j].ContributionCount
		}
		return allUsers[i].FollowerCount > allUsers[j].FollowerCount
	})

	if len(allUsers) > options.ConsiderNum {
		allUsers = allUsers[:options.ConsiderNum]
	}

	if len(allUsers) > 0 {
		minFollowers = allUsers[len(allUsers)-1].FollowerCount
	}

	return github.GithubSearchResults{
		Users:                allUsers,
		TotalUserCount:       totalUserCount,
		MinimumFollowerCount: minFollowers,
	}, nil
}

type Options struct {
	Token            string
	Locations        []string
	ExcludeLocations []string
	Amount           int
	ConsiderNum      int
	OrgsAmount       int
	PresetTitle      string
	PresetChecksum   string
}
