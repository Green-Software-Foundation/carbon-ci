resource "azurerm_network_interface" "windows1" {
  name                = var.windows_nic_name
  resource_group_name = var.rg_name
  location            = var.location

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.windows_subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_windows_virtual_machine" "main" {
  name                            = var.windows_vm_name
  resource_group_name             = var.rg_name
  location                        = var.location
  size                            = "Standard_F2"
  admin_username                  = var.admin_username
  admin_password                  = var.admin_password
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
  for_each       = { for lvm in var.other_linux_vms : lvm.name => lvm }
  source         = "../resource_modules/azurerm_virtual_machine"
  nic_name       = each.value.nic_name
  rg_name        = var.rg_name
  location       = var.location
  subnet_id      = var.linux_subnet_id
  vm_name        = each.key
  hostname       = each.value.hostname
  admin_username = var.admin_username
  admin_password = var.admin_password
}
