package main

import "fmt"

func main() {
	mergePatch := []byte(fmt.Sprintf(`{"deploy_status":%q, "message":%q}`, "fail", "error"))
	fmt.Println(fmt.Sprintf("%s", mergePatch))
}
