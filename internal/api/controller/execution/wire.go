// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package execution

import (
	"github.com/harness/gitness/internal/auth/authz"
	"github.com/harness/gitness/internal/store"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// WireSet provides a wire set for this package.
var WireSet = wire.NewSet(
	ProvideController,
)

func ProvideController(db *sqlx.DB,
	authorizer authz.Authorizer,
	executionStore store.ExecutionStore,
	pipelineStore store.PipelineStore,
	spaceStore store.SpaceStore,
) *Controller {
	return NewController(db, authorizer, executionStore, pipelineStore, spaceStore)
}
