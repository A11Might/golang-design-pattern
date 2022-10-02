package interpreter

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ExampleInterpreter 请求者(client)
func ExampleInterpreter() {
	file, _ := os.Open("./program.txt")
	buf, _ := ioutil.ReadAll(file)
	for _, line := range strings.Split(string(buf), "\n") {
		fmt.Printf("text = %s\n", line)
		node := new(ProgramNode)
		node.Parse(NewContext(line))
		fmt.Println(node.ToString())
	}

	// Output:
	// text = program end
	// [program []]
	// text = program go end
	// [program [go]]
	// text = program go right go right go right go right end
	// [program [go, right, go, right, go, right, go, right]]
	// text = program repeat 4 go right end end
	// [program [[repeat 4 [go, right]]]]
	// text = program repeat 4 repeat 3 go right go left end right end end
	// [program [[repeat 4 [[repeat 3 [go, right, go, left]], right]]]]
}
