package top

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"most-active-github-users-counter/github"
	"most-active-github-users-counter/net"
)

func chunkLocations(locations []string) [][]string {
	if len(locations) == 0 {
		return nil
	}

	// Se o preset possui ate 12 localizacoes e a string combinada e compacta (<= 180 chars),
	// executa em um unico chunk para evitar multiplas buscas desnecessarias.
	totalLen := 0
	for _, loc := range locations {
		totalLen += len(" location:") + len(loc)
	}
	if len(locations) <= 12 && totalLen <= 180 {
		return [][]string{locations}
	}

	const maxLocsPerChunk = 8
	const maxCharsPerChunk = 150

	var chunks [][]string
	var current []string
	currentLen := 0

	for _, loc := range locations {
		partLen := len(" location:") + len(loc)
		if len(current) > 0 && (len(current) >= maxLocsPerChunk || currentLen+partLen > maxCharsPerChunk) {
			chunks = append(chunks, current)
			current = nil
			currentLen = 0
		}
		current = append(current, loc)
		currentLen += partLen
	}
	if len(current) > 0 {
		chunks = append(chunks, current)
	}

	return chunks
}

func GithubTop(options Options) (github.GithubSearchResults, error) {
	var token = options.Token
	if token == "" {
		return github.GithubSearchResults{}, errors.New("Missing GITHUB token")
	}

	var excludeBuilder strings.Builder
	for _, location := range options.ExcludeLocations {
		excludeBuilder.WriteString(" -location:")
		excludeBuilder.WriteString(location)
	}
	excludeStr := excludeBuilder.String()

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

	chunks := chunkLocations(options.Locations)
	if len(chunks) == 1 {
		var locBuilder strings.Builder
		for _, loc := range chunks[0] {
			locBuilder.WriteString(" location:")
			locBuilder.WriteString(loc)
		}
		query := fmt.Sprintf("type:user%s%s", locBuilder.String(), excludeStr)
		return client.SearchUsers(github.UserSearchQuery{
			Q:        query,
			Sort:     "followers",
			Order:    "desc",
			MaxUsers: options.ConsiderNum,
		})
	}

	seen := make(map[string]bool)
	var allUsers []github.User
	totalUserCount := 0
	minFollowers := -1

	for idx, chunk := range chunks {
		if idx > 0 {
			time.Sleep(2 * time.Second)
		}

		var locBuilder strings.Builder
		for _, loc := range chunk {
			locBuilder.WriteString(" location:")
			locBuilder.WriteString(loc)
		}

		query := fmt.Sprintf("type:user%s%s", locBuilder.String(), excludeStr)
		res, err := client.SearchUsers(github.UserSearchQuery{
			Q:        query,
			Sort:     "followers",
			Order:    "desc",
			MaxUsers: options.ConsiderNum,
		})
		if err != nil {
			if len(allUsers) > 0 {
				log.Printf("warning: chunk %d/%d failed (%v); proceeding with %d users collected so far", idx+1, len(chunks), err, len(allUsers))
				break
			}
			return github.GithubSearchResults{}, err
		}

		if res.TotalUserCount > totalUserCount {
			totalUserCount = res.TotalUserCount
		}

		for _, u := range res.Users {
			if !seen[u.Login] {
				seen[u.Login] = true
				allUsers = append(allUsers, u)
			}
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
