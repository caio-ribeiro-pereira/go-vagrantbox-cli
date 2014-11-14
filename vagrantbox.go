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
  _, err := exec.Command("which", "vagrant").Output()
  if err != nil {
    fmt.Println("You need to install vagrant first...\nPlease check this site: https://www.vagrantup.com")
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
    fmt.Printf("vagrant box add %s %s\n", title, vbox.link)
    _, vb_add_err := exec.Command("vagrant", "box", "add", title, vbox.link).Output()
    if vb_add_err != nil {
      log.Fatal(vb_add_err)
      return
    }
    // vagrant init {title}
    fmt.Printf("vagrant init %s\n", title)
    _, vb_init_err := exec.Command("vagrant", "init", title).Output()
    if vb_init_err != nil {
      log.Fatal(vb_init_err)
      return
    }
    // vagrant up
    fmt.Printf("vagrant up\n")
    _, vb_up_err := exec.Command("vagrant", "up").Output()
    if vb_up_err != nil {
      log.Fatal(vb_up_err)
      return
    }
    fmt.Println("Your box is up and running...see ya!")
  }
  app.Run(os.Args)
}