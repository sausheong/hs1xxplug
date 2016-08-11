# HS1XXPlug 

TP-Link has a set of interesting WiFi smart plugs -- the [HS100](http://www.tp-link.com/en/products/details/HS100.html) which provides remote access and control for a power plug, and [HS110](http://www.tp-link.com/en/products/details/cat-5258_HS110.html) which is essentially HS100 but with energy monitoring.

Both smart plugs are only accessible through the Kasa app and there is no official documentation on how to programmatically control either of them. However, they have been reverse engineered and several of these attempts have been documented:

https://www.softscheck.com/en/reverse-engineering-tp-link-hs110/
https://georgovassilis.blogspot.sg/2016/05/controlling-tp-link-hs100-wi-fi-smart.html


`hs1xxplug` is a Go library for accessing these smart plugs, based on the information available above. I have tested them out with a HS110 only and it works perfectly fine. An example of how to use this library is:

```go
package main

import (
	"fmt"
	"github.com/sausheong/hs1xxplug"
)

func main() {
	plug := hs1xxplug.Hs1xxPlug{IPAddress: "192.168.0.196"}
	results, err := plug.MeterInfo()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(results)

}
```

