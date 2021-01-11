package playerclient

import "time"

// WaitNextStep blocks until currentStep step simulation argument is passed.
// TODO: make it asynchronous using go channels or semaphores
func (c *Client) WaitNextStep(currentStep int) {
	for c.currentTime <= currentStep {
		time.Sleep(1 * time.Millisecond)
	}
}

// WaitSynch blocks until (think) message is received.
// Can only be called once per cycle. (TODO: use sync.Cond instead)
func (c *Client) WaitSynch() {
	<-c.thinkChan
}

// WaitSight blocks until (see ...) message is received.
// Can only be called once per cycle. (TODO: use sync.Cond instead)
func (c *Client) WaitSight() {
	<-c.sightChan
}

// WaitBody blocks until (sense_body ...) message is received.
// Can only be called once per cycle. (TODO: use sync.Cond instead)
func (c *Client) WaitBody() {
	<-c.bodyChan
}
