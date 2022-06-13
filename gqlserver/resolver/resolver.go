package resolver

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var currentUsers []*user

type Resolver struct{}

type dummy struct {
	Message string
	Status  string
}

func (r *Resolver) Dummy() (*DummyResolver, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	//var responseObject Response
	//json.Unmarshal(bodyBytes, &responseObject)
	//fmt.Printf("API Response as struct %+v\n", responseObject)
	fmt.Printf("API Response as struct %+v\n", bodyBytes)

	return &DummyResolver{&dummy{
		Message: string(bodyBytes),
		Status:  "200 OK",
	}}, nil
}

type DummyResolver struct {
	dummy *dummy
}

// Message field resolver
func (r *DummyResolver) Message() *string {
	return &r.dummy.Message
}

// Status field resolver
func (r *DummyResolver) Status() *string {
	return &r.dummy.Status
}

type user struct {
	Id         int32
	Name       string
	Email      string
	profession string
}

type users struct {
	Data []user
}

type UsersResolver struct {
	users *users
}

func (r *Resolver) Users() (*UsersResolver, error) {
	var allUsers []user

	for _, u := range currentUsers {
		allUsers = append(allUsers, *u)
	}

	return &UsersResolver{&users{
		Data: allUsers,
	}}, nil
}

func (r *UsersResolver) Data() *[]*UserResolver {
	list := make([]*UserResolver, len(r.users.Data))
	for i := range r.users.Data {
		list[i] = &UserResolver{&r.users.Data[i]}
	}

	return &list
}

type UserResolver struct {
	user *user
}

func (r *Resolver) User(args struct {
	Id int32
}) (*UserResolver, error) {
	for _, u := range currentUsers {
		if u.Id == args.Id {
			return &UserResolver{u}, nil
		}
	}

	return &UserResolver{&user{}}, fmt.Errorf("User not found")
}

func (r *UserResolver) Id() *int32 {
	return &r.user.Id
}

func (r *UserResolver) Name() *string {
	return &r.user.Name
}

func (r *UserResolver) Email() *string {
	return &r.user.Email
}

func (r *UserResolver) Profession() *string {
	return &r.user.profession
}

func (r *Resolver) AddUser(args struct {
	Name       string
	Email      string
	Profession string
}) (*UserResolver, error) {
	newUser := &user{
		Id:         int32(len(currentUsers)),
		Name:       args.Name,
		Email:      args.Email,
		profession: args.Profession,
	}

	currentUsers = append(currentUsers, newUser)

	return &UserResolver{user: newUser}, nil
}
