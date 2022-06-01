package converter

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type ArtInfo struct {
	title, author string
	tags          []string
}

func MDtoHTML(filename string)(string,error){
	var mainSec string

	f,err:=os.OpenFile(filename,os.O_RDONLY,0644)
	if err!=nil{
		return "",err
	}
	defer f.Close()
	

	//前三行：文章名、作者、标签
	scn := bufio.NewScanner(f)
	line := 1
	ai := ArtInfo{}
	for ; line <= 3 && scn.Scan(); line++ {
		text := scn.Text()
		if strings.HasPrefix(text, "[[title]]") {
			ai.title = strings.TrimPrefix(text, "[[title]]")
		} else if strings.HasPrefix(text, "[[author]]") {
			ai.author = strings.TrimPrefix(text, "[[author]]")
		} else if strings.HasPrefix(text, "[[tags]]") {
			tags := strings.TrimPrefix(text, "[[tags]]")
			ai.tags = strings.Split(tags, ",")
		}
	}
	y, m, d := time.Now().Date()
	mainSec += `
		<h1 class="titleHeader">` + ai.title + `</h1>
		<div id="info">
			<img class="avatar" src="" alt="">
			<p class="author tidy" onclick="window.open('https://blog.gool.work/u/` + ai.author + `')">` + ai.author + `</p>
			<p class="modified tidy">` + strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d) + `</p>
			<img class="tagIcon" src="" alt="">
	`
	for _, v := range ai.tags {
		mainSec += `
			<div class="tags">` + v + `</div>
		`
	}
	mainSec += `</div>`

	//正餐部分
	for scn.Scan() {
		t := scn.Text()
		if t == "" {
			continue
		}
		if strings.HasPrefix(t, "# ") {
			mainSec += `
				<div class="header1">
					<img src="" alt="">
					<h1>` + strings.TrimPrefix(t, "# ") + `</h1>
				</div>
			`
		} else if strings.HasPrefix(t, "## ") {
			mainSec += `
				<h2 class="header2">` + strings.TrimPrefix(t, "## ") + `</h2>
			`
		} else if strings.HasPrefix(t, "### ") {
			mainSec += `
				<h2 class="header3">` + strings.TrimPrefix(t, "### ") + `</h2>
			`
		} else if strings.HasPrefix(t, "[[clipBlock]]") {
			linkTitle, link, _ := strings.Cut(strings.TrimPrefix(strings.TrimSuffix(t, ")"), "[[clipBlock]]["), "](")
			linkType := "unknown"
			if strings.HasSuffix(link, ".zip") {
				linkType = "zip"
			}
			mainSec += `
				<div class="clipBlock">
					<img class="` + linkType + `" src="" alt="">
					<p class="tidy" onclick="window.open('` + link + `','_blank')">` + linkTitle + `</p>
				</div>
			`
		} else if strings.HasPrefix(t, "![") {
			linkTitle, link, _ := strings.Cut(strings.TrimPrefix(strings.TrimSuffix(t, ")"), "!["), "](")
			mainSec += `
				<img class="blockImg" src="` + link + `" alt="` + linkTitle + `">
			`
		}else if strings.HasPrefix(t,"[[blockVideo]]"){
			linkTitle, link,_:=strings.Cut(strings.TrimPrefix(strings.TrimSuffix(t,")"),"[[blockVideo]]["),"](")
			mainSec+=`
				<video class="blockVideo" src="`+link+`" title="`+linkTitle+`" controls preload="auto"></video>
			`
		} else {
			mainSec += `
				<p class="normalP">` + t + `</p>
			`
		}
	}
	return mainSec,nil
}
