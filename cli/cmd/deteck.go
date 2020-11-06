/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/spf13/cobra"
)

type DeteckCommand struct {
	baseCommand
	*baseExpCommandService
}

func (dc *DeteckCommand) Init() {
	dc.command = &cobra.Command{
		Use:     "deteck",
		Aliases: []string{"k"},
		Short:   "Deteck the environment",
		Long:    "Deteck the environment is ok for chaosblade or not",
		RunE: func(cmd *cobra.Command, args []string) error {
			return spec.ReturnFail(spec.Code[spec.IllegalCommand],
				fmt.Sprintf("less TARGE to deteck"))
		},
		Example: dc.detectExample(),
	}
}

func (dc *DeteckCommand) detectExample() string {
	return "deteck os"
}
