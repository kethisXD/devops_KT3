package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strconv"
)

var tmpl = template.Must(template.New("calc").Parse(`
<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="UTF-8">
  <title>Калькулятор</title>
</head>
<body>
  <h1>КалькУлятор ( снова )</h1>
  <form method="POST" action="/">
    <input type="text" name="a" size="5" value="{{.A}}">
    <select name="op">
      <option {{if eq .Op "+"}}selected{{end}}>+</option>
      <option {{if eq .Op "-"}}selected{{end}}>-</option>
      <option {{if eq .Op "*"}}selected{{end}}>*</option>
      <option {{if eq .Op "/"}}selected{{end}}>/</option>
    </select>
    <input type="text" name="b" size="5" value="{{.B}}">
    <button type="submit">=</button>
  </form>
  {{if .Result}}
    <h2>Результат: {{.Result}}</h2>
  {{end}}
</body>
</html>
`))

type PageData struct {
    A, B, Op, Result string
}

func calc(a, b float64, op string) float64 {
    switch op {
    case "+": return a + b
    case "-": return a - b
    case "*": return a * b
    case "/":
        if b != 0 { return a / b }
    }
    return 0
}

func handler(w http.ResponseWriter, r *http.Request) {
    data := PageData{Op: "+"}
    if r.Method == http.MethodPost {
        data.A = r.FormValue("a")
        data.B = r.FormValue("b")
        data.Op = r.FormValue("op")
        af, _ := strconv.ParseFloat(data.A, 64)
        bf, _ := strconv.ParseFloat(data.B, 64)
        res := calc(af, bf, data.Op)
        data.Result = strconv.FormatFloat(res, 'f', -1, 64)
    }
    tmpl.Execute(w, data)
}

func main() {
  p := ":8080"
  http.HandleFunc("/", handler)
  fmt.Println("Сервер запущен на",p )
  http.ListenAndServe(p, nil)
}
