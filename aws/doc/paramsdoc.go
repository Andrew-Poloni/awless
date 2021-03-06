package awsdoc

func TemplateParamsDoc(templateDef, param string) (string, bool) {
	if doc, ok := manualParamsDoc[templateDef][param]; ok {
		return doc, ok
	}
	doc, ok := generatedParamsDoc[templateDef][param]
	return doc, ok
}

var manualParamsDoc = map[string]map[string]string{
	"attachalarm": {
		"name":       "The Name of the Alarm to update",
		"action-arn": "The Amazon Resource Name (ARN) of the action to execute when this alarm transitions to the ALARM state from any other state",
	},
	"attachelasticip": {
		"allow-reassociation": "Specify false to ensure the operation fails if the Elastic IP address is already associated with another resource",
	},
	"attachinstance": {
		"id":   "The ID of the Instance",
		"port": "The port on which the Instance is listenning",
	},
	"attachpolicy": {
		"arn":   "The Amazon Resource Name (ARN) of the IAM policy you want to attach",
		"user":  "The name (friendly name, not ARN) of the IAM user to attach the policy to",
		"group": "The name (friendly name, not ARN) of the IAM group to attach the policy to",
		"role":  "The name (friendly name, not ARN) of the IAM role to attach the policy to",
	},
	"attachsecuritygroup": {
		"id":       "The ID of the Security Group to add to the instance",
		"instance": "The ID of the Instance",
	},
	"checkdatabase": {
		"id":      "The ID of the RDS Database to check",
		"state":   "The state of the RDS Database to reach (available | backing-up | creating | deleting | failed | maintenance | modifying | rebooting | renaming | resetting-master-credentials | restore-error | storage-full | upgrading | not-found)",
		"timeout": "The time (in seconds) after which the check is failed",
	},
	"checkdistribution": {
		"id":      "The ID of the CloudFront Distribution to check",
		"state":   "The state of the CloudFront Distribution to reach (Deployed | InProgress | not-found)",
		"timeout": "The time (in seconds) after which the check is failed",
	},
	"checkinstance": {
		"id":      "The ID of the EC2 Instance to check",
		"state":   "The state of the EC2 Instance to reach (pending | running | shutting-down | terminated | stopping | stopped | not-found)",
		"timeout": "The time (in seconds) after which the check is failed",
	},
	"checkloadbalancer": {
		"id":      "The ID of the ELBv2 Loadbalancer to check",
		"state":   "The state of the ELBv2 Loadbalancer to reach (provisioning | active | failed | not-found)",
		"timeout": "The time (in seconds) after which the check is failed",
	},
	"checkscalinggroup": {
		"name":    "The name of the AutoScaling Group to check",
		"count":   "The number of Instances + Loadbalancers + TargetGroups in the AutoScaling Group to reach",
		"timeout": "The time (in seconds) after which the check is failed",
	},
	"checksecuritygroup": {
		"id":      "The ID of the EC2 Security Group to check",
		"state":   "The state of the EC2 Security Group to reach (unused)",
		"timeout": "The time (in seconds) after which the check is failed",
	},
	"checkvolume": {
		"id":      "The ID of the EC2 Volume to check",
		"state":   "The state of the EC2 Volume to reach (available | in-use | not-found)",
		"timeout": "The time (in seconds) after which the check is failed",
	},
	"createaccesskey": {
		"user": "The name of the user for which the access key will be generated",
	},
	"createalarm": {
		"operator":           "The arithmetic operation to use when comparing the specified statistic and threshold (GreaterThanThreshold | LessThanThreshold | LessThanOrEqualToThreshold | GreaterThanOrEqualToThreshold)",
		"statistic-function": "The statistic for the metric associated with the alarm, other than percentile (Minimum | Maximum | Sum | Average | SampleCount | pNN.NN)",
		"unit":               "The unit of measure for the statistic (Seconds | Microseconds | Milliseconds | Bytes | Kilobytes | Megabytes | Gigabytes | Terabytes | Bits | Kilobits | Megabits | Gigabits | Terabits | Percent | Count | Bytes/Second | Kilobytes/Second | Megabytes/Second | Gigabytes/Second | Terabytes/Second | Bits/Second | Kilobits/Second | Megabits/Second | Gigabits/Second | Terabits/Second | Count/Second | None)",
	},
	"createbucket": {
		"acl":  "The canned ACL to apply to the bucket (private | public-read | public-read-write | aws-exec-read | authenticated-read | bucket-owner-read | bucket-owner-full-control | log-delivery-write)",
		"name": "The name of bucket to create",
	},
	"createdatabase": {
		"autoupgrade":       "Set to true to indicate that minor version patches are applied automatically",
		"backupwindow":      "Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod (format hh24:mi-hh24:mi)",
		"dbname":            "The name of the database to create when the DB instance is created",
		"dbsecuritygroups":  "A list of DB security groups to associate with this DB instance",
		"domain":            "Specify the Active Directory Domain to create the instance in.",
		"engine":            "Provides the name of the database engine to be used for this DB instance (mysql | mariadb | oracle-se1 | oracle-se2 | oracle-se | oracle-ee | sqlserver-ee | sqlserver-se | sqlserver-ex | sqlserver-web | postgres | aurora)",
		"iamrole":           "Specify the name of the IAM role to be used when making API calls to the Directory Service",
		"license":           "License model information for this DB instance (license-included | bring-your-own-license | general-public-license)",
		"optiongroup":       "Indicates that the DB instance should be associated with the specified option group",
		"password":          "The password for the master database user",
		"public":            "'true' specifies an Internet-facing instance with a publicly resolvable DNS name, which resolves to a public IP address. 'false' specifies an internal instance with a DNS name that resolves to a private IP address",
		"parametergroup":    "The name of the DB parameter group to associate with this DB instance",
		"port":              "The port number on which the database accepts connections",
		"storagetype":       "Specifies the storage type associated with DB instance (standard | gp2 | io1)",
		"subnetgroup":       "A DB subnet group to associate with this DB instance",
		"type":              "Contains the name of the compute and memory capacity class of the DB instance (db.t1.micro | db.m1.small | db.m1.medium | db.m1.large | db.m1.xlarge | db.m2.xlarge |db.m2.2xlarge | db.m2.4xlarge | db.m3.medium | db.m3.large | db.m3.xlarge | db.m3.2xlarge | db.m4.large | db.m4.xlarge | db.m4.2xlarge | db.m4.4xlarge | db.m4.10xlarge | db.r3.large | db.r3.xlarge | db.r3.2xlarge | db.r3.4xlarge | db.r3.8xlarge | db.t2.micro | db.t2.small | db.t2.medium | db.t2.large)",
		"vpcsecuritygroups": "A list of EC2 VPC security groups to associate with this DB instance",
	},
	"createdbsubnetgroup": {
		"description": "The description for the DB subnet group",
		"name":        "The name for the DB subnet group",
		"subnets":     "The EC2 Subnet IDs for the DB subnet group",
	},
	"createdistribution": {
		"origin-domain":   "The DNS name of the Amazon S3 bucket from which you want CloudFront to get objects for this origin, for example, myawsbucket.s3.amazonaws.com",
		"certificate":     "The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate you want to use for TSL connection",
		"comment":         "Any comments you want to include about the distribution",
		"default-file":    "The object that you want CloudFront to request from your origin (for example, index.html) when a viewer requests the root URL for your distribution (http://www.example.com)",
		"domain-aliases":  "A list of CNAMEs (alternate domain names), if any, for this distribution",
		"enable":          "From this field, you can enable or disable the selected distribution",
		"forward-cookies": "Specifies which cookies to forward to the origin for this cache behavior (all | none | whitelist)",
		"forward-queries": "Indicates whether you want CloudFront to forward query strings to the origin that is associated with this cache behavior and cache based on the query string parameters (true | false)",
		"https-behaviour": "The protocol (HTTP or HTTPS) that viewers can use to access the files (allow-all | redirect-to-https | https-only)",
		"origin-path":     "An optional element that causes CloudFront to request your content from a directory in your Amazon S3 bucket or your custom origin. When you include this element, specify the directory name, beginning with a /",
		"price-class":     "The price class that corresponds with the maximum price that you want to pay for CloudFront service. If you specify PriceClass_All, CloudFront responds to requests for your objects from all CloudFront edge locations",
		"min-ttl":         "The minimum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated",
	},
	"createelasticip": {
		"domain": "Set to vpc to allocate the address for use with instances in a VPC else the address is for use with instances in EC2-Classic (vpc | ec2-classic)",
	},
	"createfunction": {
		"bucket":        "Amazon S3 bucket name where the .zip file containing your deployment package is stored. This bucket must reside in the same AWS region where you are creating the Lambda function",
		"object":        "The Amazon S3 object (the deployment package) key name you want to upload",
		"objectversion": "The Amazon S3 object (the deployment package) version you want to upload",
		"runtime":       "The runtime environment for the Lambda function you are uploading (python3.6 | python2.7 | nodejs6.10 | nodejs4.3)",
		"zipfile":       "The path toward the zip file containing your deployment package",
	},
	"creategroup": {
		"name": "The name of the group to create",
	},
	"createinstance": {
		"count": "The number of instances to launch",
		"name":  "The name of the instance to launch",
		"role":  "The name of the instance profile (role) to launch the instance with",
		"image": "The ID of the AMI of the instance to launch, which you can get by using `awless search images`",
	},
	"createkeypair": {
		"name":      "The name of the keypair to create (it will also be the name of the file stored in ~/.awless/keys)",
		"encrypted": "Set to 'true' if you want to encrypt the keypair"},
	"createlaunchconfiguration": {
		"public": "Used for groups that launch instances into a virtual private cloud (VPC). Specifies whether to assign a public IP address to each instance",
	},
	"createlistener": {
		"actiontype":  "The type of action (forward)",
		"targetgroup": "The Amazon Resource Name (ARN) of the target group",
		"certificate": "The Amazon Resource Name (ARN) of the certificate",
		"protocol":    "The protocol for connections from clients to the load balancer (TCP | HTTP | HTTPS)",
		"sslpolicy":   "The security policy that defines which ciphers and protocols are supported (ELBSecurityPolicy-2016-08 | ELBSecurityPolicy-TLS-1-2-2017-01 | ELBSecurityPolicy-TLS-1-1-2017-01 | ELBSecurityPolicy-2015-05 | ELBSecurityPolicy-TLS-1-0-2015-04)",
	},
	"createloadbalancer": {
		"scheme": "The routing range of the loadbalancer (internet-facing | internal)",
		"iptype": "The type of IP addresses used by the subnets for your load balancer: IPv4 or IPv4 and IPv6 (ipv4 | dualstack)",
	},
	"createpolicy": {
		"name":        "The friendly name of the policy",
		"description": "A friendly description of the policy",
		"effect":      "The Effect element is required and specifies whether the policy will result in an allow or an explicit deny (Allow | Deny)",
		"action":      "The Action element describes the specific action or actions that will be allowed or denied. You specify a value using a namespace that identifies a service followed by the name of the action to allow or deny (eg. sqs:SendMessage, s3:*)",
		"resource":    "The Amazon Resource Name (ARN) of the Resource element which specifies the object or objects that the policy covers",
	},
	"createqueue": {
		"delay":              "The length of time, in seconds, for which the delivery of all messages in the queue is delayed. Valid values: An integer from 0 to 900 seconds (15 minutes). The default is 0",
		"max-msg-size":       "The limit of how many bytes a message can contain before Amazon SQS rejects it. Valid values: An integer from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default is 262,144 (256 KiB)",
		"retention-period":   "The length of time, in seconds, for which Amazon SQS retains a message. Valid values: An integer from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default is 345,600 (4 days)",
		"policy":             "The queue's policy",
		"msg-wait":           "The length of time, in seconds, for which a ReceiveMessage action waits for a message to arrive. Valid values: An integer from 0 to 20 (seconds). The default is 0",
		"redrive-policy":     "The parameters for the dead letter queue functionality of the source queue",
		"visibility-timeout": "The visibility timeout for the queue. Valid values: An integer from 0 to 43,200 (12 hours). The default is 30",
	},
	"createrecord": {
		"zone":    "The ID of the hosted zone that contains the resource record sets that you want to change",
		"name":    "The name of the domain you want to perform the action on. Enter a fully qualified domain name, for example, www.example.com. You can optionally include a trailing dot",
		"type":    "The DNS record type. (A | AAAA | CNAME | MX | NAPTR | NS | PTR | SOA | SPF | SRV | TXT)",
		"value":   "The current or new DNS record value",
		"ttl":     "The resource record cache time to live (TTL), in seconds",
		"comment": "Any comments you want to include about a change batch request",
	},
	"createrole": {
		"name":              "The name of the role to create",
		"principal-account": "The ID of the account that can perform actions and access resources of the role (you can know your account ID with `awless whoami`)",
		"principal-user":    "The Amazon Resource Name (ARN) of the user that can perform actions and access resources of the role",
		"principal-service": "The AWS Service that can assume this role to perform actions and access resources of the role (e.g. 'ec2.amazonaws.com')",
		"sleep-after":       "The amount of time in seconds you want to wait after creating the role (usually used to be sure that the role creation has been propagated)",
	},
	"creates3object": {
		"bucket": "Name of the bucket to which object will be added",
		"file":   "The path toward to file to upload",
		"name":   "The name of the Object to create (by default the file name is used)",
		"acl":    "The canned ACL to apply to the object (private | public-read | public-read-write | aws-exec-read | authenticated-read | bucket-owner-read | bucket-owner-full-control | log-delivery-write)",
	},
	"createscalinggroup": {
		"healthcheck-type": "The service to use for the health checks (EC2 | ELB)",
	},
	"createscalingpolicy": {
		"adjustment-type":    "The adjustment type (ChangeInCapacity | ExactCapacity | PercentChangeInCapacity)",
		"adjustment-scaling": "The amount by which to scale, based on the specified adjustment type (e.g. '-2', '3')",
	},
	"createstack": {
		"capabilities":  "A list of values that you must specify before AWS CloudFormation can create certain stacks (CAPABILITY_IAM | CAPABILITY_NAMED_IAM)",
		"on-failure":    "Determines what action will be taken if stack creation fails (DO_NOTHING | ROLLBACK | DELETE)",
		"parameters":    "A list of Parameters that specify input parameters for the stack given using this format: key1:val1,key2:val2,...",
		"policy-file":   "The path to the file containing the stack policy body",
		"template-file": "The path to the file containing the template body with a minimum size of 1 byte and a maximum size of 51,200 bytes",
	},
	"createsubnet": {
		"name": "The 'Name' Tag for the subnet to create",
	},
	"createsubscription": {
		"endpoint": "The endpoint that you want to receive notifications. Endpoints vary by protocol: For the http or https protocol, the endpoint is a URL beginning with 'http://' or 'https://', for the email or email-json protocol, the endpoint is an email address, for the sms protocol, the endpoint is a phone number of an SMS-enabled, for the sqs protocol, the endpoint is the ARN of an Amazon SQS queue, for the application protocol, the endpoint is the EndpointArn of a mobile app and device, for the lambda protocol, the endpoint is the ARN of an AWS Lambda function.",
		"protocol": "The protocol you want to use (http | https | email | email-json | sms | sqs | lambda)",
		"topic":    "The ARN of the topic you want to subscribe to",
	},
	"createtag": {
		"resource": "The ID of the resource on which you want to add a tag",
		"key":      "The Tag key",
		"value":    "The Tag value",
	},
	"createtargetgroup": {
		"matcher": "The HTTP codes to use when checking for a successful response from a target",
	},
	"createvpc": {
		"name": "The 'Name' Tag for the VPC to create",
	},
	"createzone": {
		"comment":   "Any comments that you want to include about the hosted zone",
		"isprivate": "A value that indicates whether this is a private hosted zone",
		"vpcid":     "(Private hosted zones only) The ID of an Amazon VPC",
		"vpcregion": "(Private hosted zones only) The region in which you created an Amazon VPC",
	},
	"deleteaccesskey": {
		"id": "The ID of the access key and secret access key you want to delete",
	},
	"deletealarm": {
		"name": "The name of the alarm(s) to be deleted",
	},
	"deletebucket": {
		"name": "The name of the bucket to be deleted",
	},
	"deletedatabase": {
		"id":            "The ID of the database to be deleted",
		"skip-snapshot": "Determines whether a final DB snapshot is created before the DB instance is deleted. If true is specified, no DBSnapshot is created. If false is specified, a DB snapshot is created before the DB instance is deleted",
		"snapshot":      "The ID of the new DBSnapshot created when skip-snapshot=false",
	},
	"deletedbsubnetgroup": {
		"name": "The name of the database subnet group to be deleted",
	},
	"deletedistribution": {
		"id": "The ID of the distribution to be deleted",
	},
	"deletefunction": {
		"id": "The ID of the Lambda function to be deleted",
	},
	"deleteimage": {
		"id":               "The ID of the AMI to be deleted",
		"delete-snapshots": "Set to 'true' to also delete the snapshots created from this image",
	},
	"deleteinstance": {
		"id": "The ID(s) of the instance(s) to be deleted",
	},
	"deleteinternetgateway": {
		"id": "The ID of the Internet gateway to be deleted",
	},
	"deletekeypair": {
		"name": "The name of the key pair to be deleted",
	},
	"deletelaunchconfiguration": {
		"name": "The name of the launch configuration to be deleted",
	},
	"deleterecord": {
		"zone":  "The ID of the hosted zone that contains the resource record sets that you want to delete",
		"name":  "The name of the domain you want to perform the action on. Enter a fully qualified domain name, for example, www.example.com. You can optionally include a trailing dot",
		"type":  "The DNS record type. (A | AAAA | CNAME | MX | NAPTR | NS | PTR | SOA | SPF | SRV | TXT)",
		"value": "The DNS record value to delete",
		"ttl":   "The resource record cache time to live (TTL), in seconds",
	},
	"deleterole": {
		"name": "The name of the role to be deleted",
	},
	"deletes3object": {
		"bucket": "The name of the bucket containing the object to be deleted",
		"name":   "The name (i.e. key) of the object to be deleted",
	},
	"deletetag": {
		"resource": "The ID of the resource on which you want to remove a tag",
		"key":      "The Tag key",
		"value":    "The Tag value",
	},
	"detachalarm": {
		"name":       "The name of the alarm",
		"action-arn": "The Amazon Resource Name (ARN) to be detached of the ALARM actions",
	},
	"detachinstance": {
		"id": "The ID of the instance to be detached from target group",
	},
	"detachpolicy": {
		"arn":   "The Amazon Resource Name (ARN) of the IAM policy you want to detach",
		"user":  "The name (friendly name, not ARN) of the IAM user to detach the policy to",
		"group": "The name (friendly name, not ARN) of the IAM group to detach the policy to",
		"role":  "The name (friendly name, not ARN) of the IAM role to detach the policy to",
	},
	"detachsecuritygroup": {
		"id":       "The ID of the security group",
		"instance": "The ID of the instance to be detached",
	},
	"importimage": {
		"architecture": "The architecture of the virtual machine (i386 | x86_64)",
		"url":          "The URL to the Amazon S3-based disk image being imported. The URL can either be a https URL (https://..) or an Amazon S3 URL (s3://..)",
		"snapshot":     "The ID of the EBS snapshot to be used for importing the snapshot",
		"bucket":       "The name of the S3 bucket where the disk image is located",
		"s3object":     "The name of the S3 object where the disk image is located",
		"license":      "The license type to be used for the Amazon Machine Image (AMI) after importing (AWS | BYOL)",
		"platform":     "The operating system of the virtual machine (Windows | Linux)",
	},
	"updatebucket": {
		"name":              "The name of the bucket to update",
		"acl":               "The canned ACL to apply to the bucket (private | public-read | public-read-write | aws-exec-read | authenticated-read | bucket-owner-read | bucket-owner-full-control | log-delivery-write)",
		"public-website":    "Set to 'true' if you want to publish the content of the bucket as a public HTTP website",
		"redirect-hostname": "Hostname where HTTP requests will be redirected when publishing website",
		"index-suffix":      "A suffix that is appended to a request that is for a directory on the website endpoint (e.g. if the suffix is index.html and you make a request to samplebucket/images/ the data that is returned will be for the object with the key name images/index.html)",
		"enforce-https":     "Use HTTPS rather than HTTP when redirecting requests",
	},
	"updatedistribution": {
		"id":     "The ID of the distribution to update",
		"enable": "Enable/Disable the distribution (True | False)",
	},
	"updateinstance": {
		"type": "Changes the instance type to the specified value",
	},
	"updates3object": {
		"acl":     "The canned ACL to apply to the bucket (private | public-read | public-read-write | aws-exec-read | authenticated-read | bucket-owner-read | bucket-owner-full-control | log-delivery-write)",
		"bucket":  "The name of the bucket containing the object to be updated",
		"name":    "The name of the object to be updated",
		"version": "Used to reference a specific version of the object",
	},
	"updatesecuritygroup": {
		"id":        "The ID of the security group to be updated",
		"cidr":      "The CIDR IPv4 address range",
		"protocol":  "The IP protocol name (tcp, udp, icmp) or number. Use -1 to specify all protocols",
		"inbound":   "Set inbound to either authorize or revoke, to update the security group ingress rules (authorize | revoke)",
		"outbound":  "Set outbound to either authorize or revoke, to update the security group egress rules (authorize | revoke)",
		"portrange": "The portrange for the rule to update (any | 80 | 22-23 ...)",
	},
	"updatestack": {
		"capabilities":       "A list of values that you must specify before AWS CloudFormation can update certain stacks (CAPABILITY_IAM | CAPABILITY_NAMED_IAM)",
		"parameters":         "A list of Parameters that specify input parameters for the stack given using this format: key1:val1,key2:val2,...",
		"policy-file":        "The path to the file containing the stack policy body",
		"policy-update-file": "The path to the file containing the temporary overriding stack policy",
		"template-file":      "The path to the file containing the template body with a minimum size of 1 byte and a maximum size of 51,200 bytes",
	},
}
