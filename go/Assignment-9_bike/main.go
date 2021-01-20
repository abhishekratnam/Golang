package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Cycle struct {
	id      int `json:"id"`
	frame   int `json:"frame"`
	handle  int `json:"handle"`
	seating int `json:"seating"`
	wheels  int `json:"wheels"`
	chain   int `json:"chain"`
	price   int `json:"price"`
	// ID        uint   `json:"id"`
	// FirstName string `json:"firstname"`
	// LastName  string `json:"lastname"`
	// City      string `json:"city"`
}
type user struct {
	DataValue []*Cycle `json:"cycle"`
}

func main() {

	// NOTE: See we're using = to assign the global var
	//         	instead of := which would assign it only in this function
	db, err = gorm.Open("sqlite3", "./gorm.db")
	//db, _ = gorm.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Cycle{})

	r := gin.Default()
	r.GET("/cycle/", GetCycle)
	r.GET("/cycle/:id", GetSpecificCycle)
	r.POST("/cycle", CreateCycle)
	r.PUT("/cycle/:id", UpdateCycle)
	r.DELETE("/cycle/:id", DeleteCycle)
	r.Run(":8080")
}

func (c *user) CalPrice() *user {
	for index := range c.DataValue {
		c.DataValue[index].price = c.DataValue[index].chain + c.DataValue[index].frame + c.DataValue[index].handle + c.DataValue[index].seating + c.DataValue[index].wheels
	}
	return c
}
func DeleteCycle(c *gin.Context) {
	id := c.Params.ByName("id")
	var cycle Cycle
	d := db.Where("id = ?", id).Delete(&cycle)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateCycle(c *gin.Context) {

	var cycle Cycle
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&cycle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cycle)

	db.Save(&cycle)
	c.JSON(200, cycle)

}

func CreateCycle(c *gin.Context) {

	var cycle Cycle
	c.BindJSON(&cycle)
	db.Create(&cycle)
	c.JSON(200, cycle)
}

func GetSpecificCycle(c *gin.Context) {
	id := c.Params.ByName("id")
	var cycle Cycle
	if err := db.Where("id = ?", id).First(&cycle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cycle)
	}
}
func GetCycle(c *gin.Context) {
	var cycle []Cycle
	if err := db.Find(&cycle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cycle)
	}

}
