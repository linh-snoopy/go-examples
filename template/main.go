package main

import (
    "html/template"
    "os"
)

// Links:
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html

type Friend struct {
    Fname string
}

type Person struct {
    UserName string
    Emails   []string
    Friends  []*Friend
}

func main() {
    f1 := Friend{Fname: "minux.ma"}
    f2 := Friend{Fname: "xushiwei"}
    t := template.New("fieldname example")
    t, _ = t.Parse(`hello {{.UserName}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}
            {{with .Friends}}
            {{range .}}
                my friend name is {{.Fname}}
            {{end}}
            {{end}}
            `)
    p := Person{UserName: "Astaxie",
        Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
        Friends: []*Friend{&f1, &f2}}
    t.Execute(os.Stdout, p)
	
	///////////////////////
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} will not be outputted. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)
	
	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("Not empty pipeline if demo: {{if `anything`}} will be outputted. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)
	
	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if part {{else}} else part.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}