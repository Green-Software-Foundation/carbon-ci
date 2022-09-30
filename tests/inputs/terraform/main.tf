terraform {
  required_version = "1.2.5"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.15.1"
    }
  }
  
  backend "azurerm" {}
}

provider "azurerm" {
  features {}
}

locals {
  prefix   = "gsf"
  location = "eastus"
  username = "mruser"
  password = "mrpswrd1234!"
  
  other_linux_vms = [
    {
      name     = "${local.prefix}-vm-linux2"
      nic_name = "${local.prefix}-nic-linux2"
      hostname = "${local.prefix}-host-linux2"
    },
	{
      name     = "${local.prefix}-vm-linux3"
      nic_name = "${local.prefix}-nic-linux3"
      hostname = "${local.prefix}-host-linux3"
    }	
  ]
}

resource "azurerm_resource_group" "main" {
  name     = "${local.prefix}-rg"
  location = local.location
}

resource "azurerm_virtual_network" "main" {
  name                = "${local.prefix}-network"
  address_space       = ["10.0.0.0/22"]
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
}

resource "azurerm_subnet" "linux" {
  name                 = "${local.prefix}-linux-subnet"
  resource_group_name  = azurerm_resource_group.main.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = ["10.0.2.0/27"]
}

resource "azurerm_subnet" "windows" {
  name                 = "${local.prefix}-windows-subnet"
  resource_group_name  = azurerm_resource_group.main.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = ["10.0.2.32/27"]
}

resource "azurerm_network_interface" "linux1" {
  name                = "${local.prefix}-nic-linux1"
  resource_group_name = azurerm_resource_group.main.name
  location            = azurerm_resource_group.main.location

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.linux.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_linux_virtual_machine" "main" {
  name                            = "${local.prefix}-vm-linux1"
  resource_group_name             = azurerm_resource_group.main.name
  location                        = azurerm_resource_group.main.location
  size                            = "Standard_D2s_v3"
  admin_username                  = local.username
  admin_password                  = local.password
  disable_password_authentication = false
  network_interface_ids = [
    azurerm_network_interface.linux1.id,
  ]

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }
}

resource "azurerm_network_interface" "windows1" {
  name                = "${local.prefix}-nic-windows1"
  resource_group_name = azurerm_resource_group.main.name
  location            = azurerm_resource_group.main.location

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.windows.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_windows_virtual_machine" "main" {
  name                            = "${local.prefix}-vm-windows1"
  resource_group_name             = azurerm_resource_group.main.name
  location                        = azurerm_resource_group.main.location
  size                            = "Standard_F2"
  admin_username                  = local.username
  admin_password                  = local.password
  network_interface_ids = [
    azurerm_network_interface.windows1.id,
  ]

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2016-Datacenter"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }
}

module "other_linux_vms" {
  for_each       = { for lvm in local.other_linux_vms : lvm.name => lvm }
  source         = "./modules/azurerm_virtual_machine"
  nic_name       = each.value.nic_name
  rg_name        = azurerm_resource_group.main.name
  location       = azurerm_resource_group.main.location
  subnet_id      = azurerm_subnet.linux.id
  vm_name        = each.key
  hostname       = each.value.hostname
  admin_username = local.username
  admin_password = local.password
}
