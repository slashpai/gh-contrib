package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/google/go-github/github"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/oauth2"
)

type Contrib struct {
	apiClient *github.Client
}

func createGithubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func createContribObject(ctx context.Context, token string) *Contrib {
	return &Contrib{
		apiClient: createGithubClient(ctx, token),
	}
}

func markdownTableGenerator(header []string, rows [][]string, outputFile bool) {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(rows) // Add Bulk Data
	table.Render()
	fmt.Println(tableString.String())
	if outputFile {
		generateFile(tableString.String(), ".md")
	}
}

func generateFile(data string, extension string) {
	out, err := os.Create("output" + extension)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	out.WriteString(data)
}

func (f *Contrib) getContributionsOrg(ctx context.Context, org, author string, outputFile bool) {
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := f.apiClient.Repositories.ListByOrg(ctx, org, opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, repository := range repos {
		repo := repository.GetName()
		searchResult := f.getPullRequests(ctx, repo, org, author)
		if len(searchResult) != 0 {
			header := []string{"SL.No", "PR", "Title"}
			fmt.Printf("%s/%s\n", org, repo)
			fmt.Printf("Total PRs: %d\n", len(searchResult))
			markdownTableGenerator(header, searchResult, outputFile)
			fmt.Println()
		}
	}
}

// getContributionsRepo gets all pull requests created by a user in a repo
func (f *Contrib) getContributionsRepo(ctx context.Context, org, repo, author string, outputFile bool) {
	repository, _, err := f.apiClient.Repositories.Get(ctx, org, repo)

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	repoName := repository.GetName()
	searchResult := f.getPullRequests(ctx, repoName, org, author)

	if len(searchResult) != 0 {
		header := []string{"SL.No", "PR", "Title"}
		fmt.Printf("%s/%s\n", org, repo)
		fmt.Printf("Total PRs: %d\n", len(searchResult))
		markdownTableGenerator(header, searchResult, outputFile)
		fmt.Println()
	}
}

// getPullRequests gets all pull requests created by the user
func (f *Contrib) getPullRequests(ctx context.Context, repo, org, user string) [][]string {
	createdPullRequests := [][]string{}

	prQuery := "is:pr repo:" + org + "/" + repo + " author:" + user
	opt := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	prResults, _ := f.searchIssues(ctx, prQuery, opt)
	totalPullRequests := prResults.GetTotal()

	if totalPullRequests != 0 {
		for i, pr := range prResults.Issues {
			serialNumber := fmt.Sprintf("%v. ", i+1)
			prLink := fmt.Sprintf("[%s/%s#%v](%s)", org, repo, pr.GetNumber(), pr.GetHTMLURL())
			prTitle := pr.GetTitle()
			createdPullRequests = append(createdPullRequests, []string{serialNumber, prLink, prTitle})
		}
	}
	return createdPullRequests
}

func (f *Contrib) searchIssues(ctx context.Context, query string, opt *github.SearchOptions) (*github.IssuesSearchResult, error) {
	issuesResults, _, err := f.apiClient.Search.Issues(ctx, query, opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return issuesResults, nil
}

var flags struct {
	Token      string `help:"GitHub API token." required:""`
	OutputFile bool   `help:"Set to true if output file need to be generated."`
	ContribOrg struct {
		Org    string `help:"GitHub Org." required:""`
		Author string `help:"Author." required:""`
	} `cmd:"" help:"Github contributions of a user in a org."`

	ContribRepo struct {
		Org    string `help:"GitHub Org." required:""`
		Repo   string `help:"GitHub Repo in particular to check contributions." required:""`
		Author string `help:"Author." required:""`
	} `cmd:"" help:"Github contributions of a user to a repo of a org."`
}

func main() {
	ctx := context.Background()
	cli := kong.Parse(&flags,
		kong.Name("gh-contrib"),
		kong.Description("Count your github contributions from command line"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	switch cli.Command() {
	case "contrib-org":
		c := createContribObject(ctx, flags.Token)
		c.getContributionsOrg(ctx, flags.ContribOrg.Org, flags.ContribOrg.Author, flags.OutputFile)
	case "contrib-repo":
		c := createContribObject(ctx, flags.Token)
		c.getContributionsRepo(ctx, flags.ContribRepo.Org, flags.ContribRepo.Repo, flags.ContribRepo.Author, flags.OutputFile)
	}
}
