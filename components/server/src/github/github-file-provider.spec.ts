/**
 * Copyright (c) 2022 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { User } from "@khulnasoft/devtool-protocol";
import { ifEnvVarNotSet } from "@khulnasoft/devtool-protocol/lib/util/skip-if";
import { expect } from "chai";
import { Container, ContainerModule } from "inversify";
import { suite, retries, test, timeout, skip } from "@testdeck/mocha";
import { AuthProviderParams } from "../auth/auth-provider";
import { DevData } from "../dev/dev-data";
import { TokenProvider } from "../user/token-provider";
import { GitHubRestApi } from "./api";

import { GithubFileProvider } from "./file-provider";
import { GitHubTokenHelper } from "./github-token-helper";

@suite(timeout(10000), retries(2), skip(ifEnvVarNotSet("DEVTOOL_TEST_TOKEN_GITHUB")))
class TestFileProvider {
    static readonly AUTH_HOST_CONFIG: Partial<AuthProviderParams> = {
        id: "Public-GitHub",
        type: "GitHub",
        verified: true,
        description: "",
        icon: "",
        host: "github.com",
    };

    protected fileProvider: GithubFileProvider;
    protected user: User;
    protected container: Container;

    public before() {
        this.container = new Container();
        this.container.load(
            new ContainerModule((bind, unbind, isBound, rebind) => {
                bind(GitHubRestApi).toSelf().inSingletonScope();
                bind(AuthProviderParams).toConstantValue(TestFileProvider.AUTH_HOST_CONFIG);
                bind(GitHubTokenHelper).toSelf().inSingletonScope();
                bind(TokenProvider).toConstantValue(<TokenProvider>{
                    getTokenForHost: async () => DevData.createGitHubTestToken(),
                });
                bind(GithubFileProvider).toSelf().inSingletonScope();
            }),
        );
        this.fileProvider = this.container.get(GithubFileProvider);
        this.user = DevData.createTestUser();
    }

    @test public async testFileContent() {
        const result = await this.fileProvider.getFileContent(
            {
                repository: {
                    owner: "khulnasoft",
                    name: "devtool",
                    host: "github.com",
                    cloneUrl: "unused in test",
                },
                revision: "af51739d341bb2245598e275336ae9f730e3b41a",
            },
            this.user,
            "License.txt",
        );
        expect(result).to.not.be.undefined;
        expect(result).to.contain(`To determine under which license you may use a file from the Devtool source code,
please resort to the header of that file.`);
    }
}

module.exports = new TestFileProvider();
