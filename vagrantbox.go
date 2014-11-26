package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type vagrantbox struct {
	id          int
	description string
	provider    string
	size        string
	link        string
	title       string
}

func checkVagrant() {
	if _, err := exec.Command("which", "vagrant").Output(); err != nil {
		fmt.Println("You need to install vagrant first...")
		fmt.Println("Please check this site: https://www.vagrantup.com")
		os.Exit(0)
	}
}

func callVagrant(args ...string) {
	if _, err := exec.Command("vagrant", args...).Output(); err != nil {
		log.Fatal(err)
	}
}

func listVBoxes() []vagrantbox {
	fmt.Println("Listing all vagrant boxes...")
	doc, err := goquery.NewDocument("http://www.vagrantbox.es")
	if err != nil {
		fmt.Println("Problems to request http://www.vagrantbox.es.")
		log.Fatal(err)
	}
	re := regexp.MustCompile("[\n\t ]{1,}")
	vboxes := []vagrantbox{}
	doc.Find("#dataTable tbody tr").Each(func(i int, s *goquery.Selection) {
		data := s.Find("td")
		vbox := vagrantbox{
			id:          i,
			description: strings.TrimSpace(re.ReplaceAllString(data.Eq(0).Text(), " ")),
			provider:    data.Eq(1).Text(),
			link:        data.Eq(2).Text(),
			size:        data.Eq(3).Text(),
		}
		fmt.Printf("ID: %d\n", vbox.id)
		fmt.Printf("Description: %s\n", vbox.description)
		fmt.Printf("Size: %s MB | Provider: %s\n", vbox.size, vbox.provider)
		fmt.Printf("Link: %s\n", vbox.link)
		fmt.Println("==================================================")
		vboxes = append(vboxes, vbox)
	})
	return vboxes
}

func chooseVbox(vboxes []vagrantbox) vagrantbox {
	var id int
	fmt.Print("Choose vagrant box id to install: ")
	fmt.Scanf("%d", &id)
	vbox := vboxes[id]
	fmt.Printf("Your vagrant box is:\n%s\n", vbox.description)
	fmt.Printf("Size: %s MB | Provider: %s\n", vbox.size, vbox.provider)
	fmt.Printf("Link: %s\n", vbox.link)
	fmt.Print("Vagrant box title: ")
	fmt.Scanf("%s", &vbox.title)
	return vbox
}

func main() {
	checkVagrant()
	vboxes := listVBoxes()
	vbox := chooseVbox(vboxes)
	fmt.Println("==================================================")
	fmt.Println("Installing choosed vagrant box...")
	fmt.Printf("vagrant box add %s %s\n", vbox.title, vbox.link)
	callVagrant("box", "add", vbox.title, vbox.link)
	fmt.Printf("vagrant init %s\n", vbox.title)
	callVagrant("init", vbox.title)
	fmt.Printf("vagrant up\n")
	callVagrant("up")
	fmt.Println("Your box is up and running...see ya!")
}
