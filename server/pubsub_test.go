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

import "testing"

func TestPubSubVisitor(t *testing.T) {
	var pubsub pubSub
	//
	if pubsub.hasSubscriptions() {
		t.Errorf("should have no subscriptions")
	}
	//
	sender := newResponseSenderStub()
	sub1 := newSubscription(sender)
	pubsub.add(sub1)
	if !pubsub.hasSubscriptions() {
		t.Errorf("should have subscriptions")
	}
	if pubsub.count() != 1 {
		t.Errorf("expected 1 subscription")
	}
	//
	sub2 := newSubscription(sender)
	pubsub.add(sub2)
	if pubsub.count() != 2 {
		t.Errorf("expected 2 subscription")
	}
	//
	sub3 := newSubscription(sender)
	pubsub.add(sub3)
	if pubsub.count() != 3 {
		t.Errorf("expected 3 subscription")
	}
	//

	pubsub.publish(newOkResponse())

	//
	sub3.deactivate()
	if pubsub.count() != 2 {
		t.Errorf("expected 2 subscription")
	}
	//
	sub1.deactivate()
	if pubsub.count() != 1 {
		t.Errorf("expected 2 subscription")
	}
	//
	sub2.deactivate()
	if pubsub.count() != 0 {
		t.Errorf("expected 0 subscription")
	}
	if pubsub.hasSubscriptions() {
		t.Errorf("should have no subscriptions")
	}
}