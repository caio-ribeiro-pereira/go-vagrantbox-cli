package main

import (
  "strings"
  "regexp"
  "log"
  "fmt"
  "os"
  "os/exec"
  "github.com/codegangsta/cli"
  "github.com/PuerkitoBio/goquery"
)

type vagrantbox struct {
  description string
  provider string
  size string
  link string
}

func main() {
  _, path_err := exec.Command("export PATH=/usr/bin:/bin:/usr/local/bin:$PATH").Output()
  if path_err != nil {
    log.Fatal(path_err)
    return
  }
  _, err := exec.Command("which vagrant").Output()
  if err != nil {
    fmt.Println("You need to install vagrant first...\nPlease check this site: https://www.vagrantup.com")
    log.Fatal(err)
    return
  }
  app := cli.NewApp()
  app.Name = "vagrantbox"
  app.Usage = "CLI to install vagrant box from http://www.vagrantbox.es"
  app.Action = func(c *cli.Context) {
    fmt.Println("Please wait, we are listing...")
    doc, err := goquery.NewDocument("http://www.vagrantbox.es")
    if err != nil {
      fmt.Println("Problems to access the http://www.vagrantbox.es page.")
      log.Fatal(err)
      return
    }
    var choice int
    var title string
    vboxes := []vagrantbox{}
    re := regexp.MustCompile("[\n\t ]{1,}")
    doc.Find("#dataTable tbody tr").Each(func(i int, s *goquery.Selection) {
      data := s.Find("td")
      vbox := vagrantbox{}
      vbox.description = strings.TrimSpace(re.ReplaceAllString(data.Eq(0).Text(), " "))
      vbox.provider = data.Eq(1).Text()
      vbox.link = data.Eq(2).Text()
      vbox.size = data.Eq(3).Text()
      fmt.Printf("%d) %s\n", i, vbox.description)
      fmt.Printf("Size: %s MB | Provider: %s\n", vbox.size, vbox.provider)
      fmt.Printf("Link: %s\n", vbox.link)
      fmt.Println("==================================================")
      vboxes = append(vboxes, vbox)
    })
    fmt.Print("Choose a number to install a box: ")
    fmt.Scanf("%d", &choice)
    vbox := vboxes[choice]
    fmt.Printf("Ok! This is your box:\n%s\n", vbox.description)
    fmt.Printf("Size: %s MB | Provider: %s\n", vbox.size, vbox.provider)
    fmt.Printf("Link: %s\n", vbox.link)
    fmt.Print("What will be the box title?: ")
    fmt.Scanf("%s", &title)
    fmt.Println("==================================================")
    fmt.Println("Installing choosed box...")
    cmd := fmt.Sprintf("vagrant box add %s %s && vagrant init %s && vagrant up",title, vbox.link, title)
    out, err := exec.Command(cmd).Output()
    if err != nil {
      log.Fatal(err)
      return
    }
    fmt.Println(out)
  }
  app.Run(os.Args)
}