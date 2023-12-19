package main

// 导入gin包
import (
	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	// "golang.org/x/text/message"
	"gin/password"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type return_order_track struct {
	Id                        string `json:"idNumber"`
	Name                      string `json:"customerName"`
	Supplier_id               string `json:"supplierId"`
	Supplier_name             string `json:"supplierName"`
	Product_name              string `json:"productName"`
	Number                    string `json:"quantity"`
	Unit                      string `json:"unit"`
	Unit_price                string `json:"unitPrice"`
	Order_date                string `json:"orderDate"`
	Estimated_submission_date string `json:"estSubDate"`
	Actual_submission_date    string `json:"actSubDate"`
}
type return_restock struct {
	Supplier_id      string `json:"supplierId"`
	Supplier_name    string `json:"supplierName"`
	Supplier_contact string `json:"responsible"`
	Location         string `json:"location"`
	Product_name     string `json:"productName"`
	Detail           string `json:"specification"`
	Unit_price       string `json:"unitPrice"`
	Unit             string `json:"unit"`
	Number           string `json:"quantity"`
	Restock_date     string `json:"purchaseDate"`
}
type return_receivable struct {
	Id              string `json:"idNumber"`
	Name            string `json:"customerName"`
	Phone           string `json:"phoneNumber"`
	Address         string `json:"address"`
	Age             string `json:"age"`
	Job             string `json:"occupation"`
	Total           string `json:"amount"`
	Remaining       string `json:"pendingAmount"`
	Join_date       string `json:"dueDate"`
	Purchase_status string `json:"status"`
}
type return_customer_info struct {
	Address         string `json:"address"`
	Age             string `json:"age"`
	Name            string `json:"customerName"`
	Id              string `json:"idNumber"`
	Image           string `json:"imageSrc"`
	Job             string `json:"occupation"`
	Phone           string `json:"phoneNumber"`
	Join_date       string `json:"registrationDate"`
	Purchase_status string `json:"status"`
	Permission		string `json:"permission"`
}

func main() {
	// 初始化一个http服务对象
	r := gin.Default()
	r.SetTrustedProxies([]string{password.Ip})
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// r.Use(cors.New(config))
	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET","POST","PUT", "PATCH","OPTIONS"},
		AllowHeaders:     []string{"Authorization","X-Requested-With", "Content-Type", "Upgrade", "Origin",
        "Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers",
	},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	  }))
	//   r.Use(CORSMiddleware())
	//   r.OPTIONS("/*any", func(c *gin.Context) {
    //     c.Status(300)
    // })

	store := cookie.NewStore([]byte("password"))
	//store.Options(sessions.Options{
        //PartitionKey: "",
    	//})
	r.Use(sessions.Sessions("testsession", store))
	db, err := sql.Open("mysql", password.Db_path)
	checkErr(err)
	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	r.GET("/", func(c *gin.Context) {
		//		session := sessions.Default(c)z
		// 通过请求上下文对象Context, 直接往客户端返回一个json
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/test", func(c *gin.Context) {
		// c.SetCookie("site_cookie", "123", 3600, "/", "db_go.explosion.tw", false, true)
		// 根据cookie名字读取cookie值
		message := c.PostForm("test")
		c.JSON(200, gin.H{
			"return": message,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		id := c.PostForm("idNumber")
		get_phone := c.PostForm("phoneNumber")
		session.Set("id", id)
		// id2 := session.Get("id")
		// fmt.Println(id2)
		// c.SetCookie("id_cookie", id, 3600, "/", "db.explosion.tw", false, true)
		// fmt.Println(id,get_phone)
		err := session.Save()
		if err != nil {
			fmt.Println(err.Error())
		}
		// fmt.Println("SELECT name,phone FROM customer_info WHERE id=? and phone=?", id, get_phone)
		rows, error := db.Query("SELECT name,phone FROM customer_info WHERE id=? and phone=?", id, get_phone)
		checkErr(error)
		var name string
		var phone string
		rows.Next()
		rows.Scan(&name, &phone)
		if name != "" && phone != "" {
			// fmt.Println(name)
			c.JSON(200, gin.H{
				"result": 1,
			})
		} else {
			c.JSON(200, gin.H{
				"result": 0,
			})
		}
	})
	r.POST("/show_track_order", func(c *gin.Context) {
		// session := sessions.Default(c)
		// id := session.Get("id")
		// id2, err := c.Cookie("id_cookie")
		// checkErr(err)
		// fmt.Println(id2)
		id := c.PostForm("id")
		// fmt.Println(id)
		var permission string
		rows, err := db.Query("SELECT permission FROM customer_info WHERE id=?", id)
		checkErr(err)
		rows.Next()
		rows.Scan(&permission)
		if permission == "0"{
			rows, err := db.Query("SELECT customer_info.id,customer_info.name,customer_order_records.supplier_id,customer_order_records.supplier_name,customer_order_records.ordered_product,customer_order_records.number,customer_order_records.unit,customer_order_records.unit_price,customer_order_records.order_date,customer_order_records.estimated_submission_date,customer_order_records.actual_submission_date FROM customer_info JOIN customer_order_records ON customer_info.id = customer_order_records.id WHERE customer_info.id = ?;",id)// 
			checkErr(err)
			var array []return_order_track
			var tmp return_order_track
			for rows.Next() {
				rows.Scan(&tmp.Id, &tmp.Name, &tmp.Supplier_id, &tmp.Supplier_name, &tmp.Product_name, &tmp.Number, &tmp.Unit, &tmp.Unit_price, &tmp.Order_date, &tmp.Estimated_submission_date, &tmp.Actual_submission_date)
				array = append(array, tmp)
				// fmt.Print(name)
			}
			// fmt.Println(array)
			c.JSON(200, array)
		}else{
			rows, err := db.Query("SELECT customer_info.id,customer_info.name,customer_order_records.supplier_id,customer_order_records.supplier_name,customer_order_records.ordered_product,customer_order_records.number,customer_order_records.unit,customer_order_records.unit_price,customer_order_records.order_date,customer_order_records.estimated_submission_date,customer_order_records.actual_submission_date FROM customer_info JOIN customer_order_records ON customer_info.id = customer_order_records.id")// WHERE customer_info.id = ?;
			checkErr(err)
			var array []return_order_track
			var tmp return_order_track
			for rows.Next() {
				rows.Scan(&tmp.Id, &tmp.Name, &tmp.Supplier_id, &tmp.Supplier_name, &tmp.Product_name, &tmp.Number, &tmp.Unit, &tmp.Unit_price, &tmp.Order_date, &tmp.Estimated_submission_date, &tmp.Actual_submission_date)
				array = append(array, tmp)
				// fmt.Print(name)
			}
			// fmt.Println(array)
			c.JSON(200, array)
		}
		

	})
	r.POST("/show_restock", func(c *gin.Context) {
		// session := sessions.Default(c)
		// id := session.Get("id")
		id := c.PostForm("id")
		// fmt.Println(id)
		var permission string
		rows, err := db.Query("SELECT permission FROM customer_info WHERE id=?", id)
		checkErr(err)
		rows.Next()
		rows.Scan(&permission)
		if permission == "0"{
			rows, err := db.Query("SELECT a.supplier_id,b.supplier_name,a.supplier_contact,a.stock_location,a.ordered_product,a.detail,a.order_unit_price,a.order_unit,a.order_number,a.restock_date FROM company_procurement_info AS a JOIN supplier_info AS b ON a.supplier_id = b.supplier_id where id = ?;",id)
			checkErr(err)
			var array []return_restock
			var tmp return_restock
			for rows.Next() {
				rows.Scan(&tmp.Supplier_id, &tmp.Supplier_name, &tmp.Supplier_contact, &tmp.Location, &tmp.Product_name, &tmp.Detail, &tmp.Unit_price, &tmp.Unit, &tmp.Number, &tmp.Restock_date)
				array = append(array, tmp)
				// fmt.Print(name)
			}
			// fmt.Println(id)
			// fmt.Println(array)
			c.JSON(200, array)
		}else{
			rows, err := db.Query("SELECT a.supplier_id,b.supplier_name,a.supplier_contact,a.stock_location,a.ordered_product,a.detail,a.order_unit_price,a.order_unit,a.order_number,a.restock_date FROM company_procurement_info AS a JOIN supplier_info AS b ON a.supplier_id = b.supplier_id;")
			checkErr(err)
			var array []return_restock
			var tmp return_restock
			for rows.Next() {
				rows.Scan(&tmp.Supplier_id, &tmp.Supplier_name, &tmp.Supplier_contact, &tmp.Location, &tmp.Product_name, &tmp.Detail, &tmp.Unit_price, &tmp.Unit, &tmp.Number, &tmp.Restock_date)
				array = append(array, tmp)
				// fmt.Print(name)
			}
			// fmt.Println(id)
			// fmt.Println(array)
			c.JSON(200, array)
		}
		

	})
	r.POST("/show_accounts_receivable", func(c *gin.Context) {
		// session := sessions.Default(c)
		id := c.PostForm("id")
		// id := session.Get("id")
		var permission string
		rows, err := db.Query("SELECT permission FROM customer_info WHERE id=?", id)
		checkErr(err)
		rows.Next()
		rows.Scan(&permission)
		if permission == "0"{
			rows, err := db.Query("SELECT a.id,a.name,a.phone,a.address,a.age,a.job,b.receivable_sum,b.remaining_balance,a.purchase_status,b.should_get_date FROM customer_info AS a JOIN company_receivables_info AS b ON a.id = b.id WHERE a.id = ?;", id)
		checkErr(err)
		var array []return_receivable
		var tmp return_receivable
		for rows.Next() {
			rows.Scan(&tmp.Id, &tmp.Name, &tmp.Phone, &tmp.Address, &tmp.Age, &tmp.Job, &tmp.Total, &tmp.Remaining, &tmp.Purchase_status,&tmp.Join_date)
			array = append(array, tmp)
			// fmt.Print(name)
		}
		// fmt.Println(id)
		// fmt.Println(array)
		c.JSON(200, array)
		}else{
			rows, err := db.Query("SELECT a.id,a.name,a.phone,a.address,a.age,a.job,b.receivable_sum,b.remaining_balance,a.purchase_status,b.should_get_date FROM customer_info AS a JOIN company_receivables_info AS b ON a.id = b.id")
		checkErr(err)
		var array []return_receivable
		var tmp return_receivable
		for rows.Next() {
			rows.Scan(&tmp.Id, &tmp.Name, &tmp.Phone, &tmp.Address, &tmp.Age, &tmp.Job, &tmp.Total, &tmp.Remaining, &tmp.Purchase_status,&tmp.Join_date)
			array = append(array, tmp)
			// fmt.Print(name)
		}
		// fmt.Println(id)
		// fmt.Println(array)
		c.JSON(200, array)
		}
		
	})
	r.POST("/show_cutomer_info", func(c *gin.Context) {
		// session := sessions.Default(c)
		// id := session.Get("id")
		id := c.PostForm("id")
		rows, err := db.Query("SELECT permission FROM customer_info WHERE id=?", id)
		var permission string
		checkErr(err)
		rows.Next()
		rows.Scan(&permission)
		if permission == "0"{
			rows, err := db.Query("SELECT a.address,a.age,a.name,a.id,a.image,a.job,a.phone,a.join_date,a.purchase_status,a.permission FROM customer_info AS a WHERE a.id = ?;", id)

			checkErr(err)
		var array []return_customer_info
		var tmp return_customer_info
		for rows.Next() {
			rows.Scan(&tmp.Address,&tmp.Age,&tmp.Name,&tmp.Id,&tmp.Image,&tmp.Job,&tmp.Phone,&tmp.Join_date,&tmp.Purchase_status,&tmp.Permission)
			tmp.Image = "https://db.explosion.tw/get_customer_photo/" + tmp.Image
			// fmt.Println(tmp.Image)
			array = append(array, tmp)
			// fmt.Print(name)
		}
		// fmt.Println(id)
		// fmt.Println(array)
		c.JSON(200, array)
		}else{
			rows, err := db.Query("SELECT a.address,a.age,a.name,a.id,a.image,a.job,a.phone,a.join_date,a.purchase_status,a.permission FROM customer_info AS a;")
			checkErr(err)
		var array []return_customer_info
		var tmp return_customer_info
		for rows.Next() {
			rows.Scan(&tmp.Address,&tmp.Age,&tmp.Name,&tmp.Id,&tmp.Image,&tmp.Job,&tmp.Phone,&tmp.Join_date,&tmp.Purchase_status,&tmp.Permission)
			tmp.Image = "https://db.explosion.tw/get_customer_photo/" + tmp.Image
			// fmt.Println(tmp.Image)
			array = append(array, tmp)
			// fmt.Print(name)
		}
		// fmt.Println(id)
		// fmt.Println(array)
		c.JSON(200, array)
		}
		
		
	})
	r.POST("/create_track_order",func(c *gin.Context){
		//INSERT INTO customer_order_records (id,ordered_product,supplier_name,unit,order_date,estimated_submission_date,actual_submission_date,number,unit_price,order_amount,supplier_id) VALUES (?,?,?,?,?,?,?,?,?,?,?);
		//INSERT INTO customer_order_records (id,order_id, ordered_product, supplier_name, unit, order_date, estimated_submission_date, actual_submission_date, number, unit_price, supplier_id) VALUES (?,?,?,?,?,?,?,?,?,?,?);
		id := c.PostForm("idNumber")
		// order_id := c.PostForm("id_order")
		// name := c.PostForm("customerName")
		supplier_id := c.PostForm("supplierId")
		supplier_name := c.PostForm("supplierName")
		product_name := c.PostForm("productName")
		number := c.PostForm("quantity")
		unit := c.PostForm("unit")
		unit_price := c.PostForm("unitPrice")
		order_date := c.PostForm("orderDate")
		estimated_submission_date := c.PostForm("estSubDate")
		actual_submission_date := c.PostForm("actSubDate")
		stmt,err := db.Prepare("INSERT INTO customer_order_records (id, ordered_product, supplier_name, unit, order_date, estimated_submission_date, actual_submission_date, number, unit_price, supplier_id) VALUES (?,?,?,?,?,?,?,?,?,?);")
		//stmt, err := db.Prepare("INSERT INTO customer_order_records (id,ordered_product,supplier_name,unit,order_date,estimated_submission_date,actual_submission_date,number,unit_price,supplier_id) VALUES (?,?,?,?,?,?,?,?,?,?,?);")
		checkErr(err)
		_,err = stmt.Exec(id,product_name,supplier_name,unit,order_date,estimated_submission_date,actual_submission_date,number,unit_price,supplier_id)


	})
	r.POST("/create_restock",func(c *gin.Context) {

		id := c.PostForm("idNumber")
		supplier_id := c.PostForm("supplierId")
		supplier_name := c.PostForm("supplierName")
		supplier_contact := c.PostForm("responsible")
		location := c.PostForm("location")
		product_name := c.PostForm("productName")
		detail := c.PostForm("specification")
		unit_price := c.PostForm("unitPrice")
		unit := c.PostForm("unit")
		number := c.PostForm("quantity")
		restock_date := c.PostForm("purchaseDate")
		var s_id string
		rows,err := db.Query("SELECT supplier_id FROM supplier_info WHERE supplier_id=?",supplier_id)
		checkErr(err)
		rows.Next()
		rows.Scan(&s_id)
		if s_id == ""{
			db.Exec("INSERT INTO supplier_info (supplier_id,supplier_name) VALUES (?,?);",supplier_id,supplier_name)
		}
		//INSERT INTO company_procurement_info (id,supplier_id,supplier_contact,ordered_product,stock_location,detail,order_unit,order_number,order_unit_price,restock_date) VALUES (3,3,'contact3','product5','location000',12315,'g',1234,100.21,'2023-10-10')
		db.Exec("INSERT INTO company_procurement_info (id,supplier_id,supplier_contact,ordered_product,stock_location,detail,order_unit,order_number,order_unit_price,restock_date) VALUES (?,?,?,?,?,?,?,?,?,?)",
			id,supplier_id,supplier_contact,product_name,location,detail,unit,number,unit_price,restock_date,
	)
	})
	r.POST("/create_accounts_receivable",func (c *gin.Context)  {
		id := c.PostForm("idNumber")
		name := c.PostForm("customerName")
		// phone := c.PostForm("phoneNumber")
		// address := c.PostForm("address")
		// age := c.PostForm("age")
		// job := c.PostForm("occupation")
		total := c.PostForm("amount")
		remaining := c.PostForm("pendingAmount")
		date := c.PostForm("dueDate")
		// purchase_status := c.PostForm("status")
		// fmt.Println(id,total,remaining,name,date)
		db.Exec("INSERT INTO company_receivables_info (id,receivable_sum,remaining_balance,customer_name,should_get_date) VALUES (?,?,?,?,?)",
			id,total,remaining,name,date,
	)})
	r.POST("/create_cutomer_info",func(c *gin.Context) {
		id := c.PostForm("idNumber")
		address := c.PostForm("address")
		age := c.PostForm("age")
		name := c.PostForm("customerName")
		file,_ := c.FormFile("imageSrc")
		job := c.PostForm("occupation")
		phone := c.PostForm("phoneNumber")
		join_date := c.PostForm("registrationDate")
		purchase_status := c.PostForm("status")
		permission := c.PostForm("permission")
		// file, _ := c.FormFile("file")
		// file.Filename = time.Now().String()
		// t := time.Now().Unix()
		// tm := time.Unix(t, 0)
		// time := tm.Format("20060102030405")
		// fmt.Println(file)
		dst := "/home/explosion/db_go/" + id
		// fmt.Println(dst)
		c.SaveUploadedFile(file, dst)
		// img := dst
		db.Exec("INSERT INTO customer_info (id,name,phone,address,age,job,join_date,image,purchase_status,permission) VALUES (?,?,?,?,?,?,?,?,?,?)",
			id,name,phone,address,age,job,join_date,id,purchase_status,permission,
	)
	})
	r.POST("/update_cutomer_info",func(c *gin.Context) {
		id := c.PostForm("idNumber")
		address := c.PostForm("address")
		// age := c.PostForm("age")
		// name := c.PostForm("customerName")
		// image := c.PostForm("imageSrc")
		// job := c.PostForm("occupation")
		phone := c.PostForm("phoneNumber")
		// join_date := c.PostForm("registrationDate")
		purchase_status := c.PostForm("status")
		// permission := c.PostForm("permission")
		
		db.Exec("UPDATE customer_info SET phone = ?,address = ?,purchase_status = ? WHERE id=?",
			phone,address,purchase_status,id,
	)
	})
	r.POST("/update_accounts_receivable",func(c *gin.Context) {
		id := c.PostForm("idNumber")
		name := c.PostForm("customerName")
		// phone := c.PostForm("phoneNumber")
		// address := c.PostForm("address")
		// age := c.PostForm("age")
		// job := c.PostForm("occupation")
		total := c.PostForm("amount")
		remaining := c.PostForm("pendingAmount")
		db.Exec("UPDATE company_receivables_info SET receivable_sum = ?,remaining_balance = ?,customer_name = ? WHERE id = ?",
			total,remaining,name,id,
	)
	})
	r.POST("/is_admin",func(c *gin.Context) {
		id := c.PostForm("id")
		var permission string
		rows, err := db.Query("SELECT permission FROM customer_info WHERE id=?", id)
		checkErr(err)
		rows.Next()
		rows.Scan(&permission)
		if permission == "0"{
			c.JSON(200,gin.H{
				"result": 0,
			})
		}else{
			c.JSON(200,gin.H{
				"result": 1,
			})
		}
	})
	r.POST("/upload",func (c *gin.Context)  {
		file, _ := c.FormFile("file")
		file.Filename = time.Now().String()
		
		c.SaveUploadedFile(file, "/home/explosion/db_go/")
	})
	r.GET("/get_customer_photo/:id", func(c *gin.Context) {
		id := c.Param("id")
		imagePath := "/home/explosion/db_go/" + id
		// 使用c.File将照片发送给前端
		c.File(imagePath)
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err.Error())
	}
}

//SELECT customer_info.id,customer_info.name,customer_order_records.supplier_id,customer_order_records.supplier_name,customer_order_records.ordered_product,customer_order_records.number,customer_order_records.unit,customer_order_records.unit_price,customer_order_records.order_date,customer_order_records.estimated_submission_date FROM customer_info  JOIN customer_order_records ON customer_info.id = customer_order_records.id WHERE customer_info = 1;
//SELECT a.supplier_id,b.supplier_id,a.supplier_contact,a.stock_location,a.ordered_product,a.detail,a.order_unit_price,a.order_unit,a.order_number,a.restock_date FROM company_procurement_info AS a JOIN supplier_info AS b ON a.supplier_id = b.supplier_id;
//SELECT a.id,a.name,a.phone,a.address,a.age,a.job,b.join_date,a.purchase_status FROM customer_info AS a JOIN company_receivables_info AS b ON a.id = b.id;
//SELECT a.address,a.age,a.name,a.id,a.image,a.job,a.phone,a.join_date,a.purchase_status FROM customer_info AS a WHERE a.id = 1;

// func corsMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}