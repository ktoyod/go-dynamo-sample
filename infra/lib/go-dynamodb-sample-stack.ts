import * as cdk from '@aws-cdk/core';

import * as dynamodb from '@aws-cdk/aws-dynamodb';
import * as ec2 from '@aws-cdk/aws-ec2';
import * as ecs from '@aws-cdk/aws-ecs';
import * as elbv2 from '@aws-cdk/aws-elasticloadbalancingv2';
import * as iam from '@aws-cdk/aws-iam';
import { CfnOutput } from '@aws-cdk/core';

export class GoDynamoDBSampleStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const table = new dynamodb.Table(this, 'my-table', {
      partitionKey: { name: 'id', type: dynamodb.AttributeType.STRING },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
    });

    const vpc = new ec2.Vpc(this, 'my-vpc');

    const taskRole = new iam.Role(this, 'my-task-role', {
      assumedBy: new iam.ServicePrincipal('ecs-tasks.amazonaws.com'),
    });

    const cluster = new ecs.Cluster(this, 'my-cluster', { vpc: vpc });

    const taskDefinition = new ecs.FargateTaskDefinition(this, 'my-task-def', {
      memoryLimitMiB: 512,
      cpu: 256,
      taskRole: taskRole,
    });
    const container = taskDefinition.addContainer('my-container', {
      image: ecs.ContainerImage.fromAsset('../app'),
      logging: ecs.LogDriver.awsLogs({ streamPrefix: 'my-container-log' }),
      environment: {
        DYNAMODB_TABLE: table.tableName,
      },
    });
    container.addPortMappings({
      containerPort: 8080,
      protocol: ecs.Protocol.TCP,
    });
    const service = new ecs.FargateService(this, 'my-service', {
      cluster,
      taskDefinition,
      desiredCount: 1,
    });

    const lb = new elbv2.ApplicationLoadBalancer(this, 'my-lb', {
      vpc,
      internetFacing: true,
    });
    const allowed_ip = process.env.ALLOWED_IP || '0.0.0.0/0';
    lb.connections.allowFrom(ec2.Peer.ipv4(allowed_ip), ec2.Port.tcp(80));
    const listener = lb.addListener('my-listener', { port: 80, open: false });
    service.registerLoadBalancerTargets({
      containerName: 'my-container',
      containerPort: 8080,
      newTargetGroupId: 'my-ecs',
      listener: ecs.ListenerConfig.applicationListener(listener, {
        protocol: elbv2.ApplicationProtocol.HTTP,
      }),
    });

    table.grantFullAccess(taskRole);

    new CfnOutput(this, 'LoadBalancerDNS', {
      value: lb.loadBalancerDnsName,
    });
  }
}
