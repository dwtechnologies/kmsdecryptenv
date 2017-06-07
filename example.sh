#!/bin/bash

# Please set this to an KMS encrypted string
export TEST_KMS_DECRYPT=ENCRYPTED_TEST_VAR
echo "Encrypted value: $TEST_KMS_DECRYPT"

# Decrypt using kmsdecryptenv
export $(./kmsdecryptenv)
echo "Unencrypted value: $TEST"