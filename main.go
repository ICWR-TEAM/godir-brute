/*  

DIR BRUTE FORCE

This Tools By AHMAD DWIYAN 

Research And Development ICWR Community

*/
	

package main
import (
	
	"fmt"
	"strconv"
	"log"
	system "os"
	"bufio"
	"net/http"
	_"net/url"
	"flag"
	 "strings"

)


func main(){
	
/* Flag Arguments Declaration */

	var url, pathDir string		

	flag.StringVar(&pathDir,"f","./list_dir.txt","Specific List Path On Your Own File If Exist")
	flag.StringVar(&url,"u","http://localhost/","Specific URL Must Include HTTP or HTTPS ")

	flag.Parse()

	if (strings.Contains(url,"http") || strings.Contains(url,"https")) == false {
		fmt.Println("Please Input Valid URL See --help as Example")
		system.Exit(0)
	} 

	

 /* Run Functions */

	lines,err := readLines(pathDir)
	if err != nil{
		log.Fatalln(err)
	}
	for _ , line := range lines {

		request , errReq := getHttp(url+"/"+line)

		if errReq != nil {
			log.Fatalln(errReq)
		}
		
		if request == ""{
			continue
		}

		log.Println(" " +request)
	}

}


/* To Read Line by Line in File and returned to Array*/

func readLines(path string)([]string, error){

	file , err := system.Open(path)
	if err != nil {
		return nil,err
	} 

	defer file.Close()

	scanfile := bufio.NewScanner(file)
	scanfile.Split(bufio.ScanLines)

	var array_string []string

	for scanfile.Scan(){
		array_string = append(array_string,scanfile.Text())
	} 
	return array_string , nil
}



/* To Send Request to Server */

func getHttp(url string)(string , error){

	respond , err := http.Get(url)

	if err != nil {
		return "" , err
	}

	defer respond.Body.Close()

	var msg string

	if respond.StatusCode != 404 {

		msg = "Status Code : "+strconv.Itoa(respond.StatusCode)+" => "+url
		return msg , nil
	
	} else {

		return "" , nil
	}

}