package observer

import "fmt"

// 抽象主题
type Subject interface {
	Notify()                //通知
	State() int             //
	SetState(int)           //
	AddCallFunc(*update)    //注册
	RemoveCallFunc(*update) //注销
}

// 事件/方法(就是一个委托或中介对象，用于传递和传导动作)
type update func(int)

// 具体主题
type SubjectA struct {
	state int
	call  []*update
}

func NewSubjectA(s int) *SubjectA {
	return &SubjectA{s, []*update{}}
}
func (s *SubjectA) Notify() {
	if s == nil {
		return
	}
	for _, c := range s.call {
		(*c)(s.state)
	}
}
func (s *SubjectA) State() int {
	if s == nil {
		return 0
	}
	return s.state
}
func (s *SubjectA) SetState(i int) {
	if s == nil {
		return
	}
	s.state = i
}
func (s *SubjectA) AddCallFunc(f *update) {
	if s == nil {
		return
	}
	for _, c := range s.call {
		if c == f {
			return
		}
	}

	s.call = append(s.call, f)
}
func (s *SubjectA) RemoveCallFunc(f *update) {
	if s == nil {
		return
	}
	for i, c := range s.call {
		if c == f {
			s.call = append(s.call[:i], s.call[i+1:]...)
		}
	}
}

// 观察者接口
type Observer interface {
	Update(int)
}

// 具体观察者
type ObserverA struct {
	s     Subject
	state int
}

func NewObserverA(sa Subject, s int) *ObserverA {
	return &ObserverA{sa, s}
}

func (o *ObserverA) Update(s int) {
	if o == nil {
		return
	}
	fmt.Println("ObserverA")
}

//type ObserverB struct {
//	s     Subject
//	state int
//}
//func (o *ObserverB) Update(s int) {
//	if o == nil {
//		return
//	}
//	fmt.Println("ObserverB")
//	fmt.Println(s)
//	fmt.Println(o)
//}
//func NewObserverB(sa Subject, s int) *ObserverB {
//	return &ObserverB{sa, s}
//}
