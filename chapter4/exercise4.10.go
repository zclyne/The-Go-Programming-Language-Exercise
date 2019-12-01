// Prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"the-go-programming-language/chapter4/github"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	for _, item := range result.Items {
		if time.Since(item.CreatedAt).Hours() < 30 * 24 { // 不到一个月的issues，一个月算30天
			fmt.Printf("不到一个月的issue：%s\n", item.Title)
		} else if time.Since(item.CreatedAt).Hours() < 365 * 24 { // 不到一年的issue，一年算365天
			fmt.Printf("不到一年的issue：%s\n", item.Title)
		} else { // 超过一年的issue
			fmt.Printf("超过一年的issue：%s\n", item.Title)
		}
	}
}