package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

)

// tips: main関数内で呼ばなくても自動で先に処理される
func init() {
	// fmt.Println("Init")
}

func bazz() {
	fmt.Println("Bazz")
}

func DecalreVariable() {
	// var での宣言であれば関数の外でも呼び出せる
	var i int = 1
	var f64 float64 = 1.2
	var s string = "test"
	var t, f bool = true, false
	fmt.Println(i, f64, s, t, f)

	// default 0
	var i_emp int
	// default 0
	var f64_emp float64
	// default nil
	var s_emp string
	// default false
	var t_emp bool
	fmt.Println(i_emp, f64_emp, s_emp, t_emp)

	// 型推論してくれる
	// 関数の中でしか使えない
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
}

func convert() {
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}

func array() {
	// 指定するとappendできない
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:4])

	l := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)
}

func mapMethod() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m["apple"])
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)
}

// go は複数の変数をreturnすることができる
func add(x, y int) (int, int) {
	return x + y, x - y
}

func foo(params ...int) {
	fmt.Println(len(params), params)
}

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func ifgrammer() {
	//if のネストの書き方
	if result := by2(10); result == "ok" {
		fmt.Println("great")
	}
}

func forgrammer() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
}

func rangegrammer() {
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func switchgrammer() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("MAC")
	default:
		fmt.Println("default")
	}
}

// defer: メソッドの最後に実行される処理
// defer はファイルを閉じる処理やDBを閉じる処理など、処理の最後にしなければならないことをdefer設定しておくことで、処理忘れを防ぐことができる
func defergrammer() {
	// defer が重複した場合、したら順番に出力される
	defer fmt.Println("call defar1")
	defer fmt.Println("call defar2")
	defer fmt.Println("call defar3")
	fmt.Println("hahaha")
}

func errorhandling() {
	// ショートイニシャライゼーションは一つでもイニシャライズすればエラーにならない
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("Error")
	}
	defer file.Close()
}

// panic 発生するとプログラムが終了してしまうので、必ずケアしないといけない
func thirdPartyConnectDB() {
	// panic を自分で書くことはgoが推奨していない
	panic("Unable to connect database!")
}

func save() {
	// thirdPartyConnectDB()内でパニックが発生した後に、このdeferは実行される
	defer func() {
		// recoverがpanicのエラーをキャッチする
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

// '*int' : integerのポイント型という意味
func one(x *int) {
	*x = 1
}

// ポイントのイメージ：箱の番号。変数には番号札が渡されていて、参照する際にその番号札に書かれた数字の箱を開け、中身の値を取得する
// つまり、箱の中身の値が変更されれば、変数の値は変更となる
// "*"を使用したい場合は以下
//   - まだ中身が決まっていない変数に型定義をする際(ex. *int(integerのポイント型))
//   - ポイントの値を取得する時(ex. *x(xにはポイントが入っているので、そのポイントが指し示す値を取得する))
//
// "&"を使用したい場合は以下
//   - すでに何かを代入された変数に対して、その値のポイントを取得する際（ex. &n(n=100のとき、この100はどこのポイントに格納されているか)）
func pointer() {
	var n int = 100
	// &変数 :& はその値のポインタを出力することを意味する
	// つまり、以下のメソッド呼び出しでは、nのポインタを引数として渡している
	one(&n)
	fmt.Println(n)
	fmt.Println(&n)
}

func newgrammer() {
	var p *int = new(int)
	fmt.Printf("%T\n", p)
}

// クラスみたいなもの
// fieldを大文字で設定したら他のパッケージからもアクセス可能
// 小文字なら他のパッケージからはアクセスできない
type Vertex struct {
	X, Y int
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

// embedded: 継承みたいなもの
// field の一つに継承したいものをセットするだけでよい
type Vertex3D struct {
	Vertex
	Z int
}

func New3D(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

// Vertex というstructに紐づくメソッドみたいなもの
// Vertex に入ってる値を使って何かしらの処理をしたい時
func (v Vertex) Area() int {
	return v.X + v.Y
}

// ポイントレシーバ
// "*"を使うことで実態を書き換えれる
// Vertex のstructの値を書き換えたい時に使える
func (v *Vertex) Scale(i int) {
	v.X = i * v.X
	v.Y = i * v.Y
}

func vertexgrammer() {
	v := Vertex{1, 2}
	fmt.Println(v.Area())
	v3D := Vertex3D{Vertex{1, 2}, 4}
	v3D.Scale(2)

}

type Human interface {
	Say()
}

type Person struct {
	name string
}

func (p *Person) Say() {
	p.name = "Mr." + p.name
	fmt.Println(p.name)
}

func interfacegrammer() {
	// struct の持つメソッドを使用する場合、大抵はそのstructの値にアクセス、変更を加えたい場合が多い。そのため&をつけ、ポインタとしてmikeをイニシャライズする
	var mike Human = &Person{"mike"}
	mike.Say()
}

// タイプアサーション： メソッドの形は決まっているが、引数や戻り値にとる値の型がまだ決まっていない
func do(i interface{}) {
	// switch-type文：タイプアサーションを使ったメソッドではこの文法を使うことが多い。受け取った方によって処理を変えたいため
	// 引数で受け取った時点では型が決まっていないので、キャストしてあげる必要ある
	switch v:=i.(type){
	case int:
		ii := v
		ii *= 2
		fmt.Println(ii)
	case string:
		fmt.Println("mr"+v)
	case bool:
		fmt.Println(!v)
	default:
		fmt.Println("i dont know")
	}
}

func typeasationgrammer() {
	var i interface{} = 3
	do(i)
	do("kotaro")
	do(true)
}

func main() {
	typeasationgrammer()
}
