// +build e2e

package e2e

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"testing"
)

const (
	namespace = "argo-dataflow-system"
	baseUrl   = "http://localhost:3569"
)

var (
	restConfig          = ctrl.GetConfigOrDie()
	dynamicInterface    = dynamic.NewForConfigOrDie(restConfig)
	kubernetesInterface = kubernetes.NewForConfigOrDie(restConfig)
)

func setup(t *testing.T) {
	deletePipelines()
	waitForPipelinePodsToBeDeleted()

	t.Cleanup(portForward("testapi-0", 8378))
}

func teardown(*testing.T) {}
