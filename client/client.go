package client

import (
	"github.com/desertbit/grumble"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	zcdApp := cli.NewApp()
	zcdApp.Name = "core"
	zcdApp.Usage = "A PKI File Share System"
	zcdApp.Action = func(context *cli.Context) error {
		zcd()
		return nil
	}

	if err := zcdApp.Run(os.Args); err != nil {
		return
	}
}

func zcd() {
	app := grumble.New(&grumble.Config{
		Name:        "core",
		Description: "core Client Tools",
	})

	app.AddCommand(CreateUploadCommand())
	app.AddCommand(CreateDownloadCommand())
	app.AddCommand(CreateListCommand())
	grumble.Main(app)
}
