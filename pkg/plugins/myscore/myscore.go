package myscore

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime" // Add this import
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type myscoreplugin struct {
	handle framework.Handle
}

const Name = "MyScorePlugin"

// Name returns name of the plugin. It is used in logs, etc.
func (pl *myscoreplugin) Name() string {
	return Name
}
func (pl *myscoreplugin) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	fmt.Println("MyScorePlugin Score called")
	return 0, framework.NewStatus(framework.Success, "")
}

// New initializes a new plugin and returns it.
func New(ctx context.Context, _ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &myscoreplugin{handle: h}, nil
}
