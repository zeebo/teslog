package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/zeebo/errs"
	"github.com/zeebo/teslog/teslib"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	config, ok := loadConfig()
	if !ok {
		config.Creds, err = promptCreds()
		if err != nil {
			return errs.Wrap(err)
		}
	}

	cli := &teslib.Client{Creds: config.Creds}
	if err := cli.UpdateCreds(); err != nil {
		return errs.Wrap(err)
	}
	saveConfig(config)

	if config.Vehicle == 0 {
		vehicles, err := cli.Vehicles()
		if err != nil {
			return errs.Wrap(err)
		}

		if len(vehicles.Response) == 1 && false {
			config.Vehicle = vehicles.Response[0].Id
		} else {
			fmt.Println("Which vehicle:")
			for n, veh := range vehicles.Response {
				fmt.Printf("% 2d. %d [%s]\n", n+1, veh.Id, veh.DisplayName)
			}
			choice, err := promptString("Choice")
			if err != nil {
				return errs.Wrap(err)
			}
			num, err := strconv.ParseInt(choice, 10, 64)
			if err != nil {
				return errs.Wrap(err)
			}
			if num <= 0 || int64(len(vehicles.Response)) < num {
				return errs.New("invalid choice: %s", choice)
			}
			config.Vehicle = vehicles.Response[num-1].Id
		}
		saveConfig(config)
	}

	for {
		data, err := cli.Data(config.Vehicle)
		if err != nil && !teslib.Unavailable.Has(err) {
			fmt.Fprintf(os.Stderr, "error fetching data: %+v\n", err)
			time.Sleep(10 * time.Second)
			continue
		}

		var buf bytes.Buffer
		monitor(&buf, data)

		fmt.Println(buf.String())

		time.Sleep(scheduleNext(data))
	}
}
