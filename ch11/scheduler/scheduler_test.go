package scheduler

import (
	"cube/node"
	"cube/task"
	"cube/utils"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var nodeList = []*node.Node{
	&node.Node{Name: "test-node-1", Memory: 33554432, MemoryAllocated: 8388608, Disk: 524288000, DiskAllocated: 104857600},
	&node.Node{Name: "test-node-2", Memory: 33554432, MemoryAllocated: 16777216, Disk: 524288000, DiskAllocated: 262144000},
	&node.Node{Name: "test-node-3", Memory: 33554432, MemoryAllocated: 30408704, Disk: 524288000, DiskAllocated: 262144000},
}

func TestRoundRobinSchedulerSelectCandidateNodes(t *testing.T) {
	rrs := RoundRobin{"test-rr-scheduler", 0}

	tt := task.Task{}
	nodeList := nodeList
	got := rrs.SelectCandidateNodes(tt, nodeList)

	if !cmp.Equal(got, nodeList) {
		t.Errorf("-want/+got: \n%s", cmp.Diff(nodeList, got))
	}
}

func TestRoundRobinSchedulerScoreCandidateNodes(t *testing.T) {
	tests := []struct {
		name       string
		lastWorker int
		want       map[string]float64
	}{
		{
			name:       "node 2 scored lowest",
			lastWorker: 0,
			want: map[string]float64{
				"test-node-1": 1.0,
				"test-node-2": 0.1,
				"test-node-3": 1.0,
			},
		},
		{
			name:       "node 3 scored lowest",
			lastWorker: 1,
			want: map[string]float64{
				"test-node-1": 1.0,
				"test-node-2": 1.0,
				"test-node-3": 0.1,
			},
		},
		{
			name:       "node 0 scored lowest",
			lastWorker: 2,
			want: map[string]float64{
				"test-node-1": 0.1,
				"test-node-2": 1.0,
				"test-node-3": 1.0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rrs := RoundRobin{"test-rr-scheduler", test.lastWorker}
			task1 := task.Task{}

			candidateNodes := nodeList
			got := rrs.Score(task1, candidateNodes)
			if !cmp.Equal(got, test.want) {
				t.Errorf("-want/+got: \n%s", cmp.Diff(test.want, got))
			}
		})
	}
}

func TestRoundRobinSchedulerPickBestNode(t *testing.T) {
	tests := []struct {
		name           string
		candidateNodes []*node.Node
		scores         map[string]float64
		want           int
	}{
		{
			name:           "pick node 1 from scored candidates",
			candidateNodes: nodeList,
			scores: map[string]float64{
				"test-node-1": 0.1,
				"test-node-2": 1.0,
				"test-node-3": 1.0,
			},
			want: 0,
		},
		{
			name:           "pick node 2 from scored candidates",
			candidateNodes: nodeList,
			scores: map[string]float64{
				"test-node-1": 1.0,
				"test-node-2": 0.1,
				"test-node-3": 1.0,
			},
			want: 1,
		},
		{
			name:           "pick node 3 from scored candidates",
			candidateNodes: nodeList,
			scores: map[string]float64{
				"test-node-1": 1.0,
				"test-node-2": 1.0,
				"test-node-3": 0.1,
			},
			want: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rrs := RoundRobin{"test-rr-scheduler", 0}

			got := rrs.Pick(test.scores, test.candidateNodes)
			if !cmp.Equal(got, test.candidateNodes[test.want]) {
				t.Errorf("-want/+got: \n%s", cmp.Diff(test.candidateNodes[test.want], got))
			}
		})
	}
}

// Tests for Greedy scheduler

//var nodeList = []*node.Node{
//	&node.Node{Name: "test-node-1", Memory: 33554432, MemoryAllocated: 8388608, Disk: 524288000, DiskAllocated: 104857600},
//	&node.Node{Name: "test-node-2", Memory: 33554432, MemoryAllocated: 16777216, Disk: 524288000, DiskAllocated: 262144000},
//	&node.Node{Name: "test-node-3", Memory: 33554432, MemoryAllocated: 30408704, Disk: 524288000, DiskAllocated: 262144000},
//}

func TestGreedySchedulerScoreCandidateNodes(t *testing.T) {
	//tests := []struct {
	//	name     string
	//	mockFunc func()
	//	task     task.Task
	//	want     map[string]float64
	//}{
	//	{
	//		name: "first test",
	//		mockFunc: func() {
	//			getNodeStats = func(n *node.Node) *stats.Stats {
	//				node1Stats := *utils.GetStats(32760, 25600, 4194304, 2097152, 1766, 1, 1134, 15638, 926, 0, 19, 0, 0, 0)
	//				node1 := node.NewNode("node1:3333", "", "worker")
	//				node1.Stats = node1Stats

	//				node2Stats := *utils.GetStats(32760, 25600, 4194304, 2097152, 14696006, 1552, 2834048, 827414099, 72770, 0, 19876, 0, 0, 0)
	//				node2 := node.NewNode("node2:3333", "", "worker")
	//				node2.Stats = node2Stats

	//				node3Stats := *utils.GetStats(32760, 25600, 4194304, 2097152, 14696006, 1552, 2834048, 827414099, 72770, 0, 19876, 0, 0, 0)
	//				node3 := node.NewNode("node3:3333", "", "worker")
	//				node3.Stats = node3Stats

	//				return &stats.Stats{}
	//			}
	//		},
	//		task: task.Task{Memory: 512, Disk: 1024},
	//		want: map[string]float64{
	//			"test-node-1": 1.0,
	//			"test-node-2": 1.0,
	//			"test-node-3": 1.0,
	//		},
	//	},
	//}

	//origalGetNodeStats := getNodeStats

	//for _, test := range tests {
	//	t.Run(test.name, func(t *testing.T) {
	//		test.mockFunc()
	//		gs := Greedy{Name: "greedy-scheduler"}
	//		got := gs.Score(test.task, nodeList)
	//		if !cmp.Equal(got, nodeList) {
	//			t.Errorf("-want/+got: \n%s", cmp.Diff(nodeList, got))
	//		}
	//	})
	//diskTotal int, diskFree int, memTotal int, memUsed int, user uint64, nice uint64, sys uint64, idle uint64, iowait uint64, irq uint64, softirq uint64, steal uint64, guest uint64, guest_nice uint64//}
	//                            dT     dF     mT       mU      user  n  s     i      iow  irq sirq steal g gN
	node1Stats := *utils.GetStats(32760, 25600, 4194304, 307200, 1766, 1, 1134, 15638, 926, 0, 19, 0, 0, 0)
	node2Stats := *utils.GetStats(32760, 25600, 4194304, 2097152, 14696006, 1552, 2834048, 827414099, 72770, 0, 19876, 0, 0, 0)
	node3Stats := *utils.GetStats(32760, 25600, 4194304, 1048576, 14696006, 1552, 2834048, 827414099, 72770, 0, 19876, 0, 0, 0)

	ts1 := utils.CreateTestServer(node1Stats)
	ts2 := utils.CreateTestServer(node2Stats)
	ts3 := utils.CreateTestServer(node3Stats)

	node1 := node.Node{Name: "test-node-1", Api: ts1.URL}
	node2 := node.Node{Name: "test-node-2", Api: ts2.URL}
	node3 := node.Node{Name: "test-node-3", Api: ts3.URL}

	tt := task.Task{Memory: 512, Disk: 1024}

	want := map[string]float64{
		"test-node-1": 0.1,
		"test-node-2": 0.5,
		"test-node-3": 0.6,
	}

	gs := Greedy{Name: "greedy-scheduler"}
	got := gs.Score(tt, []*node.Node{&node1, &node2, &node3})

	if !cmp.Equal(got, want) {
		t.Errorf("-want, +got \n%s", cmp.Diff(want, got))
	}
}
func TestGreeySchedulerPickBestNode(t *testing.T) {}

// Tests for E-PVM scheduler
func TestEpvmSchedulerScoreCandidateNodes(t *testing.T) {}
func TestEpvmSchedulerPickBestNode(t *testing.T)        {}

func TestCheckTaskDisk(t *testing.T) {
	tt := task.Task{Disk: 4096}
	got := checkDisk(tt, 32760)
	if got != true {
		t.Errorf("want true, got %v", got)
	}
}

func TestSelectCandidateNodes(t *testing.T) {
	tests := []struct {
		name     string
		nodeList []*node.Node
		task     task.Task
		want     []*node.Node
	}{
		{
			name:     "select candidates returns 3 nodes",
			nodeList: nodeList,
			task:     task.Task{Memory: 512, Disk: 4096},
			want:     nodeList,
		},
		{
			name:     "select candidates returns 2 nodes",
			nodeList: nodeList,
			task:     task.Task{Memory: 8192, Disk: 4096},
			want:     []*node.Node{nodeList[0], nodeList[1]},
		},
		{
			name:     "select candidates returns 1 nodes",
			nodeList: nodeList,
			task:     task.Task{Memory: 20480, Disk: 4096},
			want:     []*node.Node{nodeList[2]},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := selectCandidateNodes(test.task, test.nodeList)
			if !cmp.Equal(test.want, got) {
				cmp.Diff(test.want, got)
			}
		})
	}
}
