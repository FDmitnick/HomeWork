/*
1、 基于 errgroup 实现一个 http server 的启动和关闭 ，
2、 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
利用 context 来检测是否存在需要退出的情况
*/

package main

import(
	"fmt"
	"golang.org/x/sync/errgroup"
	xContext "golang.org/x/net/context"
	"time"
	"os"
    "os/signal"
	"net/http"
	// "log"
)

//校验是否有协程已发生错误
func CheckGoroutineErr(ctx xContext.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func serverDebug(ctx xContext.Context, cancelFunc xContext.CancelFunc){
	go func(){
		// 检测错误发生
		err := CheckGoroutineErr(ctx)
		if err != nil{
			fmt.Println("serverDebug other context error ")
		}
	}()
	if err:=http.ListenAndServe("127.0.0.1:8003", http.DefaultServeMux); err!=nil{
		//log.Fatal(err)
		cancelFunc()
	}
	fmt.Println("ServerDebug return")
}

func serverAction(ctx xContext.Context, cancelFunc xContext.CancelFunc){
	go func(){
		// 检测错误发生
		err := CheckGoroutineErr(ctx)
		if err != nil{
			fmt.Println("serverAction other context error ")
		}
	}()
	mux:= http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "hello, Qcon!")
		cancelFunc()
	})
	fmt.Println("serverAction return")
}

func main(){
	fmt.Println(" GoRoutine_context is action ")
	ctx, cancel := xContext.WithCancel(xContext.Background())
	group, errCtx := errgroup.WithContext(ctx)

	c := make(chan os.Signal)
	signal.Notify(c)

	for index := 0; index < 3 ; index++{
		indexTemp := index 	
		group.Go(func() error{
			fmt.Println("indexTemp = ", indexTemp)
			if indexTemp == 0{

				serverDebug(errCtx, cancel)

			}else if indexTemp == 1{

				serverAction(errCtx, cancel)

			}else if indexTemp == 2{
				// 休眠1秒，用于捕获子协程2的出错
				//time.Sleep(10 * time.Second)
				select{
				case <-c:
					fmt.Println("========= get signal:")
					cancel()
					time.Sleep(5 * time.Second)
				}
			}
			return nil
		})
	}

	err := group.Wait()
	if err == nil{
		fmt.Println("over ")
	}else{
		fmt.Printf("get error: %V \n", err)
	}
}
