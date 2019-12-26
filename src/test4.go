package main


func main() {
    var i int = 100
    var u uint = uint(i)
    var f float32 = float32(i)
    println(f, u)

    str := "ABC"
    bytes := []byte(str)
    str2 := string(bytes)
    println(bytes, str2)
    var k int = 10
    var p = &k
    println(p)
    if k!=*p{
      println("다르다")
    }else{
      println("같다")
    }
    var category int = 1
    switch1(category)
    var ttt int = 50
    x(ttt)
    names()
    yyyyy()
    var hi string ="hi ㅋㅋㅋㅋ"
    say(hi)
    msg := "hello"
    say1(&msg)
    println(msg)
    msg=say2("this","is","a","book")
    println(msg);
    msg=say2("hi")
    println(msg);
    sum, fa := plus(1,2,3,4,5,6,7,8,9,1,0,2,3,4,5,2,7,)
    println(sum,  fa)
    sum2:= func(n ...int) int{
      s:=0
      for _, := range n{
        s+=i
      }
      return s
    }
    result :=sum2(1,2,3,4,5)
    println(result)



}

func switch1(a int)  {
  var name string
  switch a {
  case 1:
    name ="1번"
  case 2:
    name="2번"
  case 3:
    name="3번"
  }
  println(name)
}
func x(a int ){
  sum :=0;
  for i := 1;i <= a;i++{
    if i%2==0{
      sum++;
    }
  }
  print("짝수의 개수는  ")
  println(sum)
}
func names()  {
  names := []string{"홍길","이순","강감"}
  for index, name := range names{
    println(index,name)
  }
}
func yyyyy()  {
  var a=1
  for a<15{
    if a==5{
      a+=a
      continue
    }
    a++
    println(a)
    if a>10{
      break
    }
  }
  if a==11{
    goto END
  }
  println(a)
END:
      println("END")
      println(a)
}
func say(msg string)  {
  println(msg)
}
func say1(msg *string)  {
  println(*msg)
  *msg = "changed"
}
func say2(msg ...string) string{
  var aaaaa string =""
  for _, s := range msg{
    aaaaa+=s
  }
  return aaaaa
}
func plus(s ...int)  (int,int){
  sum:=0
  fa:=0

  for _, a := range s {
    if a%2==0{
      fa++
    }
    sum+=a
  }
  return sum,fa
}

