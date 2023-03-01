# Demo of the Carbon CI Pipeline

**You can see a sample output of running a GitHub Action using Carbon CI Pipeline Tooling in a [Pull Request comment here](https://github.com/ava-innersource/greensoftware-demo/pull/20)**

Example output:
<img width="914" alt="Github bot comment with carbon info" src="https://user-images.githubusercontent.com/26546660/214556456-03dd9073-80bb-4974-a94c-062f7461f608.png">

## Description

This is a demo usage of the Green Software Foundation's [Carbon CI Pipeline Tooling](https://github.com/Green-Software-Foundation/Carbon_CI_Pipeline_Tooling) project,
applied to GitHub Actions workflow. It is a static analysis tool for cloud infrastructure, by looking at Infra-as-Code files available in your repository.
Current version supports ARM, Pulumi and Terraform files.

The GitHub Action invoking the program can be found in [/.github/workflows/calculate_iac_carbon.yml](/.github/workflows/calculate_iac_carbon.yml).
In this demo, the Action points to the Dev branch of the main repository, available at [https://github.com/Green-Software-Foundation/carbon-ci/tree/Dev](https://github.com/Green-Software-Foundation/carbon-ci/tree/Dev).

The action looks at the [infra-as-code](https://learn.microsoft.com/en-us/devops/deliver/what-is-infrastructure-as-code) [ARM template file](https://learn.microsoft.com/en-us/azure/azure-resource-manager/templates/),
located in [/data/azuredeploy.json](/data/azuredeploy.json), which is an example code of a Resource Group containing a webapp and an Azure VM.
> NOTE: At the time of making this demo, the Carbon CI Pipeline Project looked at a limited number of Azure components (some VMs). This may have changed by now,
> reach out to the Green Software Foundation (You can find the currently appointed chair of the project [here](https://github.com/Green-Software-Foundation/carbon-ci#appointments)).

## Set up guide

### Prerequisites

In order to run this demo, you will need:
- GitHub account
- [WattTime Pro License](https://www.watttime.org/get-the-data/data-plans/) or [ElectricityMaps Key active in the regions of the demo deployment](https://api-portal.electricitymaps.com/)

If you do not have a WattTime Pro License nor an ElectricityMaps Key.
## 1. Set up GitHub Repository and GitHub Secrets

- You need a GitHub repository to run the demo in, see the [official Github Documentation](https://docs.github.com/en/get-started/quickstart/create-a-repo) for the steps.
- Make sure GitHub Actions are enabled: https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-github-actions-settings-for-a-repository
You will need to set up GitHub Repository Secrets, see the [official GitHub Documentation](https://docs.github.com/en/actions/security-guides/encrypted-secrets) for a detailed description of setting them up.
In short: In your repository, navigate to `Settings>Secrets and variables>Actions`. Select the `Secrets` tab and set up the following with your values (make sure to set all secrets, leave empty if you dont have the value):

- `ELECTRICITY_MAP_AUTH_TOKEN` - as your ElectricityMaps authentication Token SET UP AS EMPTY IF YOU DONT HAVE A KEY
- `WATT_TIME_PASS` - as your WattTime User Password SET UP AS EMPTY IF YOU DONT HAVE A KEY
- `WATT_TIME_USER` - as your WattTime Username SET UP AS EMPTY IF YOU DONT HAVE A KEY

**NOTE: You need to have at least one of these with an actual key for the demo to work**

## 2. Upload Demo files and create a new branch

You will need to copy 2 files over: [/.github/workflows/calculate_iac_carbon.yml](/.github/workflows/calculate_iac_carbon.yml) to the corresponding directory in your repository
and [/data/azuredeploy.json](/data/azuredeploy.json). **Make sure to create the missing directories if they are missing from your repository! (`.github/workflows` and `data`)**


**NOTE: By default, the GitHub action is setup to use WattTime as the Carbon Intensity Data Source. To use ElectricityMaps instead, see section 2b before continuing. Otherwise go to the next step**


Create a new branch from which you will test the demo. See the guide here: https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-and-deleting-branches-within-your-repository .
You can name it `test-branch`.


### 2b. Setting up ElectricityMaps as the Carbon Intensity Data Source for the GitHub Action

Open the [/.github/workflows/calculate_iac_carbon.yml](/.github/workflows/calculate_iac_carbon.yml) file in your repository and edit line 26: 
https://github.com/Willmish/Carbon_CI_Pipeline_Tooling/blob/292dc33f74953127011032dee69eddcff2733937/.github/workflows/calculate_iac_carbon.yml#L26
Change the value set to `CARBON_RATE_PROVIDER` from `watttime` to `electricitymap`. Commit the changes and you are all set!

## 3. Test the demo

In order to test the demo, you will need to open a Pull Request to the main branch, from the newly created `test-branch`. See the [Official Github Documentation for opening a Pull Request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request).

Once the Pull Request is open, a Github Action should start running, checking the carbon footprint of your infrastructure. Once its done, refresh the Pull Request page and you will see a comment from a GitHub bot,
stating the Carbon Footprint. You can also [view the GitHub Action logs for more details on how the app ran](https://docs.github.com/en/actions/monitoring-and-troubleshooting-workflows/using-workflow-run-logs).

<img width="914" alt="Github bot comment with carbon info" src="https://user-images.githubusercontent.com/26546660/214556456-03dd9073-80bb-4974-a94c-062f7461f608.png">
