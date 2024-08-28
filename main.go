package main

import (
	"context"
	"fmt"
	"os"

	// "regexp"

	"github.com/google/go-github/v63/github"
)

func get_all_branches(token string, org_name string, repo_name string) map[string]bool {
	branches := make(map[string]bool)

	client := github.NewClient(nil).WithAuthToken(token)
	var current_page int = 0
	for {

		opt := &github.BranchListOptions{Protected: github.Bool(false), ListOptions: github.ListOptions{PerPage: 10, Page: current_page}}

		branches_per_page, resp, err := client.Repositories.ListBranches(context.Background(), org_name, repo_name, opt)

		if err != nil {
			panic(err)
		}
		for _, name := range branches_per_page {
			branches[*name.Name] = true
		}

		if resp.NextPage == 0 {
			break
		} else {
			current_page = resp.NextPage
		}
		fmt.Println("got ", len(branches), "branches")
	}

	return branches
}

func panic_on_err(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	org, ok := os.LookupEnv("GH_CLEANER_ORG")

	if !ok {
		panic("GH_CLEANER_TARGET not set")
	}

	target, ok := os.LookupEnv("GH_CLEANER_REPO")

	if !ok {
		panic("GH_CLEANER_TARGET not set")
	}

	token, ok := os.LookupEnv("GH_CLEANER_TOKEN")

	if !ok {
		panic("GH_CLEANER_token not set")
	}

	fmt.Println("cleaning", target)

	branches := get_all_branches(token, org, target)

	for branch_name, _ := range branches {
		fmt.Println(branch_name)
		client := github.NewClient(nil).WithAuthToken(token)
		fmt.Println("will delete branch ", "heads/"+branch_name)
		_, err := client.Git.DeleteRef(context.Background(), org, target, "heads/"+branch_name)
		fmt.Println("deleted branch", branch_name)
		panic_on_err(err)
	}

	fmt.Println("success!")
}
