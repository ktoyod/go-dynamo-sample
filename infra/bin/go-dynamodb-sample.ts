#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { GoDynamoDBSampleStack } from '../lib/go-dynamodb-sample-stack';

const app = new cdk.App();
new GoDynamoDBSampleStack(app, 'InfraStack', {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
});
