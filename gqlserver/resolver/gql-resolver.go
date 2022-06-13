package resolver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GqlResolver struct {
	dummy *dummy
}

type todo struct {
	Id     int32
	Title  string
	Body   string
	Status string
}

type todos struct {
	Data []todo
}

type TodosResolver struct {
	todos *todos
}

func (r *Resolver) Todos() (*TodosResolver, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/todos", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return &TodosResolver{&todos{}}, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &TodosResolver{&todos{}}, fmt.Errorf(err.Error())
	}

	var data todos
	json.Unmarshal(bodyBytes, &data)

	return &TodosResolver{&todos{
		Data: data.Data,
	}}, nil
}

func (r *TodosResolver) Data() *[]*TodoResolver {
	list := make([]*TodoResolver, len(r.todos.Data))
	for i := range r.todos.Data {
		list[i] = &TodoResolver{&r.todos.Data[i]}
	}

	return &list
}

type TodoResolver struct {
	todo *todo
}

func (r *Resolver) Todo(args struct {
	Id int32
}) (*TodoResolver, error) {
	client := &http.Client{}
	url := "http://localhost:8080/todo/" + fmt.Sprint(args.Id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &TodoResolver{&todo{}}, fmt.Errorf(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return &TodoResolver{&todo{}}, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	currentTodo := &todo{}
	json.Unmarshal(bodyBytes, &currentTodo)

	return &TodoResolver{currentTodo}, nil
}

func (r *TodoResolver) Id() *int32 {
	return &r.todo.Id
}

func (r *TodoResolver) Title() *string {
	return &r.todo.Title
}

func (r *TodoResolver) Body() *string {
	return &r.todo.Body
}

func (r *TodoResolver) Status() *string {
	return &r.todo.Status
}

func (r *Resolver) AddTodo(args struct {
	Title  string
	Body   string
	Status string
}) (*TodoResolver, error) {
	postBody, _ := json.Marshal(map[string]string{
		"title":  args.Title,
		"body":   args.Body,
		"status": args.Status,
	})
	requestBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/todo", requestBody)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return &TodoResolver{&todo{}}, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &TodoResolver{&todo{}}, fmt.Errorf(err.Error())
	}

	currentTodo := &todo{}
	json.Unmarshal(bodyBytes, &currentTodo)

	return &TodoResolver{currentTodo}, nil
}
