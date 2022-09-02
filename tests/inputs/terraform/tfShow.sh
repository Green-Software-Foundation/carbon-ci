#!/bin/bash
terraform show -json plan.tfplan > plan.json
exit 0