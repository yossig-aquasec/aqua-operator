## General 

This guide explains how to deploy and use the Aqua Security Operator to manage Aqua's deployments in an OpenShift 4.x enviornemnt. You can use the Operator to deploy Aqua Entperiese or any of its components -
* Server (aka “console”)
* Database (optional; you can map an external database as well) 
* Gateway 
* Enforcer
* Scanner
* KubeEnforcer

Use the Aqua Operator to: 
* Easily deploy Aqua Enterprise on OpenShift
* Manage and scale up Aqua security components with additional replicas
* Assign metadata tags to Aqua Enterprise components
* Easily add and delete Aqua components like Scanner daemons, Kube-Enforcers and Enforcers
	
You can find all Aqua's Operator CRs and their properties at [Custom Resources](../deploy/crds). 
	   
## Prerequisites 

Make sure you have a license and access to the Aqua registry. To obtain a license, please contact Aqua Security at https://www.aquasec.com/about-us/contact-us/.

It is advised that you read about the [Aqua Environment and Configuration](https://docs.aquasec.com/docs/purpose-of-this-section) and [Aqua's sizing guide](https://docs.aquasec.com/docs/sizing-guide) before deploying and using the Operator. 

## Types of Aqua Opertaor
Aqua Security maintans three types of Operators:
* **Marketplace** - The marketplace operator is purchased through Red Hat Marketplace.
* **Community** - Aqua's official Operator. It typically represents the latest and newest version of the Operator. 
* **Certified** - Aqua's official Operator vetted and certified by RedHat. The certified Operator is based on the latest community Operator. It is packaged in a mode that allows it to work in disconnected networks, and contains UBI images as part of the package.  

## Deploying the Aqua Operator
1. Create a new namespace/project called "aqua" for the Aqua deployment.

2. Install the Aqua Operator from Red Hat's OperatorHub and add it to the "aqua" namespace. 

3. Create a secret for the database password
```
oc create secret generic aqua-database-password --from-literal=db-password=<password> -n aqua
```

4. To work with the community Operator, you need to create a registry secret to Aqua's images registry. Aqua's registry credentials are identical to the username and password for Aqua's supprot portal (https://success.aquasec.com.) -
```bash
oc create secret docker-registry aqua-registry --docker-server=registry.aquasec.com --docker-username=<AQUA_USERNAME> --docker-password=<AQUA_PASSWORD> --docker-email=<user email> -n aqua
```


## Deploying Aqua Enterprise using Custom Resources
The Aqua Operator includes a few CRDs to allow you to deploy Aqua in different configurations. Before you create your deployment CR, please review commons CR examples in the section *CR Examples* below.


**[AquaCSP CRD](../deploy/crds/operator.aquasec.com_aquacsps_crd.yaml)** provides the fastest methods to deploy Aqua Enterprise in a single cluster. AquaCSP defines how to deploy the Server, Gateway, Aqua Enforcer, and KubeEnforcer in the target cluster. Please see the [example CR](../deploy/crds/operator_v1alpha1_aquacsp_cr.yaml) for the listing of all fields and configurations.
* You can set the enforcement mode using the ```.spec.enforcer.enforceMode``` property in the CR file.
* You can deploy a Route by setting the  ```.spec.route``` property to "true".
* The default service type for the Console and Gateway is ClusterIP. You can change the service type in the CR.
* You can choose to deploy a different version of Aqua CSP by setting the ```.spec.infra.version```  property or change the image ```.spec.<<server/gateway/database>>.image.tag```.
* You can choose to use an external database by providing the ```.spec.externalDB```  property details.
* You can omit the Enforcer and KubeEnforcer components by removing them from the CR.
* You can add server/gateway environment variables with ```.spec.<<serverEnvs/gatewayEnvs>>``` (same convention of name value as k8s deployment).
* You can define the server/gateway resources requests/limits with ```.spec.<<server/gateway>>.resources```
* You can define mTLS by mounting the files (probably kept in a k8s secret) with ```.spec.<<server/gateway>>.volumes``` ```.spec.<<server/gateway>>.volumeMounts``` and adding the relevant environment variables.


The **[AquaServer CRD](../deploy/crds/operator_v1alpha1_aquaserver_cr.yaml)**, **[AquaDatabase CRD](../deploy/crds/operator_v1alpha1_aquadatabase_cr.yaml)**, and **[AquaGateway CRD](../deploy/crds/operator_v1alpha1_aquagateway_cr.yaml)** are used for advanced configurations where the server components are deployed across multiple clusters.

**[AquaEnforcer CRD](../deploy/crds/operator_v1alpha1_aquaenforcer_cr.yaml)** is used to deploy the Aqua Enforcer in any cluster. Please see the [example CR](../deploy/crds/operator_v1alpha1_aquaenforcer_cr.yaml) for the listing of all fields and configurations.
* You need to provide a token to identify the Aqua Enforcer.
* You can set the target Gateway using the ```.spec.gateway.host```and ```.spec.gateway.port``` properties.
* You can choose to deploy a different version of the Aqua Enforcer by setting the ```.spec.deploy.image.tag``` property. 
    If you choose to run old/custom Aqua Enforcer version, you must set ```.spec.common.allowAnyVersion``` .
* You can add environment variables using ```.spec.env```.
* You can define the enforcer resources requests/limits using ```.spec.deploy.resources```.
* You can define mTLS by mounting the files (probably kept in a k8s secret) with ```.spec.deploy.volumes``` ```.spec.deploy.volumeMounts``` and adding the relevant environment variables.

**[AquaKubeEnforcer CRD](../deploy/crds/operator_v1alpha1_aquakubeenforcer_cr.yaml)** is used to deploy the KubeEnforcer in your target cluster. Please see the [example CR](../deploy/crds/operator_v1alpha1_aquakubeenforcer_cr.yaml) for the listing of all fields and configurations.
* You need to provide a token to identify the KubeEnforcer to the Aqua Server.
* You can set the target Gateway using the ```.spec.config.gateway_address```  property.
* You can choose to deploy a different version of the KubeEnforcer by setting the ```.spec.deploy.image.tag``` property.
    If you choose to run old/custom Aqua KubeEnforcer version, you must set ```.spec.allowAnyVersion``` .
* You can add environment variables using ```.spec.env```.
* You can define the kube-enforcer resources requests/limits using ```.spec.deploy.resources```.
* You can define mTLS by mounting the files (probably kept in a k8s secret) with ```.spec.deploy.volumes``` ```.spec.deploy.volumeMounts``` and adding the relevant environment variables.

**[AquaScanner CRD](../deploy/crds/operator_v1alpha1_aquascanner_cr.yaml)** is used to deploy the Aqua Scanner in any cluster. Please see the [example CR](../deploy/crds/operator_v1alpha1_aquascanner_cr.yaml) for the listing of all fields and configurations.
* You need to set the target Aqua Server using the ```.spec.login.host```  property.
* You need to provide the ```.spec.login.username``` and ```.spec.login.password``` to authenticate with the Aqua Server.
* You can set ``.spec.login.tlsNoVerify`` if you connect scanner to HTTPS server, and don't want to use mTLS verification.  
* You can choose to deploy a different version of the Aqua Scanner by setting the ```.spec.image.tag``` property.
    If you choose to run old/custom Aqua Scanner version, you must set ```.spec.common.allowAnyVersion``` .


## Operator Upgrades ##
**Major versions** - When switching from an older operator channel to this channel,
the operator will update the Aqua components to this channel Aqua version.

**Minor versions** - For the certified operator, the Aqua operator is using the "Seamless Upgrades" mechanism.
You can set the upgrade approval strategy in the operator subscription to either "Automatic" or "Manual".
If you choose "Automatic", the OLM will automatically upgrade to the latest operator version in this channel
with the latest Aqua components images. If you choose the "Manual" approval strategy, the OLM will
only notify you there is an update available, and will upgrade the operator version and Aqua components
only after a manual approval given.

*If you choose "Manual" approval strategy, we recommend that before approving, make sure you are indeed
upgrading to the most updated version. you can do that by deleting the suggested InstallPlan, once the 
suggested InstallPlan is deleted, the OLM will generate a new InstallPlan to the latest version available
in this channel*

For community operator, you can upgrade minor version by changing the relevant CRs 
```.spec.infra.version``` or ```.spec.<<NAME>>.image.tag```

	
## CR Examples ##

#### Example: Deploying the Aqua Server with an Aqua Enforcer and KubeEnforcer (all in one CR)

```yaml
---
apiVersion: operator.aquasec.com/v1alpha1
kind: AquaCsp
metadata:
  name: aqua
  namespace: aqua
spec:
  infra:                                    
    serviceAccount: "aqua-sa"               
    namespace: "aqua"                       
    version: "6.2"                          
    requirements: true                      
  common:
    imagePullSecret: "aqua-registry"        # Optional: If already created image pull secret then mention in here
    dbDiskSize: 10
    databaseSecret:                         # Optional: If already created database secret then mention in here
      key: "db-password"
      name: "aqua-database-password"      
  database:                                 
    replicas: 1                            
    service: "ClusterIP"
    image:
      registry: "registry.aquasec.com"
      repository: "database"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always                    
  gateway:                                  
    replicas: 1                             
    service: "ClusterIP"
    image:
      registry: "registry.aquasec.com"
      repository: "gateway"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always                     
  server:                                   
    replicas: 1                             
    service: "LoadBalancer" 
    image:
      registry: "registry.aquasec.com"
      repository: "console"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always 
  enforcer:                                 # Optional: If defined, the Operator will create the default Aqua Enforcer 
    enforcerMode: false                     # Defines whether the default Enforcer will work in "Enforce" (true) or "Audit Only" (false) mode
  kubeEnforcer:                             # Optional: If defined, the Operator will create a KubeEnforcer
    registry: "registry.aquasec.com"        
    tag: "<<IMAGE TAG>>" 
  route: true                               # Optional: If defined and set to true, the Operator will create a Route to enable access to the console
```

If you haven't used the "route" option in the Aqua CSP CR, you should define a Route manually to enable external access to the Aqua Server (Console).

#### Example: Simple deployment of the Aqua Server 

```yaml
---
apiVersion: operator.aquasec.com/v1alpha1
kind: AquaCsp
metadata:
  name: aqua
  namespace: aqua
spec:
  infra:                                    
    serviceAccount: "aqua-sa"               
    namespace: "aqua"                       
    version: "6.2"                          
    requirements: true                      
  common:
    imagePullSecret: "aqua-registry"        # Optional: If already created image pull secret then mention in here
    dbDiskSize: 10
    databaseSecret:                         # Optional: If already created database secret then mention in here
      key: "db-password"
      name: "aqua-database-password"      
  database:                                 
    replicas: 1                            
    service: "ClusterIP"
    image:
      registry: "registry.aquasec.com"
      repository: "database"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always                    
  gateway:                                  
    replicas: 1                             
    service: "ClusterIP"
    image:
      registry: "registry.aquasec.com"
      repository: "gateway"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always                     
  server:                                   
    replicas: 1                             
    service: "LoadBalancer" 
    image:
      registry: "registry.aquasec.com"
      repository: "console"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always  
  route: true                               # Optional: If defined and set to true, the Operator will create a Route to enable access to the console
```

If you haven't used the "route" option in the Aqua CSP CR, you should define a Route manually to enable external access to the Aqua Server (Console).

#### Example: Deploying Aqua Enterprise with split database

"Split database" means there is a separate database for audit-related data: 
```yaml
---
apiVersion: operator.aquasec.com/v1alpha1
kind: AquaCsp
metadata:
  name: aqua
  namespace: aqua
spec:
  infra:                                    
    serviceAccount: "aqua-sa"               
    namespace: "aqua"                       
    version: "6.2"                          
    requirements: true                      
  common:
    imagePullSecret: "aqua-registry"        # Optional: If already created image pull secret then mention in here
    dbDiskSize: 10
    databaseSecret:                         # Optional: If already created database secret then mention in here
      key: "db-password"
      name: "aqua-database-password"
    splitDB: true      
  database:                                 
    replicas: 1                            
    service: "ClusterIP"
    image:
      registry: "registry.aquasec.com"
      repository: "database"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always                    
  gateway:                                  
    replicas: 1                             
    service: "ClusterIP"
    image:
      registry: "registry.aquasec.com"
      repository: "gateway"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always                     
  server:                                   
    replicas: 1                             
    service: "LoadBalancer" 
    image:
      registry: "registry.aquasec.com"
      repository: "console"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always  
  route: true                               # Optional: If defined and set to true, the Operator will create a Route to enable access to the console
```

#### Example: Deploying Aqua Enterprise with an external database

```yaml
---
apiVersion: operator.aquasec.com/v1alpha1
kind: AquaCsp
metadata:
  name: aqua
  namespace: aqua
spec:
  infra:                                    
    serviceAccount: "aqua-sa"               
    namespace: "aqua"                       
    version: "6.2"                          
    requirements: true                      
  common:
    imagePullSecret: "aqua-registry"        # Optional: If already created image pull secret then mention in here
    dbDiskSize: 10      
  externalDb:
    host: "<<EXTERNAL DATABASE IP>>"
    port: "<<EXTERNAL DATABASE PORT>>"
    username: "<<EXTERNAL DATABASE USER NAME>>"
    password: "<<EXTERNAL DATABASE PASSWORD>>"    # Optional: you can specify the database password secret in common.databaseSecret                     
  gateway:                                  
    replicas: 1                             
    service: "ClusterIP"
    image:
      registry: "registry.aquasec.com"
      repository: "gateway"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always                     
  server:                                   
    replicas: 1                             
    service: "LoadBalancer" 
    image:
      registry: "registry.aquasec.com"
      repository: "console"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always  
  route: true                               # Optional: If defined and set to true, the Operator will create a Route to enable access to the console
```

### Example: Deploying Aqua Enterprise with a split external database

```yaml
---
apiVersion: operator.aquasec.com/v1alpha1
kind: AquaCsp
metadata:
  name: aqua
  namespace: aqua
spec:
  infra:                                    
    serviceAccount: "aqua-sa"               
    namespace: "aqua"                       
    version: "6.2"                          
    requirements: true                      
  common:
    imagePullSecret: "aqua-registry"        # Optional: If already created image pull secret then mention in here
    dbDiskSize: 10
    splitDB: true      
  externalDb:
    host: "<<EXTERNAL DATABASE IP>>"
    port: "<<EXTERNAL DATABASE PORT>>"
    username: "<<EXTERNAL DATABASE USER NAME>>"
    password: "<<EXTERNAL DATABASE PASSWORD>>"    # Optional: you can specify the database password secret in common.databaseSecret
  auditDB:
    information:
      host: "<<AUDIT EXTERNAL DB IP>>"
      port: "<<AUDIT EXTERNAL DB PORT>>"
      username: "<<AUDIT EXTERNAL DB USER NAME>>"
      password: "<<AUDIT EXTERNAL DB PASSWORD>>"  # Optional: you can specify the database password secret in auditDB.secret
    secret:                                       # Optional: the secret that hold the audit database password. will create one if not provided
      key: 
      name:                     
  gateway:                                  
    replicas: 1                             
    service: "ClusterIP"
    image:
      registry: "registry.aquasec.com"
      repository: "gateway"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always                     
  server:                                   
    replicas: 1                             
    service: "LoadBalancer" 
    image:
      registry: "registry.aquasec.com"
      repository: "console"
      tag: "<<IMAGE TAG>>"
      pullPolicy: Always  
  route: true                               # Optional: If defined and set to true, the Operator will create a Route to enable access to the console
```

#### Example: Deploying Aqua Enforcer(s)

If you haven't deployed any Aqua Enforcers, or if you want to deploy additional Enforcers, follow the instructions [here](https://github.com/aquasecurity/aqua-operator/blob/master/deploy/crds/operator_v1alpha1_aquaenforcer_cr.yaml).

This is an example of a simple Enforcer deployment: 
```yaml
---
apiVersion: operator.aquasec.com/v1alpha1
kind: AquaEnforcer
metadata:
  name: aqua
spec:
  infra:                                    
    serviceAccount: "aqua-sa"                
    version: "6.2"                          # Optional: auto generate to latest version
  common:
    imagePullSecret: "aqua-registry"        # Optional: if already created image pull secret then mention in here
  deploy:                                   # Optional: information about Aqua Enforcer deployment
    image:                                  # Optional: take the default value and version from infra.version
      repository: "enforcer"                # Optional: default = enforcer
      registry: "registry.aquasec.com"      # Optional: default = registry.aquasec.com
      tag: "<<IMAGE TAG>>"                  # Optional: default = 5.3
      pullPolicy: "IfNotPresent"            # Optional: default = IfNotPresent
  gateway:                                  # Required: data about the gateway address
    host: aqua-gateway
    port: 8443
  token: "<<your-token>>"                   # Required: The Enforcer group token can use an existing secret instead (you can create a token from the Aqua console)
```

#### Example: Deploying the KubeEnforcer

Here is an example of a KubeEnforcer deployment:
```yaml
apiVersion: operator.aquasec.com/v1alpha1
kind: AquaKubeEnforcer
metadata:
  name: aqua
spec:
  infra:
    version: '6.2'
    serviceAccount: aqua-kube-enforcer-sa
  config:
    gateway_address: 'aqua-gateway.aqua:8443'     # Required: provide <<AQUA GW IP OR DNS: AQUA GW PORT>>
    cluster_name: aqua-secure                     # Required: provide your cluster name
    imagePullSecret: aqua-registry                # Required: provide the imagePullSecret name
  deploy:
    service: ClusterIP
    image:
      registry: "registry.aquasec.com"
      tag: "<<KUBE_ENFORCER_TAG>>"
      repository: kube-enforcer
      pullPolicy: Always
  token: <<KUBE_ENFORCER_GROUP_TOKEN>>            # Optional: The KubeEnforcer group token (if not provided manual approval will be required)    
```

#### Example: Deploy the Aqua Scanner

You can deploy more Scanners; here is an example:
```yaml
apiVersion: operator.aquasec.com/v1alpha1
kind: AquaScanner
metadata:
  name: aqua
  namespace: aqua
spec:
  infra:
    serviceAccount: aqua-sa
    version: '6.2'
  deploy:
    replicas: 1
    image:
      registry: "registry.aquasec.com"
      repository: "scanner"
      tag: "<<IMAGE TAG>>"
  login:
    username: "<<YOUR AQUA USER NAME>>"
    password: "<<YOUR AQUA USER PASSWORD>>"
    host: 'http://aqua-server:8080'    #Required: provide <<(http:// or https://)Aqua Server IP or DNS: Aqua Server port>>
```