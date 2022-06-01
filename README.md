# myM2H - a Markdown to HTML converter for my blog website.
为我的博客生成html元素，搭配模板以是博客撰写更加便捷
中文说明请参考[sidewis/myM2HDocumentation](https://doc.gool.work/myM2HDocumentation/)

## Installation Guide
### Install go in your local machine of run it as a executable (which requires remote server deployed).
For Linux users, all you need to do is to download the source.zip file on which tag you prefer. Or just copy
```
https://github.com/jcbl1/myM2H.git
```
and run
```bash
git clone https://github.com/jcbl1/myM2H.git
```
in your terminal.
### Install from source code
Let's first check your go version by running
```bash
go version
```
which will properly result something like
```bash
$ go version
go version go1.18.2 linux/amd64
```
At this spot, you head to the directory in which the source.zip file is downloaded. And unzip the file to a folder named "myM2H".
Go ahead to the target directory. And run 
```bash
go get ./...&&go install -f
```
This will not produce any output but don't worry. That is what was supposed to be.
Here you go. You have installed myM2H on your machine and are ready to go.



## Easy Tutorial
With myM2H, you can do pretty much everything related to Markdown-to-HTML format converting work. And in this tutorial, we are just focusing on some easy points of it.

### Convert [lorem].md to [lorem].html 
![Screenshot showing what a classic markdown file looks like](https://gool.work/r/ss2022-06-01-001.png)

