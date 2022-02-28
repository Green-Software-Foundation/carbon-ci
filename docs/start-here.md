# Documentation Overview

This repository provides a sample of patterns and practices for Avanade projects.
```
Update this description.
DELETE THIS COMMENT
```

## Related

- **[How to contribute](../CONTRIBUTING.md)**
- **[Our code of conduct](../CODE_OF_CONDUCT.md)**

## Folder Structure
```
Add additional folders that your project uses
DELETE THIS COMMENT
```
- docs - documentation, process flows and state diagrams, and technical architecture.
- test - automated tests
- .vscode - configuration and recommendations for the VS Code IDE.
- .devcontainer - default configuration for GitHub codespaces or containerized development.

# Getting started
## Installation

### Software Installation for Node
```
Delete this section if your project doesn't use node
DELETE THIS COMMENT
```

1. Open the `.env.template` file, and update with your settings - save it as `.env`
2. Run `npm install` to ensure you have all dependencies.
3. Run the development server with `npm run dev`
4. Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Python Tooling
```
Delete this section if your project doesn't use python
DELETE THIS COMMENT
```

The tooling uses Python, and we recommend using a conda environment when installing requirements, for example:

```bash
$ conda create --name distribdata python=3.9 -y
$ conda activate distribdata
$ pip install -r requirements-dev.txt
```

### Software Installation for Python
```
Delete this section if your project doesn't use python
DELETE THIS COMMENT
```

1. Create a separate Python environment for your installation, and activate it. You have two options:

   a. _Use a Conda distribution_

   If you are using a distribution of conda, you may want to create a new conda environment, rather than use venv:

   `conda create --name distribdata python=3.9 -y`

   b. _Use a Python virtual environment_

   On Windows, you may need to use `python` command where there are references to the `python3` command.

   On linux, you may need to run `sudo apt-get install python3-venv` first.

   ```bash
   $ python3 -m venv env
   $ source env/bin/activate
   $ pip3 install -r requirements-dev.txt
   ```

2. Install the required dependencies in your new Python environment.

   ```bash
   $ pip3 install -r requirements-dev.txt
   ```

   The `requirements.txt` file can be used alone if you don't intend to develop further.

### Running the SQL code directly
```
Delete this section if your project doesn't have SQL
DELETE THIS COMMENT
```

You can use the [SQL Server](https://marketplace.visualstudio.com/items?itemName=ms-mssql.mssql) extension for VS Code to run the SQL ledger examples directly. On non-windows machines, make sure you have drivers for MS SQL installed.

On a mac, you can run the following commands:

```
brew install unixodbc
brew tap microsoft/mssql-release https://github.com/Microsoft/homebrew-mssql-release
brew update
brew install msodbcsql mssql-tools
```