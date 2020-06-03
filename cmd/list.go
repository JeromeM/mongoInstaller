/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// Variables declarations
const url string = "https://www.mongodb.com/download-center/community"
const serverSelect string = ".dl-server-select"

// Types declarations
type List []string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available versions",
	Long: `List available versions on MongoDB website.`,
	Run: func(cmd *cobra.Command, args []string) {
		listMongo(args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listMongo(args []string) {
	list := List{}
	list.fill()

	fmt.Println(list)
}

func (list *List) fill() {

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(serverSelect + " option").Each(func(i int, selection *goquery.Selection) {
		optionValue, _ := selection.Attr("value")
		if val, _ := strconv.Atoi(optionValue); val > -1 {
			*list = append(*list, selection.Text())
		}
	})
}

func (list List) String() string {
	var listStr string = ""
	for _, version := range list {
		listStr += version + "\n"
	}
	return listStr
}