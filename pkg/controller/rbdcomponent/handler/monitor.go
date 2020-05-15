package handler

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"

	rainbondv1alpha1 "github.com/goodrain/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	"github.com/goodrain/rainbond-operator/pkg/util/commonutil"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MonitorName name for rbd-monitor.
var MonitorName = "rbd-monitor"

type monitor struct {
	ctx        context.Context
	client     client.Client
	etcdSecret *corev1.Secret

	component *rainbondv1alpha1.RbdComponent
	cluster   *rainbondv1alpha1.RainbondCluster
	labels    map[string]string

	pvcParametersRWO *pvcParameters
	storageRequest   int64
}

var _ ComponentHandler = &monitor{}
var _ StorageClassRWOer = &monitor{}

// NewMonitor returns a new rbd-monitor handler.
func NewMonitor(ctx context.Context, client client.Client, component *rainbondv1alpha1.RbdComponent, cluster *rainbondv1alpha1.RainbondCluster) ComponentHandler {
	return &monitor{
		ctx:    ctx,
		client: client,

		component:      component,
		cluster:        cluster,
		labels:         LabelsForRainbondComponent(component),
		storageRequest: getStorageRequest("MONITOR_DATA_STORAGE_REQUEST", 21),
	}
}

func (m *monitor) Before() error {
	secret, err := etcdSecret(m.ctx, m.client, m.cluster)
	if err != nil {
		return fmt.Errorf("failed to get etcd secret: %v", err)
	}
	m.etcdSecret = secret

	if err := setStorageCassName(m.ctx, m.client, m.component.Namespace, m); err != nil {
		return err
	}

	return nil
}

func (m *monitor) Resources() []interface{} {
	return []interface{}{
		m.statefulset(),
		m.serviceForMonitor(),
	}
}

func (m *monitor) After() error {
	return nil
}

func (m *monitor) ListPods() ([]corev1.Pod, error) {
	return listPods(m.ctx, m.client, m.component.Namespace, m.labels)
}

func (m *monitor) SetStorageClassNameRWO(pvcParameters *pvcParameters) {
	m.pvcParametersRWO = pvcParameters
}

func (m *monitor) statefulset() interface{} {
	claimName := "data" // unnecessary
	promDataPVC := createPersistentVolumeClaimRWO(m.component.Namespace, claimName, m.pvcParametersRWO, m.labels, m.storageRequest)

	args := []string{
		"--advertise-addr=$(POD_IP):9999",
		"--alertmanager-address=$(POD_IP):9093",
		"--storage.tsdb.path=/prometheusdata",
		"--storage.tsdb.no-lockfile",
		"--storage.tsdb.retention=7d",
		fmt.Sprintf("--log.level=%s", m.component.LogLevel()),
		"--etcd-endpoints=" + strings.Join(etcdEndpoints(m.cluster), ","),
	}
	volumeMounts := []corev1.VolumeMount{
		{
			Name:      claimName,
			MountPath: "/prometheusdata",
		},
	}
	var volumes []corev1.Volume
	if m.etcdSecret != nil {
		volume, mount := volumeByEtcd(m.etcdSecret)
		volumeMounts = append(volumeMounts, mount)
		volumes = append(volumes, volume)
		args = append(args, etcdSSLArgs()...)
	}

	ds := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      MonitorName,
			Namespace: m.component.Namespace,
			Labels:    m.labels,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: m.component.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: m.labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   MonitorName,
					Labels: m.labels,
				},
				Spec: corev1.PodSpec{
					ImagePullSecrets:              imagePullSecrets(m.component, m.cluster),
					TerminationGracePeriodSeconds: commonutil.Int64(0),
					ServiceAccountName:            "rainbond-operator",
					Containers: []corev1.Container{
						{
							Name:            MonitorName,
							Image:           m.component.Spec.Image,
							ImagePullPolicy: m.component.ImagePullPolicy(),
							Env: []corev1.EnvVar{
								{
									Name: "POD_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "status.podIP",
										},
									},
								},
							},
							Args:         args,
							VolumeMounts: volumeMounts,
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: resource.MustParse("128Mi"),
									corev1.ResourceCPU:    resource.MustParse("0m"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: resource.MustParse("2048Mi"),
									corev1.ResourceCPU:    resource.MustParse("500m"),
								},
							},
						},
					},
					Volumes: volumes,
				},
			},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{*promDataPVC},
		},
	}

	return ds
}

func (m *monitor) serviceForMonitor() interface{} {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      MonitorName,
			Namespace: m.component.Namespace,
			Labels:    m.labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 9999,
					TargetPort: intstr.IntOrString{
						IntVal: 9999,
					},
				},
			},
			Selector: m.labels,
		},
	}

	return svc
}
