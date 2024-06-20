package exercises

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type DBResult struct {
	Items []*DataBase
}

type DataBase struct {
	Item  string
	Price float32
}

const DBURL = "http://localhost:8000/list"

var dbList = template.Must(template.New("dbList").Parse(`
<table>
<tr style='text-align: left'>
<th>Item</th>
<th>Price</th>
</tr>
{{range .Items}}
<tr>
<td>{{.Item}}$</td>
<td>{{.Price}}$</td>
</tr>
{{end}}
</table>
`))

func ListItem() (*DBResult, error) {
	resp, err := http.Get(DBURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var res DBResult
	var result map[string]float32
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	for item, value := range result {
		res.Items = append(res.Items, &DataBase{Item: item, Price: value})
	}
	return &res, nil
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	result, err := ListItem()
	fmt.Println(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Search query failed: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := dbList.Execute(w, result); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template: %s", err), http.StatusInternalServerError)
	}
}

func Ex712() {
	http.HandleFunc("/", dbHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
