package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type PersonController struct {
	beego.Controller
}

type Person struct {
	ID        string   `json:"id,omitempty";form:"-"`
	Firstname string   `json:"firstname,omitempty";form:"firstname"`
	Lastname  string   `json:"lastname,omitempty";form:"lastname"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty";form:"city"`
	State string `json:"state,omitempty";form:"state"`
}

var people []Person

func (controller *PersonController) CreatePerson() {
	firstname := controller.Ctx.Request.Form.Get("firstname")
	lastname := controller.Ctx.Request.Form.Get("lastname")
	addressCity := controller.Ctx.Request.Form.Get("address-city")
	addressState := controller.Ctx.Request.Form.Get("address-state")

	newId := strconv.FormatInt(int64(len(people)+1), 10)

	people = append(people, Person{ID: newId, Firstname: firstname, Lastname: lastname, Address: &Address{City: addressCity, State: addressState}})

	controller.GetPeople()
}

func (controller *PersonController) GetPeople() {
	controller.Data["json"] = people
	controller.ServeJSON()
}

func (controller *PersonController) GetPerson() {
	id := controller.Ctx.Input.Param(":id")

	fmt.Printf("Llegó el id %s\n", id)

	var personReturn []Person

	for _, value := range people {
		if value.ID == id {
			personReturn = append(personReturn, value)
			break
		}
	}

	controller.Data["json"] = personReturn
	controller.ServeJSON()
}

func (controller *PersonController) UpdatePerson() {
	id := controller.Ctx.Input.Param(":id")

	firstname := controller.Ctx.Request.Form.Get("firstname")
	lastname := controller.Ctx.Request.Form.Get("lastname")
	addressCity := controller.Ctx.Request.Form.Get("address-city")
	addressState := controller.Ctx.Request.Form.Get("address-state")

	for i, value := range people {
		if value.ID == id {
			people[i].Firstname = firstname
			people[i].Lastname = lastname
			people[i].Address = &Address{
				City:  addressCity,
				State: addressState,
			}
			break
		}
	}

	//for _, value := range people {
	//	fmt.Println(value)
	//}

	controller.GetPeople()
}

func (controller *PersonController) DeletePerson() {
	id := controller.Ctx.Input.Param(":id")

	for index, value := range people {
		if value.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	controller.GetPeople()
}

func main() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	beego.Router("/people", &PersonController{}, "post:CreatePerson")
	beego.Router("/people", &PersonController{}, "get:GetPeople")
	beego.Router("/people/:id:int", &PersonController{}, "get:GetPerson")
	beego.Router("/people/:id:int", &PersonController{}, "put:UpdatePerson")
	beego.Router("/people/:id:int", &PersonController{}, "delete:DeletePerson")
	beego.Run()
}
