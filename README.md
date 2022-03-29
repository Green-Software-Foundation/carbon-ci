# Carbon_CI_Pipeline_Tooling

## Scope

This project aims to build the capability to calculate the carbon emissions of an application via static analysis of the code and any infrastructure as code files in a repository.

This tooling covers the pipeline from a git commit, up to deployment, in a typical continuous integration/continuous deployment process.

## Appointments

- Fergus Kidd (Avanade)

## Projects

- [Innovation WG](https://greensoftware.foundation/working-groups/innovation)

## Resources

- [Slack Channel](https://greensoftwarefdn.slack.com/archives/C038YRLD2NN) (Members Only)

## Getting Started

Welcome to the Innovation Working Group.

This working group is part of the Green Software Foundation. It is open to any member of the Foundation, for more details regarding the foundation and the application form to join please go to https://greensoftware.foundation.

- Make sure you have joined the GSF slack and have introduced yourself in the Innovatation slack channel.
- We have weekly meetings over zoom, ensure you have access to the above meeting schedule.
- Familiarize yourself with our scope above and take a look at our current projects in flight - links are above.
- If you are having any problems with access please reach out to smcilroy@contractor.linuxfoundation.org or helpdesk@greensoftware.io.

## Copyright

Innovation WG projects are copyrighted under [Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/).

## Patent

No Patent License. No patent licenses are granted for the Draft Deliverables or Approved Deliverables developed by this Working Group.

## License

Innovation WG projects are licensed under the MIT License - see the [License.md](license/innovation-wg-license.md)file for details

## Dataset

CDLA-Permissive-1.0

## Collaborating With the WG

1. Create a [new Issue](https://github.com/Green-Software-Foundation/standards_wg/issues/new)
2. Discuss Issue with WG --> Create PR if required
3. PR to be submitted against the **DEV feature branch**
4. PR discussed with the WG. If agreed, the WG Chair will merge into **DEV Feature branch**

<figure>
	<img src="images/single-trunk-branch.svg" alt="GSF Single-Trunk Based Branch Flow">
	<figcaption></figcaption>
</figure>

5. See [The Way we Work](https://github.com/Green-Software-Foundation/standards_wg/blob/main/the_way_we_work.md) for futher details.

## Help

helpdesk@greensoftware.io

# carbon-measurement

> A repository to demonstrate ways to measure carbon as part of a CI/CD pipeline.

[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)
![GitHub issues](https://img.shields.io/github/issues/ava-innersource/carbon-measurement)
![GitHub](https://img.shields.io/github/license/ava-innersource/carbon-measurement)
![GitHub Repo stars](https://img.shields.io/github/stars/ava-innersource/carbon-measurement?style=social)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](https://avanade.github.io/code-of-conduct/)
[![Incubating InnerSource](https://img.shields.io/badge/Incubating-Ava--Maturity-%23FF5800?labelColor=yellow)](https://avanade.github.io/maturity-model/)

## Overview

This repository contains examples to demonstrate how to measure carbon during the development process.

Considered approaches include:

1. A GitHub Action to run during deployment
2. A GitHub Action to run on a pull request
3. A Pre-commit application to run every time a commit is made

## Licensing

carbon-measurement is available under the [MIT Licence](./LICENCE). It will be donated to the Green Software Foundation if the [Carbon Pipeline Tooling](https://github.com/Green-Software-Foundation/innovation_wg/issues/18) proposal is accepted.

## Solutions Referenced

- [GitHub Actions for Azure](https://docs.microsoft.com/en-us/azure/developer/github/github-actions?WT.mc_id=AI-MVP-5004204)
- [GitHub Actions for .NET](https://docs.microsoft.com/en-us/dotnet/devops/github-actions-overview?WT.mc_id=AI-MVP-5004204)
- [Create a VM](https://docs.microsoft.com/en-us/azure/templates/microsoft.compute/virtualmachines?WT.mc_id=AI-MVP-5004204)
- [Deploy an ARM template](https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/deploy-portal?WT.mc_id=AI-MVP-5004204)
- [Git Hooks](https://githooks.com/)
- [Electricitymap API](https://static.electricitymap.org/api/docs/index.html)
- [Creating a Docker Action](https://docs.github.com/en/actions/creating-actions/creating-a-docker-container-action)

## Documentation

Limited documentation has been created, as this is a work in progress. Further documentation will be added to the `docs` folder, along with setup instructions.

## Contact

Feel free to [raise an issue on GitHub](https://github.com/ava-innersource/carbon-measurement/issues), or see our [security disclosure](./SECURITY.md) policy.

## Contributing

Contributions are welcome. See information on [contributing](./CONTRIBUTING.md), as well as our [code of conduct](https://avanade.github.io/code-of-conduct/).

If you're happy to follow these guidelines, then check out the [getting started](./docs/start-here.md) guide.

## Who is Avanade?

[Avanade](https://www.avanade.com) is the leading provider of innovative digital, cloud and advisory services, industry solutions and design-led experiences across the Microsoft ecosystem.
