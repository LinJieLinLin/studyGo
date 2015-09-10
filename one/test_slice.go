package main
import "fmt"
func main(){
	fmt.Println("Start")
	arr:=[]int64{1,2}
	fmt.Println(arr)
	arr = append(arr[1:],arr...)
	b:=arr[1:]
	fmt.Println(b)
	b[0]=10
	fmt.Println(arr)
	交换位置(arr,0,1)
}
func 交换位置(arr []int64,i,j int64){
	fmt.Println(arr)
	arr[i],arr[j] = arr[j],arr[i]
	fmt.Println(arr)
}