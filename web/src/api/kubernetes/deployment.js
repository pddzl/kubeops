import service from '@/utils/request'

// 获取deployment list
export const getDeploymentList = (data) => {
  return service({
    url: '/deployment/getDeploymentList',
    method: 'post',
    data
  })
}

// 获取deployment in raw
export const getDeploymentRaw = (data) => {
  return service({
    url: '/deployment/getDeploymentRaw',
    method: 'post',
    data
  })
}

// 获取deployment in detail
export const getDeploymentDetail = (data) => {
  return service({
    url: '/deployment/getDeploymentDetail',
    method: 'post',
    data
  })
}

// 获取deployment关联的replicaset
export const getNewReplicaSet = (data) => {
  return service({
    url: '/deployment/getNewReplicaSet',
    method: 'post',
    data
  })
}

// 删除deployment
export const deleteDeployment = (data) => {
  return service({
    url: '/deployment/deleteDeployment',
    method: 'post',
    data
  })
}
