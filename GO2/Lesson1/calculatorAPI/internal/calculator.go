/*
## Задача № 1
Написать API для указанных маршрутов(endpoints)
"/info"   // Информация об API
"/first"  // Случайное число
"/second" // Случайное число
"/add"    // Сумма двух случайных чисел
"/sub"    // Разность
"/mul"    // Произведение
"/div"    // Деление

*результат вернуть в виде JSON

"math/rand"
number := rand.Intn(100)
! не забудьте про Seed()


GET http://127.0.0.1:1234/first

GET http://127.0.0.1:1234/second

GET http://127.0.0.1:1234/add
GET http://127.0.0.1:1234/sub
GET http://127.0.0.1:1234/mul
GET http://127.0.0.1:1234/div
GET http://127.0.0.1:1234/info
*/

package calculator

import (
	"math"
	"math/rand"
)

type Calculator struct {
}

func New() *Calculator {
	calc := &Calculator{}
	return calc
}

func (c *Calculator) Info() string {

	res := "<h1>Калькулятор API</h1>"
	res = res + "<h2>Маршруты:</h2>"
	res = res + "<dl>"
	res = res + "<dt><a href=\"/info\">\"/info\"</a></dt><dd>Информация об API</dd>"
	res = res + "<dt><a href=\"/first\">\"/first\"</a></dt><dd>Случайное число</dd>"
	res = res + "<dt><a href=\"/second\">\"/second\"</a></dt><dd>Случайное число</dd>"
	res = res + "<dt><a href=\"/add\">\"/add\"</a></dt><dd>Сумма двух случайных чисел</dd>"
	res = res + "<dt><a href=\"/sub\">\"/sub\"</a></dt><dd>Разность</dd>"
	res = res + "<dt><a href=\"/mul\">\"/mul\"</a></dt><dd>Произведение</dd>"
	res = res + "<dt><a href=\"/div\">\"/div\"</a></dt><dd>Деление</dd>"
	res = res + "</dl>"
	return res
}
func (c *Calculator) random() int {
	//rand.Seed(time.Now().UnixNano()) //Deprecated
	return rand.Intn(100)
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Sub(a, b int) int {
	return a - b
}

func (c *Calculator) Mul(a, b int) int {
	return a * b
}

func (c *Calculator) Div(a, b int) float64 {
	if b == 0 {
		return math.MaxFloat64
	}
	return float64(a) / float64(b)
}

func (c *Calculator) First() int {
	return c.random()
}
func (c *Calculator) Second() int {
	return c.random()
}
