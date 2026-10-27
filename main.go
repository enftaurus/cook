package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const leetcodeEndpoint = "https://leetcode.com/graphql"

var count = 0

type ProfileResponse struct {
	Data struct {
		MatchedUser *struct {
			Username string `json:"username"`
			Profile  struct {
				RealName string `json:"realName"`
				AboutMe  string `json:"aboutMe"`
			} `json:"profile"`
			SubmitStats struct {
				AcceptedSubmissions []struct {
					Difficulty string `json:"difficulty"`
					Count      int    `json:"count"`
				} `json:"acSubmissionNum"`
			} `json:"submitStats"`
		} `json:"matchedUser"`
	} `json:"data"`
}

func fetchLeetCodeProfile(
	client *http.Client,
	username string,
) error {
	const query = `
		query userProfile($username: String!) {
			matchedUser(username: $username) {
				username
				profile {
					realName
					aboutMe
				}
				submitStats {
					acSubmissionNum {
						difficulty
						count
					}
				}
			}
		}
	`

	payload := map[string]any{
		"query": query,
		"variables": map[string]string{
			"username": username,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		leetcodeEndpoint,
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://leetcode.com/")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return fmt.Errorf(
			"rate limited (429); Retry-After: %s",
			resp.Header.Get("Retry-After"),
		)
	}

	if resp.StatusCode == http.StatusForbidden {
		return fmt.Errorf("access denied (403)")
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result ProfileResponse
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	user := result.Data.MatchedUser
	if user == nil {
		fmt.Printf("%s: profile unavailable\n", username)
		return nil
	}

	fmt.Printf("\nUsername: %s\n", user.Username)
	fmt.Printf("Name: %s\n", user.Profile.RealName)

	for _, stat := range user.SubmitStats.AcceptedSubmissions {
		fmt.Printf("  %-8s %d accepted\n", stat.Difficulty, stat.Count)
	}

	return nil
}

// FetchTopProfiles fetches the public profiles of the 10 usernames
// extracted from the leaderboard you provided.
func FetchTopProfiles() {
	usernames := []string{
		"404personality",
		"wrench420",
		"MKhubaibZafar",
		"hamza_aziz_ai",
		"Shailesh_M_J",
		"this_is_derek",
		"liyangzi6p626",
		"workinghours",
		"vOerSaT5rb",
		"exe_cutor",
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	for i, username := range usernames {
		fmt.Printf(
			"\n[%d/%d] Fetching %s\n",
			i+1,
			len(usernames),
			username,
		)

		err := fetchLeetCodeProfile(client, username)
		if err != nil {
			fmt.Printf("Error for %s: %v\n", username, err)

			// Stop if LeetCode denies access or rate-limits requests.
			if strings.Contains(err.Error(), "429") ||
				strings.Contains(err.Error(), "403") {
				fmt.Println("Stopping further requests.")
				return
			}
			count++
		}

		if i < len(usernames)-1 {

		}
	}
}
func main() {
	for i := 0; i < 10; i++ {
		FetchTopProfiles()

	}
	fmt.Println("total count:", count)
}
