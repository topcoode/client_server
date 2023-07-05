// package main

// import (
// 	"fmt"
// 	"frontend/Database"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq"
// )

// type Data struct {
// 	MCC            int    `form:"mcc" db:"mcc" `
// 	MNC            int    `form:"mnc" db:"mnc"`
// 	IMEI           int    `form:"imei" db:"imei" `
// 	EquimentStatus string `form:"equimentstatus" db:"equimentstatus"`
// }

// func main() {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("C:/Program Files/Go/src/html/template//*.html")

// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", nil)
// 		fmt.Println(c)

// 	})

// 	router.POST("/plmnierpost", PlmnierPOST)
// 	router.GET("/plmnierget", PlmnierGET)
// 	router.PUT("/plmnierupdate", PlmnierUPDATE)
// 	router.DELETE("/plmnierdelete", PlmnierDELETE)
// 	router.Run(":8080")

// }
// func PlmnierPOST(c *gin.Context) {
// 	//PLMNEIR-----------
// 	c.HTML(http.StatusOK, "plmneir.html", gin.H{})
// 	data := &Data{}
// 	if err := c.ShouldBind(data); err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	fmt.Println("plmneir")
// 	db, err := Database.Dbconnection()
// 	insertQuery := `INSERT INTO plmneir (mcc,mnc,imei,equimentstatus) VALUES ($1,$2,$3,$4)`
// 	_, err = db.Exec(insertQuery, data.MCC, data.MNC, data.IMEI, data.EquimentStatus)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	c.String(http.StatusOK, "User registered successfully!")

// }
// func PlmnierGET(c *gin.Context) {
// 	GetUserdata()
// }
// func GetUserdata() []Data {
// 	var user []Data
// 	db, _ := Database.Dbconnection()
// 	rows, err := db.Query("SELECT * FROM plmneir")
// 	if err != nil {
// 		fmt.Println("Failed to execute the query:", err)
// 	}
// 	for rows.Next() {
// 		var result Data
// 		err := rows.Scan(&result.MCC, &result.MNC, &result.IMEI, &result.EquimentStatus)
// 		if err != nil {
// 			fmt.Println("Failed to scan row:", err)
// 		}
// 		user = append(user, result)
// 		fmt.Println("user.....>", user)

// 	}

//		if err = rows.Err(); err != nil {
//			fmt.Println("Error retrieving query results:", err)
//		}
//		//c.JSON(http.StatusOK, user)
//		return user
//	}
//
//	func PlmnierUPDATE(c *gin.Context) {
//		var user Data
//		if err := c.BindJSON(&user); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		db, err := Database.Dbconnection()
//		if err != nil {
//			fmt.Println("error in the db connection!!")
//		}
//		query := `UPDATE plmneir SET mcc=$1, mnc=$2 WHERE imei=$3;`
//		_, err = db.Exec(query, user.MCC, user.MNC, user.IMEI)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//
//	func PlmnierDELETE(c *gin.Context) {
//		var user Data
//		if err := c.BindJSON(&user); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		db, err := Database.Dbconnection()
//		if err != err {
//			fmt.Println(err)
//		}
//		query := `DELETE FROM plmneir WHERE imei = $1;`
//		_, err = db.Exec(query, user.IMEI)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
package main

import "fmt"

type Struct1 struct {
	Data1 string
	Data2 string
}
type Struct2 struct {
	Data3   string
	Data4   string
	struct1 []Struct1
}

func main() {
	Struct1data := []Struct1{
		{Data1: "asdfasf", Data2: "asdsaf"}, {Data1: "asfaf", Data2: "sdafas"},
	}
	datavar := Struct2{
		Data3:   "asff",
		Data4:   "adsfas",
		struct1: Struct1data,
	}
	fmt.Println(datavar)
	for _, data := range datavar.struct1 {
		fmt.Println(data)
	}

}
