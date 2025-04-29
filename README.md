# Qdrant DBaaS Prototype

### Project Structure
```bash
pkg/
├── provider/                         # Interfaces and cloud provider implementations
│   ├── provider.go                   # Common interfaces: Provider, ClusterConfig, etc.
│   ├── models/                       # Shared provider models (optional, if exists)
│   ├── amazon/                       # AWS-specific implementation
│   │   ├── amazon.go                 # AWSProvider (manages services)
│   │   ├── models.go                 # InfrastructureInfo, ClusterInfo
│   │   ├── access/                   # IAM access management
│   │   │   └── access.go
│   │   ├── controlplane/             # EKS Control Plane management
│   │   │   └── controlplane.go
│   │   ├── network/                  # VPC and subnet management
│   │   │   └── network.go
│   │   ├── nodegroups/               # EC2 Node Groups management
│   │   │   └── nodegroups.go
│   │   ├── inputs/                   # Input configuration structures
│   │   │   └── controlplane.go
│   │   ├── config.go                 # AWS provider config model
│   │   └── errors.go                 # AWS-specific errors
│   ├── google/                       # GCP stub
│   │   └── google.go
│   └── microsoft/                    # Azure stub
│       └── microsoft.go
│
├── factory/                          # Provider factory
│   ├── factory.go
│   └── errors.go
│
├── helm/                             # Helm management abstraction
│   ├── client/                       # Helm client wrapper
│   │   └── client.go
│   ├── chart/                        # Chart loading utilities
│   │   └── load.go
│   ├── install/                      # Helm installation logic
│   │   ├── builder.go
│   │   └── install.go
│   ├── uninstall/                    # Helm uninstallation logic
│   │   └── uninstall.go
│   └── errors.go                     # Helm-related errors
│
├── service/                          # Higher-level services
│   ├── helm_service.go               # HelmService abstraction
│   └── cluster_service.go            # ClusterService: create infra + deploy apps
│
└── qdrant/                           # Qdrant-specific defaults
    ├── defaults.go                   # Default Qdrant environment values
    └── values.go                     # Qdrant values model (Env, Resources)
```

### Call Stack
```bash
ClusterService.CreateAndDeployCluster(cloud, clusterCfg, createOpts)
├── factory.NewProvider(cloud, clusterCfg)
│   ├── amazon.New()  # or google.New(), microsoft.New()
│   └── returns provider.Provider
│
├── provider.Provider.CreateCluster(clusterCfg, createOpts)
│   ├── (AWS) network.Service.Create(clusterCfg)
│   │   └── returns NetworkInfo (VPC, Subnets)
│   ├── (AWS) controlplane.Service.Create(clusterCfg)
│   │   └── returns ControlPlaneInfo (EKS API server)
│   ├── (AWS) nodegroups.Service.Create(clusterCfg)
│   │   └── returns NodeGroupsInfo (EC2 instances)
│   └── (AWS) access.Service.Create(clusterCfg)
│       └── returns AccessInfo (IAM roles for access)
│
├── Collects InfrastructureInfo
│
├── helmService.Install(InstallRequest)
│   ├── helm/client.NewClient()         # Initializes Helm client
│   ├── helm/chart.Load(chartPath)      # Loads Helm chart
│   ├── helm/install.NewBuilder()       # Prepares installation action
│   └── install.Run(chart, values)      # Deploys Helm release
│
└── Qdrant successfully deployed into the created cluster
```

### Usage Example
```bash
helmSvc := helmservice.NewHelmService(...)
clusterSvc := service.NewClusterService(helmSvc)

err := clusterSvc.CreateAndDeployCluster("aws", provider.ClusterConfig{
    Name:       "test-cluster",
    Region:     "us-east-1",
    NodeCount:  3,
    NodeType:   "t3.medium",
    K8sVersion: "1.32",
}, provider.CreateOptions{
    CreateNetwork:      true,
    CreateControlPlane: true,
    CreateNodeGroups:   true,
    CreateAccess:       true,
})
if err != nil {
    panic(err)
}
```