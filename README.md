# confcat
 a super lightweight configuration management package that supports configuration hot loading.

# features
+ super lightweight
+ configuration hot load
+ default value support

# usage
``` go
package main

import (
	"fmt"
    "time"
    "log"

	"github/Andrewpqc/confcat/config"
)

func main() {
	c, err := config.NewConfig("test.conf")
	if err != nil {
		log.Fatal("error to new config:%v", err)
	}

	for {
		fmt.Println(c.GetString("host", "127.0.0.1"))
		fmt.Println(c.GetInt("port", 235))
		fmt.Println(c.GetFloat("P", 3.14))
		time.Sleep(3 * time.Second)
	}
}
```
the content of `test.conf`:
```conf
# host
host=120.77.220.255

# port
port=1526

# secret_key
secret_key=this is a demo secret key

# p
P=3.14159
```
start the program and edit and save `test.conf`,then you will see the output:
```
120.77.220.255
1526
3.14159
120.77.220.255
1526
3.14159
120.77.220.255
1526
3.14159
detect config file:test.conf changed,reloading...
reload successed!
1.1.1.1
1526
3.14159
1.1.1.1
1526
3.14159
```
Enjoy.