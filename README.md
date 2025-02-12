
# goAvatar

Go package for quick generation of random pixelized avatars

![examples](https://github.com/ethercod3/goAvatar/blob/main/examples/examples.png?raw=true)

# Example usage

```go
package main

import (
	"github.com/ethercod3/goAvatar"
	"log"
)

func main() {
	options := goAvatar.AvatarOptions{
		Dimensions: 5,
		FileSizePx: 500,
	}
	img := goAvatar.GenerateAvatar(options)
	file, err := os.Create("./avatar.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, img)
}
```
