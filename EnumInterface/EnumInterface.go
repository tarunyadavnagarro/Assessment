package main
import "fmt"

func Name(name string) string{
       return name
}
func Mass(radius int) int{
       if radius<1{
              fmt.Println("invalid radius")
              return 0
       }
 return radius*radius*314
}

type Planet interface {
     Name() string
     Mass() int
}
type pluto struct {
              name string
              radius int
}
type mercury struct {
              name string
              radius int
}
type earth struct {
              name string
              radius int
}
type jupitor struct {
              name string
              radius int
}

func (p pluto) Name() string{
       return p.name
}
func (p pluto) Mass() int{
       if p.radius<1{
              fmt.Println("invalid radius")
              return 0
       }
 return p.radius*p.radius*314
}
func (m mercury) Name() string{
    return m.name
}
func (m mercury) Mass() int{
       if m.radius<1{
              fmt.Println("invalid radius")
              return 0
       }
 return m.radius*m.radius*314
}
func (e earth) Name() string{
       return e.name
}
func (e earth) Mass() int{
       if e.radius<1{
              fmt.Println("invalid radius")
              return 0
       }
 return e.radius*e.radius*314
}
func (j jupitor) Name() string{
       return j.name
}
func (j jupitor) Mass() int{
       if j.radius<1{
              fmt.Println("invalid radius")
              return 0
       }
 return j.radius*j.radius*314
}
func main() {
     p := pluto{name:"Pluto",radius:3}
     m := mercury{name:"mercury",radius:1}
     e := earth{name:"earth",radius:19}
     j := jupitor{name:"jupitor",radius:6044}

     fmt.Println(" Name of planet is :", p.Name())
     fmt.Println("and mass in Kg is 10^19 X ",p.Mass())
     fmt.Println(" Name of planet is :",m.Name())
     fmt.Println("and mass in Kg is 10^21 X ",m.Mass())
     fmt.Println(" Name of planet is :",e.Name())
     fmt.Println("and mass in Kg is 10^21 X ",e.Mass())
     fmt.Println(" Name of planet is :",j.Name())
     fmt.Println("and mass in Kg is 10^21 X ",j.Mass())
             
}