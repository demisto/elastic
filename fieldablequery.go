// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

// NameQuery represents a query which is attached to a specific name.
type NameQuery interface {
	// SetName sets query name.
	SetName(name string)
	// Name returns the query name.
	Name() string
}
