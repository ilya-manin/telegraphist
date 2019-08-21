# Telegraphist
[![CI status](https://github.com/xamut/telegraphist/workflows/CI/badge.svg)](https://github.com/xamut/telegraphist/actions)

Bindings for Telegram Bot API on Go

## Documentation

* Official Telegram Bot API [documentation](https://core.telegram.org/bots/api)

## Usage

```golang
package main

import (
	"encoding/json"
	"fmt"

	Telegraphist "github.com/xamut/telegraphist"
)

func main() {
	telegraphist, err = Telegraphist.New(&Telegraphist.Config{
		BotToken: "<YOUR_BOT_TOKEN>",
	})

	if err != nil {
		fmt.Println(err)
	}

	user, err := telegraphist.GetMe()
	if err != nil {
		fmt.Println(err)
	}

	userJSON, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println(string(userJSON))
}

```

## License

Telegraphist is released under the [MIT License](https://opensource.org/licenses/MIT).
