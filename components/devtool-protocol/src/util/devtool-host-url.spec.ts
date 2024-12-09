/**
 * Copyright (c) 2020 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import * as chai from "chai";
import { suite, test } from "@testdeck/mocha";
import { DevtoolHostUrl } from "./devtool-host-url";
const expect = chai.expect;

@suite
export class DevtoolHostUrlTest {
    @test public parseWorkspaceId_hosts_withEnvVarsInjected() {
        const actual = new DevtoolHostUrl(
            "https://gray-grasshopper-nfbitfia.ws-eu02.devtool-staging.com/#passedin=test%20value/https://github.com/khulnasoft/devtool-test-repo",
        ).workspaceId;
        expect(actual).to.equal("gray-grasshopper-nfbitfia");
    }

    @test public async testWithoutWorkspacePrefix() {
        expect(
            new DevtoolHostUrl("https://3000-moccasin-ferret-155799b3.ws-eu02.devtool-staging.com/")
                .withoutWorkspacePrefix()
                .toString(),
        ).to.equal("https://devtool-staging.com/");
    }

    @test public async testWithoutWorkspacePrefix2() {
        expect(new DevtoolHostUrl("https://devtool-staging.com/").withoutWorkspacePrefix().toString()).to.equal(
            "https://devtool-staging.com/",
        );
    }

    @test public async testWithoutWorkspacePrefix3() {
        expect(
            new DevtoolHostUrl("https://gray-rook-5523v5d8.ws-dev.my-branch-1234.staging.devtool-dev.com/")
                .withoutWorkspacePrefix()
                .toString(),
        ).to.equal("https://my-branch-1234.staging.devtool-dev.com/");
    }

    @test public async testWithoutWorkspacePrefix4() {
        expect(
            new DevtoolHostUrl("https://my-branch-1234.staging.devtool-dev.com/").withoutWorkspacePrefix().toString(),
        ).to.equal("https://my-branch-1234.staging.devtool-dev.com/");
    }

    @test public async testWithoutWorkspacePrefix5() {
        expect(
            new DevtoolHostUrl("https://abc-nice-brunch-4224.staging.devtool-dev.com/")
                .withoutWorkspacePrefix()
                .toString(),
        ).to.equal("https://abc-nice-brunch-4224.staging.devtool-dev.com/");
    }

    @test public async testWithoutWorkspacePrefix6() {
        expect(
            new DevtoolHostUrl("https://gray-rook-5523v5d8.ws-dev.abc-nice-brunch-4224.staging.devtool-dev.com/")
                .withoutWorkspacePrefix()
                .toString(),
        ).to.equal("https://abc-nice-brunch-4224.staging.devtool-dev.com/");
    }
}
module.exports = new DevtoolHostUrlTest();
