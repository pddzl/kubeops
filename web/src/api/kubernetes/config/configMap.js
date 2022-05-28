import service from '@/utils/request'

// 获取configMap list
export const getConfigMapList = (data) => {
  return service({
    url: '/configMap/getConfigMapList',
    method: 'post',
    data
  })
}

// 获取configMap编排
export const getConfigMapRaw = (data) => {
  return service({
    url: '/configMap/getConfigMapRaw',
    method: 'post',
    data
  })
}

// 获取configMap详情
export const getConfigMapDetail = (data) => {
  return service({
    url: '/configMap/getConfigMapDetail',
    method: 'post',
    data
  })
}

// 删除configMap
export const deleteConfigMap = (data) => {
  return service({
    url: '/configMap/deleteConfigMap',
    method: 'post',
    data
  })
}
