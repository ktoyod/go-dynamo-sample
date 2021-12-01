import {
  expect as expectCDK,
  matchTemplate,
  MatchStyle,
} from '@aws-cdk/assert';
import * as cdk from '@aws-cdk/core';
import * as GoDynamoDBSample from '../lib/go-dynamodb-sample-stack';

test('Empty Stack', () => {
  const app = new cdk.App();
  // WHEN
  const stack = new GoDynamoDBSample.GoDynamoDBSampleStack(app, 'MyTestStack');
  // THEN
  expectCDK(stack).to(
    matchTemplate(
      {
        Resources: {},
      },
      MatchStyle.EXACT
    )
  );
});
