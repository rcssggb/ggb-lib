package playerclient

import "time"

// WaitNextStep blocks until currentStep step simulation argument is passed.
// TODO: make it asynchronous using go channels or semaphores
func (c *Client) WaitNextStep(currentStep int) {
	for c.currentTime <= currentStep {
		time.Sleep(1 * time.Millisecond)
	}
}
