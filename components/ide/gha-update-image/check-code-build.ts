// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

import { $ } from "bun";
import { parseArgs } from "util";
import { appendGitHubOutput, pathToWorkspaceYaml, readWorkspaceYaml } from "./lib/common";

$.nothrow();

const workspaceYamlInfo = await readWorkspaceYaml();
const rawWorkspaceYaml = workspaceYamlInfo.rawText;
const workspaceYaml = workspaceYamlInfo.parsedObj;

const { values } = parseArgs({
    args: Bun.argv,
    options: {
        branch: {
            type: "string",
        },
    },
    strict: true,
    allowPositionals: true,
});

const inputs = {
    branch: values.branch,
};

// example: bun run check-code-build.ts --branch gp-code/release/1.89
const main = async () => {
    if (!inputs.branch || !inputs.branch.startsWith("gp-code/release/")) {
        throw new Error(`invalid branch ${inputs.branch}, expected something like \`gp-code/release/1.90\``);
    }
    const commit =
        await $`curl -H 'Accept: application/vnd.github.VERSION.sha' https://api.github.com/repos/khulnasoft/openvscode-server/commits/${inputs.branch}`.text();

    const version = JSON.parse(
        await $`curl https://raw.githubusercontent.com/khulnasoft/openvscode-server/${commit}/package.json`.text(),
    ).version;

    console.log("fetch khulnasoft/openvscode-server with " + inputs.branch, { commit, version });

    if (workspaceYaml.defaultArgs.codeVersion === version) {
        console.error("code version is the same, no need to update");
        return;
    }
    console.log(
        `found different version ${version} (than ${workspaceYaml.defaultArgs.codeVersion}) with commit:${commit} (than ${workspaceYaml.defaultArgs.codeCommit})`,
    );
    const newYaml = rawWorkspaceYaml
        .replace(workspaceYaml.defaultArgs.codeCommit, commit)
        .replace(workspaceYaml.defaultArgs.codeVersion, version);

    await Bun.write(pathToWorkspaceYaml, newYaml);
    await appendGitHubOutput(`codeVersion=${version}`)
    await appendGitHubOutput(`codeCommit=${commit}`)
};

await main();
