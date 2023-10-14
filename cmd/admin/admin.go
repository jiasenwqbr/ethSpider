package admin

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"os"
)

var log = logging.Logger("admin-cmd")

var AdminCmd = &cli.Command{
	Name:  "admin",
	Usage: "Commands for remotely taking admin related actions",
	Subcommands: []*cli.Command{
		loginCmd,
		userCmd,
	},
}

var loginCmd = &cli.Command{
	Name:  "login",
	Usage: "Login to remote client",
	Action: func(cctx *cli.Context) error {
		ethUrl := os.Getenv("ETH_URL")
		log.Warn(ethUrl)

		return nil
	},
}

var userCmd = &cli.Command{
	Name:  "users",
	Usage: "Create user, edit permission, or delete",
	Subcommands: []*cli.Command{
		createCmd,
		listCmd,
	},
}

var createCmd = &cli.Command{
	Name:  "create",
	Usage: "Create user by admin",
	Action: func(cctx *cli.Context) error {
		log.Info("create")
		return nil
	},
}

var listCmd = &cli.Command{
	Name:  "list",
	Usage: "Show user info",
	Action: func(cctx *cli.Context) error {
		log.Info("list")
		return nil
	},
}
