import service from '@/utils/request'

// 获取pod list
export const getReplicaSetList = (data) => {
  return service({
    url: '/replicaSet/getReplicaSetList',
    method: 'post',
    data
  })
}

// 获取replicaSet in raw
export const getReplicaSetRaw = (data) => {
  return service({
    url: '/replicaSet/getReplicaSetRaw',
    method: 'post',
    data
  })
}

// 获取replicaSet in detail
export const getReplicaSetDetail = (data) => {
  return service({
    url: '/replicaSet/getReplicaSetDetail',
    method: 'post',
    data
  })
}

// 获取replicaSet pods
export const getReplicaSetPods = (data) => {
  return service({
    url: '/replicaSet/getReplicaSetPods',
    method: 'post',
    data
  })
}
