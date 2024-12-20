## Features

- **Dev environments as code** - Devtool is like infrastructure-as-code, but for your development environment. Devtool defines your editor extensions and requires dependencies in a declarative [`.devtool.yml` configuration](https://www.devtool.io/docs/introduction/devtool-tutorial/2-configure-your-devtool-yml). Spinning up dev environments is easily repeatable and reproducible empowering you to automate, version-control, and share dev environments across your team.
- [**Prebuilt dev environments**](https://www.devtool.io/docs/configure/projects/prebuilds) - Devtool continuously prebuilds all your git branches similar to a CI server. Control how Devtool pre-configures and initializes environments before you even start a workspace through tasks commands in your .devtool.yml. No more watching apt-get or npm install again. 
- [**Secure**](https://www.devtool.io/security) - Each Devtool workspace or prebuild runs on a secured single-use container providing fast startup times without compromising on security. Devtool generates SLSA level 1 compliant provenance. Devtool is also GDPR and SOC2 compliant. And, of course, Devtool is open-source and available for review by everyone.
- **Workspaces based on Docker** - Devtool instantly starts a container in the cloud based on an (optional) [Docker image](https://www.devtool.io/docs/config-docker/). If you’re already using Docker, you can easily re-use your Docker file. 
- **GitLab, GitHub, and Bitbucket integration** - Devtool seamlessly [integrates](https://www.devtool.io/docs/configure/authentication) into your workflow and works with all major Git hosting platforms including GitHub, GitLab, and Bitbucket.
- **Integrated code reviews** - with Devtool you can do native code reviews on any PR/MR. No need to switch contexts anymore and clutter your local machine with your colleagues' PR/MR.
- **Collaboration** - invite team members to your dev environment or snapshot of any state of your dev environment to share it with your team asynchronously.
**Professional & customizable developer experience** - a Devtool workspace gives you the same capabilities as your Linux machine - pre-configured and optimized for your development workflow. Install any [VS Code extension](https://www.devtool.io/docs/references/ides-and-editors/vscode-extensions) with one click on a user and/or team level. You can also bring your [dotfiles](https://www.devtool.io/docs/configure/user-settings/dotfiles#dotfiles) and customize your dev environment as you like.


## Getting Started

- **Browser**: 
    - Using Devtool dashboard [devtool.io/new](https://devtool.io/new).
    - Add `devtool.io/# `as a prefix to any of your GitHub/ GitLab/ Bitbucket repository, like [this](https://devtool.io/#https://github.com/khulnasoft/template-typescript-react)
- **CLI**: You can also [install Devtool CLI](https://www.devtool.io/docs/references/devtool-cli#installation) and create your first workspace directly from your terminal :)


## Documentation

All documentation can be found on [www.devtool.io/docs](https://www.devtool.io/docs). For example, see [Devtool tutorial](https://www.devtool.io/docs/introduction/devtool-tutorial) and check the following helpful resources:
  - [Workspace Lifecycle](https://www.devtool.io/docs/configure/workspaces/workspace-lifecycle)
  - [Configure repositories](https://www.devtool.io/docs/configure/repositories)
  - [Organizations](https://www.devtool.io/docs/configure/orgs)
  - [IDE & Editors support](https://www.devtool.io/docs/references/ides-and-editors)
  - [Video screencasts](https://www.devtool.io/screencasts)
  - [Devtool samples](https://github.com/devtool-samples)


## Questions

For questions and support please use [Devtool community Discord](https://www.devtool.io/chat). Join the conversation, and connect with other community members. 💬
You can also follow [@devtool](https://twitter.com/devtool) for announcements and updates from our team.

For enterprise deployment and customized solutions, please explore our [**Enterprise offerings**](https://www.devtool.io/contact/enterprise-self-serve) to get started with a setup that meets your organization's needs.

## Issues

The issue tracker is used for tracking bug reports and feature requests for the Devtool open source project as well as planning current and future development efforts. 🗺️

You can upvote popular feature requests or create a new one.


## Related Projects

During the development of Devtool, we also developed some of our infrastructure toolings to make development easier and more efficient. To this end, we've developed many open-source projects including:
- [Workspace images](https://github.com/khulnasoft/workspace-images): Ready to use docker images for Devtool workspaces
- [OpenVSCode Server](https://github.com/khulnasoft/openvscode-server): Run the latest VS Code on a remote machine accessed through a browser
- [Devtool browser extension](https://github.com/khulnasoft/browser-extension): It adds a Devtool button to the configured GitLab, GitHub and Bitbucket installations
- [Leeway](https://github.com/khulnasoft/leeway) - A heavily caching build system
- [Dazzle](https://github.com/khulnasoft/dazzle) - An experimental Docker image builder
- [Werft](https://github.com/csweichel/werft) - A Kubernetes native CI system

## Code of Conduct

We want to create a welcoming environment for everyone interested in contributing to Devtool or participating in discussions with the Devtool community.
This project has adopted the [Contributor Covenant Code of Conduct](https://github.com/khulnasoft/.github/blob/main/CODE_OF_CONDUCT.md), [version 2.0](https://www.contributor-covenant.org/version/2/0/code_of_conduct/).
