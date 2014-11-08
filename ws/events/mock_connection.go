package events

// mockConnection allows us to mock a connection with different
// error states. Set the different error states to simulate an error.
type mockConnection struct {
	sendErr  error // The send error to return
	closeErr error // The close error to return
}

// Send doesn't do anything with the message. It returns the value of
// sendErr as its error.
func (c *mockConnection) Send(msg interface{}) error {
	return c.sendErr
}

// Close doesn't do anything other than return closeErr.
func (c *mockConnection) Close() error {
	return c.closeErr
}
