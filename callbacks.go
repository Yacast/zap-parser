package zapparser

import (
	"fmt"
	"sync/atomic"
)

// OnClose adds a callback on parser close
func (p *Parser) OnClose(callback func()) error {
	if atomic.LoadUint32(p.running) != 0 {
		return fmt.Errorf("cannot set callback while parser running")
	}
	p.onClose = append(p.onClose, callback)
	return nil
}

// OnEntry adds a callback on entry found
func (p *Parser) OnEntry(callback func(*Entry)) error {
	if atomic.LoadUint32(p.running) != 0 {
		return fmt.Errorf("cannot set callback while parser running")
	}
	p.onEntry = append(p.onEntry, callback)
	return nil
}

// OnError adds a callback on parsing line error
func (p *Parser) OnError(callback func(error)) error {
	if atomic.LoadUint32(p.running) != 0 {
		return fmt.Errorf("cannot set callback while parser running")
	}
	p.onError = append(p.onError, callback)
	return nil
}
