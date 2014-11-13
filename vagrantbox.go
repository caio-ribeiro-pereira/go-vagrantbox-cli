package main

import (
  "log"
  "fmt"
  "os"
  "github.com/codegangsta/cli"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  app := cli.NewApp()
  app.Name = "vagrantbox"
  app.Usage = "CLI to list and install vagrant box from http://www.vagrantbox.es"
  app.Commands = []cli.Command{
    {
      Name: "list",
      ShortName: "ls",
      Usage: "list all boxes",
      Action: func(c *cli.Context) {
        doc, err := goquery.NewDocument("http://www.vagrantbox.es")
        if err != nil {
          log.Fatal(err)
        }
        fmt.Println("listing...")
        doc.Find("#dataTable tbody tr").Each(func(i int, s *goquery.Selection) {
          description := s.Get(0).Text()
          provider := s.Get(1).Text()
          fmt.Println("%d) %s = %s", i, description, provider)
        })
      },
    },
    {
      Name: "install",
      ShortName: "inst",
      Usage: "install a box given a title and url",
      Action: func(c *cli.Context) {
        if len(c.Args()) == 2 {
          fmt.Println("installing...", c.Args()[0], c.Args()[1])  
        } else {
          fmt.Println("A title and url is needed!")
        }
      },
    },
  }
  app.Run(os.Args)
}