// Ultimate Provisioner: UP cmd
// Copyright (c) 2019 Stephen Cheng and contributors

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package impl

import "github.com/stephencheng/up/model/stack"

var (
	InstanceName string
	Dryrun       bool
	TaskStack    = stack.New("task")
	StepStack    = stack.New("step")
)

func SetInstanceName(id string) {
	if id != "" {
		InstanceName = id
	} else {
		InstanceName = "nonamed"
	}
}

func SetDryrun() {
	Dryrun = true
}
