package state

import "fmt"

/*
	设计思想：
		Context: 拥有多种状态的对象(struct), 状态属性为State
		State: 封装特定状态行为的interface
		ConcreteState: 具体的状态，继承接口State，不同的状态执行Context的不同行为
*/
//context对象
type Context struct {
	State 		ActionState
	HealthValue int
}

//账号的行为
func (a *Context) View() {
	a.State.View()
}
func (a *Context) Comment() {
	a.State.Comment()
}
func (a *Context) Create() {
	a.State.Create()
}

func (a *Context) SetHealth(value int) {
	a.HealthValue = value
	a.changestate()
}

func (a *Context) changestate() {
	if a.HealthValue < 0 {
		a.State = &ClosedState{}
	} else if a.HealthValue > 10 {
		a.State = &NormalState{}
	} else if a.HealthValue <10 && a.HealthValue > 0 {
		a.State = &RestrictedState{}
	}
}

func NewContext(health int) *Context {
	con := &Context{HealthValue: health}
	con.changestate()
	return con
}
//因为含有多种状态，状态类型不确定，这里需要声明State接口，多态原则
type ActionState interface {
	View()
	Comment()
	Create()
}
/*
具体的状态，继承接口
*/
//Normal state
type NormalState struct {}

func (n *NormalState) View() {
	fmt.Println("view normal")
}
func (n *NormalState) Comment() {
	fmt.Println("comment normal")
}
func (n *NormalState) Create() {
	fmt.Println("create normal")
}
//Restricted State
type RestrictedState struct {}

func (n *RestrictedState) View() {
	fmt.Println("view Restricted")
}
func (n *RestrictedState) Comment() {
	fmt.Println("comment Restricted")
}
func (n *RestrictedState) Create() {
	fmt.Println("create Restricted")
}
//Closed State
type ClosedState struct {}

func (n *ClosedState) View() {
	fmt.Println("view closed")
}
func (n *ClosedState) Comment() {
	fmt.Println("comment closed")
}
func (n *ClosedState) Create() {
	fmt.Println("create closed")
}
