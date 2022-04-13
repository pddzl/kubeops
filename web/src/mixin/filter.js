// filter

// 1. namespace

const statusNsMap = {
  'Active': 'success',
}

export const statusNsFilter = (status) => {
  return statusNsMap[status] || 'info'
}

// 2. node

const statusNodeTypeMap = {
  'True': 'success',
  'Unknown': 'danger',
}

export const statusNodeTypeFilter = (status) => {
  return statusNodeTypeMap[status] || 'info'
}

const statusNodeMap = {
  'True': 'Ready',
  'Unknown': 'NotReady'
}

export const statusNodeFilter = (status) => {
  return statusNodeMap[status] || 'Unknown'
}

// 3. pod

const statusPodMap = {
  'Running': 'success'
}

export const statusPodFilter = (status) => {
  return statusPodMap[status] || 'info'
}

// 4. deploymentDetail -> replicaSet

const statusRsMap = {
  'True': 'success',
  'False': 'danger',
  'Unknown': 'info',
}

export const statusRsFilter = (status) => {
  return statusRsMap[status] || 'info'
}
