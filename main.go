package main

import (
	"fmt"
	"strconv"

	"github.com/go-pg/pg/v10"
)

func ExampleDB_Model() {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	/* err := createSchema(db)
	    if err != nil {
	        panic(err)
	    }

	    user1 := &User{
	        Name:   "admin",
	        Emails: []string{"admin1@admin", "admin2@admin"},
	    }
	    _, err = db.Model(user1).Insert()
	    if err != nil {
	        panic(err)
	    }

	    _, err = db.Model(&User{
	        Name:   "root",
	        Emails: []string{"root1@root", "root2@root"},
	    }).Insert()
	    if err != nil {
	        panic(err)
	    }

	    story1 := &Story{
	        Title:    "Cool story",
	        AuthorId: user1.Id,
	    }
	    _, err = db.Model(story1).Insert()
	    if err != nil {
	        panic(err)
	    }

	    // Select user by primary key.
	    user := &User{Id: user1.Id}
	    err = db.Model(user).WherePK().Select()
	    if err != nil {
	        panic(err)
	    }

	    // Select all users.
	    var users []User
	    err = db.Model(&users).Select()
	    if err != nil {
	        panic(err)
	    }

	    // Select story and associated author in one query.
	    story := new(Story)
	    err = db.Model(story).
	        Relation("Author").
	        Where("story.id = ?", story1.Id).
	        Select()
	    if err != nil {
	        panic(err)
	    }

	    fmt.Println(user)
	    fmt.Println(users)
	    fmt.Println(story)
	    // Output: User<1 admin [admin1@admin admin2@admin]>
	    // [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	    // Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
	}

	// createSchema creates database schema for User and Story models.
	func createSchema(db *pg.DB) error {
	    models := []interface{}{
	        (*User)(nil),
	        (*Story)(nil),
	    }

	    for _, model := range models {
	        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
	            Temp: true,
	        })
	        if err != nil {
	            return err
	        }
	    }
	    return nil
	}
	*/
}

type Group struct {
	Name  string
	Users []User
}

type Direction struct {
	Name string
}
type User struct {
	Age       int
	Direction Direction
}

type Zal struct {
	Name   string
	Area   uint
	Groups []Group
}

type Studio struct {
	Name string
	Zals []Zal
}

func (s *Studio) AddZall(Name string, Area uint) {
	zal := Zal{Name: Name, Area: Area}
	s.Zals = append(s.Zals, zal)

}

type Cat struct {
	Name       string
	Weight     int
	JumpHeight int
	TailLenght int
}

func (c *Cat) Meow() {
	fmt.Println("Mew " + c.Name)
}
func (c *Cat) WeightCat() {
	fmt.Println("WeightCat " + c.Name + " " + strconv.Itoa(c.Weight))
}
func (c *Cat) CheckJump(ChekedHight int) {
	if c.JumpHeight >= ChekedHight {
		fmt.Println(c.Name + " запрыгнет")
	} else {
		fmt.Println(c.Name + " упадет")
	}
}

func (c *Cat) TailLenghtCat(TailLenghtMin int, TailLenghtMax int) {
	if c.TailLenght > TailLenghtMin && c.TailLenght < TailLenghtMax {
		fmt.Println(c.Name + "Хвот в допусках")
	} else {
		fmt.Println(c.Name + "Хвот НЕ в допусках")
	}
}
func main() {
	/*	s1 := Studio{Name: "первая студия"}
			s1.AddZall("Зал 1", 100)
			s1.AddZall("Зал 2", 200)

			fmt.Println(s1.Zals[0])
		barsik := Cat}
		pushok := Cat{Name: "pushok"}
		barsik.Meow()
		pushok.Meow()*/
	barsik := Cat{Weight: 10, Name: "barsik", JumpHeight: 5, TailLenght: 20}
	pushok := Cat{Weight: 8, Name: "pushok", JumpHeight: 15, TailLenght: 17}

	barsik.CheckJump(10)
	pushok.CheckJump(10)

	barsik.TailLenghtCat(222, 223)

}
