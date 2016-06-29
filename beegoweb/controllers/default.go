package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type BlaController struct {
	beego.Controller
}

type TabController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["AdminName"] = "talset"
	c.TplName = "index.tpl"
}

func (c *BlaController) Get() {
	c.Data["AdminName"] = "talset"
	c.TplName = "json.tpl"
	mystruct := "{bla: foo}"
  c.Data["jsonp"] = mystruct
  c.ServeJSONP()
	}

	func (c *TabController) Get() {
		c.TplName = "tab.tpl"
    // Array
		intarray := [6]int{2, 3, 5, 7, 11, 13}
		c.Data["intarray"] = intarray

		// Slice
		intslice := []int{3, 5, 7, 13}
		intslice = append(intslice, 22, 12, 56)
		c.Data["intslice"] = intslice

    // Slice of slice
		sliceofslice := [][]string{
			[]string{"1a", "2a", "3a"},
			[]string{"1b", "2b", "3b"},
			[]string{"1c", "2c", "3c"},
		}
		c.Data["sliceofslice"] = sliceofslice

		// map
		mymap := make(map[string]int)
		mymap["foo"] = 2
		mymap["bar"] = 4
		mymap["bla"] = 9
		c.Data["map"] = mymap

		// structmap
		type Vertex struct {
			Lat, Long float64
	  }

		var structmap = map[string]Vertex{
			"Bell Labs": {40.68433, -74.39967},
			"Google":    {37.42202, -122.08408},
		}
		c.Data["structmap"] = structmap

	}
