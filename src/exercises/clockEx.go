package exercises

import (
	"fmt"
)

type Clock struct {
	Hour, Minute, Second int
}

func (c Clock) String() string {
	return fmt.Sprintf("%v:%v:%v", c.Hour, c.Minute, c.Second)
}

func (c *Clock) AddMinutes(minutes int) {
	c.Minute += minutes
	c.Hour += int(c.Minute / 60)
	c.Minute = c.Minute % 60
}

func (c *Clock) SubstractMinutes(minutes int) {
	c.Hour -= int(minutes / 60)
	c.Minute -= minutes % 60
}

func (c *Clock) Equals(c1 *Clock) bool {
	if c.Hour != c1.Hour || c.Minute != c1.Minute || c.Second != c1.Second {
		return false
	}
	return true
}

func AddClocks(c1 Clock, c2 Clock) (c3 Clock) {
	c3.Hour = c1.Hour + c2.Hour
	c3.Minute = c1.Minute + c2.Minute
	c3.Second = c1.Second + c2.Second

	c3.Minute += int(c3.Second / 60)
	c3.Second = c3.Second % 60

	c3.Hour += int(c3.Minute / 60)
	c3.Minute = c3.Minute % 60
	return c3
}

//func main() {
//	var c1, c2 *exercises.Clock = &exercises.Clock{02, 02, 03}, &exercises.Clock{02, 02, 03}
//	if c1.Equals(c2) {
//		fmt.Println("Equal")
//	} else {
//		fmt.Println("Not Equal")
//	}
//	c3 := exercises.AddClocks(*c1, *c2)
//	fmt.Println(c3)
//	c1.AddMinutes(3)
//	fmt.Println(c1)
//	c1.SubstractMinutes(180)
//	fmt.Println(c1)
//}
