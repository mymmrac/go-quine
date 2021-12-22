package main

import (
	"fmt"
	"strings"
)

func main() {
	code := "package main\n\nimport (\n	\"fmt\"\n	\"strings\"\n)\n\nfunc main() {\n	code := \"?\"\n	escapedCode := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(\n		code, `\\`, `\\\\`), `\"`, `\\\"`), \"\\n\", `\\n`)\n	codeSamples := strings.Split(code, string(rune(63)))\n	fmt.Print(codeSamples[0], escapedCode, codeSamples[1])\n}\n"
	escapedCode := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
		code, `\`, `\\`), `"`, `\"`), "\n", `\n`)
	codeSamples := strings.Split(code, string(rune(63)))
	fmt.Print(codeSamples[0], escapedCode, codeSamples[1])
}
