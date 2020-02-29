package playerclient

import "time"

// WaitNextStep blocks until next step simulation.
// TODO: make it asynchronous using go channels or semaphores
func (c *Client) WaitNextStep(currentStep int) {
	for c.currentTime <= currentStep {
		time.Sleep(1 * time.Millisecond)
	}
}
