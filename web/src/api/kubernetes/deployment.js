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
