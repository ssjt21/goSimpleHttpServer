
package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
	"net/http"
)

//判断路径是否存在
func  IsExists(path string) bool {
	workdir, err:=os.Lstat(path);
	if (err != nil) {
		fmt.Println("[Error] Working folder path parameter does not exist! "+path);
		return false;
	}
	//println(workdir.IsDir());
	if(! workdir.IsDir()){
		fmt.Println("[Error] The given parameter is not a folder! "+path)
		return false;
	}
	return true;

}

func StartServer(workDir *string, port *int){

		//defer func() {
		//	if err:=recover();err!=nil{
		//		fmt.Println("[Error] Current port in use! "+strconv.Itoa(*port));
		//	}
		//}()
		println("[Info] Http File Server start: http://127.0.0.1:"+strconv.Itoa(*port)+" .....");
		println("[Info] Current Working folder is: "+*workDir);
		http.Handle("/",http.FileServer(http.FileSystem(http.Dir(*workDir))));
		err:=http.ListenAndServe(":"+strconv.Itoa(*port),nil);
		if err!=nil {
			fmt.Println("[Error] Current port in use! "+strconv.Itoa(*port));
		}

}

func main(){

	var port=flag.Int("p",8000,"Listenning port,default value is 8000.");
	var workDir=flag.String("d",".","Working directory path,default value is '.' .");
	flag.Parse();



	//端口判断1. 数字  2.范围1-65535
	if *port>65535 || *port < 0 {
		fmt.Println("[Error] Please input valid value for listen port [1<=port<=65535]");
		os.Exit(1);
	}
	//路径判断

	if !IsExists(*workDir) {
		os.Exit(2);
	}
	//启动server
	StartServer(workDir,port);


	//http.ListenAndServe("127.0.0.1:80",);

	//panic("[Error] Current port in use! "+strconv.Itoa(*port));


}

