---
sidebar_position: 6
sidebar_label: "Fast Failover"
---

# Application Fast Failover

When the stateful application (i.e. Pod with HwameiStor volume) runs into a problem, especially caused by the node issue, 
it's important to reschedule the Pod to another health node and keep running.

However, due to the design of the Kubernetes' StatefulSet and Deployment, 
it will wait a long time (e.g. 5 mins) before rescheduling the Pod. 
Especially, it will never reschedule the Pod automatically for the StatefulSet Pod.
This will cause the application stop, and even cause a huge business loss.

HwameiStor provides a feature of fast failover to solve this problem. When identifying the application issue, 
it will reschedule the Pod immediately without waiting for a very long time. 
HwameiStor will fail the Pod over to another health node, and ensure the required data volumes are also located at the node. 
So, the application can continue to work. 

# How to use

HwameiStor provides the fast failover considering the two cases:

* Node Failure  
  
  When a node fails, all the Pods on this node can't work any more。As to the Pod using HwameiStor volume，
  it's necessary to reschedule to another health node with the associated data volume replica.
  User can trigger the fast failover for this node by:
  ```
  Add a label to this node:

  kubectl label node <nodeName> hwameistor.io/failover=start

  When the fast failover completes, the label will be modified as:

  hwameistor.io/failover=completed
  ```
  
* Pod Failure

  When a Pod fails, user can trigger the fast failover for it by:
  ```
  Add a lable to this Pod:

  kubectl label pod <podName> hwameistor.io/failover=start

  When the fast failover completes, the old Pod will be deleted and then the new one will be created on a new node.
  ```
