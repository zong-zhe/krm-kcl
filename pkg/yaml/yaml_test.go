package yaml

import "testing"

const (
	workload = `

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.4
  creationTimestamp: null
  name: containerizedworkloads.core.oam.dev
spec:
  group: core.oam.dev
  names:
    kind: ContainerizedWorkload
    listKind: ContainerizedWorkloadList
    plural: containerizedworkloads
    singular: containerizedworkload
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: A ContainerizedWorkload is a workload that runs OCI containers.
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: A ContainerizedWorkloadSpec defines the desired state of a
            ContainerizedWorkload.
          properties:
            arch:
              description: CPUArchitecture required by this workload.
              enum:
              - i386
              - amd64
              - arm
              - arm64
              type: string
            containers:
              description: Containers of which this workload consists.
              items:
                description: A Container represents an Open Containers Initiative
                  (OCI) container.
                properties:
                  args:
                    description: Arguments to be passed to the command run by this
                      container.
                    items:
                      type: string
                    type: array
                  command:
                    description: Command to be run by this container.
                    items:
                      type: string
                    type: array
                  config:
                    description: ConfigFiles that should be written within this container.
                    items:
                      description: A ContainerConfigFile specifies a configuration
                        file that should be written within a container.
                      properties:
                        path:
                          description: Path within the container at which the configuration
                            file should be written.
                          type: string
                        value:
                          description: Value that should be written to the configuration
                            file.
                          type: string
                      required:
                      - path
                      - value
                      type: object
                    type: array
                  env:
                    description: Environment variables that should be set within this
                      container.
                    items:
                      description: A ContainerEnvVar specifies an environment variable
                        that should be set within a container.
                      properties:
                        name:
                          description: Name of the environment variable. Must be composed
                            of valid Unicode letter and number characters, as well
                            as _ and -.
                          pattern: ^[-_a-zA-Z0-9]+$
                          type: string
                        value:
                          description: Value of the environment variable.
                          type: string
                      required:
                      - name
                      - value
                      type: object
                    type: array
                  image:
                    description: Image this container should run. Must be a path-like
                      or URI-like representation of an OCI image. May be prefixed
                      with a registry address and should be suffixed with a tag.
                    type: string
                  imagePullSecret:
                    description: ImagePullSecret specifies the name of a Secret from
                      which the credentials required to pull this container's image
                      can be loaded.
                    type: string
                  livenessProbe:
                    description: A LivenessProbe assesses whether this container is
                      alive. Containers that fail liveness probes will be restarted.
                    properties:
                      exec:
                        description: Exec probes a container's health by executing
                          a command.
                        properties:
                          command:
                            description: Command to be run by this probe.
                            items:
                              type: string
                            type: array
                        required:
                        - command
                        type: object
                      failureThreshold:
                        description: FailureThreshold specifies how many consecutive
                          probes must fail in order for the container to be considered
                          healthy.
                        format: int32
                        type: integer
                      httpGet:
                        description: HTTPGet probes a container's health by sending
                          an HTTP GET request.
                        properties:
                          httpHeaders:
                            description: HTTPHeaders to send with the GET request.
                            items:
                              description: A HTTPHeader to be passed when probing
                                a container.
                              properties:
                                name:
                                  description: Name of this HTTP header. Must be unique
                                    per probe.
                                  type: string
                                value:
                                  description: Value of this HTTP header.
                                  type: string
                              required:
                              - name
                              - value
                              type: object
                            type: array
                          path:
                            description: Path to probe, e.g. '/healthz'.
                            type: string
                          port:
                            description: Port to probe.
                            format: int32
                            type: integer
                        required:
                        - path
                        - port
                        type: object
                      initialDelaySeconds:
                        description: InitialDelaySeconds after a container starts
                          before the first probe.
                        format: int32
                        type: integer
                      periodSeconds:
                        description: PeriodSeconds between probes.
                        format: int32
                        type: integer
                      successThreshold:
                        description: SuccessThreshold specifies how many consecutive
                          probes must success in order for the container to be considered
                          healthy.
                        format: int32
                        type: integer
                      tcpSocket:
                        description: TCPSocketProbe probes a container's health by
                          connecting to a TCP socket.
                        properties:
                          port:
                            description: Port this probe should connect to.
                            format: int32
                            type: integer
                        required:
                        - port
                        type: object
                      timeoutSeconds:
                        description: TimeoutSeconds after which the probe times out.
                        format: int32
                        type: integer
                    type: object
                  name:
                    description: Name of this container. Must be unique within its
                      workload.
                    type: string
                  ports:
                    description: Ports exposed by this container.
                    items:
                      description: A ContainerPort specifies a port that is exposed
                        by a container.
                      properties:
                        containerPort:
                          description: Port number. Must be unique within its container.
                          format: int32
                          type: integer
                        name:
                          description: Name of this port. Must be unique within its
                            container. Must be lowercase alphabetical characters.
                          pattern: ^[a-z]+$
                          type: string
                        protocol:
                          description: Protocol used by the server listening on this
                            port.
                          enum:
                          - TCP
                          - UDP
                          type: string
                      required:
                      - containerPort
                      - name
                      type: object
                    type: array
                  readiessProbe:
                    description: A ReadinessProbe assesses whether this container
                      is ready to serve requests. Containers that fail readiness probes
                      will be withdrawn from service.
                    properties:
                      exec:
                        description: Exec probes a container's health by executing
                          a command.
                        properties:
                          command:
                            description: Command to be run by this probe.
                            items:
                              type: string
                            type: array
                        required:
                        - command
                        type: object
                      failureThreshold:
                        description: FailureThreshold specifies how many consecutive
                          probes must fail in order for the container to be considered
                          healthy.
                        format: int32
                        type: integer
                      httpGet:
                        description: HTTPGet probes a container's health by sending
                          an HTTP GET request.
                        properties:
                          httpHeaders:
                            description: HTTPHeaders to send with the GET request.
                            items:
                              description: A HTTPHeader to be passed when probing
                                a container.
                              properties:
                                name:
                                  description: Name of this HTTP header. Must be unique
                                    per probe.
                                  type: string
                                value:
                                  description: Value of this HTTP header.
                                  type: string
                              required:
                              - name
                              - value
                              type: object
                            type: array
                          path:
                            description: Path to probe, e.g. '/healthz'.
                            type: string
                          port:
                            description: Port to probe.
                            format: int32
                            type: integer
                        required:
                        - path
                        - port
                        type: object
                      initialDelaySeconds:
                        description: InitialDelaySeconds after a container starts
                          before the first probe.
                        format: int32
                        type: integer
                      periodSeconds:
                        description: PeriodSeconds between probes.
                        format: int32
                        type: integer
                      successThreshold:
                        description: SuccessThreshold specifies how many consecutive
                          probes must success in order for the container to be considered
                          healthy.
                        format: int32
                        type: integer
                      tcpSocket:
                        description: TCPSocketProbe probes a container's health by
                          connecting to a TCP socket.
                        properties:
                          port:
                            description: Port this probe should connect to.
                            format: int32
                            type: integer
                        required:
                        - port
                        type: object
                      timeoutSeconds:
                        description: TimeoutSeconds after which the probe times out.
                        format: int32
                        type: integer
                    type: object
                  resources:
                    description: Resources required by this container
                    properties:
                      cpu:
                        description: CPU required by this container.
                        properties:
                          required:
                            description: Required CPU count. 1.0 represents one CPU
                              core.
                            type: string
                        required:
                        - required
                        type: object
                      extended:
                        description: Extended resources required by this container.
                        items:
                          description: ExtendedResource required by a container.
                          properties:
                            name:
                              description: Name of the external resource. Resource
                                names are specified in kind.group/version format,
                                e.g. motionsensor.ext.example.com/v1.
                              type: string
                            required:
                              anyOf:
                              - type: integer
                              - type: string
                              description: Required extended resource(s), e.g. 8 or
                                "very-cool-widget"
                              x-kubernetes-int-or-string: true
                          required:
                          - name
                          - required
                          type: object
                        type: array
                      gpu:
                        description: GPU required by this container.
                        properties:
                          required:
                            description: Required GPU count.
                            type: string
                        required:
                        - required
                        type: object
                      memory:
                        description: Memory required by this container.
                        properties:
                          required:
                            description: Required memory.
                            type: string
                        required:
                        - required
                        type: object
                      volumes:
                        description: Volumes required by this container.
                        items:
                          description: VolumeResource required by a container.
                          properties:
                            accessMode:
                              description: AccessMode of this volume; RO (read only)
                                or RW (read and write).
                              enum:
                              - RO
                              - RW
                              type: string
                            disk:
                              description: Disk requirements of this volume.
                              properties:
                                ephemeral:
                                  description: Ephemeral specifies whether an external
                                    disk needs to be mounted.
                                  type: boolean
                                required:
                                  description: Required disk space.
                                  type: string
                              required:
                              - required
                              type: object
                            mountPath:
                              description: MouthPath at which this volume will be
                                mounted within its container.
                              type: string
                            name:
                              description: Name of this volume. Must be unique within
                                its container.
                              type: string
                            sharingPolicy:
                              description: SharingPolicy of this volume; Exclusive
                                or Shared.
                              enum:
                              - Exclusive
                              - Shared
                              type: string
                          required:
                          - mountPath
                          - name
                          type: object
                        type: array
                    required:
                    - cpu
                    - memory
                    type: object
                required:
                - image
                - name
                type: object
              type: array
            osType:
              description: OperatingSystem required by this workload.
              enum:
              - linux
              - windows
              type: string
          required:
          - containers
          type: object
        status:
          description: A ContainerizedWorkloadStatus represents the observed state
            of a ContainerizedWorkload.
          properties:
            conditions:
              description: Conditions of the resource.
              items:
                description: A Condition that may apply to a resource.
                properties:
                  lastTransitionTime:
                    description: LastTransitionTime is the last time this condition
                      transitioned from one status to another.
                    format: date-time
                    type: string
                  message:
                    description: A Message containing details about this condition's
                      last transition from one status to another, if any.
                    type: string
                  reason:
                    description: A Reason for this condition's last transition from
                      one status to another.
                    type: string
                  status:
                    description: Status of this condition; is it currently True, False,
                      or Unknown?
                    type: string
                  type:
                    description: Type of this condition. At most one of each condition
                      type may apply to a resource at any point in time.
                    type: string
                required:
                - lastTransitionTime
                - reason
                - status
                - type
                type: object
              type: array
            resources:
              description: Resources managed by this containerised workload.
              items:
                description: A ResourceReference refers to an resource managed by
                  an OAM resource.
                properties:
                  apiVersion:
                    description: APIVersion of the referenced resource.
                    type: string
                  kind:
                    description: Kind of the referenced resource.
                    type: string
                  name:
                    description: Name of the referenced resource.
                    type: string
                  uid:
                    description: UID of the referenced resource.
                    type: string
                required:
                - apiVersion
                - kind
                - name
                type: object
              type: array
          type: object
      type: object
  version: v1alpha2
  versions:
  - name: v1alpha2
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
	v1Crd = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com
spec:
  group: stable.example.com
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        # openAPIV3Schema is the schema for validating custom objects.
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                cronSpec:
                  type: string
                image:
                  type: string
                replicas:
                  type: integer
  scope: Namespaced
  names:
    plural: crontabs
    singular: crontab
    kind: CronTab
    shortNames:
      - ct
`
	v1beta1Crd = `
---
# Deprecated in v1.16 in favor of apiextensions.k8s.io/v1
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: crontabs.stable.example.com
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: stable.example.com
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                cronSpec:
                  type: string
                image:
                  type: string
                replicas:
                  type: integer
  # either Namespaced or Cluster
  scope: Namespaced
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: crontabs
    # singular name to be used as an alias on the CLI and for display
    singular: crontab
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: CronTab
    # shortNames allow shorter string to match your resource on the CLI
    shortNames:
      - ct
  preserveUnknownFields: false
`
)

func TestSplitDocuments(t *testing.T) {
	crds := v1Crd + v1beta1Crd
	files, _ := SplitDocuments(crds)
	if len(files) != 2 {
		t.Errorf("splitDocuments failed. expected 2, got %d", len(files))
	}
}

func TestSplitYamlStreamDocuments(t *testing.T) {
	crds := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foo.example.com
spec:
  group: example.com
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                field1:
                  type: string
  scope: Namespaced
  names:
    plural: foo
    singular: foo
    kind: Foo
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: bar.example.com
spec:
  group: example.com
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                field2:
                  type: string
  scope: Namespaced
  names:
    plural: bar
    singular: bar
    kind: Bar
`
	files, _ := SplitDocuments(crds)
	if len(files) != 2 {
		t.Errorf("splitDocuments failed. expected 2, got %d", len(files))
	}
	if files[0] != `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foo.example.com
spec:
  group: example.com
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                field1:
                  type: string
  scope: Namespaced
  names:
    plural: foo
    singular: foo
    kind: Foo` {
		panic(files[0])
	}
	if files[1] != `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: bar.example.com
spec:
  group: example.com
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                field2:
                  type: string
  scope: Namespaced
  names:
    plural: bar
    singular: bar
    kind: Bar
` {
		panic(files[1])
	}
}
