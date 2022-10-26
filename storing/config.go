package storing

import (
	"fmt"
	"time"
)

var dbDriver string = "sqlite3"

var now = time.Now()
var TwitDbFileName string = fmt.Sprintf("twitter-%d%02d%02d%02d%02d.db", int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
