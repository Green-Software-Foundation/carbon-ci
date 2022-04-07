# carbon_ci_pipeline_tooling

> Calculate the carbon emissions of an application via static analysis of the code and any infrastructure as code files in a repository.

[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)
![GitHub issues](https://img.shields.io/github/issues/ava-innersource/carbon-measurement)
![GitHub](https://img.shields.io/github/license/ava-innersource/carbon-measurement)
![GitHub Repo stars](https://img.shields.io/github/stars/ava-innersource/carbon-measurement?style=social)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](https://avanade.github.io/code-of-conduct/)
[![Incubating InnerSource](https://img.shields.io/badge/Incubating-Ava--Maturity-%23FF5800?labelColor=yellow)](https://avanade.github.io/maturity-model/)

## Scope

This project aims to build the capability to calculate the carbon emissions of an application via static analysis of the code and any infrastructure as code files in a repository.

This tooling covers the pipeline from a git commit, up to deployment, in a typical continuous integration/continuous deployment process.

Considered approaches include:

1. A GitHub Action to run during deployment
2. A GitHub Action to run on a pull request
3. A Pre-commit application to run every time a commit is made

## Appointments

- Fergus Kidd (Avanade)

## Solutions Referenced

- [GitHub Actions for Azure](https://docs.microsoft.com/en-us/azure/developer/github/github-actions?WT.mc_id=AI-MVP-5004204)
- [GitHub Actions for .NET](https://docs.microsoft.com/en-us/dotnet/devops/github-actions-overview?WT.mc_id=AI-MVP-5004204)
- [Create a VM](https://docs.microsoft.com/en-us/azure/templates/microsoft.compute/virtualmachines?WT.mc_id=AI-MVP-5004204)
- [Deploy an ARM template](https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/deploy-portal?WT.mc_id=AI-MVP-5004204)
- [Git Hooks](https://githooks.com/)
- [Electricitymap API](https://static.electricitymap.org/api/docs/index.html)
- [Creating a Docker Action](https://docs.github.com/en/actions/creating-actions/creating-a-docker-container-action)

## Documentation

Limited documentation has been created, as this is not yet an approved GSF deliverable. Further documentation will be added to the `docs` folder, along with setup instructions.

## Working with the project

Contributions are welcome. See information on [contributing](./CONTRIBUTING.md).

If you're happy to follow these guidelines, then check out the [getting started](./docs/start-here.md) guide.

This project is part of the Green Software Foundation. It is open to any member of the Foundation, for more details regarding the foundation and the application form to join please go to https://greensoftware.foundation.

### Contributing and our ways of working

1. Create a [new Issue](https://github.com/Green-Software-Foundation/carbon_ci_pipeline_tooling/issues/new)
2. Discuss Issue with group --> Create PR if required
3. PR to be submitted against the **DEV feature branch**
4. PR discussed with the group. If agreed, the chair and maintainer will merge into **DEV Feature branch**

<figure>
	<img src="images/single-trunk-branch.svg" alt="GSF Single-Trunk Based Branch Flow">
	<figcaption></figcaption>
</figure>

5. See [The Way we Work](https://github.com/Green-Software-Foundation/standards_wg/blob/main/the_way_we_work.md) for futher details.

## Resources

- [Slack Channel](https://greensoftwarefdn.slack.com/archives/C038YRLD2NN) (Members Only)

## Licensing

carbon_ci_pipeline_tooling is available under the [MIT Licence](./LICENCE).

## Copyright

Innovation WG projects are copyrighted under [Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/).

## Patent

No Patent License. No patent licenses are granted for the Draft Deliverables or Approved Deliverables developed by this Working Group.

## Dataset

CDLA-Permissive-1.0

## Help

helpdesk@greensoftware.io
