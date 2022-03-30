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
