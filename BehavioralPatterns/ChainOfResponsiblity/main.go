package main

import "fmt"

type department interface {
	execute(*patient)
	setNext(department)
}

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}



type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")

	}
	fmt.Println("Doctor checking patient")
}

func (d *doctor) setNext(next department) {
	d.next = next
}



type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
}

func main() {


	doctor := &doctor{}

	reception := &reception{}
	reception.setNext(doctor)

	Pat := &patient{name: "Bob"}
	//Pat2 := &patient{name: "Tom"}

	//Patient visiting
	reception.execute(Pat)
	//reception.execute(Pat)


}