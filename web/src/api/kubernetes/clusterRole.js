import service from '@/utils/request'

// 获取clusterRole list
export const getClusterRoleList = (data) => {
  return service({
    url: '/clusterRole/getClusterRoleList',
    method: 'post',
    data
  })
}

// 获取clusterRole编排
export const getClusterRoleRaw = (data) => {
  return service({
    url: '/clusterRole/getClusterRoleRaw',
    method: 'post',
    data
  })
}

// 获取clusterRole详情
export const getClusterRoleDetail = (data) => {
  return service({
    url: '/clusterRole/getClusterRoleDetail',
    method: 'post',
    data
  })
}
