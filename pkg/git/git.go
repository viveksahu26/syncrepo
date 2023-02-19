package git

import (
	"context"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type Services interface {
	Clone(ctx context.Context, project, url string, username string, token string) (*git.Repository, error)
	Checkout(ctx context.Context, repo *git.Repository, branch string, commitHash string) error
	CreateRemote(ctx context.Context, repo *git.Repository, remoteName string, url string) (*git.Remote, error)
	Push(ctx context.Context, repo *git.Repository, remoteName, username, token string) error
}

// struct gitService will implement all methods of GitServices.
type GitService struct{}

// Clone: It clones the repository.
func (g *GitService) Clone(ctx context.Context, path, repoURL, username, token string) (*git.Repository, error) {
	// add logs: cloning repo

	// cloning requires username and password for private repos
	repo, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:  repoURL,
		Auth: &http.BasicAuth{Username: string(username), Password: token},
	})
	if err != nil {
		// Add logs: "error cloning repository
		return nil, err
	}

	// Add logs: succesfully repo cloned
	return repo, nil
}

// Push: pushes main repo changes to followers repo
func (g *GitService) Push(ctx context.Context, repo *git.Repository, remoteName, username, token string) error {
	// Add log: pushing to remote repo

	err := repo.Push(&git.PushOptions{
		RemoteName: remoteName,
		Auth: &http.BasicAuth{
			Username: username,
			Password: token,
		},
		// push forcefully
		Force: true,
	})
	if err != nil {
		// Add logs: "error pushing to remote

		return err
	}
	return nil
}

func (g *GitService) CreateRemote(ctx context.Context, repo *git.Repository, remoteName, repourl string) (*git.Remote, error) {
	remote, err := repo.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{repourl},
	})
	if err != nil {
		// Add logs: "error creGitServiceating remote

		return nil, err
	}
	return remote, nil
}

func (g *GitService) Checkout(ctx context.Context, repo *git.Repository, branch, commitHash string) error {
	workingTree, err := repo.Worktree()
	if err != nil {
		// Add Logs: "error getting worktree"
	}
	err = workingTree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(string(branch)),
	})
	if err != nil {
		// Add Logs: error checking out branch
	}
	return nil
}

func InitGitServices() Services {
	return &GitService{}
}
