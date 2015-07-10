/*
 * Copyright (c) 2015, Chris Benedict <chrisbdaemon@gmail.com>
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * Redistributions of source code must retain the above copyright notice, this
 * list of conditions and the following disclaimer.
 *
 * Redistributions in binary form must reproduce the above copyright notice,
 * this list of conditions and the following disclaimer in the documentation
 * and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package trap

import (
	"fmt"
	"strconv"

	"github.com/chrisbdaemon/beartrap/config"
	"github.com/chrisbdaemon/beartrap/config/validate"
	"github.com/chrisbdaemon/beartrap/trap/tcptrap"
)

// Interface defines the interface all traps adhere to
type Interface interface {
	Validate() []error
	Start() error
}

// BaseTrap holds data common to all trap types
type BaseTrap struct {
	Severity int
	params   config.Params
}

// New takes in a params object and returns a trap
func New(params config.Params) (Interface, error) {
	baseTrap := new(BaseTrap)
	var trap Interface

	baseTrap.params = params

	switch params["type"] {
	case "tcp":
		trap = tcptrap.New(params, baseTrap)
	default:
		return nil, fmt.Errorf("Unknown trap type")
	}

	// will validate later *crosses fingers*
	baseTrap.Severity, _ = strconv.Atoi(params["severity"])

	return trap, nil
}

// Validate performs validation on the parameters of the trap
func (trap *BaseTrap) Validate() []error {
	errors := []error{}

	switch err := validate.Int(trap.params["severity"]); {
	case err != nil:
		errors = append(errors, fmt.Errorf("Invalid severity: %s", err))
	case trap.Severity < 0:
		errors = append(errors, fmt.Errorf("Severity cannot be negative"))
	}

	return errors
}
