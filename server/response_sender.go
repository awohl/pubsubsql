/* Copyright (C) 2013 CompleteDB LLC.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with PubSubSQL.  If not, see <http://www.gnu.org/licenses/>.
 */

package pubsubsql

// responseSender is responsible for channeling reponses to client connection
type responseSender struct {
	sender           chan response // channel to publish responses to
	connectionId     uint64
	connectionStoper *Stoper
}

// factory
func newResponseSenderStub(connectionId uint64) *responseSender {
	return &responseSender{
		sender:           make(chan response, config.CHAN_RESPONSE_SENDER_BUFFER_SIZE),
		connectionId:     connectionId,
		connectionStoper: NewStoper(),
	}
}

func (this *responseSender) send(res response) bool {
	select {
	case this.sender <- res:
		if !this.connectionStoper.Stoped() {
			return true
		}
		debug("sender is stoped")
	case <-this.connectionStoper.GetChan():
		debug("connection is closed")
	default:
		logerror("sender queue is full connection: ", this.connectionId)
		// notify client connection that it needs to close due to inability to 
		// send responses in a timely manner
		this.connectionStoper.Stop(0)
	}
	return false
}

func (this *responseSender) tryRecv() response {
	select {
	case res := <-this.sender:
		return res
	default:
		return nil
	}
	return nil
}

func (this *responseSender) recv() response {
	return <-this.sender
}
