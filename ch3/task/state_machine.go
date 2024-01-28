package task

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

//type StateMachine struct {
//	StartState   *State
//	CurrentState *State
//}
//
//func (s *StateMachine) GetCurrentState() State {
//	return *s.CurrentState
//}
//
//func (s *StateMachine) SetPending() error {
//	if s.CurrentState == nil {
//		s.CurrentState = *Pending
//	}
//}
//
//func (s *StateMachine) SetScheduled(currentState State) error {
//	if currentState == Pending {
//		s.CurrentState = Scheduled
//		return nil
//	}
//
//	return fmt.Errorf("cannot transition from %s to Scheduled", currentState)
//}
//
//func (s *StateMachine) SetRunning(currentState State) error {
//	if Contains(stateTransitionMap[currentState], Running) {
//		s.CurrentState = Running
//		return nil
//	}
//
//	return fmt.Errorf("cannot transition from %s to Running", currentState)
//}
//
//func (s *StateMachine) SetCompleted(currentState State) error {
//	if Contains(stateTransitionMap[currentState], Completed) {
//		s.CurrentState = Completed
//		return nil
//	}
//
//	return fmt.Errorf("cannot transition from %s to Completed", currentState)
//}
//
//func (s *StateMachine) SetFailed(currentState State) error {
//	if Contains(stateTransitionMap[currentState], Completed) {
//		s.CurrentState = Failed
//		return nil
//	}
//
//	return fmt.Errorf("cannot transition from %s to Failed", currentState)
//}

var stateTransitionMap = map[State][]State{
	Pending:   []State{Scheduled},
	Scheduled: []State{Running, Failed},
	Running:   []State{Running, Completed, Failed},
	Completed: []State{},
	Failed:    []State{},
}

func Contains(states []State, state State) bool {
	for _, s := range states {
		if s == state {
			return true
		}
	}
	return false
}
