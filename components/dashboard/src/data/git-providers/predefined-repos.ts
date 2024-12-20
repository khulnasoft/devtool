/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

export const PREDEFINED_REPOS = [
    {
        url: "https://github.com/devtool-demos/voting-app",
        repoName: "demo-docker",
        description: "A fully configured demo with Docker Compose, Redis and Postgres",
        repoPath: "github.com/devtool-demos/voting-app",
    },
    {
        url: "https://github.com/devtool-demos/spring-petclinic",
        repoName: "demo-java",
        description: "A fully configured demo with Java, Maven and Spring Boot",
        repoPath: "github.com/devtool-demos/spring-petclinic",
    },
] as const;
