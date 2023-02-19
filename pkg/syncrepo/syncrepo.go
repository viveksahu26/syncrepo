package syncrepo

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/viveksahu26/syncrepo/config"
	"github.com/viveksahu26/syncrepo/pkg/git"
	"github.com/viveksahu26/syncrepo/types"
)

type syncRepoService struct {
	GitService git.Services
}

func (s syncRepoService) SyncRepo(ctx context.Context, event *types.PushEventGitlab) error {
	// get the reference and split them. For example: "refs/heads/master"
	ref := strings.Split(event.Ref, "/")

	// get the branch
	targetBranch := ref[len(ref)-1]

	// check the push event is for main branch or not
	if targetBranch != "main" {
		// push happened to branch other than main or master
		// so no need to respond
		// Add Log: TODO
		return nil
	}
	// so push event is on right branch.

	// get SyncRepoConfig
	syncrepoconfig := config.SyncRepoConfig{}

	// Now clone that repository in memory
	tempRepoDir, err := ioutil.TempDir("", event.Repository.Name)
	fmt.Println("tempRepoDir: ", tempRepoDir)
	if err != nil {
		// TODO: Add logs error creating tempdi
	}

	// once the use of this temp dir being completed,
	// then remove it or clean up
	defer CleanTempDir(tempRepoDir)

	// Now clone that repo in tempRepoDir path
	repository, err := s.GitService.Clone(ctx, tempRepoDir, syncrepoconfig.RepoURL, "oauth2", syncrepoconfig.Token)
	if err != nil {
		return err
	}

	err = s.GitService.Checkout(ctx, repository, "main", event.CheckoutSha)
	if err != nil {
		return err
	}

	for _, followerConfig := range syncrepoconfig.FollowersRepoConfig {

		// create a remote in cloned repo
		_, err := s.GitService.CreateRemote(ctx, repository, followerConfig.RemoteName, followerConfig.RepoURL)
		if err != nil {
			return err
		}

		// push the changes to follower repo. or sync main repo with follower repo
		err = s.GitService.Push(ctx, repository, followerConfig.RemoteName, followerConfig.UserName, followerConfig.Token)
		if err != nil {
			return err
		}
	}
	return err
}

func CleanTempDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		// TODO: Add Logs error removing tempdir"
	}
}

func InitSyncRepoServices(gitService git.Services) *syncRepoService {
	return &syncRepoService{
		GitService: gitService,
	}
}
