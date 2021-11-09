// +build disruptive

package e2e

import (
	"testing"
	"time"

	"github.com/openshift/csi-driver-shared-resource/test/framework"
)

func TestBasicThenDriverRestartThenChangeShare(t *testing.T) {
	testArgs := &framework.TestArgs{
		T: t,
	}
	TestBasicThenDriverRestartThenChangeShareWithReadOnlyMount(t)
	framework.CreateTestNamespace(testArgs)
	defer framework.CleanupTestNamespaceAndClusterScopedResources(testArgs)
	basicShareSetupAndVerification(testArgs)

	t.Logf("%s: initiating csi driver restart", time.Now().String())
	framework.RestartDaemonSet(testArgs)
	t.Logf("%s: csi driver restart complete, check test pod", time.Now().String())
	testArgs.TestDuration = 30 * time.Second
	testArgs.SearchString = "invoker"
	framework.ExecPod(testArgs)

	t.Logf("%s: now changing share", time.Now().String())
	framework.ChangeShare(testArgs)
	testArgs.SearchString = "ca.crt"
	framework.ExecPod(testArgs)
}


func TestBasicThenDriverRestartThenChangeShareWithReadOnlyMount(t *testing.T) {
	testArgs := &framework.TestArgs{
		T: t,
	}
	testArgs.ReadOnly = true
	prep(testArgs)
	framework.CreateTestNamespace(testArgs)
	defer framework.CleanupTestNamespaceAndClusterScopedResources(testArgs)
	basicShareSetupAndVerification(testArgs)

	t.Logf("%s: initiating csi driver restart", time.Now().String())
	framework.RestartDaemonSet(testArgs)
	t.Logf("%s: csi driver restart complete, check test pod", time.Now().String())
	testArgs.TestDuration = 30 * time.Second
	testArgs.SearchString = "invoker"
	framework.ExecPod(testArgs)

	t.Logf("%s: now changing share", time.Now().String())
	framework.ChangeShare(testArgs)
	testArgs.SearchString = "ca.crt"
	framework.ExecPod(testArgs)
}
