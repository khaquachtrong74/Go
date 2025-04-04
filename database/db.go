package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	UserName string
	PassWord string
}

func ConnectDatabase() (*sql.DB, error) {
	fmt.Println("Truy xuất thông tin đăng nhập")
	dsn := "TIK:123@tcp(127.0.0.1:3306)/login"
	fmt.Println("Truy xuất vào database")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	err = db.Ping()
	fmt.Println("Ping vào database")
	if err != nil {
		fmt.Println("Ping thất bại")
		db.Close()
		return nil, err
	}
	return db, nil
}
func ServeFormHandle(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "<h1>Trang chủ</h1><a href='/login'>Đăng nhập</a>")
	http.ServeFile(w, r, path)

}
func SubmitLoginHandler(w http.ResponseWriter, r *http.Request, path string) {
	switch r.Method {
	case http.MethodPost:
		fmt.Println("Processing login form data (POST request)")
		var submitData Account
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error Parse Form: %v", err), http.StatusInternalServerError)
			return
		}
		submitData.UserName = string(r.PostFormValue("account"))
		submitData.PassWord = r.PostFormValue("passw")
		checkLogin(w, submitData.UserName, submitData.PassWord)
	case http.MethodGet:
		fmt.Println("Serving login form (GET request)")
		http.ServeFile(w, r, path)
	default:
		fmt.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Phương thức không được phép", http.StatusMethodNotAllowed)
	}
}
func SubmitRegisterHandle(w http.ResponseWriter, r *http.Request, path string) {
	switch r.Method {
	case http.MethodPost:
		var submitData Account
		var err error

		db, err := ConnectDatabase()
		if err != nil {
			fmt.Println("Error Connect database")
			return
		}
		err = r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error Parse Form: %v", err), http.StatusInternalServerError)
			return
		}
		submitData.UserName = string(r.PostFormValue("account"))
		submitData.PassWord = r.PostFormValue("passw")
		if submitData.UserName == "" || submitData.PassWord == "" {
			fmt.Println("Can't not empty input")
			return
		}
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(submitData.PassWord), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Error system for hashing passowrd")
			return
		}
		user_password := string(hashPassword)
		var dummy string
		selectQuery := "SELECT username FROM users WHERE username=?"
		if err := db.QueryRow(selectQuery, submitData.UserName).Scan(&dummy); err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Tài khoản chưa đăng ký")
			} else {
				fmt.Println("Lỗi hệ thống khi kiểm tra tên đăng nhập")
				return
			}
		} else {
			fmt.Println("Tài khoản đã tồn tại, vui lòng chọn tên khác!")
			return
		}
		insertQuery := "INSERT INTO users(username, password_hash) VALUES(?,?)"
		result, err := db.Exec(insertQuery, submitData.UserName, user_password)
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				if mysqlErr.Number == 1062 { // Lỗi trùng khoá Unique
					http.Error(w, "Tên đăng nhập đã tồn tại", http.StatusBadRequest)
					fmt.Println("Lỗi Tên đăng nhập đã tồn tại")
					return
				}
			}
			fmt.Println("Lỗi Insert vào database")
		}
		// connect and save on database
		rowAffected, err := result.RowsAffected()
		if rowAffected == 1 {
			fmt.Println("Register success")
			fmt.Fprintf(w, "<a href='/'>Quay lại trang chủ</a>")
		} else {
			fmt.Fprintf(w, "<h1>Đăng ký thất bại, vui lòng</h1>")
			fmt.Fprintf(w, "<a href='/register'>Quay lại trang Đăng ký</a>")
		}
	case http.MethodGet:
		fmt.Println("Serving register form (GET METHOD)")
		http.ServeFile(w, r, path)
	default:
		fmt.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Phương thức không được phép", http.StatusMethodNotAllowed)
	}

}
func checkLogin(w http.ResponseWriter, sumittedUsername string, sumittedUserPassword string) (bool, error) {
	var usr Account
	db, err := ConnectDatabase()
	if err != nil {
		fmt.Println("Lỗi kết nối vào database")
		return false, err
	}
	const Select = "SELECT username, password_hash FROM users WHERE username=?"
	//user_password_hash, err := bcrypt.GenerateFromPassword([]byte(sumittedUserPassword), bcrypt.DefaultCost)
	//	if err != nil {
	//		fmt.Println("Error system for hashing password user")
	//		return false, err
	//	}
	//	sumittedUserPassword = string(user_password_hash)
	err = db.QueryRow(Select, sumittedUsername).Scan(&usr.UserName, &usr.PassWord)
	fmt.Println("Sau khi truy cap vao database!")
	// Scan use pointer
	if err != nil {
		fmt.Println("Có vẻ như bạn chưa đăng ký, Vui lòng ấn link dưới đây!")
		fmt.Fprintln(w, "<h1>Trang đăng ký</h1><a href='/register'>Đăng ký</a>")
		return false, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(usr.PassWord), []byte(sumittedUserPassword)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			fmt.Println("Sai mat khau cho user")
			return false, err
		} else {
			fmt.Println("Lỗi hệ thống")
			return false, err
		}
	}
	fmt.Println("Đăng nhập thành công")
	w.WriteHeader(http.StatusOK) // Set status code 200 OK
	fmt.Fprintf(w, "<h1>Đã nhận dữ liệu thành công!</h1>")
	fmt.Fprintf(w, "<p>Dữ liệu bạn gửi: %s</p>", sumittedUsername)
	fmt.Fprintf(w, "<a href='/'>Quay lại trang chủ</a>")
	return true, nil

}
