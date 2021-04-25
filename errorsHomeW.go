package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"project/panic"
)
var db *sql.DB

func connectMysql() error{
	//数据库连接
	var err error
	// 数据源语法："用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db,err =sql.Open("mysql","root:ZMZzmz761028@(127.0.0.1:3306)/errorTest")
	if err != nil{
		fmt.Println(err)
		errWarp := errors.Wrap(err, "open 连接失败")
		return  errWarp
	}else{
		//fmt.Println("open connect success ")
	}

	err =db.Ping()
    if err != nil{
		fmt.Println("ping 链接失败")
		errWarp := errors.Wrap(err, "ping 连接失败")
		return  errWarp
	}else{
		//fmt.Println("connect Ping success ")
	}

	return  nil
}

// 获取到sql 句柄之后，对取到的数据进行查询
func readSql(DB *sql.DB) error{
	// 查询数据
	_, errWarp := DB.Query("select id, name from users where id = ?", 1)
	if errWarp != nil {
		return errors.Wrap(errWarp, "Query")
	}

	// 查询数据数量
	return nil
}

// 获取到sql 句柄之后，对取到的数据进行查询
func writeSql(DB *sql.DB) error{
	test := 1
	for i:= 0; i<100 ; i++ {
		if test == 10{
			test = 1
		}
		sql := fmt.Sprintf(
			"insert into error_test(id, age, money) values (%d, %d, %d)",
			i,
			test,
			i,
		)

		_, err := DB.Exec(sql)
		if err!=nil{
			fmt.Printf("Exec failed \n");
			errWarp := errors.Wrap(err, "Exec 失败")
			return errWarp
		}else{
			//fmt.Printf("Exec succes \n");
		}
		test++
	}

	// 查询数据数量
	return nil
}

func main(){

	var err error

	// 捕捉错误
	defer func(){
		if err=recover(); err!=nil{
			fmt.Printf(" have a panic action , need recover \n")
			fmt.Printf("stack err trace: %+v :%T, %v \n", err, errors.Cause(err), errors.Cause(err))
			panic.PrintStack()
            //os.Exit(1)
		}
	}()
	
	// 链接数据库
	err = connectMysql()
		fmt.Printf("stack err trace: %+v :%T, %v \n", err, errors.Cause(err), errors.Cause(err))
	}


	DB,err =sql.Open("mysql","root:ZMZzmz761028@(127.0.0.1:3306)/errorTest")
	if err != nil{
		fmt.Println(err)
		errWarp := errors.Wrap(err, "open 连接失败")
		return  errWarp
	}else{
		//fmt.Println("open connect success ")
	}


	err = writeSql(DB)
	if err != nil{
		fmt.Printf("stack err trace: %+v :%T, %v \n", err, errors.Cause(err), errors.Cause(err))
	}

	err = readSql(DB)
	if err != nil{
		fmt.Printf("stack err trace: %+v :%T, %v \n", err, errors.Cause(err), errors.Cause(err))
	}
}


