# aws-closest-region
Automatically select the closest AWS region and reduce latency 🔥

[![Build Status](https://travis-ci.org/mtojek/aws-closest-region.svg?branch=master)](https://travis-ci.org/mtojek/aws-closest-region)

Status: **Done**

Find the **closest AWS region** to your customers and **reduce the latency** while accessing the service endpoint. The application can be used in both - **local development** environment and **remote production hosts**.

```aws-closest-region``` can automatically discover new launched regions!

## Features

 * support **100% AWS regions** in all non-gov partitions (incl. China)
 * support **100% AWS services** (also brand new ones)
 * **quick integration** with shell scripts (``AWS_DEFAULT_REGION=`aws-closest-region` ``)
 * automatically detect next launched endpoints (use AWS SDK models)
 * verbose mode to show all latencies

## Snapshots

<img src="https://github.com/mtojek/aws-closest-region/blob/master/snapshots/snapshot-1.png" alt="Screenshot Desktop" width="720px" />


