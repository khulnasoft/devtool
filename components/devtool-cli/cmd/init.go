// Copyright (c) 2020 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	"github.com/khulnasoft/devtool/devtool-cli/pkg/devtool"
	"github.com/khulnasoft/devtool/devtool-cli/pkg/devtoollib"
	"github.com/khulnasoft/devtool/devtool-cli/pkg/utils"
	protocol "github.com/khulnasoft/devtool/devtool-protocol"
	"github.com/khulnasoft/devtool/supervisor/api"
)

var (
	interactive = false
)

// initCmd initializes the workspace's .devtool.yml file
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a Devtool configuration for this project.",
	Long: `
Create a Devtool configuration for this project.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		cfg := devtoollib.DevtoolFile{}
		if interactive {
			if err := askForDockerImage(ctx, &cfg); err != nil {
				return err
			}
			if err := askForPorts(&cfg); err != nil {
				return err
			}
			if err := askForTask(&cfg); err != nil {
				return err
			}
		} else {
			cfg.AddPort(3000)
			cfg.AddTask("echo 'start script'", "echo 'init script'")
		}

		d, err := yaml.Marshal(cfg)
		if err != nil {
			return err
		}
		if !interactive {
			wsInfo, err := devtool.GetWSInfo(ctx)
			if err != nil {
				return fmt.Errorf("failed to get workspace info: %w", err)
			}
			defaultImage, err := getDefaultWorkspaceImage(ctx, wsInfo)
			if err != nil {
				fmt.Printf("failed to get organization default workspace image: %v\n", err)
				fmt.Println("fallback to devtool default")
				defaultImage = ""
			}
			yml := ""
			if defaultImage != "" {
				yml = yml + fmt.Sprintf("# Image of workspace. Learn more: https://www.devtool.io/docs/configure/workspaces/workspace-image\nimage: %s\n\n", defaultImage)
			}
			yml = yml + `# List the start up tasks. Learn more: https://www.devtool.io/docs/configure/workspaces/tasks
tasks:
  - name: Script Task
    init: echo 'init script' # runs during prebuild => https://www.devtool.io/docs/configure/projects/prebuilds
    command: echo 'start script'

# List the ports to expose. Learn more: https://www.devtool.io/docs/configure/workspaces/ports
ports:
  - name: Frontend
    description: Port 3000 for the frontend
    port: 3000
    onOpen: open-preview

# Learn more from ready-to-use templates: https://www.devtool.io/docs/introduction/getting-started/quickstart
`
			d = []byte(yml)
		} else {
			fmt.Printf("\n\n---\n%s", d)
		}

		if _, err = os.Stat(".devtool.yml"); err == nil {
			prompt := promptui.Prompt{
				IsConfirm: true,
				Label:     ".devtool.yml file already exists, overwrite?",
			}
			if _, err = prompt.Run(); err != nil {
				fmt.Printf("Not overwriting .devtool.yml file. Aborting.\n")
				return GpError{Silence: true, Err: err, OutCome: utils.Outcome_Success}
			}
		}

		if err = os.WriteFile(".devtool.yml", d, 0644); err != nil {
			return err
		}

		// open .devtool.yml and Dockerfile
		if v, ok := cfg.Image.(devtoollib.DevtoolImage); ok {
			if _, err = os.Stat(v.File); os.IsNotExist(err) {
				if err = os.WriteFile(v.File, []byte(`FROM devtool/workspace-full

USER devtool

# Install custom tools, runtime, etc. using apt-get
# For example, the command below would install "bastet" - a command line tetris clone:
#
# RUN sudo apt-get -q update && \
#     sudo apt-get install -yq bastet && \
#     sudo rm -rf /var/lib/apt/lists/*
#
# More information: https://www.devtool.io/docs/config-docker/
`), 0644); err != nil {
					return err
				}
			}

			err = openCmd.RunE(cmd, []string{v.File})
			if err != nil {
				return err
			}
		}
		return openCmd.RunE(cmd, []string{".devtool.yml"})
	},
}

func getDefaultWorkspaceImage(ctx context.Context, wsInfo *api.WorkspaceInfoResponse) (string, error) {
	client, err := devtool.ConnectToServer(ctx, wsInfo, []string{
		"function:getDefaultWorkspaceImage",
	})
	if err != nil {
		return "", err
	}
	defer client.Close()

	res, err := client.GetDefaultWorkspaceImage(ctx, &protocol.GetDefaultWorkspaceImageParams{
		WorkspaceID: wsInfo.WorkspaceId,
	})
	if err != nil {
		return "", err
	}
	if res.Source == protocol.WorkspaceImageSourceInstallation {
		return "", nil
	}
	return res.Image, nil
}

func isRequired(input string) error {
	if input == "" {
		return errors.New("Cannot be empty")
	}
	return nil
}

func ask(lbl string, def string, validator promptui.ValidateFunc) (string, error) {
	scslbl := strings.Trim(strings.Split(lbl, "(")[0], " ")
	prompt := promptui.Prompt{
		Label:    lbl,
		Validate: validator,
		Default:  def,
		Templates: &promptui.PromptTemplates{
			Success: fmt.Sprintf("%s: ", scslbl),
		},
	}
	return prompt.Run()
}

func askForDockerImage(ctx context.Context, cfg *devtoollib.DevtoolFile) error {
	prompt := promptui.Select{
		Label: "Workspace Docker image",
		Items: []string{"default", "custom image", "docker file"},
		Templates: &promptui.SelectTemplates{
			Selected: "Workspace Image: {{ . }}",
		},
	}
	chce, _, err := prompt.Run()
	if err != nil {
		return err
	}

	if chce == 0 {
		wsInfo, err := devtool.GetWSInfo(ctx)
		if err != nil {
			return fmt.Errorf("failed to get workspace info: %w", err)
		}
		defaultImage, err := getDefaultWorkspaceImage(ctx, wsInfo)
		if err != nil {
			return fmt.Errorf("failed to get organization default workspace image: %w", err)
		}
		cfg.SetImageName(defaultImage)
		return nil
	}
	if chce == 1 {
		nme, err := ask("Image name", "", isRequired)
		if err != nil {
			return err
		}
		cfg.SetImageName(nme)
		return nil
	}

	// configure docker file
	dockerFile, err := ask("Dockerfile path", ".devtool.Dockerfile", isRequired)
	if err != nil {
		return err
	}
	ctxtPath, err := ask("Docker context path (enter to skip)", "", nil)
	if err != nil {
		return err
	}
	cfg.SetImage(devtoollib.DevtoolImage{
		File:    dockerFile,
		Context: ctxtPath,
	})
	return nil
}

func parsePorts(input string) ([]int32, error) {
	if input == "" {
		return []int32{}, nil
	}
	prts := strings.Split(input, ",")
	rst := make([]int32, 0)
	for _, prt := range prts {
		if pv, err := strconv.ParseUint(strings.TrimSpace(prt), 10, 16); err != nil {
			return nil, err
		} else {
			rst = append(rst, int32(pv))
		}
	}
	return rst, nil
}

func askForPorts(cfg *devtoollib.DevtoolFile) error {
	input, err := ask("Expose Ports (comma separated)", "", func(input string) error {
		if _, err := parsePorts(input); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	prts, err := parsePorts(input)
	if err != nil {
		return err
	}

	for _, pv := range prts {
		cfg.AddPort(pv)
	}
	return nil
}

func askForTask(cfg *devtoollib.DevtoolFile) error {
	input, err := ask("Startup task (enter to skip)", "", nil)
	if err != nil {
		return err
	}
	if input != "" {
		cfg.AddTask(input)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "walk me through an interactive setup.")
}
