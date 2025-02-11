
# goAvatar

Go package for quick generation of random pixelized avatars



# Example usage

```go
package main

import (
	"github.com/ethercod3/goAvatar"
)

func main() {
	options := goAvatar.AvatarOptions{
		Filepath:   "./avatar.png",
		Dimensions: 5,
		FileSizePx: 500,
	}
	goAvatar.GenerateAvatar(options)
}
```
