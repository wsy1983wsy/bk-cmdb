/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package app

import (
	"os"

	"configcenter/src/tools/cmdb_ctl/app/config"
	"configcenter/src/tools/cmdb_ctl/cmd"

	"github.com/spf13/cobra"
)

func Run() error {
	config.Conf = new(config.Config)

	rootCmd := &cobra.Command{
		Use: os.Args[0],
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	config.Conf.AddFlags(rootCmd)
	rootCmd.AddCommand(cmd.NewLogCommand())
	rootCmd.AddCommand(cmd.NewExitCommand())

	return rootCmd.Execute()
}